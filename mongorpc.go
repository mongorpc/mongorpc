package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}
