package socket

import (
	"net"
)

type Socket struct {
	Connection net.Conn
}

const (
	ServerHost = "localhost"
	ServerPort = "9988"
	ServerType = "tcp"
)

func (s *Socket) Connect() error {
	var err error

	s.Connection, err = net.Dial(ServerType, ServerHost+":"+ServerPort)

	return err
}

func (s *Socket) Disconnect() error {
	return s.Connection.Close()
}

func (s *Socket) Write(w string) error {
	_, err := s.Connection.Write([]byte(w))

	return err
}

func (s *Socket) Read() (string, error) {
	buffer := make([]byte, 1024)
	mLen, err := s.Connection.Read(buffer)

	return string(buffer[:mLen]), err
}
