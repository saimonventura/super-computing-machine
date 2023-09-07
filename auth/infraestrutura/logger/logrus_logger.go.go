package logger

import (
	"super-computing-machine/auth/domain"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() domain.Logger {
	logger := logrus.New()
	return &LogrusLogger{logger}
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}
