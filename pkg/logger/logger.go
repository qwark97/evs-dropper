package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	instance *log.Logger
}

func New() *Logger {
	l := log.Default()
	logger := Logger{instance: l}
	return &logger
}

func (l Logger) Debug(format string, v ...any) {
	adjustedFormat := fmt.Sprintf("DBG %s\n", format)
	l.instance.Printf(adjustedFormat, v...)
}

func (l Logger) Info(format string, v ...any) {
	adjustedFormat := fmt.Sprintf("INF %s\n", format)
	l.instance.Printf(adjustedFormat, v...)
}

func (l Logger) Error(format string, v ...any) {
	adjustedFormat := fmt.Sprintf("ERR %s\n", format)
	l.instance.Printf(adjustedFormat, v...)
}
