package openmetry

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	prometheusGolang "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	metric2 "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"log"
	"net/http"
	"time"
)

var (
	ApiRequestDuration = prometheusGolang.NewHistogramVec(
		prometheusGolang.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Duration of API requests in seconds",
			Buckets: prometheusGolang.DefBuckets,
		},
		[]string{"method", "endpoint", "status"},
	)

	ApiRequestCount = prometheusGolang.NewCounterVec(
		prometheusGolang.CounterOpts{
			Name: "api_request_count",
			Help: "Total number of API requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	FunctionDuration = prometheusGolang.NewHistogramVec(
		prometheusGolang.HistogramOpts{
			Name:    "function_execution_time_seconds",
			Help:    "Duration of function execution in seconds",
			Buckets: prometheusGolang.DefBuckets, // Default Prometheus buckets
		},
		[]string{"function"},
	)
)

func TrackExecutionTime[Input any, Output any](funcName string, request Input, fn func(input Input) (Output, error)) (Output, error) {
	start := time.Now()
	result, err := fn(request) // Execute the function
	duration := time.Since(start).Seconds()

	// Record execution time in Prometheus
	FunctionDuration.WithLabelValues(funcName).Observe(duration)
	return result, err
}

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c) // Call the next handler
		duration := time.Since(start).Seconds()

		// Get request info
		method := c.Request().Method
		endpoint := c.Path() // Use path template like "/users/:id"
		status := c.Response().Status

		// Record metrics
		ApiRequestDuration.WithLabelValues(method, endpoint, http.StatusText(status)).Observe(duration)
		ApiRequestCount.WithLabelValues(method, endpoint, http.StatusText(status)).Inc()

		return err
	}
}

func RegisterMetric(cs ...prometheusGolang.Collector) {
	prometheusGolang.MustRegister(cs...)
}

type openTelemetryProvider struct {
	meterProvider *metric.MeterProvider
	meter         metric2.Meter
	config        MetricsConfig
}

func (o *openTelemetryProvider) NewCounterInt64(name string, description string) (metric2.Int64Counter, error) {
	counter, err := o.meter.Int64Counter(
		name,
		metric2.WithDescription(description),
	)
	if err != nil {
		return nil, err
	}
	return counter, nil
}

func (o *openTelemetryProvider) SetupMetrics() (error, func()) {

	otel.SetMeterProvider(o.meterProvider)
	http.Handle(fmt.Sprintf("/%s", o.config.Path), promhttp.Handler())
	go func() {
		log.Println(fmt.Sprintf("Serving metrics at :%d/%s", o.config.Port, o.config.Path))
		_ = http.ListenAndServe(fmt.Sprintf(":%d", o.config.Port), nil)
	}()

	return nil, func() {
		_ = o.meterProvider.Shutdown(context.Background())
	}

}

type MetricsConfig struct {
	Port int
	Path string
}
type IMetricProvider interface {
	SetupMetrics() (error, func())
	NewCounterInt64(name string, description string) (metric2.Int64Counter, error)
}

func NewOpenTelemetryProvider(config MetricsConfig) IMetricProvider {
	if config.Path == "" {
		config.Path = "metrics"
	}
	if config.Port == 0 {
		config.Port = 2223
	}
	exporter, err := prometheus.New()
	if err != nil {
		return nil
	}
	meterProvider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := meterProvider.Meter("echo-metric")

	otel.SetMeterProvider(meterProvider)
	return &openTelemetryProvider{
		meterProvider: meterProvider,
		meter:         meter,
		config:        config,
	}
}

type MetricClient int8

const OpenTelemetryClient MetricClient = 1

func MetricFactory(client MetricClient, config MetricsConfig) IMetricProvider {
	switch client {
	case OpenTelemetryClient:
		return NewOpenTelemetryProvider(config)
	default:
		return NewOpenTelemetryProvider(config)
	}
}

func LabelMetricInt64(counter metric2.Int64Counter, label string, value string) {
	labels := metric2.WithAttributes(
		attribute.KeyValue{
			Key:   attribute.Key(label),
			Value: attribute.StringValue(value),
		},
	)
	counter.Add(context.Background(), 1, labels)
}
