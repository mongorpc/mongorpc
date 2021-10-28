package mongorpc

import (
	"context"

	"github.com/mongorpc/mongorpc/proto"
)

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
func (srv *MongoRPCServer) DeleteDocument(ctx context.Context, in *proto.DeleteDocumentRequest) (*proto.DeleteDocumentResponse, error) {

	// Delete document
	// err := deleteDocument(in.Collection, in.Id)
	// if err != nil {
	// 	return nil, err
	// }

	// Return document
	return &proto.DeleteDocumentResponse{}, nil
}
