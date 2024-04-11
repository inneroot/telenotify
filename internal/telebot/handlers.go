package telebot

import (
	"fmt"
	"log/slog"
	"strconv"

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

	telebot.Handle("/subscribe", func(c tele.Context) error {
		recipient := c.Recipient()
		id, err := strconv.Atoi(recipient.Recipient())
		if err != nil {
			logger.Error("error converting recipient id", slog.String("error", err.Error()))
			telebot.Send(recipient, "error")
			return err
		}
		SubscribedUsers[id] = true
		subsStr := fmt.Sprintln(SubscribedUsers)
		slog.Info("subscribed user", slog.String("subs", subsStr))
		telebot.Send(recipient, "subscribed")
		return nil
	})
}
