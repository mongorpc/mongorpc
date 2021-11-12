package interceptor

import (
	"context"

	gotrue "github.com/netlify/gotrue/api"
)

func (i *Interceptor) Permissions(ctx context.Context, claims *gotrue.GoTrueClaims, req interface{}) error {

	return nil
}

// func NewPermissionValidator() error {
// 	oso, err := oso.NewOso()
// 	if err != nil {
// 		return err
// 	}
// 	oso.RegisterClass(reflect.TypeOf(RequestPayload{}), nil)

// 	if err := oso.LoadFiles([]string{"main.polar"}); err != nil {
// 		return err
// 	}

// 	return nil
// }

// type RequestPayload struct {
// 	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
// 	// The collection to query
// 	Collection *string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
// 	// The documentid to query
// 	DocumentId *string `protobuf:"bytes,3,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
// 	// Limit the number of documents returned
// 	Limit *int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
// 	// Skip the first n documents
// 	Skip *int32 `protobuf:"varint,4,opt,name=skip,proto3" json:"skip,omitempty"`
// 	// Filter to apply to the query
// 	Filter map[string]map[string]interface{} `protobuf:"bytes,5,rep,name=filter,proto3" json:"filter,omitempty"`
// 	// Sort to apply to the query
// 	Sort []map[string]map[string]interface{} `protobuf:"bytes,6,rep,name=sort,proto3" json:"sort,omitempty"`
// 	// The document to insert
// 	Document interface{} `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
// }

// func ValidatePayload(req interface{}) (*RequestPayload, error) {
// 	payload := &RequestPayload{}
// 	switch r := req.(type) {
// 	case *proto.ListDocumentsRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.Limit = &r.Limit
// 		payload.Skip = &r.Skip
// 		payload.Filter = mongorpc.DecodeFilter(r.Filter)
// 		payload.Sort = mongorpc.DecodeSortToInterface(r.Sort)

// 	case *proto.GetDocumentRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.DocumentId = &r.DocumentId

// 	case *proto.CreateDocumentRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.Document = mongorpc.Decode(r.Document)

// 	case *proto.UpdateDocumentRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.DocumentId = &r.DocumentId
// 		payload.Document = mongorpc.Decode(r.Document)

// 	case *proto.DeleteDocumentRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.DocumentId = &r.DocumentId

// 	case *proto.CountDocumentsRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.Filter = mongorpc.DecodeFilter(r.Filter)

// 	case *proto.ListenRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection
// 		payload.Filter = mongorpc.DecodeFilter(r.Filter)
// 		payload.Sort = mongorpc.DecodeSortToInterface(r.Sort)

// 	case *proto.CreateIndexRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.ListIndexesRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.DeleteIndexRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.ReindexRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.CollectionStatsRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.CreateCollectionRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.RenameCollectionRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.DeleteCollectionRequest:
// 		payload.Database = r.Database
// 		payload.Collection = &r.Collection

// 	case *proto.HealthCheckRequest:
// 		break
// 	case *proto.Empty:
// 		break
// 	default:
// 		return nil, nil
// 	}
// 	return payload, nil
// }
