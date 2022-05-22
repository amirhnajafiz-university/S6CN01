package client

import (
	"encoding/json"
	"net"

	"github.com/amirhnajafiz/packet-monitoring/internal/client/telemetry"
	"github.com/amirhnajafiz/packet-monitoring/internal/protocol"
)

type Client struct {
	Metric telemetry.Metrics
}

func (c Client) Start(cfg Config) {
	c.Metric = telemetry.NewMetrics()

	telemetry.NewServer(cfg.Telemetry).Start()

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

		c.Metric.Requests.Add(1)

		go c.processClient(connection)
	}
}

func (c Client) processClient(connection net.Conn) {
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

		c.Metric.CPUUtilization.Observe(p.CPUUtilization)
		c.Metric.CPULoad.Observe(p.CPULoad)
		c.Metric.FreeMemory.Observe(float64(p.FreeMemory))
		c.Metric.UsedMemory.Observe(float64(p.UsedMemory))
		c.Metric.NetByteSend.Observe(float64(p.NetByteSend))
		c.Metric.NetByteRec.Observe(float64(p.NetByteRec))
	}
}
