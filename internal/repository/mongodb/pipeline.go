// Package mongodb provides a fluent pipeline builder.
package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// PipelineBuilder provides a fluent interface for building aggregation pipelines.
type PipelineBuilder struct {
	stages bson.A
}

// NewPipelineBuilder creates a new pipeline builder.
func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{
		stages: bson.A{},
	}
}

// Match adds a $match stage.
func (p *PipelineBuilder) Match(filter bson.D) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$match", Value: filter}})
	return p
}

// MatchQuery adds a $match stage from a QueryBuilder.
func (p *PipelineBuilder) MatchQuery(q *QueryBuilder) *PipelineBuilder {
	return p.Match(q.Filter())
}

// Project adds a $project stage.
func (p *PipelineBuilder) Project(fields bson.D) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$project", Value: fields}})
	return p
}

// Group adds a $group stage.
func (p *PipelineBuilder) Group(id interface{}, accumulators bson.D) *PipelineBuilder {
	groupDoc := bson.D{{Key: "_id", Value: id}}
	groupDoc = append(groupDoc, accumulators...)
	p.stages = append(p.stages, bson.D{{Key: "$group", Value: groupDoc}})
	return p
}

// Sort adds a $sort stage.
func (p *PipelineBuilder) Sort(sort bson.D) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$sort", Value: sort}})
	return p
}

// SortAsc adds ascending sort on a field.
func (p *PipelineBuilder) SortAsc(field string) *PipelineBuilder {
	return p.Sort(bson.D{{Key: field, Value: 1}})
}

// SortDesc adds descending sort on a field.
func (p *PipelineBuilder) SortDesc(field string) *PipelineBuilder {
	return p.Sort(bson.D{{Key: field, Value: -1}})
}

// Limit adds a $limit stage.
func (p *PipelineBuilder) Limit(n int64) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$limit", Value: n}})
	return p
}

// Skip adds a $skip stage.
func (p *PipelineBuilder) Skip(n int64) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$skip", Value: n}})
	return p
}

// Unwind adds an $unwind stage.
func (p *PipelineBuilder) Unwind(path string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$unwind", Value: path}})
	return p
}

// UnwindPreserve adds an $unwind stage that preserves null/empty arrays.
func (p *PipelineBuilder) UnwindPreserve(path string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: path},
		{Key: "preserveNullAndEmptyArrays", Value: true},
	}}})
	return p
}

// Lookup adds a $lookup stage.
func (p *PipelineBuilder) Lookup(from, localField, foreignField, as string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: from},
		{Key: "localField", Value: localField},
		{Key: "foreignField", Value: foreignField},
		{Key: "as", Value: as},
	}}})
	return p
}

// LookupPipeline adds a $lookup stage with a pipeline.
func (p *PipelineBuilder) LookupPipeline(from, as string, let bson.D, pipeline bson.A) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: from},
		{Key: "let", Value: let},
		{Key: "pipeline", Value: pipeline},
		{Key: "as", Value: as},
	}}})
	return p
}

// AddFields adds an $addFields stage.
func (p *PipelineBuilder) AddFields(fields bson.D) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$addFields", Value: fields}})
	return p
}

// ReplaceRoot adds a $replaceRoot stage.
func (p *PipelineBuilder) ReplaceRoot(newRoot interface{}) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: newRoot}}}})
	return p
}

// Count adds a $count stage.
func (p *PipelineBuilder) Count(field string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$count", Value: field}})
	return p
}

// Sample adds a $sample stage.
func (p *PipelineBuilder) Sample(size int64) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: size}}}})
	return p
}

// Out adds an $out stage.
func (p *PipelineBuilder) Out(collection string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$out", Value: collection}})
	return p
}

// Merge adds a $merge stage.
func (p *PipelineBuilder) Merge(into string, on []string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$merge", Value: bson.D{
		{Key: "into", Value: into},
		{Key: "on", Value: on},
	}}})
	return p
}

// Facet adds a $facet stage.
func (p *PipelineBuilder) Facet(facets map[string]bson.A) *PipelineBuilder {
	facetDoc := bson.D{}
	for name, pipeline := range facets {
		facetDoc = append(facetDoc, bson.E{Key: name, Value: pipeline})
	}
	p.stages = append(p.stages, bson.D{{Key: "$facet", Value: facetDoc}})
	return p
}

// Build returns the built pipeline.
func (p *PipelineBuilder) Build() bson.A {
	return p.stages
}

// Append appends raw stages.
func (p *PipelineBuilder) Append(stages ...bson.D) *PipelineBuilder {
	for _, s := range stages {
		p.stages = append(p.stages, s)
	}
	return p
}
