package mux

import (
	"net/http"

	"go-college/api/openapi"
	"go-college/internal/infra/metrics"

	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitHttpMux(log *zerolog.Logger, metricInst metrics.Metrics) *http.ServeMux {
	mux := http.NewServeMux()

	swaggerFS := http.Dir("./api/openapi")
	mux.Handle("GET /swagger/swagger.json", http.StripPrefix("/swagger/", http.FileServer(swaggerFS)))
	mux.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.json"),
	))

	if metricInst != nil {
		mux.Handle("/metrics", metricInst.HTTPHandler())
		mux.Handle("/debug/vars", metricInst.HTTPHandler())
	}

	_ = openapi.SwaggerInfo

	return mux
}
