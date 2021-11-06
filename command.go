package mongorpc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// RunCommand executes the given command against the database.
func (srv *MongoRPC) RunDatabaseCommand(ctx context.Context, database string, command interface{}) (*mongo.SingleResult, error) {
	result := srv.DB.Database(database).RunCommand(ctx, command)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}
