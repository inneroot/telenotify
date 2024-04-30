package telebot

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/inneroot/telenotify/internal/config"
	"github.com/inneroot/telenotify/internal/repository"
	tele "gopkg.in/telebot.v3"
)

var (
	SubscribedUsers = make(map[int]bool)
)

type Bot struct {
	log     *slog.Logger
	telebot *tele.Bot
	repo    repository.SubscriberRepository
}

func MustInit(ctx context.Context, logger *slog.Logger, repo repository.SubscriberRepository) *Bot {
	bot, err := New(ctx, logger, repo)
	if err != nil {
		logger.Error("failed to start telebot", "err", err)
		os.Exit(1)
	}
	return bot
}

func New(ctx context.Context, logger *slog.Logger, repo repository.SubscriberRepository) (*Bot, error) {
	log := logger.With(slog.String("module", "telebot"))

	pref := tele.Settings{
		Token:  config.GetTgToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	telebot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	setHandlers(ctx, log, telebot, repo)

	return &Bot{
		log:     log,
		repo:    repo,
		telebot: telebot,
	}, nil
}

func (b *Bot) Run() {
	go func() {
		b.telebot.Start()
	}()
	b.log.Info("telebot started")
}

func (b *Bot) Stop() {
	b.log.Info("stopping telebot")
	b.telebot.Stop()
}

func (b *Bot) NotifySubscribed(ctx context.Context, message string) error {
	const op = "NotifySubscribed"
	b.log.Info(op)

	ids, err := b.repo.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	for _, id := range ids {
		chat, err := b.telebot.ChatByID(int64(id))
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		} else {
			b.telebot.Send(chat, message)
		}
	}
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
