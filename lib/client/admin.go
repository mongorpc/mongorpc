package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
)

func (client *MongoRPCClient) GetDatabases(ctx context.Context) ([]string, error) {
	result, err := client.admin.ListDatabases(ctx, &mongorpc.Empty{})
	if err != nil {
		return nil, err
	}
	return decoder.Decode(result).([]string), nil
}

func (database *Database) Delete(ctx context.Context) error {
	_, err := database.client.admin.DropDatabase(ctx, &mongorpc.DropDatabaseRequest{
		Database: database.name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) CreateCollection(ctx context.Context, name string) (*Collection, error) {
	_, err := database.client.admin.CreateCollection(ctx, &mongorpc.CreateCollectionRequest{
		Database:   database.name,
		Collection: name,
	})
	if err != nil {
		return nil, err
	}
	return &Collection{
		parent: database,
		name:   name,
		client: database.client,
	}, nil
}

func (collection *Collection) Delete(ctx context.Context) error {
	_, err := collection.client.admin.DropCollection(ctx, &mongorpc.DropCollectionRequest{
		Database:   collection.parent.name,
		Collection: collection.name,
	})
	if err != nil {
		return err
	}
	return nil
}
