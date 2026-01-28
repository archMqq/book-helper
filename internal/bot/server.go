package bot

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v4"
)

var (
	ErrInternalServer = errors.New("Наблюдается ошибка на стороне сервера. Повторите попытку позже.")
	ErrServerIsBusy   = errors.New("В настоящий момент сервер слишком занят. Повторите попытку позже.")
)

type State int

const (
	StateNone State = iota
	StateRegistered
	StateWaitAuthors
	StateWaitGenres
	StateSearching
)

type server struct {
	bot         *tele.Bot
	router      *mux.Router
	logger      *logrus.Logger
	userService service.UserService
	recService  service.RecService
	userStates  map[int64]State
}

func newServer(b *tele.Bot, userService service.UserService, recService service.RecService) *server {
	return &server{
		router:      mux.NewRouter(),
		logger:      logger.Init(),
		bot:         b,
		userService: userService,
		recService:  recService,
		userStates:  make(map[int64]State),
	}
}

func (s *server) helloHandle(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userID := c.Sender().ID
	username := c.Sender().FirstName

	if _, ok := s.userStates[userID]; !ok {

		err := s.userService.CreateUser(ctx, userID, username)
		if errors.Is(err, context.DeadlineExceeded) {
			s.logger.Warn("timeout creating user", "userID", userID)
			return c.Send(ErrServerIsBusy)
		}
		if errors.Is(err, domain.ErrUserExists) {
			return c.Send("С возвращением " + username)
		} else if err != nil {
			s.logger.Error("unknown err ", err.Error())
			return c.Send(ErrInternalServer.Error())
		}

		return c.Send(c.Text())
	} else {
		return c.Send("С возвращением " + username)
	}

}

func (s *server) getBooksHandle(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	userID := c.Sender().ID
	pref, err := s.userService.GetPreferences(ctx, userID)
	if err != nil {
		if err == domain.ErrDatabaseRequest {
			s.logger.Error(fmt.Errorf("error db request: %w", err))
			return c.Send(ErrInternalServer)
		}
		if errors.Is(err, context.DeadlineExceeded) {
			s.logger.Warn("timeout preferences getting", "userID", userID)
			return c.Send(ErrServerIsBusy)
		}
	}

	rec, err := s.recService.GetBooks(ctx, pref)
	if err != nil {
		if err == domain.ErrGptIsDown {
			s.logger.Warn(err)
			return c.Send("Ошибка сервера при запросе рекомендаций.")
		}

		if errors.Is(err, context.DeadlineExceeded) {
			s.logger.Warn("timeout gpt asking")
			return c.Send(ErrServerIsBusy)
		}
	}

	return c.Send(rec)
}

func (s *server) saveAuthorsHandle(c tele.Context) error {
	userID := c.Sender().ID
	s.userStates[userID] = StateWaitAuthors

	return c.Send("Перечислите ваших любимых авторов одним сообщением через запятую. \n\nПример:\"Братья стругацкие, Гоголь, Пушкин\"\n\n Если таковые отсутствуют или не хотите их указывать, отправьте \"-\"")
}

func (s *server) saveGenresHandle(c tele.Context) error {
	userID := c.Sender().ID
	s.userStates[userID] = StateWaitGenres

	return c.Send("Перечислите ваши любимые жанры одним сообщением через запятую. \n\nПример:\"Фантастика, романы, детективы\"\n\n Если таковые отсутствуют или не хотите их указывать, отправьте \"-\"")
}

func (s *server) textHandle(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userID := c.Sender().ID

	switch s.userStates[userID] {
	case StateWaitAuthors:
		text := c.Text()

		authors := strings.Split(text, ",")
		if len(authors) == 1 && strings.TrimSpace(authors[0]) == "-" {
			c.Send("Жаль, что у вас нет любимых авторов. Однако, наши рекомендации обязательно помогут вам их найти.")
		} else {
			err := s.userService.SaveAuthors(ctx, userID, authors)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					s.logger.Warn("timeout authors saving", "authors", authors)
					return c.Send(ErrServerIsBusy)
				}

				s.logger.Error(err)
				return c.Send(ErrInternalServer)
			}
		}

		s.userStates[userID] = StateNone
		return c.Send("Думаю теперь нам стоит заполнить любимые жанры. Для этого ннапиши \"/genres\"")
	case StateWaitGenres:
		text := c.Text()

		genres := strings.Split(text, ", ")

		if len(genres) == 1 && strings.TrimSpace(genres[0]) == "-" {
			c.Send("Жаль. Я надеюсь это временно")
		} else {
			err := s.userService.SaveGenres(ctx, userID, genres)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					s.logger.Warn("timeout genres saving", "genres", genres)
					return c.Send(ErrServerIsBusy)
				}

				s.logger.Error(err)
				return c.Send(ErrInternalServer)
			}
		}
		s.userStates[userID] = StateNone
		return c.Send("Теперь можете посмотреть что мы можем вам предложить написав \"/recommend\"")

	}
	return c.Send("Прости, я ещё не научился поддерживать диалог.")
}
