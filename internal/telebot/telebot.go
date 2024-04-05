package telebot

import (
	"context"
	"log/slog"
	"time"

	"github.com/inneroot/telenotify/internal/config"
	tele "gopkg.in/telebot.v3"
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

	log.Info("success: bot start")
	telebot.Start()
	return nil
}
