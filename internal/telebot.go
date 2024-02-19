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

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/start", func(c tele.Context) error {
		recipient := c.Recipient()
		log.Println(recipient.Recipient())
		b.Send(recipient, string(recipient.Recipient()))
		return nil
	})

	b.Start()

	log.Println("bot is live")
}
