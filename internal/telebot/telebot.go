package telebot

import (
	"context"
	"log/slog"
	"time"

	"github.com/inneroot/telenotify/internal/config"
	tele "gopkg.in/telebot.v3"
)

var (
	SubscribedUsers = make(map[int]bool)
)

func Run(ctx context.Context, logger *slog.Logger) error {
	log := logger.With(slog.String("module", "telebot"))
	log.Info("starting telegram bot")

	pref := tele.Settings{
		Token:  config.GetTgToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	telebot, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	setHandlers(log, telebot)

	go func() {
		for {
			time.Sleep(20 * time.Second)
			log.Info("sending update")
			for sub := range SubscribedUsers {
				chat, err := telebot.ChatByID(int64(sub))
				if err != nil {
					telebot.Send(chat, "update error")
					log.Error("update error", "error", err.Error())
				}
				telebot.Send(chat, "update")
			}
		}
	}()
	log.Info("success: bot start")
	telebot.Start()
	return nil
}
