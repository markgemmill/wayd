package services

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var _logger *slog.Logger = application.DefaultLogger(slog.LevelDebug)

type LoggerService struct {
	log *slog.Logger
}

func (s *LoggerService) Name() string {
	return "LoggerService"
}

func (s *LoggerService) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	s.log.Debug(fmt.Sprintf("LoggerService.OnStartup... %s", options.Name))
	return nil
}

func (s *LoggerService) OnShutdown() error {
	s.log.Debug("LoggerService.OnShutdown...")
	return nil
}

func (l *LoggerService) Logger() *slog.Logger {
	return l.log
}

func (l *LoggerService) Debug(message string) {
	l.log.Debug(message)
}

func (l *LoggerService) Info(message string) {
	l.log.Info(message)
}

func (l *LoggerService) Warn(message string) {
	l.log.Warn(message)
}

func (l *LoggerService) Error(message string) {
	l.log.Error(message)
}

func (l *LoggerService) Fatal(message string) {

	l.log.Error(message)
	os.Exit(1)
}

func ApplicationLogger(logFile io.Writer) *LoggerService {
	logHandler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	_logger = slog.New(logHandler)
	return &LoggerService{
		log: _logger,
	}
}
