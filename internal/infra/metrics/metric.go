package metrics

import (
	"expvar"
	"strconv"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

type MetricsOptions struct {
	Enabled bool
}

type Metrics interface {
	RecordHttpRequestDuration(method, path string, status int, duration time.Duration)
	RecordHttpRequest(method, path string, status int)
	HTTPHandler() http.Handler
}

type metricsImpl struct {
	log 			*zerolog.Logger
	pool 			*pgxpool.Pool
	httpDuration 	*prometheus.HistogramVec
	httpRequests 	*prometheus.CounterVec
}

var (
	reg 							*prometheus.Registry
	onceMetrics 					sync.Once
	metricsInst 					Metrics
	httpRequestDurationBuckets = 	[]float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
)

func GetRegistry() *prometheus.Registry {
	return reg
}

func InitMetrics(log * zerolog.Logger, pool *pgxpool.Pool) Metrics {
	onceMetrics.Do(func() {
		reg = prometheus.NewRegistry()

		reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		reg.MustRegister(collectors.NewGoCollector())

		httpDuration := prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_request_duration_seconds",
				Help: "HTTP request duratoin in seconds",
				Buckets: httpRequestDurationBuckets,
			},
			[]string{"method", "path", "status"},
		)

		if err := reg.Register(httpDuration); err != nil {
			log.Warn().Err(err).Msg("Failed to register httpDuration metric")
		}

		httpRequests := prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_server_requests_total",
				Help: "Total number of HTTP requests by method, path and status",
			},
			[]string{"method", "path", "status"},
		)
		

		if err := reg.Register(httpRequests); err != nil {
			log.Warn().Err(err).Msg("failed to register httpRequests metric")
		}

		metricsInst = &metricsImpl{
			log: log,
			pool: pool,
			httpDuration: httpDuration,
			httpRequests: httpRequests,
		}

		if pool != nil {
			if err := reg.Register(NewPgxPoolCollector(pool)); err != nil {
				log.Warn().Err(err).Msg("Failed to register pgx collector")
			}
		}
	})

	return metricsInst
}

func (m *metricsImpl) RecordHttpRequestDuration(method, path string, status int, duration time.Duration) {
	m.httpDuration.WithLabelValues(method, path, strconv.Itoa(status)).Observe(duration.Seconds())
}

func (m *metricsImpl) RecordHttpRequest(method, path string, status int) {
	m.httpRequests.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
}

func (m *metricsImpl) HTTPHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/debug/vars" {
			expvar.Handler().ServeHTTP(w, r)
			return
		}

		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	})
}
