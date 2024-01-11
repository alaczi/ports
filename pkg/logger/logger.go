package logger

import "log"

type Logger interface {
	Log(v ...any)
	Logf(format string, v ...any)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(v ...any) {
	log.Print(v...)
}

func (c *ConsoleLogger) Logf(format string, v ...any) {
	log.Printf(format, v...)
}