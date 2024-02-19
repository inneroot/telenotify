package telebot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func Run(telegramToken string) {
	log.Println("starting telegram bot")

	pref := tele.Settings{
		Token:  telegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	telebot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	setHandlers(telebot)

	telebot.Start()
}
