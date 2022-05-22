package agent

import (
	"fmt"
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
		for {
			a.busyJob()
			fmt.Println(a.createStatus())
		}
	}()
}
