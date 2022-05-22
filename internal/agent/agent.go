package agent

import (
	"encoding/json"
	"fmt"
	"github.com/amirhnajafiz/packet-monitoring/internal/agent/socket"
	"time"

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

	p.CPUUtilization, err = a.Util.cpuUtilization()
	if err != nil {
		return p, err
	}

	p.CPULoad, err = a.Util.cpuLoad()
	if err != nil {
		return p, err
	}

	p.FreeMemory, err = a.Util.freeMemory()
	if err != nil {
		return p, err
	}

	p.UsedMemory, err = a.Util.usedMemory()
	if err != nil {
		return p, err
	}

	p.NetByteSend, p.NetByteRec, err = a.Util.getNetInfo()
	if err != nil {
		return p, err
	}

	return p, nil
}

func (a Agent) Start() {
	go func() {
		c, err := makeConnection(socket.Config{
			ServerHost: "localhost",
			ServerPort: "8080",
			ServerType: "tcp",
		})
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

			_ = c.Write(b)
		}
	}()
}
