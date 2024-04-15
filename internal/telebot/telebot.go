package telebot

import (
	"context"
	"log/slog"
	"time"

	"github.com/inneroot/telenotify/internal/config"
	"github.com/inneroot/telenotify/internal/repository"
	tele "gopkg.in/telebot.v3"
)

var (
	SubscribedUsers = make(map[int]bool)
)

func Run(ctx context.Context, logger *slog.Logger, repo repository.SubscriberRepository) error {
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

	setHandlers(ctx, log, telebot, repo)

	go fakeUpdates(ctx, log, telebot, repo)
	log.Info("success: bot start")
	telebot.Start()
	return nil
}

func fakeUpdates(ctx context.Context, logger *slog.Logger, telebot *tele.Bot, repo repository.SubscriberRepository) {
	logger.Info("fakeUpdates start")
	for {
		time.Sleep(20 * time.Second)
		logger.Info("sending update")
		ids, err := repo.GetAll(ctx)
		if err != nil {
			logger.Debug("fakeUpdates", "error", err.Error())
			continue
		}
		for _, id := range ids {
			chat, err := telebot.ChatByID(int64(id))
			if err != nil {
				logger.Error("update error", "error", err.Error())
			} else {
				telebot.Send(chat, "update")
			}
		}
	}
}
