// Package observability provides Prometheus metrics for MongoRPC.
package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// RequestsTotal counts total gRPC requests.
	RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mongorpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "status"},
	)

	// RequestDuration measures request latency.
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "mongorpc_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	// ActiveConnections tracks active connections.
	ActiveConnections = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "mongorpc_active_connections",
			Help: "Number of active connections",
		},
	)

	// MongoDBOperations counts MongoDB operations.
	MongoDBOperations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mongorpc_mongodb_operations_total",
			Help: "Total MongoDB operations",
		},
		[]string{"operation", "database", "collection"},
	)

	// MongoDBErrors counts MongoDB errors.
	MongoDBErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mongorpc_mongodb_errors_total",
			Help: "Total MongoDB errors",
		},
		[]string{"operation", "database", "collection"},
	)

	// DocumentsProcessed counts documents processed.
	DocumentsProcessed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mongorpc_documents_processed_total",
			Help: "Total documents processed",
		},
		[]string{"operation"},
	)
)

// RecordRequest records a request for metrics.
func RecordRequest(method, status string, durationSeconds float64) {
	RequestsTotal.WithLabelValues(method, status).Inc()
	RequestDuration.WithLabelValues(method).Observe(durationSeconds)
}

// RecordMongoOp records a MongoDB operation.
func RecordMongoOp(op, database, collection string, err error) {
	MongoDBOperations.WithLabelValues(op, database, collection).Inc()
	if err != nil {
		MongoDBErrors.WithLabelValues(op, database, collection).Inc()
	}
}
