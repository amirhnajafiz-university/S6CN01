package agent

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Util struct {
}

// cpu utilization
func (_ Util) cpuUtilization() ([]float64, error) {
	return cpu.Percent(time.Second, false)
}

// cpu load
func (_ Util) cpuLoad() (float64, error) {
	info, err := load.Avg()

	return info.Load1, err
}

// free memory size
func (_ Util) freeMemory() (uint64, error) {
	memInfo, err := mem.VirtualMemory()

	return memInfo.Free, err
}

// used memory size
func (_ Util) usedMemory() (uint64, error) {
	memInfo, err := mem.VirtualMemory()

	return memInfo.Used, err
}

// net info
func (_ Util) getNetInfo() (uint64, uint64, error) {
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
