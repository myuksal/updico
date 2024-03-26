package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MetricRpcCallLatency = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: MetricRpcCallLatencyName,
		Help: MetricRpcCallLatencyHelp,
	})
)
