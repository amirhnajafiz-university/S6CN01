package socket

type Config struct {
	ServerHost string `koanf:"host"`
	ServerPort string `koanf:"port"`
	ServerType string `koanf:"type"`
}
