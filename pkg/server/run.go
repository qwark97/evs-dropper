package server

import (
	"github.com/qwark97/evs-dropper/pkg/formatters"
	"github.com/qwark97/evs-dropper/pkg/listeners"
	"github.com/qwark97/evs-dropper/pkg/logger"
	"github.com/qwark97/evs-dropper/pkg/presenters"
)

const (
	ADDR = "localhost"
	PORT = 8080

	NATS_ADDR = "localhost"
	NATS_PORT = 4222
)

func Run() error {
	logger := logger.New()
	logger.Info("initialized the logger")

	listenerConf := listeners.NatsConf{
		Addr: ADDR,
		Port: NATS_PORT,
	}
	listener := listeners.NewNats(logger, listenerConf)
	if err := listener.Connect(); err != nil {
		return err
	}
	defer listener.Disconnect()

	formatter := formatters.New()
	presenter := presenters.NewStdout()

	server := new(logger, listener, formatter, presenter)
	server.registerEndpoints()
	return server.serve(ADDR, PORT)
}
