package telebot

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func Run() {
	log.Println("starting telegram bot")

	pref := tele.Settings{
		Token:  os.Getenv("TELEBOTTOKEN"),
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

	b.Start()
	log.Println("bot is live")
}
