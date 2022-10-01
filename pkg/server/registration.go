package server

const (
	// paths
	PING = "/ping"
)

func (s *srv) registerEndpoints() {
	s.mux.HandleFunc(PING, s.ping)

	s.logger.Info("registered routes")
}
