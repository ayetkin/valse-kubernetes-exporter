package _errors

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Errors interface {
	ErrorWithFields(message string, fields logrus.Fields)
	WarnWithFields(message string, fields logrus.Fields)
	InfoWithFields(message string, fields logrus.Fields)
	ErrorWithException(message string, err error)
	WarnWithException(message string, err error)
}

func Handle(logger *logrus.Logger) Errors {
	return &errors{log: logger}
}

func NewLogger() *logrus.Logger {
	return newLogger()
}

func newLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{FullTimestamp: true},
		Level:     logrus.DebugLevel,
	}
}

type errors struct {
	log *logrus.Logger
}

func (e *errors) ErrorWithFields(message string, fields logrus.Fields) {
	e.log.WithFields(fields).Error(message)
}

func (e *errors) WarnWithFields(message string, fields logrus.Fields) {
	e.log.WithFields(fields).Warn(message)
}

func (e *errors) InfoWithFields(message string, fields logrus.Fields) {
	e.log.WithFields(fields).Info(message)
}

func (e *errors) ErrorWithException(message string, err error) {
	e.log.WithField("exception", err).Error(message)
}

func (e *errors) WarnWithException(message string, err error) {
	e.log.WithField("exception", err).Warn(message)
}
