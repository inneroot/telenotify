package telebot

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/inneroot/telenotify/internal/repository"
	tele "gopkg.in/telebot.v3"
)

func setHandlers(ctx context.Context, logger *slog.Logger, telebot *tele.Bot, repo repository.SubscriberRepository) {
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
		repo.Add(ctx, id)
		telebot.Send(recipient, "subscribed")
		return nil
	})
}
