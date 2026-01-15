package bot

import tele "gopkg.in/telebot.v4"

func initHandlers(srv *server) {

	srv.bot.Handle("/start", func(c tele.Context) error {
		userID := c.Sender().ID
		userName := c.Sender().FirstName

		u := srv.store.User().Register(userID, userName)
		if u == nil {
			return c.Send("С возвращением " + userName)
		}

		return c.Send(c.Text())
	})
}
