package logger

import "log"

func New() *log.Logger {
	l := log.Default()

	return l
}
