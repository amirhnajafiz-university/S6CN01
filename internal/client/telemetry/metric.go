package telemetry

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "packet_monitoring"
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

func newCounter(counterOpts prometheus.CounterOpts) prometheus.Counter {
	ev := prometheus.NewCounter(counterOpts)

	if err := prometheus.Register(ev); err != nil {
		panic(err)
	}

	return ev
}

func newHistogram(histogramOpts prometheus.HistogramOpts) prometheus.Histogram {
	ev := prometheus.NewHistogram(histogramOpts)

	if err := prometheus.Register(ev); err != nil {
		panic(err)
	}

	return ev
}

func NewMetrics() Metrics {
	return Metrics{
		Requests: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "total_requests",
			Help:        "total number of requests sent to prometheus client",
			ConstLabels: nil,
		}),
		CPUUtilization: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "cpu_utilization",
			Help:        "average cpu utilization of system",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
		CPULoad: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "cpu_load",
			Help:        "average cpu load of system",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
		FreeMemory: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "free_memory",
			Help:        "free memory size",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
		UsedMemory: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "used_memory",
			Help:        "in used memory size",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
		NetByteSend: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "net_byte_send",
			Help:        "net bytes send",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
		NetByteRec: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "net_byte_rec",
			Help:        "net bytes receive",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
	}
}
