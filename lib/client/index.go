package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
)

type Index struct {
	Name   string
	Unique bool
	Keys   []*IndexKey
}

type IndexKey struct {
	Field     string
	Direction IndexDirection
}

type IndexDirection int32

const (
	IndexDirection_ASCENDING  IndexDirection = 0
	IndexDirection_DESCENDING IndexDirection = 1
)

func (collection *Collection) Indexes(ctx context.Context) (interface{}, error) {
	resp, err := collection.client.admin.ListIndexes(ctx, &mongorpc.ListIndexesRequest{
		Database:   collection.parent.name,
		Collection: collection.name,
	})
	if err != nil {
		return nil, err
	}
	return decoder.Decode(resp), nil
}

func (collection *Collection) CreateIndex(ctx context.Context, index Index) error {
	keys := make([]*mongorpc.IndexKey, len(index.Keys))
	for i, key := range index.Keys {
		direction := mongorpc.IndexDirection_ASCENDING
		if key.Direction == IndexDirection_DESCENDING {
			direction = mongorpc.IndexDirection_DESCENDING
		}
		keys[i] = &mongorpc.IndexKey{
			Field:     key.Field,
			Direction: direction,
		}
	}
	i := &mongorpc.Index{
		Name:   index.Name,
		Unique: index.Unique,
		Keys:   keys,
	}
	_, err := collection.client.admin.CreateIndex(ctx, &mongorpc.CreateIndexRequest{
		Database:   collection.parent.name,
		Collection: collection.name,
		Index:      i,
	})
	if err != nil {
		return err
	}
	return nil
}

func (collection *Collection) DeleteIndex(ctx context.Context, index string) error {
	_, err := collection.client.admin.DropIndex(ctx, &mongorpc.DropIndexRequest{
		Database:   collection.parent.name,
		Collection: collection.name,
		Index:      index,
	})
	if err != nil {
		return err
	}
	return nil
}
