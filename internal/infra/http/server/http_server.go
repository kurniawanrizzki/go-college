package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"go-college/internal/middleware"
	"go-college/internal/service"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type HttpServerOptions struct {
	AppName      string
	Mode         string
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
	MaxBodyBytes int64
}

var (
	onceServer     = sync.Once{}
	onceMetricsSrv = sync.Once{}
	httpServerInst *http.Server
	metricsServer  *http.Server
	handler        http.Handler
)

func InitHttpServer(logger *zerolog.Logger, opt *HttpServerOptions, mw middleware.Middleware, svc *service.Service, mux *http.ServeMux) *http.Server {
	onceServer.Do(func() {
		serverPort := fmt.Sprintf(":%d", opt.Port)

		handler = mux

		maxBodyBytes := opt.MaxBodyBytes
		if maxBodyBytes == 0 {
			maxBodyBytes = 1 << 20
		}

		if maxBodyBytes > 0 {
			bodyLimitedHandler := handler
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
				bodyLimitedHandler.ServeHTTP(w, r)
			})
		}

		handler = mw.CORS()(handler)
		handler = mw.Handler()(handler)
		if opt.AppName != "" {
			handler = otelhttp.NewHandler(handler, opt.AppName)
		}

		httpServerInst = &http.Server{
			Addr:         serverPort,
			WriteTimeout: opt.WriteTimeout,
			ReadTimeout:  opt.ReadTimeout,
			IdleTimeout:  opt.IdleTimeout,
			Handler:      handler,
		}
	})
	return httpServerInst
}

func InitMetricsServer(logger *zerolog.Logger, opt *HttpServerOptions, handler http.Handler) *http.Server {
	onceMetricsSrv.Do(func() {
		serverPort := fmt.Sprintf(":%d", opt.Port)

		maxBodyBytes := opt.MaxBodyBytes
		if maxBodyBytes == 0 {
			maxBodyBytes = 1 << 20
		}

		var wrappedHandler = handler
		if maxBodyBytes > 0 {
			bodyLimitedHandler := wrappedHandler
			wrappedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
				bodyLimitedHandler.ServeHTTP(w, r)
			})
		}

		metricsServer = &http.Server{
			Addr:         serverPort,
			WriteTimeout: opt.WriteTimeout,
			ReadTimeout:  opt.ReadTimeout,
			IdleTimeout:  opt.IdleTimeout,
			Handler:      wrappedHandler,
		}
	})

	return metricsServer
}
