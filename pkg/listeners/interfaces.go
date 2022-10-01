package listeners

type ILogger interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Error(format string, v ...any)
}
