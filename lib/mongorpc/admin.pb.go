// Copyright 2021 MongoRPC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: mongorpc/admin.proto

package mongorpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IndexDirection int32

const (
	IndexDirection_ASCENDING  IndexDirection = 0
	IndexDirection_DESCENDING IndexDirection = 1
)

// Enum value maps for IndexDirection.
var (
	IndexDirection_name = map[int32]string{
		0: "ASCENDING",
		1: "DESCENDING",
	}
	IndexDirection_value = map[string]int32{
		"ASCENDING":  0,
		"DESCENDING": 1,
	}
)

func (x IndexDirection) Enum() *IndexDirection {
	p := new(IndexDirection)
	*p = x
	return p
}

func (x IndexDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_mongorpc_admin_proto_enumTypes[0].Descriptor()
}

func (IndexDirection) Type() protoreflect.EnumType {
	return &file_mongorpc_admin_proto_enumTypes[0]
}

func (x IndexDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexDirection.Descriptor instead.
func (IndexDirection) EnumDescriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{0}
}

type DropDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *DropDatabaseRequest) Reset() {
	*x = DropDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropDatabaseRequest) ProtoMessage() {}

func (x *DropDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DropDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{0}
}

func (x *DropDatabaseRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

type CreateCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database   string                   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Collection string                   `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	Options    *CreateCollectionOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CreateCollectionRequest) Reset() {
	*x = CreateCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionRequest) ProtoMessage() {}

func (x *CreateCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionRequest.ProtoReflect.Descriptor instead.
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCollectionRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *CreateCollectionRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *CreateCollectionRequest) GetOptions() *CreateCollectionOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type CreateCollectionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size             int64  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Max              int64  `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	Validator        *Value `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	ValidationLevel  string `protobuf:"bytes,4,opt,name=validationLevel,proto3" json:"validationLevel,omitempty"`
	ValidationAction string `protobuf:"bytes,5,opt,name=validationAction,proto3" json:"validationAction,omitempty"`
	StorageEngine    *Value `protobuf:"bytes,6,opt,name=storageEngine,proto3" json:"storageEngine,omitempty"`
}

func (x *CreateCollectionOptions) Reset() {
	*x = CreateCollectionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionOptions) ProtoMessage() {}

func (x *CreateCollectionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionOptions.ProtoReflect.Descriptor instead.
func (*CreateCollectionOptions) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCollectionOptions) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateCollectionOptions) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *CreateCollectionOptions) GetValidator() *Value {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *CreateCollectionOptions) GetValidationLevel() string {
	if x != nil {
		return x.ValidationLevel
	}
	return ""
}

func (x *CreateCollectionOptions) GetValidationAction() string {
	if x != nil {
		return x.ValidationAction
	}
	return ""
}

func (x *CreateCollectionOptions) GetStorageEngine() *Value {
	if x != nil {
		return x.StorageEngine
	}
	return nil
}

type DropCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database   string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *DropCollectionRequest) Reset() {
	*x = DropCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropCollectionRequest) ProtoMessage() {}

func (x *DropCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropCollectionRequest.ProtoReflect.Descriptor instead.
func (*DropCollectionRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{3}
}

func (x *DropCollectionRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DropCollectionRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

type ListCollectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *ListCollectionsRequest) Reset() {
	*x = ListCollectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsRequest) ProtoMessage() {}

func (x *ListCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsRequest.ProtoReflect.Descriptor instead.
func (*ListCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{4}
}

func (x *ListCollectionsRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

type ListIndexesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database   string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *ListIndexesRequest) Reset() {
	*x = ListIndexesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIndexesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIndexesRequest) ProtoMessage() {}

func (x *ListIndexesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIndexesRequest.ProtoReflect.Descriptor instead.
func (*ListIndexesRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{5}
}

func (x *ListIndexesRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *ListIndexesRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

type CreateIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database   string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	Index      *Index `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *CreateIndexRequest) Reset() {
	*x = CreateIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIndexRequest) ProtoMessage() {}

func (x *CreateIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIndexRequest.ProtoReflect.Descriptor instead.
func (*CreateIndexRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{6}
}

func (x *CreateIndexRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *CreateIndexRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *CreateIndexRequest) GetIndex() *Index {
	if x != nil {
		return x.Index
	}
	return nil
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the index
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The keys to index
	Keys []*IndexKey `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	// The unique flag
	Unique bool `protobuf:"varint,3,opt,name=unique,proto3" json:"unique,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{7}
}

func (x *Index) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index) GetKeys() []*IndexKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Index) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

type IndexKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field to index
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The direction to index
	Direction IndexDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=mongorpc.IndexDirection" json:"direction,omitempty"`
}

func (x *IndexKey) Reset() {
	*x = IndexKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexKey) ProtoMessage() {}

func (x *IndexKey) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexKey.ProtoReflect.Descriptor instead.
func (*IndexKey) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{8}
}

func (x *IndexKey) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *IndexKey) GetDirection() IndexDirection {
	if x != nil {
		return x.Direction
	}
	return IndexDirection_ASCENDING
}

type DropIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database   string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	Index      string `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *DropIndexRequest) Reset() {
	*x = DropIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongorpc_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropIndexRequest) ProtoMessage() {}

func (x *DropIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongorpc_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropIndexRequest.ProtoReflect.Descriptor instead.
func (*DropIndexRequest) Descriptor() ([]byte, []int) {
	return file_mongorpc_admin_proto_rawDescGZIP(), []int{9}
}

func (x *DropIndexRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DropIndexRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *DropIndexRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

var File_mongorpc_admin_proto protoreflect.FileDescriptor

var file_mongorpc_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63,
	0x1a, 0x14, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x13, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xfb,
	0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x12, 0x2d, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x28, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0x53, 0x0a, 0x15,
	0x44, 0x72, 0x6f, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x34, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0x5b, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4b, 0x65,
	0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x22,
	0x58, 0x0a, 0x08, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x36, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x10, 0x44, 0x72, 0x6f,
	0x70, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2a,
	0x2f, 0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x32, 0x8a, 0x04, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x52, 0x50, 0x43, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x31, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a,
	0x0e, 0x44, 0x72, 0x6f, 0x70, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x44, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x44, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x72, 0x6f, 0x70,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x3b, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mongorpc_admin_proto_rawDescOnce sync.Once
	file_mongorpc_admin_proto_rawDescData = file_mongorpc_admin_proto_rawDesc
)

func file_mongorpc_admin_proto_rawDescGZIP() []byte {
	file_mongorpc_admin_proto_rawDescOnce.Do(func() {
		file_mongorpc_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_mongorpc_admin_proto_rawDescData)
	})
	return file_mongorpc_admin_proto_rawDescData
}

var file_mongorpc_admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mongorpc_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_mongorpc_admin_proto_goTypes = []interface{}{
	(IndexDirection)(0),             // 0: mongorpc.IndexDirection
	(*DropDatabaseRequest)(nil),     // 1: mongorpc.DropDatabaseRequest
	(*CreateCollectionRequest)(nil), // 2: mongorpc.CreateCollectionRequest
	(*CreateCollectionOptions)(nil), // 3: mongorpc.CreateCollectionOptions
	(*DropCollectionRequest)(nil),   // 4: mongorpc.DropCollectionRequest
	(*ListCollectionsRequest)(nil),  // 5: mongorpc.ListCollectionsRequest
	(*ListIndexesRequest)(nil),      // 6: mongorpc.ListIndexesRequest
	(*CreateIndexRequest)(nil),      // 7: mongorpc.CreateIndexRequest
	(*Index)(nil),                   // 8: mongorpc.Index
	(*IndexKey)(nil),                // 9: mongorpc.IndexKey
	(*DropIndexRequest)(nil),        // 10: mongorpc.DropIndexRequest
	(*Value)(nil),                   // 11: mongorpc.Value
	(*Empty)(nil),                   // 12: mongorpc.Empty
}
var file_mongorpc_admin_proto_depIdxs = []int32{
	3,  // 0: mongorpc.CreateCollectionRequest.options:type_name -> mongorpc.CreateCollectionOptions
	11, // 1: mongorpc.CreateCollectionOptions.validator:type_name -> mongorpc.Value
	11, // 2: mongorpc.CreateCollectionOptions.storageEngine:type_name -> mongorpc.Value
	8,  // 3: mongorpc.CreateIndexRequest.index:type_name -> mongorpc.Index
	9,  // 4: mongorpc.Index.keys:type_name -> mongorpc.IndexKey
	0,  // 5: mongorpc.IndexKey.direction:type_name -> mongorpc.IndexDirection
	12, // 6: mongorpc.MongoRPCAdmin.ListDatabases:input_type -> mongorpc.Empty
	1,  // 7: mongorpc.MongoRPCAdmin.DropDatabase:input_type -> mongorpc.DropDatabaseRequest
	2,  // 8: mongorpc.MongoRPCAdmin.CreateCollection:input_type -> mongorpc.CreateCollectionRequest
	4,  // 9: mongorpc.MongoRPCAdmin.DropCollection:input_type -> mongorpc.DropCollectionRequest
	5,  // 10: mongorpc.MongoRPCAdmin.ListCollections:input_type -> mongorpc.ListCollectionsRequest
	6,  // 11: mongorpc.MongoRPCAdmin.ListIndexes:input_type -> mongorpc.ListIndexesRequest
	7,  // 12: mongorpc.MongoRPCAdmin.CreateIndex:input_type -> mongorpc.CreateIndexRequest
	10, // 13: mongorpc.MongoRPCAdmin.DropIndex:input_type -> mongorpc.DropIndexRequest
	11, // 14: mongorpc.MongoRPCAdmin.ListDatabases:output_type -> mongorpc.Value
	12, // 15: mongorpc.MongoRPCAdmin.DropDatabase:output_type -> mongorpc.Empty
	12, // 16: mongorpc.MongoRPCAdmin.CreateCollection:output_type -> mongorpc.Empty
	12, // 17: mongorpc.MongoRPCAdmin.DropCollection:output_type -> mongorpc.Empty
	11, // 18: mongorpc.MongoRPCAdmin.ListCollections:output_type -> mongorpc.Value
	11, // 19: mongorpc.MongoRPCAdmin.ListIndexes:output_type -> mongorpc.Value
	11, // 20: mongorpc.MongoRPCAdmin.CreateIndex:output_type -> mongorpc.Value
	12, // 21: mongorpc.MongoRPCAdmin.DropIndex:output_type -> mongorpc.Empty
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mongorpc_admin_proto_init() }
func file_mongorpc_admin_proto_init() {
	if File_mongorpc_admin_proto != nil {
		return
	}
	file_mongorpc_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mongorpc_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropDatabaseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropCollectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIndexesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIndexRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mongorpc_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropIndexRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mongorpc_admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mongorpc_admin_proto_goTypes,
		DependencyIndexes: file_mongorpc_admin_proto_depIdxs,
		EnumInfos:         file_mongorpc_admin_proto_enumTypes,
		MessageInfos:      file_mongorpc_admin_proto_msgTypes,
	}.Build()
	File_mongorpc_admin_proto = out.File
	file_mongorpc_admin_proto_rawDesc = nil
	file_mongorpc_admin_proto_goTypes = nil
	file_mongorpc_admin_proto_depIdxs = nil
}