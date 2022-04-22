package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	MetricsNamespace = "httpserver"
)

var (
	functionLatency = CreateExecutionTimeMetric(MetricsNamespace,
		"postHandle Time")

)

func Register() {
	err := prometheus.Register(functionLatency)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Help: help,
		Name: "prometheus_test",
		Buckets: prometheus.ExponentialBuckets(0.01, 2, 15),
	}, []string{"test"})
}

func NewTimer() *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: functionLatency,
		start: now,
	}
}

func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
}

