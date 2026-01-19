package bot

import (
	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v4"
)

type server struct {
	bot    *tele.Bot
	logger *logrus.Logger
	store  store.Store
	router *mux.Router
}

func newServer(b *tele.Bot, store store.Store) *server {
	return &server{
		router: mux.NewRouter(),
		logger: logger.Init(),
		bot:    b,
		store:  store,
	}
}
