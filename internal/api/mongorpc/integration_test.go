//go:build integration

package mongorpc_test

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
	mongorpcapi "github.com/mongorpc/mongorpc/internal/api/mongorpc"
	"github.com/mongorpc/mongorpc/internal/repository/mongodb"
)

var (
	testDB     *mongodb.Client
	testServer *mongorpcapi.Server
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get MongoDB URI from environment or use default (standalone on port 27018)
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:password@localhost:27018/?authSource=admin"
	}

	var err error
	testDB, err = mongodb.New(ctx, mongoURI, "mongorpc_test")
	if err != nil {
		panic("Failed to connect to MongoDB: " + err.Error())
	}

	testServer = mongorpcapi.NewServer(testDB)

	// Run tests
	code := m.Run()

	// Cleanup
	_ = testDB.Database().Drop(context.Background())
	_ = testDB.Close(context.Background())

	os.Exit(code)
}

func TestCreateDocument_Integration(t *testing.T) {
	ctx := context.Background()

	// Create a document
	req := &mongorpcv1.CreateDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Document: &mongorpcv1.Document{
			Fields: map[string]*mongorpcv1.Value{
				"name": {ValueType: &mongorpcv1.Value_StringValue{StringValue: "Alice"}},
				"age":  {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 30}},
			},
		},
	}

	resp, err := testServer.CreateDocument(ctx, req)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}

	if resp.WriteResult.InsertedCount != 1 {
		t.Errorf("expected 1 inserted, got %d", resp.WriteResult.InsertedCount)
	}

	// Verify _id was returned
	idField := resp.Document.Fields["_id"]
	if idField == nil {
		t.Error("expected _id field in response")
	}
}

func TestGetDocument_Integration(t *testing.T) {
	ctx := context.Background()

	// First create a document
	createReq := &mongorpcv1.CreateDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Document: &mongorpcv1.Document{
			Fields: map[string]*mongorpcv1.Value{
				"name":  {ValueType: &mongorpcv1.Value_StringValue{StringValue: "Bob"}},
				"email": {ValueType: &mongorpcv1.Value_StringValue{StringValue: "bob@example.com"}},
			},
		},
	}

	createResp, err := testServer.CreateDocument(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}

	// Get the created document
	idValue := createResp.Document.Fields["_id"].GetObjectIdValue()
	getReq := &mongorpcv1.GetDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Id:         idValue,
	}

	getResp, err := testServer.GetDocument(ctx, getReq)
	if err != nil {
		t.Fatalf("GetDocument failed: %v", err)
	}

	if !getResp.Found {
		t.Error("expected document to be found")
	}

	if getResp.Document.Fields["name"].GetStringValue() != "Bob" {
		t.Errorf("expected name 'Bob', got %s", getResp.Document.Fields["name"].GetStringValue())
	}
}

func TestUpdateDocument_Integration(t *testing.T) {
	ctx := context.Background()

	// Create a document
	createReq := &mongorpcv1.CreateDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Document: &mongorpcv1.Document{
			Fields: map[string]*mongorpcv1.Value{
				"name": {ValueType: &mongorpcv1.Value_StringValue{StringValue: "Charlie"}},
				"age":  {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 25}},
			},
		},
	}

	createResp, err := testServer.CreateDocument(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}

	// Update the document
	idValue := createResp.Document.Fields["_id"].GetObjectIdValue()
	updateReq := &mongorpcv1.UpdateDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Id:         idValue,
		Update: &mongorpcv1.UpdateSpec{
			UpdateType: &mongorpcv1.UpdateSpec_Operators{
				Operators: &mongorpcv1.UpdateOperators{
					Set: map[string]*mongorpcv1.Value{
						"age": {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 26}},
					},
				},
			},
		},
	}

	updateResp, err := testServer.UpdateDocument(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateDocument failed: %v", err)
	}

	if updateResp.WriteResult.MatchedCount != 1 {
		t.Errorf("expected 1 matched, got %d", updateResp.WriteResult.MatchedCount)
	}
	if updateResp.WriteResult.ModifiedCount != 1 {
		t.Errorf("expected 1 modified, got %d", updateResp.WriteResult.ModifiedCount)
	}
}

