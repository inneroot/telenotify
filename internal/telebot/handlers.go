package telebot

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

func setHandlers(telebot *tele.Bot) {
	telebot.Handle("/ping", func(c tele.Context) error {
		return c.Send("Pong!")
	})

	telebot.Handle("/start", func(c tele.Context) error {
		recipient := c.Recipient()
		log.Println(recipient.Recipient())
		telebot.Send(recipient, string(recipient.Recipient()))
		return nil
	})
}
