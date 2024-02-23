package telebot

import (
	"context"
	"log"
	"time"

	"github.com/inneroot/telenotify/internal/config"
	tele "gopkg.in/telebot.v3"
)

func Run(ctx context.Context) error {
	log.Println("starting telegram bot")

	pref := tele.Settings{
		Token:  config.GetTgToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	telebot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return err
	}

	setHandlers(telebot)

	telebot.Start()
	return nil
}
