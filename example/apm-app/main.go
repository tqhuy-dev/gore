package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tqhuy-dev/gore/go_monitor/openmetry"
	"github.com/tqhuy-dev/gore/utilities"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	trace2 "go.opentelemetry.io/otel/trace"
	"net/http"
	"sync"
	"time"
)

type SpanType int8

const (
	Function      SpanType = 1
	Database      SpanType = 2
	CloudProvider SpanType = 3
	Queue         SpanType = 4
	API           SpanType = 5
	Grpc          SpanType = 6
	SSE           SpanType = 7
	Socket        SpanType = 8
)

type StatusSpan int8

const (
	Success StatusSpan = 1
	Failure StatusSpan = 2
	Warning StatusSpan = 3
)

//docker run --rm -d --name jaeger \
//-e COLLECTOR_OTLP_ENABLED=true \
//-p 16686:16686 -p 4317:4317 \
//my-jaeger:custom

const TypeOfSpan = attribute.Key("span_type")
const StatusOfSpan = attribute.Key("span_status")
const SourceTrackId = attribute.Key("source_track_id")
const CallerTrackId = attribute.Key("caller_track_id")

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
			semconv.ServiceNamespaceKey.String("go-ns"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func getTraceID(ctx context.Context) string {
	spanCtx := trace2.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return ""
	}
	return spanCtx.TraceID().String()
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
		//spanCtx := trace2.NewSpanContext(trace2.SpanContextConfig{
		//	TraceID:    trace2.TraceID{},
		//	TraceFlags: trace2.FlagsSampled,
		//	Remote:     true,
		//})
		tracer := otel.Tracer(utilities.ToString(ctx.Value("order_code")))

		ctx, span := tracer.Start(ctx, "read-api")
		defer span.End()
		readFunctionNested(ctx)
		time.Sleep(100 * time.Millisecond)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"value": 1,
		})
	})

	e.GET("/group", func(c echo.Context) error {
		ctx := context.WithValue(context.Background(), "trace_flow", "group")
		tracer := otel.Tracer("group")
		ctx, span := tracer.Start(ctx, "trace-group")
		span.SetAttributes(SourceTrackId.String("1.0.1"))
		ctx = context.WithValue(ctx, "track_id", "1.0.1")
		defer span.End()
		RunTraceGroup(ctx)
		time.Sleep(100 * time.Millisecond)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"value": 1,
		})
	})

	e.Logger.Fatal(e.Start(":8081"))
}

func readFunctionNested(ctx context.Context) {
	tracer := otel.Tracer(utilities.ToString(ctx.Value("order_code")))
	_, span := tracer.Start(ctx, "readFunctionNested")
	time.Sleep(50 * time.Millisecond)
	defer span.End()
	span.SetAttributes(StatusOfSpan.Int64(int64(Success)))
	span.SetAttributes(TypeOfSpan.Int64(int64(Function)))
}

func RunTraceGroup(ctx context.Context) {
	prevTrackId := ctx.Value("track_id").(string)
	tracer := otel.Tracer(utilities.ToString(ctx.Value("trace_flow")))
	ctxSpan, span := tracer.Start(ctx, "RunTraceGroup")
	ctxSpan = context.WithValue(ctxSpan, "track_id", "1.0.2")
	span.SetAttributes(SourceTrackId.String("1.0.2"))
	span.SetAttributes(CallerTrackId.String(prevTrackId))
	time.Sleep(50 * time.Millisecond)
	defer span.End()

	RunTraceGroupSerial1(ctxSpan)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		RunTraceGroupWg1(ctxSpan)
	}()
	go func() {
		defer wg.Done()
		RunTraceGroupWg2(ctxSpan)
	}()

	wg.Wait()
}

func RunTraceGroupSerial1(ctx context.Context) {
	prevTrackId := ctx.Value("track_id").(string)
	tracer := otel.Tracer(utilities.ToString(ctx.Value("trace_flow")))
	ctx, span := tracer.Start(ctx, "RunTraceGroupSerial1")
	ctx = context.WithValue(ctx, "track_id", "1.0.3")
	span.SetAttributes(SourceTrackId.String("1.0.3"))
	span.SetAttributes(CallerTrackId.String(prevTrackId))
	time.Sleep(50 * time.Millisecond)
	defer span.End()
}

func RunTraceGroupWg1(ctx context.Context) {
	prevTrackId := ctx.Value("track_id").(string)
	tracer := otel.Tracer(utilities.ToString(ctx.Value("trace_flow")))
	_, span := tracer.Start(ctx, "RunTraceGroupWg1")
	span.SetAttributes(SourceTrackId.String("1.0.4"))
	span.SetAttributes(CallerTrackId.String(prevTrackId))
	time.Sleep(50 * time.Millisecond)
	defer span.End()
}

func RunTraceGroupWg2(ctx context.Context) {
	prevTrackId := ctx.Value("track_id").(string)
	tracer := otel.Tracer(utilities.ToString(ctx.Value("trace_flow")))
	_, span := tracer.Start(ctx, "RunTraceGroupWg2")
	span.SetAttributes(SourceTrackId.String("1.0.5"))
	span.SetAttributes(CallerTrackId.String(prevTrackId))
	time.Sleep(50 * time.Millisecond)
	defer span.End()
}
