package bot

import (
	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v4"
)

type server struct {
	bot     *tele.Bot
	logger  *logrus.Logger
	service service.UserService
	router  *mux.Router
}

func newServer(b *tele.Bot, service service.UserService) *server {
	return &server{
		router:  mux.NewRouter(),
		logger:  logger.Init(),
		bot:     b,
		service: service,
	}
}
