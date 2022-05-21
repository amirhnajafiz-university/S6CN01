package agent

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Agent struct {
}

func (_ Agent) cpuUtilization() ([]float64, error) {
	return cpu.Percent(time.Second, false)
}

func (_ Agent) cpuLoad() (float64, error) {
	info, err := load.Avg()

	return info.Load1, err
}

// free memory size
func (_ Agent) freeMemory() (uint64, error) {
	memInfo, err := mem.VirtualMemory()

	return memInfo.Free, err
}

// used memory size
func (_ Agent) usedMemory() (uint64, error) {
	memInfo, err := mem.VirtualMemory()

	return memInfo.Used, err
}

// net info
func (_ Agent) getNetInfo() (uint64, uint64, error) {
	var (
		sent uint64 = 0
		rec  uint64 = 0
	)

	info, err := net.IOCounters(true)
	for _, v := range info {
		sent += v.BytesSent
		rec += v.BytesRecv
	}

	return sent, rec, err
}

func Do() {

}
