package telemetry

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "packet-monitoring"
	Subsystem = "agents"
)

// Metrics has all the client metrics.
type Metrics struct {
	CPUUtilization prometheus.Histogram `json:"CPUUtilization,omitempty"`
	CPULoad        prometheus.Histogram `json:"CPULoad,omitempty"`
	FreeMemory     prometheus.Histogram `json:"freeMemory,omitempty"`
	UsedMemory     prometheus.Histogram `json:"usedMemory,omitempty"`
	NetByteSend    prometheus.Histogram `json:"netByteSend,omitempty"`
	NetByteRec     prometheus.Histogram `json:"netByteRec,omitempty"`
	Requests       prometheus.Counter   `json:"requests,omitempty"`
}
