package logger

import (
	"haiken/entity"
	"log"
)

type LoggerImpl struct{}

func NewLogger() entity.Logger {
	return &LoggerImpl{}
}

const (
	LogPrefixError = "[Error]"
	LogPrefixWarn  = "[Warnning]"
	LogPrefixInfo  = "[Info]"
	LogPrefixDebug = "[Debug]"
)

func (l *LoggerImpl) Errorf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixError)
	log.Printf(format, args...)
}
func (l *LoggerImpl) Warnf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixWarn)
	log.Printf(format, args...)
}
func (l *LoggerImpl) Infof(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixInfo)
	log.Printf(format, args...)
}
func (l *LoggerImpl) Debugf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixDebug)
	log.Printf(format, args...)
}
