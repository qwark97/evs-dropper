package server

import (
	"encoding/json"
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

func (s *srv) startListening(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s.logger.Debug("HTTP %s %s", r.Method, r.URL)

	if err := s.listener.StartListening(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResp{
			Status: http.StatusInternalServerError,
			Reason: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *srv) stopListening(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s.logger.Debug("HTTP %s %s", r.Method, r.URL)

	err := s.listener.StopListening()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResp{
			Status: http.StatusInternalServerError,
			Reason: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
