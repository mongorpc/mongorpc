package lib

import (
	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srv *MongoRPCServer) Listen(in *mongorpc.ListenRequest, stream mongorpc.MongoRPC_ListenServer) error {

	// Pipeline
	// TODO: check if pipeline decoder is working
	var pipeline interface{}
	if in.Match != nil {
		pipeline = decoder.Decode(in.Match)
	}

	// Options
	opts := options.ChangeStream()
	if in.Options != nil {
		if in.Options.FullDocument {
			opts.SetFullDocument(options.UpdateLookup)
		}
	}

	// MongoDB Change Streams
	changes, err := srv.DB.Database(in.Database).Collection(in.Collection).Watch(stream.Context(), pipeline, opts)
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

		// send change to client
		stream.Send(&mongorpc.ListenResponse{})
	}
	return nil
}
