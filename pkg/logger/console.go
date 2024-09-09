package logger

import (
	"github.com/dusted-go/logging/prettylog"
	"go.uber.org/fx"
	"log/slog"
)

type LoggerProvider struct {
	Logger *slog.Logger
}

func NewLoggerProvider() *LoggerProvider {
	return &LoggerProvider{
		Logger: slog.New(prettylog.NewHandler(nil)),
	}
}

func (l *LoggerProvider) GetConsoleLogger() *slog.Logger {
	return l.Logger
}

var Module = fx.Options(
	fx.Provide(NewLoggerProvider),
)
