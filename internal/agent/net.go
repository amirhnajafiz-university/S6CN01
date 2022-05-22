package agent

import (
	"time"

	"github.com/amirhnajafiz/packet-monitoring/internal/agent/socket"
)

func makeConnection(cfg socket.Config) (socket.Socket, error) {
	s := socket.Socket{
		Cfg: cfg,
	}

	err := s.Connect()
	if err != nil {
		return s, err
	}

	return s, nil
}

func errorHandling(cfg socket.Config) socket.Socket {
	for {
		s, err := makeConnection(cfg)
		if err == nil {
			return s
		}

		time.Sleep(1 * time.Second)
	}
}
