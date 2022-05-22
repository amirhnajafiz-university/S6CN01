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

// CpuUtilization cpu utilization
func (_ Util) CpuUtilization() (float64, error) {
	var total float64

	cores, err := cpu.Percent(time.Second, false)

	for _, core := range cores {
		total += core
	}

	total /= float64(len(cores))

	return total, err
}

// CpuLoad cpu load
func (_ Util) CpuLoad() (float64, error) {
	info, err := load.Avg()

	return info.Load1, err
}

// FreeMemory free memory size
func (_ Util) FreeMemory() (uint64, error) {
	memInfo, err := mem.VirtualMemory()

	return memInfo.Free, err
}

// UsedMemory used memory size
func (_ Util) UsedMemory() (uint64, error) {
	memInfo, err := mem.VirtualMemory()

	return memInfo.Used, err
}

// GetNetInfo net info
func (_ Util) GetNetInfo() (uint64, uint64, error) {
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
