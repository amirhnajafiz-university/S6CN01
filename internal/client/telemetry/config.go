package telemetry

type Config struct {
	Address string `koanf:"address"`
	Enabled bool   `koanf:"enabled"`
}
