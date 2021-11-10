package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB change events rpc stream handler
func (srv *MongoRPC) Listen(in *proto.ListenRequest, stream proto.MongoRPC_ListenServer) error {

	// TODO: pass operation type and filter in mongo pipeline
	// pipeline := bson.D{
	// 	{
	// 		"$match", bson.D{
	// 			{"operationType", "insert"},
	// 			{"fullDocument.duration", bson.D{
	// 				{"$gt", 30},
	// 			}},
	// 		},
	// 	},
	// }

	// MongoDB Change Streams
	changes, err := srv.DB.Database(in.Database).Collection(in.Collection).Watch(stream.Context(), mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return err
	}
	defer changes.Close(stream.Context())

	// Iterate over changes
	for changes.Next(stream.Context()) {

		// Get change document
		var change map[string]interface{}
		err := changes.Decode(&change)
		if err != nil {
			return err
		}

		// operation type of change event
		operationType := change["operationType"].(string)

		// Get full document
		fullDocument := change["fullDocument"]

		// send change to client
		stream.Send(&proto.ListenResponse{
			Operation: operationType,
			Document:  Encode(fullDocument),
		})

	}
	return nil
}
