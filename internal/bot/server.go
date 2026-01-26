package bot

import (
	"github.com/archMqq/book-helper/internal/domain"
	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v4"
)

type server struct {
	bot         *tele.Bot
	router      *mux.Router
	logger      *logrus.Logger
	userService service.UserService
	recService  service.RecService
}

func newServer(b *tele.Bot, userService service.UserService, recService service.RecService) *server {
	return &server{
		router:      mux.NewRouter(),
		logger:      logger.Init(),
		bot:         b,
		userService: userService,
		recService:  recService,
	}
}

func (s *server) helloHandle(c tele.Context) error {
	userID := c.Sender().ID
	username := c.Sender().FirstName

	err := s.userService.CreateUser(userID, username)
	if err != nil {
		return c.Send("С возвращением " + username)
	}

	return c.Send(c.Text())
}

func (s *server) getBooksHandle(c tele.Context) error {
	userID := c.Sender().ID
	pref, err := s.userService.GetPreferences(userID)
	if err == domain.ErrDatabaseRequest {
		return c.Send("Наблюдается ошибка на стороне сервера. Повторите попытку позже.")
	}

	rec, err := s.recService.GetBooks(pref)
	if err == domain.ErrGptIsDown {
		return c.Send("Ошибка сервера при запросе рекомендаций.")
	}

	return c.Send(rec)
}
