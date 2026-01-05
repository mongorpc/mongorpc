// Package mongodb provides a fluent query builder.
package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// QueryBuilder provides a fluent interface for building queries.
type QueryBuilder struct {
	filter     bson.D
	projection bson.D
	sort       bson.D
	limit      int64
	skip       int64
}

// NewQueryBuilder creates a new query builder.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		filter: bson.D{},
	}
}

// Where adds a filter condition.
func (q *QueryBuilder) Where(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: value})
	return q
}

// Eq adds an equality condition.
func (q *QueryBuilder) Eq(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: value})
	return q
}

// Ne adds a not-equal condition.
func (q *QueryBuilder) Ne(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$ne", Value: value}}})
	return q
}

// Gt adds a greater-than condition.
func (q *QueryBuilder) Gt(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$gt", Value: value}}})
	return q
}

// Gte adds a greater-than-or-equal condition.
func (q *QueryBuilder) Gte(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$gte", Value: value}}})
	return q
}

// Lt adds a less-than condition.
func (q *QueryBuilder) Lt(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$lt", Value: value}}})
	return q
}

// Lte adds a less-than-or-equal condition.
func (q *QueryBuilder) Lte(field string, value interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$lte", Value: value}}})
	return q
}

// In adds an in-array condition.
func (q *QueryBuilder) In(field string, values ...interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$in", Value: values}}})
	return q
}

// NotIn adds a not-in-array condition.
func (q *QueryBuilder) NotIn(field string, values ...interface{}) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$nin", Value: values}}})
	return q
}

// Regex adds a regex condition.
func (q *QueryBuilder) Regex(field, pattern string, options string) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{
		{Key: "$regex", Value: pattern},
		{Key: "$options", Value: options},
	}})
	return q
}

// Exists adds an exists condition.
func (q *QueryBuilder) Exists(field string, exists bool) *QueryBuilder {
	q.filter = append(q.filter, bson.E{Key: field, Value: bson.D{{Key: "$exists", Value: exists}}})
	return q
}

// Select adds projection fields.
func (q *QueryBuilder) Select(fields ...string) *QueryBuilder {
	for _, f := range fields {
		q.projection = append(q.projection, bson.E{Key: f, Value: 1})
	}
	return q
}

// Exclude excludes projection fields.
func (q *QueryBuilder) Exclude(fields ...string) *QueryBuilder {
	for _, f := range fields {
		q.projection = append(q.projection, bson.E{Key: f, Value: 0})
	}
	return q
}

// OrderBy adds ascending sort.
func (q *QueryBuilder) OrderBy(field string) *QueryBuilder {
	q.sort = append(q.sort, bson.E{Key: field, Value: 1})
	return q
}

// OrderByDesc adds descending sort.
func (q *QueryBuilder) OrderByDesc(field string) *QueryBuilder {
	q.sort = append(q.sort, bson.E{Key: field, Value: -1})
	return q
}

// Limit sets the limit.
func (q *QueryBuilder) Limit(n int64) *QueryBuilder {
	q.limit = n
	return q
}

// Skip sets the skip.
func (q *QueryBuilder) Skip(n int64) *QueryBuilder {
	q.skip = n
	return q
}

// And combines with AND.
func (q *QueryBuilder) And(builders ...*QueryBuilder) *QueryBuilder {
	conditions := bson.A{q.filter}
	for _, b := range builders {
		conditions = append(conditions, b.filter)
	}
	q.filter = bson.D{{Key: "$and", Value: conditions}}
	return q
}

// Or combines with OR.
func (q *QueryBuilder) Or(builders ...*QueryBuilder) *QueryBuilder {
	conditions := bson.A{q.filter}
	for _, b := range builders {
		conditions = append(conditions, b.filter)
	}
	q.filter = bson.D{{Key: "$or", Value: conditions}}
	return q
}

// Build returns the built query components.
func (q *QueryBuilder) Build() (filter, projection, sort bson.D, limit, skip int64) {
	return q.filter, q.projection, q.sort, q.limit, q.skip
}

// Filter returns just the filter.
func (q *QueryBuilder) Filter() bson.D {
	return q.filter
}
