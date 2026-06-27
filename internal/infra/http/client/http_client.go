package httpclient

import (
	"net"
	"net/http"
	"time"

	"go-college/internal/preference"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/circuitbreaker"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
	"github.com/failsafe-go/failsafe-go/timeout"
	"github.com/rs/zerolog"
)

type HttpClientOptions struct {
	CircuitBreaker        CircuitBreakerOptions
	KeepAlive             time.Duration
	MaxIdleConnsPerHost   int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	DialTimeout           time.Duration
	TLSHandshakeTimeout   time.Duration
	ResponseHeaderTimeout time.Duration
	ExpectContinueTimeout time.Duration
	Timeout               time.Duration
	MaxIdleConns          int
	Enabled               bool
	DisableCompression    bool
}

type CircuitBreakerOptions struct {
	MaxRetries int
	BackoffMin time.Duration
	BackofMax  time.Duration
}

func InitHttpClient(log *zerolog.Logger, opt *HttpClientOptions) *http.Client {
	if !opt.Enabled {
		log.Debug().Msg("HTTP client is disabled")
		return nil
	}

	baseTransport := &http.Transport{
		MaxIdleConns:          opt.MaxIdleConns,
		MaxIdleConnsPerHost:   opt.MaxIdleConnsPerHost,
		MaxConnsPerHost:       opt.MaxConnsPerHost,
		IdleConnTimeout:       opt.IdleConnTimeout,
		DisableCompression:    opt.DisableCompression,
		TLSHandshakeTimeout:   opt.TLSHandshakeTimeout,
		ResponseHeaderTimeout: opt.ResponseHeaderTimeout,
		ExpectContinueTimeout: opt.ExpectContinueTimeout,
		DialContext: (&net.Dialer{
			Timeout:   opt.DialTimeout,
			KeepAlive: opt.KeepAlive,
		}).DialContext,
	}

	transport := baseTransport
	timeoutPolicy := createTimeoutPolicy(log)
	circuitBreakerPolicy := createCircuitBreakerPolicy(log)
	retryPolicy := createRetryPolicy(log, &opt.CircuitBreaker)

	wrappedTransport := failsafehttp.NewRoundTripper(
		transport,
		circuitBreakerPolicy,
		retryPolicy,
		timeoutPolicy,
	)

	client := &http.Client{
		Transport: wrappedTransport,
		Timeout:   opt.Timeout,
	}

	log.Info().Msg("HTTP client initialized")

	return client
}

func getTraceInfo(e failsafe.ExecutionInfo) (traceId, spanId string) {
	ctx := e.Context()
	if ctx == nil {
		return "", ""
	}

	if v, ok := ctx.Value(preference.CONTEXT_KEY_LOG_TRACE_ID).(string); ok && v != "" {
		traceId = v
	}

	if v, ok := ctx.Value(preference.CONTEXT_KEY_LOG_SPAN_ID).(string); ok && v != "" {
		spanId = v
	}
	return traceId, spanId
}

func withTraceInfo(event *zerolog.Event, e failsafe.ExecutionInfo) *zerolog.Event {
	traceId, spanId := getTraceInfo(e)
	if traceId != "" {
		event = event.Str(string(preference.CONTEXT_KEY_LOG_TRACE_ID), traceId)
	}
	if spanId != "" {
		event = event.Str(string(preference.CONTEXT_KEY_LOG_SPAN_ID), spanId)
	}
	return event
}

func createTimeoutPolicy(log *zerolog.Logger) failsafe.Policy[*http.Response] {
	return timeout.NewBuilder[*http.Response](3 * time.Second).
		OnTimeoutExceeded(func(e failsafe.ExecutionDoneEvent[*http.Response]) {
			withTraceInfo(log.Info(), e).Msg("Request timed out")
		}).Build()
}

func createCircuitBreakerPolicy(log *zerolog.Logger) failsafe.Policy[*http.Response] {
	return circuitbreaker.NewBuilder[*http.Response]().
		HandleIf(func(response *http.Response, err error) bool {
			return response != nil && response.StatusCode == http.StatusServiceUnavailable
		}).
		WithDelayFunc(failsafehttp.DelayFunc).
		OnStateChanged(func(event circuitbreaker.StateChangedEvent) {
			log.Info().
				Str("old_state", event.OldState.String()).
				Str("new_state", event.NewState.String()).
				Msg("Circuit breaker state changed")
		}).Build()
}

func createRetryPolicy(log *zerolog.Logger, opt *CircuitBreakerOptions) failsafe.Policy[*http.Response] {
	return retrypolicy.NewBuilder[*http.Response]().
		WithMaxRetries(opt.MaxRetries).
		WithBackoff(opt.BackoffMin, opt.BackofMax).
		HandleIf(shouldRetry).
		OnRetry(func(e failsafe.ExecutionEvent[*http.Response]) {
			event := log.Warn()

			if err := e.LastError(); err != nil {
				event = event.Err(err)
			}

			event = withTraceInfo(event, e)
			event.Msg(getRetryReason(e.LastResult()))
		}).
		OnRetryScheduled(func(e failsafe.ExecutionScheduledEvent[*http.Response]) {
			withTraceInfo(
				log.Info().Int("attempt", e.Attempts()).Dur("delay", e.Delay),
				e,
			).Msg("Retry scheduled")
		}).
		Build()
}

func shouldRetry(response *http.Response, _ error) bool {
	return response == nil || response.StatusCode == http.StatusServiceUnavailable
}

func getRetryReason(response *http.Response) string {
	if response == nil {
		return "Retry attempt: response is nil"
	}

	return "Retry attempt: response status is 503 Service Unavailable"
}
