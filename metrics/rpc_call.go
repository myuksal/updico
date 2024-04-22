package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	JsonRpcHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: CollectMetricJsonRpcName,
		Help: CollectMetricJsonRpcHelp,
	}, []string{CollectMetricJsonRpcVector_1},
	)
)
