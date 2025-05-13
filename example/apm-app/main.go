package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tqhuy-dev/gore/go_monitor/openmetry"
	"github.com/tqhuy-dev/gore/utilities"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"net/http"
	"time"
)

//docker run --rm -d --name jaeger \
//-e COLLECTOR_OTLP_ENABLED=true \
//-p 16686:16686 -p 4317:4317 \
//jaegertracing/all-in-one:latest

func initTracer() (*trace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("go-apm-service"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		panic(err)
	}
	orderCode := "2222"
	defer func() { _ = tp.Shutdown(context.Background()) }()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(openmetry.PrometheusMiddleware)
	e.GET("/read", func(c echo.Context) error {
		ctx := context.WithValue(context.Background(), "order_code", orderCode)
		tracer := otel.Tracer(utilities.ToString(ctx.Value("order_code")))

		ctx, span := tracer.Start(ctx, "read-api")
		defer span.End()
		readFunctionNested(ctx)
		time.Sleep(1 * time.Second)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"value": 1,
		})
	})
	e.GET("/write/:data", func(c echo.Context) error {
		ctx := context.WithValue(context.Background(), "order_code", orderCode)
		tracer := otel.Tracer(utilities.ToString(ctx.Value("order_code")))
		ctx, span := tracer.Start(ctx, "write-api")
		defer span.End()
		writeFunctionNested(ctx)
		time.Sleep(1 * time.Second)
		data := c.Param("data")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	})
	e.GET("/error", func(c echo.Context) error {
		data := c.Param("data")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	})
	e.Logger.Fatal(e.Start(":8081"))
}

func readFunctionNested(ctx context.Context) {
	tracer := otel.Tracer(utilities.ToString(ctx.Value("order_code")))
	_, span := tracer.Start(ctx, "readFunctionNested")
	time.Sleep(500 * time.Millisecond)
	defer span.End()
}

func writeFunctionNested(ctx context.Context) {
	tracer := otel.Tracer(utilities.ToString(ctx.Value("order_code")))
	_, span := tracer.Start(ctx, "writeFunctionNested")
	time.Sleep(500 * time.Millisecond)
	defer span.End()
}
