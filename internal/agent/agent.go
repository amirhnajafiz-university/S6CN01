package agent

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/amirhnajafiz/packet-monitoring/internal/agent/socket"
	"github.com/amirhnajafiz/packet-monitoring/internal/protocol"
)

type Agent struct {
	ID       int
	BusyTime int
	Util     Util
}

func (a Agent) busyJob() {
	fmt.Printf("[%d]Agent is working ...\n", a.ID)
	time.Sleep(time.Duration(a.BusyTime) * time.Second)
}

func (a Agent) createStatus() (*protocol.Protocol, error) {
	var (
		p   *protocol.Protocol
		err error
	)

	p.CPUUtilization, err = a.Util.CpuUtilization()
	if err != nil {
		return p, err
	}

	p.CPULoad, err = a.Util.CpuLoad()
	if err != nil {
		return p, err
	}

	p.FreeMemory, err = a.Util.FreeMemory()
	if err != nil {
		return p, err
	}

	p.UsedMemory, err = a.Util.UsedMemory()
	if err != nil {
		return p, err
	}

	p.NetByteSend, p.NetByteRec, err = a.Util.GetNetInfo()
	if err != nil {
		return p, err
	}

	return p, nil
}

func (a Agent) Start(cfg socket.Config) {
	go func() {
		c, err := makeConnection(cfg)
		if err != nil {
			return
		}

		for {
			a.busyJob()

			p, e := a.createStatus()
			if e != nil {
				continue
			}

			b, _ := json.Marshal(p)

			for {
				er := c.Write(b)
				if er != nil {
					continue
				}

				if er == net.ErrClosed {
					c = errorHandling(cfg)
				}

				ack, er := c.Read()
				if er != nil {
					continue
				}

				if ack == "1" {
					break
				}
			}
		}
	}()
}
