package telemetry

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server contains information about metrics server.
type Server struct {
	srv     *http.ServeMux
	address string
}

// NewServer creates a new monitoring server.
func NewServer(cfg Config) Server {
	var srv *http.ServeMux

	if cfg.Enabled {
		srv = http.NewServeMux()
		srv.Handle("/metrics", promhttp.Handler())
	}

	return Server{
		address: cfg.Address,
		srv:     srv,
	}
}

// Start creates and run a metric server for prometheus in new go routine.
func (s Server) Start() {
	go func() {
		_ = http.ListenAndServe(s.address, s.srv)
	}()
}
