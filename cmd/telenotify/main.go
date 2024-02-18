package main

import (
	telebot "github.com/inneroot/telenotify/internal"
	"github.com/inneroot/telenotify/pkg/logger"
)

func main() {
	logger.SetLogger()
	telebot.Run()
}
