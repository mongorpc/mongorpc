package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// gRPC server address
	address = "localhost:27051"
)

// Example MongoRPC Client
type ExampleClient struct {
	ctx      context.Context
	mongorpc proto.MongoRPCClient
}

// ClientInterceptor is a gRPC interceptor that adds the access token to the request
type ClientInterceptor struct {
	accessToken string
}

// UnaryClientInterceptor is a gRPC interceptor that adds the access token to the request
func (interceptor *ClientInterceptor) UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if interceptor.accessToken != "" {
		md.Set("authorization", interceptor.accessToken)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// StreamClientInterceptor is a gRPC interceptor that adds the access token to the request
func (interceptor ClientInterceptor) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if interceptor.accessToken != "" {
		md.Set("authorization", interceptor.accessToken)
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	return streamer(ctx, desc, cc, method, opts...)
}

func main() {
	interceptor := &ClientInterceptor{
		accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJ1aWQiOiI1YjEwODcxZS0zYmMwLTExZWMtOGQzZC0wMjQyYWMxMzAwMDMifQ.PlVY78uCmkzzKVG3tOvQu9aeXnMOImUzG6b_lygHH2U",
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(interceptor.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(interceptor.StreamClientInterceptor),
	)
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	c := proto.NewMongoRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Create a new client
	e := ExampleClient{
		ctx:      ctx,
		mongorpc: c,
	}

	// e.ListCollections()
	// e.ListDocuments()
	e.DocumentByID()
	// e.CreateDocument()
	// e.CreateIndex()
	// e.DeleteIndex()
	// e.ListIndexes()

	e.CollectionStats()

	// var waitGroup sync.WaitGroup

	// // listen document changes
	// stream, err := e.mongorpc.Listen(ctx, &proto.ListenRequest{
	// 	Database:   "sample_mflix",
	// 	Collection: "sample_movies",
	// })

	// if err != nil {
	// 	logrus.Fatalf("could not listen: %v", err)
	// }

	// waitGroup.Add(1)
	// go func() {
	// 	for {
	// 		resp, err := stream.Recv()
	// 		if err != nil {
	// 			logrus.Fatalf("could not listen: %v", err)
	// 		}
	// 		logrus.Printf("Changes: %s", resp.Document)
	// 	}
	// }()

	// waitGroup.Wait()
}

// collection stats
func (c *ExampleClient) CollectionStats() {

	res, err := c.mongorpc.CollectionStats(c.ctx, &proto.CollectionStatsRequest{
		Database:   "sample_mflix",
		Collection: "movies",
	})
	if err != nil {
		logrus.Fatalf("could not get collection stats: %v", err)
	}
	logrus.Printf("Collection Stats: %s", res)
}

// delete index by name in collection
func (c *ExampleClient) DeleteIndex() {
	res, err := c.mongorpc.DeleteIndex(c.ctx, &proto.DeleteIndexRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		Name:       "title_index",
	})
	if err != nil {
		logrus.Fatalf("could not delete index: %v", err)
	}
	logrus.Printf("Index: %s", res)
}

// list all indexes for a collection
func (c *ExampleClient) ListIndexes() {
	res, err := c.mongorpc.ListIndexes(c.ctx, &proto.ListIndexesRequest{
		Database:   "sample_mflix",
		Collection: "movies",
	})

	if err != nil {
		logrus.Fatalf("could not get indexes: %v", err)
	}

	logrus.Printf("Indexes: %s", res.Indexes)
}

// create index
func (c *ExampleClient) CreateIndex() {
	keys := []*proto.IndexKey{}
	keys = append(keys, &proto.IndexKey{
		Field: "title",
	})

	result, err := c.mongorpc.CreateIndex(c.ctx, &proto.CreateIndexRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		Index: &proto.Index{
			Name: "title_index",
			Keys: keys,
		},
	})
	if err != nil {
		logrus.Fatalf("could not create index: %v", err)
	}
	logrus.Printf("Index: %s", result)
}

// list all collections
func (c *ExampleClient) ListCollections() {
	r, err := c.mongorpc.ListCollections(c.ctx, &proto.ListCollectionsRequest{
		Database: "sample_mflix",
	})
	if err != nil {
		logrus.Fatalf("could not get collection: %v", err)
	}
	logrus.Printf("Collection: %s", r.Collections)
}

// get document
func (c *ExampleClient) DocumentByID() {
	doc, err := c.mongorpc.GetDocument(c.ctx, &proto.GetDocumentRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		DocumentId: "573a1390f29313caabcd4135",
	})
	if err != nil {
		logrus.Fatalf("could not get document: %v", err)
	}
	logrus.Printf("Document: %s", doc.Document)
}

// list all documents
func (c *ExampleClient) ListDocuments() {
	// sort by year
	sort := []*proto.Sort{}
	sort = append(sort, &proto.Sort{
		Field:     "year",
		Ascending: true,
	})

	// filter by year
	filter := []*proto.Filter{}
	filter = append(filter, &proto.Filter{
		Operator: &proto.Filter_Equal{
			Equal: &proto.Equal{
				Field: "type",
				Value: &proto.Value{
					Type: &proto.Value_StringValue{
						StringValue: "movie",
					},
				},
			},
		},
	})

	documents, err := c.mongorpc.ListDocuments(c.ctx, &proto.ListDocumentsRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		Limit:      10,
		Sort:       sort,
		Filter:     filter,
	}, grpc.MaxCallRecvMsgSize(1024*1024*1024))
	if err != nil {
		logrus.Fatalf("could not get documents: %v", err)
	}
	logrus.Printf("Documents: %s", documents.Documents)
}

// create document
func (c *ExampleClient) CreateDocument() {

	// Movie document
	movie := Movie{
		Title: "The Shawshank Redemption",
		Year:  "1994",
	}

	// Encode the movie to JSON
	var result map[string]interface{}
	data, err := json.Marshal(movie)
	if err != nil {
		logrus.Fatalf("could not marshal movie: %v", err)
	}

	// Decode the JSON to a map
	err = json.Unmarshal(data, &result)
	if err != nil {
		logrus.Fatalf("could not unmarshal movie: %v", err)
	}

	// create movie document
	insertResp, err := c.mongorpc.CreateDocument(c.ctx, &proto.CreateDocumentRequest{
		Database:   "sample_mflix",
		Collection: "sample_movies",
		Document: &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: mongorpc.EncodeMap(result),
			},
		},
	})
	if err != nil {
		logrus.Fatalf("could not create document: %v", err)
	}
	logrus.Printf("Document: %s", insertResp.DocumentId)
}

// Movie struct
type Movie struct {
	// Title of the movie
	Title string `json:"title"`
	// Year of the movie
	Year string `json:"year"`
}
