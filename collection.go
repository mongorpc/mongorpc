package mongorpc

// func (server *MongoRPCServer) ListCollections(ctx context.Context, in *proto.ListCollectionsRequest) (*proto.ListCollectionsResponse, error) {
// 	logrus.Printf("ListCollections: %v", in)
// 	cur, err := server.DB.Database(in.Database).ListCollections(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var results bson.A
// 	for cur.Next(ctx) {
// 		var result bson.D
// 		err := cur.Decode(&result)
// 		if err != nil {
// 			return nil, err
// 		}
// 		results = append(results, result)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return nil, err
// 	}
// 	defer cur.Close(ctx)

// 	temporaryBytes, err := json.Marshal(results)
// 	if err != nil {
// 		logrus.Error(err)
// 		return nil, err
// 	}
// 	var jsonDocuments []interface{}
// 	err = json.Unmarshal(temporaryBytes, &jsonDocuments)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &proto.ListCollectionsResponse{
// 		Collections: &proto.ArrayValue{
// 			Values: ArrayToValue(jsonDocuments),
// 		},
// 	}, err
// }
