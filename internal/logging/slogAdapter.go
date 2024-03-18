package logging

import (
	"log/slog"
	"os"
)

type SlogAdapter struct {
	logger *slog.Logger
}

func NewSlogAdapter() Logger {
    return &SlogAdapter{
        logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
    }
}

func (s *SlogAdapter) Debug(msg string, args ...interface{}) {
    s.logger.Debug(msg, args...)
}

func (s *SlogAdapter) Info(msg string, args ...interface{}) {
    s.logger.Info(msg, args...)
}

func (s *SlogAdapter) Warn(msg string, args ...interface{}) {
    s.logger.Warn(msg, args...)
}

func (s *SlogAdapter) Error(msg string, args ...interface{}) {
    s.logger.Error(msg, args...)
}