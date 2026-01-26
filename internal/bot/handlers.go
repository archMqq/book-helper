package bot

func initHandlers(srv *server) {
	srv.bot.Handle("/start", srv.helloHandle)

	srv.bot.Handle("/recommend", srv.getBooksHandle)
}
