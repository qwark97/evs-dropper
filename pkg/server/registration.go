package server

const (
	// paths
	PING            = "/ping"
	START_LISTENING = "/start-listening"
	STOP_LISTENING  = "/stop-listening"
)

func (s *srv) registerEndpoints() {
	s.mux.HandleFunc(PING, s.ping)
	s.mux.HandleFunc(START_LISTENING, s.startListening)
	s.mux.HandleFunc(STOP_LISTENING, s.stopListening)

	s.logger.Info("registered routes")
}
