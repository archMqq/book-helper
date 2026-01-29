package bot

import (
	"database/sql"
	"log"
	"time"

	"github.com/archMqq/book-helper/internal/bot/config"
	"github.com/archMqq/book-helper/internal/bot/repository"
	"github.com/archMqq/book-helper/internal/bot/services"
	_ "github.com/lib/pq"
	tele "gopkg.in/telebot.v4"
)

func Start(cfg *config.Config) {
	b := initBot(cfg)
	db, err := newDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUser(db)
	userService := services.NewUserService(userRepo)

	srv := newServer(b, userService)
	initHandlers(srv)

	b.Start()
}

func initBot(cfg *config.Config) *tele.Bot {
	settings := tele.Settings{
		Token:  cfg.TGToken,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(settings)
	if err != nil {
		log.Fatalf("%s %s", "bot init err:", err)
	}

	return b
}

func newDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
