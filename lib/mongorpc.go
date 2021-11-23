package lib

import (
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRPCServer struct {
	mongorpc.UnimplementedMongoRPCServer
	DB *mongo.Client
}
