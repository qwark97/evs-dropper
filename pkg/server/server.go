package server

import (
	"fmt"
	"log"
	"net/http"
)

type srv struct {
	logger *log.Logger

	listener  IListener
	formatter IFormatter
	presenter IPresenter

	mux *http.ServeMux
}

func new(logger *log.Logger, l IListener, f IFormatter, p IPresenter) *srv {

	m := http.NewServeMux()

	s := srv{
		logger:    logger,
		listener:  l,
		formatter: f,
		presenter: p,
		mux:       m,
	}
	return &s
}

func (s *srv) serve(addr string, port int) error {
	fullAddr := fmt.Sprintf("%s:%d", addr, port)
	s.logger.Printf("start serving at: %s", fullAddr)
	err := http.ListenAndServe(fullAddr, s.mux)
	return err
}

func (s *srv) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\n"))
}
