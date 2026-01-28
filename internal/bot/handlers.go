package bot

import "gopkg.in/telebot.v4"

func initHandlers(srv *server) {
	srv.bot.Handle("/start", srv.helloHandle)

	srv.bot.Handle("/recommend", srv.getBooksHandle)
	srv.bot.Handle("/authors", srv.saveAuthorsHandle)
	srv.bot.Handle("/genres", srv.saveGenresHandle)

	srv.bot.Handle(telebot.OnText, srv.textHandle)
}
