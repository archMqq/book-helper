package bot

import (
	"errors"
	"strings"

	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v4"
)

var (
	ErrInternalServer = errors.New("Наблюдается ошибка на стороне сервера. Повторите попытку позже.")
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
	userID := c.Sender().ID
	username := c.Sender().FirstName

	if _, ok := s.userStates[userID]; !ok {

		err := s.userService.CreateUser(userID, username)
		if err == domain.ErrUserExists {
			return c.Send("С возвращением " + username)
		} else if err != nil {
			s.logger.Error(err)
			return c.Send(ErrInternalServer)
		}

		return c.Send(c.Text())
	} else {
		return c.Send("С возвращением " + username)
	}

}

func (s *server) getBooksHandle(c tele.Context) error {
	userID := c.Sender().ID
	pref, err := s.userService.GetPreferences(userID)
	if err == domain.ErrDatabaseRequest {
		return c.Send(ErrInternalServer)
	}

	rec, err := s.recService.GetBooks(pref)
	if err == domain.ErrGptIsDown {
		s.logger.Warn(err)
		return c.Send("Ошибка сервера при запросе рекомендаций.")
	}

	return c.Send(rec)
}

func (s *server) saveAuthorsHandle(c tele.Context) error {
	userID := c.Sender().ID
	s.userStates[userID] = StateWaitAuthors

	return c.Send("Перечислите ваших любимых авторов одним сообщением через запятую. Если таковые отсутствуют, отправьте \"-\"")
}

func (s *server) textHandle(c tele.Context) error {
	userID := c.Sender().ID

	switch s.userStates[userID] {
	case StateWaitAuthors:
		text := c.Text()

		authors := strings.Split(text, ",")
		if len(authors) == 1 && strings.TrimSpace(authors[0]) == "-" {
			c.Send("Жаль, что у вас нет любимых авторов. Однако, наши рекомендации обязательно помогут вам их найти.")
		} else {
			err := s.userService.SaveAuthors(userID, authors)
			if err != nil {
				s.logger.Error(err)
			}
		}

		return c.Send("Думаю теперь нам стоит заполнить любимые жанры. Для этого ннапиши \"/genres\"")
	}
	return c.Send("Прости, я ещё не научился поддерживать диалог.")
}
