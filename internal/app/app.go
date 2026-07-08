package app

import (
	"flag"
	"log"
	"net/http"
	"time"

	"go-college/internal/config"
	httpHandler "go-college/internal/handler"
	"go-college/internal/infra/database"
	"go-college/internal/infra/grace"
	httpclient "go-college/internal/infra/http/client"
	httpmux "go-college/internal/infra/http/mux"
	httpserver "go-college/internal/infra/http/server"
	"go-college/internal/infra/logger"
	"go-college/internal/infra/metrics"
	"go-college/internal/infra/query"
	"go-college/internal/infra/tracer"
	"go-college/internal/middleware"
	"go-college/internal/repository"
	"go-college/internal/service"
	"go-college/internal/util"
)

const (
	DefaultMinJitter = 100
	DefaultMaxJitter = 2000
)

func Run() {
	minJitter, maxJitter := parseFlags()
	sleepWithJitter(minJitter, maxJitter)

	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	log := logger.InitLogger(conf.Logger)

	var tracerInst tracer.Tracer
	if conf.Tracer.Enabled {
		tracerInst = tracer.InitTracer(log, conf.Tracer)
		if tracerInst != nil {
			defer tracerInst.Stop()
		}
	}

	sql0 := database.InitDB(log, conf.Database.Postgres)
	defer sql0.Close()

	httpClient := httpclient.InitHttpClient(log, conf.HTTP.Client)
	queryLoader := query.InitQueryLoader(log, conf.Queries)

	repo := repository.InitRepository(log, sql0, queryLoader)
	svc := service.InitService(log, repo, httpClient)

	var metricsInst metrics.Metrics
	if conf.Metric.Enabled {
		metricsInst = metrics.InitMetrics(log, sql0)
	}

	mw := middleware.InitMiddleware(log, conf.Middleware, conf.Tracer.Enabled, metricsInst)
	httpMux := httpmux.InitHttpMux(log, metricsInst)
	httpHandler.InitHttpHandler(httpMux, mw, svc, sql0)
	httpServer := httpserver.InitHttpServer(log, conf.HTTP.Server, mw, svc, httpMux)

	var metricsServer *http.Server
	if conf.Metric.Enabled && conf.HTTP.MetricsServer != nil {
		metricsServer = httpserver.InitMetricsServer(log, conf.HTTP.MetricsServer, metricsInst.HTTPHandler())
	}

	app := grace.InitGrace(log, httpServer, metricsServer, conf.App.ShutdownTimeout)
	app.Serve()
}

func parseFlags() (minJitter, maxJitter int) {
	flag.IntVar(&minJitter, "minSleep", DefaultMinJitter, "min. sleep duration during app initialization")
	flag.IntVar(&maxJitter, "maxSleep", DefaultMaxJitter, "max. sleep duration during app initialization")
	flag.Parse()

	return
}

func sleepWithJitter(low, high int) {
	if low < 1 {
		low = DefaultMinJitter
	}

	if high < 1 || high < low {
		high = DefaultMaxJitter
	}

	rnd := util.RandomInt(high-low) + low
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	log.Printf("%d ms sleep during initialization", rnd)
}
