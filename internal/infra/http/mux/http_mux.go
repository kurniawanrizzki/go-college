package mux

import (
	"go-college/internal/infra/metrics"
	"net/http"

	"github.com/rs/zerolog"
)


func InitHttpMux(log *zerolog.Logger, metricInst metrics.Metrics) *http.ServeMux {
	mux := http.NewServeMux()
	
	if metricInst != nil {
		mux.Handle("/metrics", metricInst.HTTPHandler())
		mux.Handle("/debug/vars", metricInst.HTTPHandler())
	}

	return mux
}


