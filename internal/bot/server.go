package bot

import (
	"github.com/archMqq/book-helper/internal/store"
	tele "gopkg.in/telebot.v4"
)

type server struct {
	bot   *tele.Bot
	store store.Store
}

func newServer(b *tele.Bot, store store.Store) *server {
	return &server{
		bot:   b,
		store: store,
	}
}
