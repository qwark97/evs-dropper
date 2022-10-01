package server

import (
	"fmt"
	"net/http"
)

type srv struct {
	logger ILogger

	listener  IListener
	formatter IFormatter
	presenter IPresenter

	mux *http.ServeMux
}

func new(logger ILogger, l IListener, f IFormatter, p IPresenter) *srv {

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
	s.logger.Info("start serving at: %s", fullAddr)
	err := http.ListenAndServe(fullAddr, s.mux)
	return err
}

func (s *srv) ping(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("HTTP %s %s", r.Method, r.URL)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\n"))
}
