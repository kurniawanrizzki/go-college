package grace

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type App interface {
	Serve()
}

type app struct {
	log        *zerolog.Logger
	httpServer *http.Server
	metricSrv  *http.Server
	timeout    time.Duration
}

type AppOptions struct {
	Name            string
	Version         string
	Environment     string
	ShutdownTimeout time.Duration
}

var onceGrace = &sync.Once{}

func InitGrace(log *zerolog.Logger, httpServer *http.Server, metricSrv *http.Server, timeout time.Duration) App {
	var gs *app

	onceGrace.Do(func() {
		gs = &app{
			log:        log,
			httpServer: httpServer,
			metricSrv:  metricSrv,
			timeout:    timeout,
		}
	})

	return gs
}

func (g *app) Serve() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	var wg sync.WaitGroup
	listenConfig := net.ListenConfig{
		KeepAlive: 30 * time.Second,
	}

	httpListener, err := listenConfig.Listen(context.Background(), "tcp", g.httpServer.Addr)

	if err != nil {
		g.log.Error().
			Err(err).
			Str("addr", g.httpServer.Addr).
			Msg("Failed to listen for HTTP server")
	}

	g.log.Info().
		Msg("HTTP server already started at " + g.httpServer.Addr)

	wg.Go(func() {
		if serveErr := g.httpServer.Serve(httpListener); serveErr != nil {
			g.log.Error().
				Err(serveErr).
				Msg("HTTP server error")
		}
	})

	var metricsListener net.Listener

	if g.metricSrv != nil {
		metricsListener, err = listenConfig.Listen(context.Background(), "tcp", g.metricSrv.Addr)

		if err != nil {
			g.log.Error().
				Err(err).
				Str("addr", g.metricSrv.Addr).
				Msg("Failed to listen for metrics HTTP server")

			return
		}

		g.log.Info().
			Msg("Metrics HTTP server already started at " + g.metricSrv.Addr)

		wg.Go(func() {
			if serveErr := g.metricSrv.Serve(metricsListener); serveErr != nil && serveErr != http.ErrServerClosed {
				g.log.Error().Err(serveErr).Msg("Metrics HTTP server error")
			}
		})
	}

	<-signalCh
	g.log.Debug().Msg("Received shutdown signal, gracefully shutting down...")

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), g.timeout)
	defer cancelShutdown()

	if err := g.httpServer.Shutdown(shutdownCtx); err != nil {
		g.log.Error().Err(err).Msg("HTTP server shutdown error")
	}

	if g.metricSrv != nil {
		if err := g.metricSrv.Shutdown(shutdownCtx); err != nil {
			g.log.Error().Err(err).Msg("Metrics HTTP shutdown error")
		}
	}

	wg.Wait()
	g.log.Debug().Msg("Shutdown complete")
}
