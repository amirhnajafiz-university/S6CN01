package client

import "github.com/amirhnajafiz/packet-monitoring/internal/client/telemetry"

type Config struct {
	Telemetry  telemetry.Config `koanf:"telemetry"`
	ServerHost string           `koanf:"host"`
	ServerPort string           `koanf:"port"`
	ServerType string           `koanf:"type"`
}
