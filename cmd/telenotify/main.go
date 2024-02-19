package main

import (
	"log"
	"os"

	telebot "github.com/inneroot/telenotify/internal"
	"github.com/inneroot/telenotify/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.SetLogger()

	err := godotenv.Load("telegram-token.env")
	if err != nil {
		log.Println("no telegram-token.env provided", err)
	}
	token, ok := os.LookupEnv("TELEBOTTOKEN")
	if !ok {
		log.Fatal("TELEBOTTOKEN env must be provided")
	}

	telebot.Run(token)
}
