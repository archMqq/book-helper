package bot

import (
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
