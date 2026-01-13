package bot

import (
	"log"
	"time"

	"github.com/archMqq/book-helper/internal/config"
	tele "gopkg.in/telebot.v4"
)

func Start(cfg *config.Config) {
	b := initBot(cfg)

	InitHandlers(b)

	b.Start()
}

func initBot(cfg *config.Config) *tele.Bot {
	pref := tele.Settings{
		Token:  cfg.TGToken,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("%s %s", "bot init err:", err)
	}

	return b
}
