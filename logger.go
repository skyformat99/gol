package gol

import (
	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/adapter/file"
)

// Logger ...
type Logger interface {
	Debug(i ...interface{})
	Debugf(format string, i ...interface{})

	Info(i ...interface{})
	Infof(format string, i ...interface{})

	Warn(i ...interface{})
	Warnf(format string, i ...interface{})

	Error(i ...interface{})
	Errorf(format string, i ...interface{})

	Critical(i ...interface{})
	Criticalf(format string, i ...interface{})

	SetLevel(LogLevel)
	SetOption(LogOption)
	AddLogAdapter(adapter.Adapter)
}

// NewLogger create a Logger
func NewLogger(level LogLevel) Logger {
	logger := &gollog{
		level:   level,
		option:  LstdFlags,
		logChan: make(chan string, 1024),
	}
	logger.AddLogAdapter(file.NewConsoleAdapter())

	go logger.msgPump()
	return logger
}
