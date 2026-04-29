---
title: OpenTelemetry
icon: fa-satellite-dish
primary: "#F5A623"
lang: go
---

## fa-layer-group Concepts (Signals)

```go
Signals:
  Traces  -> distributed request flow
  Metrics -> numeric measurements over time
  Logs    -> timestamped event records

Every signal shares:
  Resource + Attributes (who/where)
  Context + Propagation (cross-service)
```

## fa-cogs Tracer Provider

```go
import "go.opentelemetry.io/otel"
import "go.opentelemetry.io/otel/sdk/trace"

func initTracer() (*trace.TracerProvider, error) {
    exp, err := otlptrace.New(context.Background(),
        otlptracegrpc.NewClient(
            otlptracegrpc.WithEndpoint("localhost:4317"),
            otlptracegrpc.WithInsecure(),
        ),
    )
    if err != nil {
        return nil, err
    }
    tp := trace.NewTracerProvider(trace.WithBatcher(exp))
    otel.SetTracerProvider(tp)
    return tp, nil
}
```

## fa-project-diagram Spans & Context

```go
tracer := otel.Tracer("svc")
ctx, span := tracer.Start(context.Background(), "handleRequest",
    trace.WithSpanKind(trace.SpanKindServer),
)
defer span.End()

childCtx, child := tracer.Start(ctx, "dbQuery")
child.SetStatus(codes.Error, "conn refused")
child.End()
```

## fa-tags Span Attributes & Events

```go
span.SetAttributes(
    attribute.String("http.method", "GET"),
    attribute.Int("http.status_code", 200),
    attribute.String("db.system", "postgresql"),
)

span.AddEvent("cache_miss", trace.WithAttributes(
    attribute.String("key", "user:123"),
))
span.RecordError(err, trace.WithAttributes(
    attribute.String("retry", "true"),
))
```

## fa-exchange-alt Propagators

```go
import (
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel"
)

otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
    propagation.TraceContext{},
    propagation.Baggage{},
))

carrier := propagation.HeaderCarrier(req.Header)
ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
```

## fa-tachometer-alt Meter Provider

```go
import "go.opentelemetry.io/otel/sdk/metric"

func initMeter() (*metric.MeterProvider, error) {
    exp, err := otlpmetric.New(context.Background(),
        otlpmetricgrpc.NewClient(
            otlpmetricgrpc.WithEndpoint("localhost:4317"),
            otlpmetricgrpc.WithInsecure(),
        ),
    )
    if err != nil {
        return nil, err
    }
    return metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(exp))), nil
}
```

## fa-chart-bar Counters & Histograms

```go
meter := otel.Meter("svc")

counter, _ := meter.Int64Counter("http.requests",
    metric.WithDescription("Total HTTP requests"),
)
counter.Add(ctx, 1, metric.WithAttributes(
    attribute.String("method", "GET"),
))

hist, _ := meter.Float64Histogram("http.duration",
    metric.WithUnit("ms"),
)
hist.Record(ctx, 42.5, metric.WithAttributes(
    attribute.String("route", "/api/users"),
))
```

## fa-scroll Logger Provider

```go
import "go.opentelemetry.io/otel/log/global"

lp := log.NewLoggerProvider(log.WithProcessor(
    log.NewBatchProcessor(otlplog.NewClient(
        otlploggrpc.WithEndpoint("localhost:4317"),
        otlploggrpc.WithInsecure(),
    )),
))
global.SetLoggerProvider(lp)

logger := global.Logger("svc")
var severity log.Severity = log.SeverityInfo
logger.Emit(ctx, log.Record{
    Body:       log.StringValue("request completed"),
    Severity:   severity,
    Attributes: []log.KeyValue{log.String("status", "ok")},
})
```

## fa-server Resource & Attributes

```go
import "go.opentelemetry.io/otel/sdk/resource"
import semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

res, _ := resource.Merge(
    resource.Default(),
    resource.NewWithAttributes(
        semconv.SchemaURL,
        semconv.ServiceName("api-server"),
        semconv.ServiceVersion("1.2.0"),
        attribute.String("env", "prod"),
    ),
)
tp := trace.NewTracerProvider(
    trace.WithResource(res),
    trace.WithBatcher(exp),
)
```

## fa-upload OTLP Exporter

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318

exporters:
  otlp:
    endpoint: jaeger:4317
    tls:
      insecure: true
  prometheus:
    endpoint: "0.0.0.0:8889"

service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [otlp]
    metrics:
      receivers: [otlp]
      exporters: [prometheus]
```

## fa-random Sampling

```go
import "go.opentelemetry.io/otel/sdk/trace"

type ratioSampler struct{ ratio float64 }

func (s ratioSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
    if rand.Float64() < s.ratio {
        return trace.SamplingResult{
            Decision: trace.RecordAndSample,
        }
    }
    return trace.SamplingResult{Decision: trace.Drop}
}

func (s ratioSampler) Description() string { return "RatioSampler" }

tp := trace.NewTracerProvider(
    trace.WithSampler(trace.TraceIDRatioBased(0.1)),
)
```

## fa-suitcase Baggage

```go
import "go.opentelemetry.io/otel/baggage"

mem, _ := baggage.NewMember("tenant.id", "acme")
b, _ := baggage.New(mem)
ctx = baggage.ContextWithBaggage(ctx, b)

b = baggage.FromContext(ctx)
tenant := b.Member("tenant.id").Value()
```

## fa-rocket Go SDK Setup

```go
func setupOTel() (shutdown func(), err error) {
    res, _ := resource.Merge(resource.Default(),
        resource.NewWithAttributes(semconv.SchemaURL,
            semconv.ServiceName("my-service"),
        ),
    )
    tp := trace.NewTracerProvider(
        trace.WithResource(res),
        trace.WithBatcher(otlptrace.NewClient()),
    )
    mp := metric.NewMeterProvider(
        metric.WithResource(res),
        metric.WithReader(metric.NewPeriodicReader(
            otlpmetric.NewClient(),
        )),
    )
    otel.SetTracerProvider(tp)
    otel.SetMeterProvider(mp)
    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{}, propagation.Baggage{},
    ))
    return func() { tp.Shutdown(context.Background()); mp.Shutdown(context.Background()) }, nil
}
```

## fa-network-wired Collector Config

```yaml
receivers:
  otlp:
    protocols:
      grpc: { endpoint: "0.0.0.0:4317" }
      http: { endpoint: "0.0.0.0:4318" }

processors:
  batch:
    timeout: 5s
    send_batch_size: 1024
  memory_limiter:
    check_interval: 1s
    limit_mib: 512

exporters:
  otlp/jaeger:
    endpoint: jaeger:4317
    tls: { insecure: true }
  prometheus:
    endpoint: "0.0.0.0:8889"
  elasticsearch:
    endpoints: ["http://es:9200"]
    logs_index: otel-logs

service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [otlp/jaeger]
    metrics:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [prometheus]
    logs:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [elasticsearch]
```
