// Package observability provides OpenTelemetry tracing for MongoRPC.
package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/mongorpc/mongorpc"

// Tracer returns the global tracer for MongoRPC.
func Tracer() trace.Tracer {
	return otel.Tracer(tracerName)
}

// InitTracer initializes the OpenTelemetry tracer.
func InitTracer(serviceName, serviceVersion string) (func(context.Context) error, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}

// SpanAttributes returns common span attributes for gRPC methods.
func SpanAttributes(method, database, collection string) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("rpc.method", method),
		attribute.String("db.system", "mongodb"),
		attribute.String("db.name", database),
		attribute.String("db.mongodb.collection", collection),
	}
}

// StartSpan starts a new span with the given name.
func StartSpan(ctx context.Context, name string, attrs ...attribute.KeyValue) (context.Context, trace.Span) {
	return Tracer().Start(ctx, name, trace.WithAttributes(attrs...))
}
