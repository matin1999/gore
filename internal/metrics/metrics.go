package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	DatabaseErrorCounter         *prometheus.CounterVec
}

func InitMetrics() *Metrics {

	m := &Metrics{
		DatabaseErrorCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "",
				Name:      "",
				Help:      "",
			},
			[]string{"status_code", "channel"},
		),
	}

	prometheus.MustRegister(m.DatabaseErrorCounter)
	return m
}

func StartMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		panic(err)
	}
}
