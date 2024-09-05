package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
	RequestsDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
	RateSourceTotalRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_source_requests_total",
			Help: "Total number of requests to rate source",
		},
		[]string{"source", "status"},
	)
)

func RunServer() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("0.0.0.0:9000", nil) // nolint:all
	if err != nil {
		return
	}
}
