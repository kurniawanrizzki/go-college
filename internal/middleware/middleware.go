package middleware

import (
	"go-college/internal/infra/metrics"
	appErr "go-college/internal/model/errors"
	"go-college/internal/preference"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Middleware interface {
	Handler() func(http.Handler) http.Handler
	CORS() func(http.Handler) http.Handler
}

type middleware struct {
	log            *zerolog.Logger
	opt            *MiddlewareOptions
	publicPaths    map[string]bool
	limit          int
	period         time.Duration
	tracingEnabled bool
	metrics        metrics.Metrics
}

type MiddlewareOptions struct {
	PublicPaths []string           `yaml:"public_paths"`
	RateLimiter RateLimiterOptions `yaml:"rate_limiter"`
}

type RateLimiterOptions struct {
	Command string `yaml:"command"`
	Limit   int    `yaml:"limit"`
}

const TimeFormat = time.RFC3339

var (
	onceMiddleware = &sync.Once{}
	middlewareInst Middleware

	timeDict = map[string]time.Duration{
		"S": time.Second,
		"M": time.Minute,
		"H": time.Hour,
		"D": time.Hour * 24,
	}
)

func InitMiddleware(log *zerolog.Logger, opt *MiddlewareOptions, tracingEnabled bool, metricsInst metrics.Metrics) Middleware {
	onceMiddleware.Do(func() {
		limit := opt.RateLimiter.Limit
		period, err := parsePeriod(opt.RateLimiter.Command)

		if err != nil {
			log.Panic().Err(err).Send()
		}

		publicPathsMap := make(map[string]bool, len(opt.PublicPaths))

		for _, p := range opt.PublicPaths {
			publicPathsMap[p] = true
		}

		middlewareInst = &middleware{
			log:            log,
			opt:            opt,
			publicPaths:    publicPathsMap,
			limit:          limit,
			period:         period,
			tracingEnabled: tracingEnabled,
			metrics:        metricsInst,
		}
	})

	return middlewareInst
}

func parsePeriod(command string) (time.Duration, error) {
	parts := strings.Split(command, "-")

	if len(parts) != 2 {
		return 0, appErr.New(preference.FormatError, appErr.CodeHTTPBadRequest)
	}

	unit, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, appErr.New("invalid rate limit format", appErr.CodeHTTPBadRequest)
	}

	unitKey := strings.ToUpper(parts[1])

	if t, ok := timeDict[unitKey]; ok {
		return time.Duration(unit) * t, nil
	}

	return 0, appErr.New(preference.FormatError, appErr.CodeHTTPBadRequest)
}
