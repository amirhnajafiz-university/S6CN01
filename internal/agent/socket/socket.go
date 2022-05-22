package socket

import (
	"net"
)

type Socket struct {
	Connection net.Conn
	Cfg        Config
}

func (s *Socket) Connect() error {
	var err error

	addr := s.Cfg.ServerHost + ":" + s.Cfg.ServerPort

	s.Connection, err = net.Dial(s.Cfg.ServerType, addr)

	return err
}

func (s *Socket) Disconnect() error {
	return s.Connection.Close()
}

func (s *Socket) Write(w []byte) error {
	_, err := s.Connection.Write(w)

	return err
}

func (s *Socket) Read() (string, error) {
	buffer := make([]byte, 1024)
	mLen, err := s.Connection.Read(buffer)

	return string(buffer[:mLen]), err
}
