package gokit

import (
	"github.com/MrLYC/gokit/builtins"
)

// Logger :
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

var logger Logger

// SetLogger :
func SetLogger(log Logger) {
	logger = log
}

// GetLogger :
func GetLogger() Logger {
	return logger
}

func init() {
	SetLogger(builtins.NewLogger())
}
