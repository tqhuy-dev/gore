package main

import (
	"context"
	"github.com/labstack/echo/v4/middleware"
	"github.com/s-platform/gore/go_monitor/openmetry"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

func init() {
	openmetry.RegisterMetric(openmetry.FunctionDuration, openmetry.ApiRequestDuration, openmetry.ApiRequestCount)
}

func main() {
	// Initialize metrics
	metricProvider := openmetry.MetricFactory(openmetry.OpenTelemetryClient, openmetry.MetricsConfig{
		Port: 2223,
		Path: "metrics",
	})
	err, shutdown := metricProvider.SetupMetrics()
	if err != nil {
		log.Fatalf("Failed to initialize metrics: %v", err)
	}
	defer shutdown()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(openmetry.PrometheusMiddleware)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, OpenTelemetry!")
	})

	e.GET("/fast-track", func(c echo.Context) error {
		return c.String(http.StatusOK, "Fast tracking")
	})
	type TempUpdate struct {
		Type string
	}

	counterIn64, _ := metricProvider.NewCounterInt64("order_metrics", "order metrics description")

	e.POST("/update", func(c echo.Context) error {
		var request TempUpdate
		if err := c.Bind(&request); err != nil {
			return err
		}
		result, err := openmetry.TrackExecutionTime("update_order", request.Type, func(input string) (string, error) {
			res := UpdateOrder(input)
			openmetry.LabelMetricInt64(counterIn64, "order_type", input)
			return res, nil
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, result)
	})

	e.GET("/hello", func(c echo.Context) error {
		meter := otel.Meter("example-meter")
		counter, _ := meter.Float64Counter("http_requests_total")

		// Simulate work
		counter.Add(context.Background(), 1)
		time.Sleep(100 * time.Millisecond)

		return c.String(http.StatusOK, "Hello, Metrics!")
	})

	// Start the Echo server
	e.Logger.Fatal(e.Start(":8081"))
}

func UpdateOrder(request string) string {
	return request
}
