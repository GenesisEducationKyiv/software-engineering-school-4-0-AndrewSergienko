package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	EmailsSentTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "emails_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"status"},
	)
	EmailsSentLastTime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "emails_last_time",
			Help: "Last time email was sent",
		},
		[]string{"status"},
	)
)

func RunServer() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("0.0.0.0:9000", nil) // nolint:all
	if err != nil {
		return
	}
}
