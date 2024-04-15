package telebot

import (
	"context"
	"log/slog"

	"github.com/inneroot/telenotify/internal/repository"
	tele "gopkg.in/telebot.v3"
)

func setHandlers(ctx context.Context, logger *slog.Logger, telebot *tele.Bot, repo repository.SubscriberRepository) {
	telebot.Handle("/ping", func(c tele.Context) error {
		recipient := c.Recipient()
		logger.Debug("/ping", "recipient", recipient.Recipient())
		return c.Send("Pong!")
	})

	telebot.Handle("/start", func(c tele.Context) error {
		recipient := c.Recipient()
		logger.Debug("/start", "recipient", recipient.Recipient())
		telebot.Send(recipient, recipient.Recipient())
		return nil
	})

	telebot.Handle("/subscribe", func(c tele.Context) error {
		recipient := c.Recipient()
		chatID := c.Message().Chat.ID
		repo.Add(ctx, chatID)
		telebot.Send(recipient, "subscribed")
		return nil
	})

	telebot.Handle("/unsubscribe", func(c tele.Context) error {
		recipient := c.Recipient()
		chatID := c.Message().Chat.ID
		repo.Del(ctx, chatID)
		telebot.Send(recipient, "unsubscribed")
		return nil
	})
}
