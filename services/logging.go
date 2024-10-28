package services

import (
	"log/slog"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var _logger = application.DefaultLogger(slog.LevelDebug)

type Logger struct {
	log *slog.Logger
}

func (l *Logger) Logger() *slog.Logger {
	return l.log
}

func (l *Logger) Debug(message string) {
	l.log.Debug(message)
}

func ApplicationLogger() *Logger {
	return &Logger{
		log: _logger,
	}
}
