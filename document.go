package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
)

// Get document by id from collection in database
func (srv *MongoRPCServer) GetDocument(ctx context.Context, in *proto.GetDocumentRequest) (*proto.GetDocumentResponse, error) {

	// Get document
	// doc, err := getDocument(in.Collection, in.Id)
	// if err != nil {
	// 	return nil, err
	// }

	// Return document
	return &proto.GetDocumentResponse{
		// Document: doc,
	}, nil
}

// List documents from collection in database
func (srv *MongoRPCServer) ListDocuments(ctx context.Context, in *proto.ListDocumentsRequest) (*proto.ListDocumentsResponse, error) {

	// Get documents
	// docs, err := listDocuments(in.Collection)
	// if err != nil {
	// 	return nil, err
	// }

	// Return documents
	return &proto.ListDocumentsResponse{
		// Documents: docs,
	}, nil
}

// Create document in collection in database
func (srv *MongoRPCServer) CreateDocument(ctx context.Context, in *proto.CreateDocumentRequest) (*proto.CreateDocumentResponse, error) {

	// Create document
	// doc, err := createDocument(in.Collection, in.Document)
	// if err != nil {
	// 	return nil, err
	// }

	// Return document
	return &proto.CreateDocumentResponse{
		// Document: doc,
	}, nil
}

// Update document in collection in database
func (srv *MongoRPCServer) UpdateDocument(ctx context.Context, in *proto.UpdateDocumentRequest) (*proto.UpdateDocumentResponse, error) {

	// Update document
	// doc, err := updateDocument(in.Collection, in.Id, in.Document)
	// if err != nil {
	// 	return nil, err
	// }

	// Return document
	return &proto.UpdateDocumentResponse{
		// Document: doc,
	}, nil
}

// Delete document from collection in database
func (srv *MongoRPCServer) DeleteDocument(ctx context.Context, in *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {

	// Delete document
	res, err := srv.DB.Database(in.Collection).Collection(in.Collection).DeleteOne(ctx, in.DocumentId)
	if err != nil {
		return nil, err
	}

	// Return document
	return &proto.DeleteDocumentResponse{
		DeletedCount: res.DeletedCount,
	}, nil
}
