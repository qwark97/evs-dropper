package server

type IListener interface {
}
type IFormatter interface {
}
type IPresenter interface {
}
type ILogger interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Error(format string, v ...any)
}
