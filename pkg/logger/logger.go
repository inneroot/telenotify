package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/inneroot/telenotify/internal/config"
)

func SetLogger() *slog.Logger {
	level := slog.LevelInfo
	if config.IsDev() {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
	logger.Info(fmt.Sprintf("Logger level: %s", level.String()), "ENV", os.Getenv("ENV"))
	return logger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
