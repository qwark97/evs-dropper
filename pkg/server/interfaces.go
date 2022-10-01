package server

type IListener interface {
	StartListening()
	StopListening()
	DumpInfo()
	StreamTraffic()
	CleanData()
	Connect() error
	Disconnect()
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
