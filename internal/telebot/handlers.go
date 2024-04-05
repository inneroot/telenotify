package telebot

import (
	"log/slog"

	tele "gopkg.in/telebot.v3"
)

func setHandlers(logger *slog.Logger, telebot *tele.Bot) {
	telebot.Handle("/ping", func(c tele.Context) error {
		recipient := c.Recipient()
		logger.Info("/ping", "recipient", recipient.Recipient())
		return c.Send("Pong!")
	})

	telebot.Handle("/start", func(c tele.Context) error {
		recipient := c.Recipient()
		logger.Info("/start", "recipient", recipient.Recipient())
		telebot.Send(recipient, recipient.Recipient())
		return nil
	})
}
