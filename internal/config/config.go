package config

import (
	"os"

	"go-college/internal/infra/database"
	app "go-college/internal/infra/grace"
	httpclient "go-college/internal/infra/http/client"
	httpserver "go-college/internal/infra/http/server"
	"go-college/internal/infra/logger"
	"go-college/internal/infra/metrics"
	"go-college/internal/infra/query"
	"go-college/internal/infra/tracer"
	"go-college/internal/middleware"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App        *app.AppOptions
	HTTP       HTTPConfig
	Database   DatabseConfig
	Logger     *logger.LoggerOptions
	Middleware *middleware.MiddlewareOptions
	Queries    *query.QueriesOptions
	Tracer     *tracer.TracerOptions
	Metric     *metrics.MetricsOptions
}

type HTTPConfig struct {
	Server        *httpserver.HttpServerOptions `yaml:"server"`
	MetricsServer *httpserver.HttpServerOptions `yaml:"metrics_server"`
	Client        *httpclient.HttpClientOptions `yaml:"client"`
}

type DatabseConfig struct {
	Postgres *database.DatabaseOptions
}

func InitConfig() (*Config, error) {
	data, err := os.ReadFile("./configs/config.yaml")
	if err != nil {
		return nil, err
	}

	expanded := expandEnv(string(data))

	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func expandEnv(s string) string {
	return os.Expand(s, func(name string) string {
		if val := os.Getenv(name); val != "" {
			return val
		}

		return ""
	})
}
