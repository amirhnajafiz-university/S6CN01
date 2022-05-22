package config

import (
	"github.com/amirhnajafiz/packet-monitoring/internal/agent/socket"
	"github.com/amirhnajafiz/packet-monitoring/internal/client"
	"github.com/amirhnajafiz/packet-monitoring/internal/client/telemetry"
)

func Default() Config {
	return Config{
		Agent: socket.Config{
			ServerHost: "localhost",
			ServerPort: "8080",
			ServerType: "tcp",
		},
		Client: client.Config{
			ServerHost: "localhost",
			ServerPort: "8080",
			ServerType: "tcp",
		},
		Telemetry: telemetry.Config{
			Address: "1224",
			Enabled: true,
		},
	}
}
