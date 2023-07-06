package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	currentCount = 0

	// The Prometheus metric that will be exposed.
	httpHitsForEntireApplication = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_app_http_hit_total",
			Help: "Total number of http hits of entire application.",
		},
	)
	httpHits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_app_http_hit_total_for_endpoints",
			Help: "Total number of http hits.",
		},
		[]string{"endpoint", "method"},
	)
	prometheusRegistry = prometheus.NewRegistry()
)

func init() {
	prometheusRegistry.MustRegister(httpHits, httpHitsForEntireApplication)

}

func MetricsHandler() http.Handler {
	return promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
}

func Count(endpoint string, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.WithLabelValues(endpoint, r.Method).Inc()
		httpHitsForEntireApplication.Inc()
		f(w, r) // original function call
	}
}
