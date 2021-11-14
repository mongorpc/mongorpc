package mongorpc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase(ctx context.Context, mongoURI string) (*mongo.Client, error) {
	// connect to mongodb
	database, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// ping mongodb to check if it's up
	err = database.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return database, nil
}