func TestDeleteDocument_Integration(t *testing.T) {
	ctx := context.Background()

	// Create a document
	createReq := &mongorpcv1.CreateDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Document: &mongorpcv1.Document{
			Fields: map[string]*mongorpcv1.Value{
				"name": {ValueType: &mongorpcv1.Value_StringValue{StringValue: "ToDelete"}},
			},
		},
	}

	createResp, err := testServer.CreateDocument(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}

	// Delete the document
	idValue := createResp.Document.Fields["_id"].GetObjectIdValue()
	deleteReq := &mongorpcv1.DeleteDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Id:         idValue,
	}

	deleteResp, err := testServer.DeleteDocument(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteDocument failed: %v", err)
	}

	if deleteResp.WriteResult.DeletedCount != 1 {
		t.Errorf("expected 1 deleted, got %d", deleteResp.WriteResult.DeletedCount)
	}

	// Verify document is gone
	getReq := &mongorpcv1.GetDocumentRequest{
		Database:   "mongorpc_test",
		Collection: "users",
		Id:         idValue,
	}

	getResp, _ := testServer.GetDocument(ctx, getReq)
	if getResp.Found {
		t.Error("expected document to be deleted")
	}
}

func TestCountDocuments_Integration(t *testing.T) {
	ctx := context.Background()
	collName := "count_test"

	// Insert multiple documents
	coll := testDB.Client().Database("mongorpc_test").Collection(collName)
	_, _ = coll.InsertMany(ctx, []interface{}{
		bson.M{"x": 1},
		bson.M{"x": 2},
		bson.M{"x": 3},
	})
	defer coll.Drop(ctx)

	req := &mongorpcv1.CountDocumentsRequest{
		Database:   "mongorpc_test",
		Collection: collName,
	}

	resp, err := testServer.CountDocuments(ctx, req)
	if err != nil {
		t.Fatalf("CountDocuments failed: %v", err)
	}

	if resp.Count != 3 {
		t.Errorf("expected count 3, got %d", resp.Count)
	}
}

func TestInsertMany_Integration(t *testing.T) {
	ctx := context.Background()
	collName := "insert_many_test"

	defer testDB.Client().Database("mongorpc_test").Collection(collName).Drop(ctx)

	req := &mongorpcv1.InsertManyRequest{
		Database:   "mongorpc_test",
		Collection: collName,
		Documents: []*mongorpcv1.Document{
			{Fields: map[string]*mongorpcv1.Value{"n": {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 1}}}},
			{Fields: map[string]*mongorpcv1.Value{"n": {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 2}}}},
			{Fields: map[string]*mongorpcv1.Value{"n": {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 3}}}},
		},
	}

	resp, err := testServer.InsertMany(ctx, req)
	if err != nil {
		t.Fatalf("InsertMany failed: %v", err)
	}

	if len(resp.InsertedIds) != 3 {
		t.Errorf("expected 3 inserted ids, got %d", len(resp.InsertedIds))
	}
	if resp.WriteResult.InsertedCount != 3 {
		t.Errorf("expected 3 inserted count, got %d", resp.WriteResult.InsertedCount)
	}
}

func TestDeleteMany_Integration(t *testing.T) {
	ctx := context.Background()
	collName := "delete_many_test"

	// Insert documents
	coll := testDB.Client().Database("mongorpc_test").Collection(collName)
	_, _ = coll.InsertMany(ctx, []interface{}{
		bson.M{"status": "pending"},
		bson.M{"status": "pending"},
		bson.M{"status": "done"},
	})
	defer coll.Drop(ctx)

	req := &mongorpcv1.DeleteManyRequest{
		Database:   "mongorpc_test",
		Collection: collName,
		// Delete all (empty filter)
	}

	resp, err := testServer.DeleteMany(ctx, req)
	if err != nil {
		t.Fatalf("DeleteMany failed: %v", err)
	}

	if resp.WriteResult.DeletedCount != 3 {
		t.Errorf("expected 3 deleted, got %d", resp.WriteResult.DeletedCount)
	}
}
