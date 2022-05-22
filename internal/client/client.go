package client

import (
	"encoding/json"
	"net"

	"github.com/amirhnajafiz/packet-monitoring/internal/protocol"
)

func Start(cfg Config) {
	addr := cfg.ServerHost + ":" + cfg.ServerPort

	server, err := net.Listen(cfg.ServerType, addr)
	if err != nil {
		panic(err)
	}

	defer func(server net.Listener) {
		_ = server.Close()
	}(server)

	for {
		connection, er := server.Accept()
		if er != nil {
			panic(er)
		}

		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	for {
		var p protocol.Protocol

		buffer := make([]byte, 1024)

		mLen, err := connection.Read(buffer)
		if err != nil {
			_, err = connection.Write([]byte("0"))

			continue
		}

		err = json.Unmarshal(buffer[:mLen], &p)
		if err != nil {
			_, err = connection.Write([]byte("0"))

			continue
		}

		_, err = connection.Write([]byte("1"))
	}
}
