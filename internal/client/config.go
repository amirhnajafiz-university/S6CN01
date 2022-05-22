package client

import "github.com/amirhnajafiz/packet-monitoring/internal/client/telemetry"

type Config struct {
	Telemetry  telemetry.Config
	ServerHost string
	ServerPort string
	ServerType string
}
