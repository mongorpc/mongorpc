package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// gRPC server address
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	c := proto.NewMongoRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Get List of Collections
	listCollections(err, c, ctx)
	// list documents
	err, movie, result, data := listDocuments(err, c, ctx)
	// create document
	insertResp := createDocument(err, c, ctx, result)
	// update document
	updateDocument(movie, data, err, result, c, ctx, insertResp)

}

func listCollections(err error, c proto.MongoRPCClient, ctx context.Context) {
	r, err := c.ListCollections(ctx, &proto.ListCollectionsRequest{
		Database: "sample_mflix",
	})
	if err != nil {
		logrus.Fatalf("could not get collection: %v", err)
	}
	logrus.Printf("Collection: %s", r.Collections)

	// get document
	doc, err := c.GetDocument(ctx, &proto.GetDocumentRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		DocumentId: "573a1390f29313caabcd4135",
	})
	if err != nil {
		logrus.Fatalf("could not get document: %v", err)
	}
	logrus.Printf("Document: %s", doc.Document)
}

func listDocuments(err error, c proto.MongoRPCClient, ctx context.Context) (error, Movie, map[string]interface{}, []byte) {
	documents, err := c.ListDocuments(ctx, &proto.ListDocumentsRequest{
		Database:   "sample_mflix",
		Collection: "movies",
		Limit:      1,
	}, grpc.MaxCallRecvMsgSize(1024*1024*1024))
	if err != nil {
		logrus.Fatalf("could not get documents: %v", err)
	}
	logrus.Printf("Documents: %s", documents.Documents)

	movie := Movie{
		Title: "The Shawshank Redemption",
		Year:  "1994",
	}

	var result map[string]interface{}

	data, err := json.Marshal(movie)
	if err != nil {
		logrus.Fatalf("could not marshal movie: %v", err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		logrus.Fatalf("could not unmarshal movie: %v", err)
	}
	return err, movie, result, data
}

func updateDocument(movie Movie, data []byte, err error, result map[string]interface{}, c proto.MongoRPCClient, ctx context.Context, insertResp *proto.CreateDocumentResponse) {
	movie.Title = randomMovieTitle()
	movie.Year = "1972"
	data, err = json.Marshal(movie)
	if err != nil {
		logrus.Fatalf("could not marshal movie: %v", err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		logrus.Fatalf("could not unmarshal movie: %v", err)
	}

	updateResp, err := c.UpdateDocument(ctx, &proto.UpdateDocumentRequest{
		Database:   "sample_mflix",
		Collection: "moviesx",
		DocumentId: insertResp.DocumentId,
		Document: &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: mongorpc.EncodeMap(result),
			},
		},
	})
	if err != nil {
		logrus.Fatalf("could not update document: %v", err)
	}
	logrus.Printf("Document: %s", updateResp)
}

func createDocument(err error, c proto.MongoRPCClient, ctx context.Context, result map[string]interface{}) *proto.CreateDocumentResponse {

	insertResp, err := c.CreateDocument(ctx, &proto.CreateDocumentRequest{
		Database:   "sample_mflix",
		Collection: "moviesx",
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
	return insertResp
}

type Movie struct {
	Title string `json:"name"`
	Year  string `json:"year"`
}

func randomMovieTitle() string {
	movies := []string{
		"The Shawshank Redemption",
		"The Godfather",
		"The Godfather: Part II",
		"The Dark Knight",
		"12 Angry",
		"Schindler's List",
		"Pulp Fiction",
		"The Lord of the Rings: The Return of the King",
		"Fight Club",
		"The Lord of the Rings: The Fellowship of the Ring",
		"Forrest Gump",
		"Inception",
		"Star Wars: Episode V - The Empire Strikes Back",
		"The Lord of the Rings: The Two Towers",
		"One Flew Over the Cuckoo's Nest",
		"In the Name of the Father",
		"Goodfellas",
		"Star Wars: Episode IV - A New Hope",
		"Seven Samurai",
		"The Matrix",
		"City of God",
		"Se7en",
		"The Silence of the Lambs",
		"It's a Wonderful Life",
		"Life Is Beautiful",
	}
	return movies[rand.Intn(len(movies))]
}
