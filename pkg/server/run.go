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
)

func Run() error {
	logger := logger.New()
	logger.Println("initialized the logger")

	listener := listeners.NewNats()
	formatter := formatters.New()
	presenter := presenters.NewStdout()

	server := new(logger, listener, formatter, presenter)
	server.registerEndpoints()
	return server.serve(ADDR, PORT)
}
