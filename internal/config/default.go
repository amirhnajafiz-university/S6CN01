package config

import (
	"github.com/amirhnajafiz/packet-monitoring/internal/agent/socket"
	"github.com/amirhnajafiz/packet-monitoring/internal/client"
	"github.com/amirhnajafiz/packet-monitoring/internal/client/telemetry"
)

func Default() Config {
	return Config{
		NumberOfAgents: 4,
		Agent: socket.Config{
			ServerHost: "localhost",
			ServerPort: "8080",
			ServerType: "tcp",
		},
		Client: client.Config{
			Telemetry: telemetry.Config{
				Address: "localhost:1224",
				Enabled: true,
			},
			ServerHost: "localhost",
			ServerPort: "8080",
			ServerType: "tcp",
		},
	}
}
