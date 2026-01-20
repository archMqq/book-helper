package bot

import tele "gopkg.in/telebot.v4"

func initHandlers(srv *server) {
	srv.bot.Handle("/start", func(c tele.Context) error {
		userID := c.Sender().ID
		username := c.Sender().FirstName

		err := srv.userService.CreateUser(userID, username)
		if err != nil {
			return c.Send("С возвращением " + username)
		}

		return c.Send(c.Text())
	})
}
