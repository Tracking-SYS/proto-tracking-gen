// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: services/product.proto

package services

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	entities "github.com/lk153/proto-tracking-gen/go/entities"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SingleProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SingleProductRequest) Reset() {
	*x = SingleProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleProductRequest) ProtoMessage() {}

func (x *SingleProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleProductRequest.ProtoReflect.Descriptor instead.
func (*SingleProductRequest) Descriptor() ([]byte, []int) {
	return file_services_product_proto_rawDescGZIP(), []int{0}
}

func (x *SingleProductRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SingleProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *entities.ProductInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SingleProductResponse) Reset() {
	*x = SingleProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleProductResponse) ProtoMessage() {}

func (x *SingleProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleProductResponse.ProtoReflect.Descriptor instead.
func (*SingleProductResponse) Descriptor() ([]byte, []int) {
	return file_services_product_proto_rawDescGZIP(), []int{1}
}

func (x *SingleProductResponse) GetData() *entities.ProductInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Page  uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_services_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProductRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ProductRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*entities.ProductInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_services_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductResponse) GetData() []*entities.ProductInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProductCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Page  uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ProductCreateRequest) Reset() {
	*x = ProductCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreateRequest) ProtoMessage() {}

func (x *ProductCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreateRequest.ProtoReflect.Descriptor instead.
func (*ProductCreateRequest) Descriptor() ([]byte, []int) {
	return file_services_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductCreateRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProductCreateRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ProductCreateRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ProductCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *entities.ProductInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProductCreateResponse) Reset() {
	*x = ProductCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreateResponse) ProtoMessage() {}

func (x *ProductCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreateResponse.ProtoReflect.Descriptor instead.
func (*ProductCreateResponse) Descriptor() ([]byte, []int) {
	return file_services_product_proto_rawDescGZIP(), []int{5}
}

func (x *ProductCreateResponse) GetData() *entities.ProductInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_services_product_proto protoreflect.FileDescriptor

var file_services_product_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a,
	0x14, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x30,
	0x92, 0x41, 0x2d, 0x0a, 0x1a, 0x2a, 0x0b, 0x47, 0x65, 0x74, 0x20, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x32, 0x0b, 0x47, 0x65, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32,
	0x0f, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x37, 0x37, 0x37, 0x34, 0x35, 0x38, 0x34, 0x36, 0x7d,
	0x22, 0x4e, 0x0a, 0x15, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xd8, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x6e, 0x92, 0x41, 0x6b,
	0x0a, 0x2b, 0x2a, 0x13, 0x47, 0x65, 0x74, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x4f, 0x66, 0x20,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0x14, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x3c, 0x7b,
	0x22, 0x69, 0x64, 0x73, 0x22, 0x3a, 0x5b, 0x37, 0x37, 0x37, 0x34, 0x35, 0x38, 0x34, 0x36, 0x2c,
	0x37, 0x37, 0x37, 0x34, 0x37, 0x30, 0x37, 0x32, 0x2c, 0x37, 0x37, 0x37, 0x35, 0x30, 0x35, 0x35,
	0x38, 0x5d, 0x2c, 0x20, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3a, 0x20, 0x32, 0x30, 0x7d, 0x22, 0x48, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xda, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x6a, 0x92, 0x41, 0x67, 0x0a, 0x20, 0x2a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0x43, 0x7b,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x20, 0x31, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x3a, 0x20, 0x32, 0x30,
	0x30, 0x30, 0x2c, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20,
	0x31, 0x7d, 0x22, 0x4e, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0xf5, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x67,
	0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x79, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0xbc, 0x03, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6b, 0x31, 0x35, 0x33, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x92, 0x41, 0xfe, 0x02, 0x12, 0x2d, 0x0a, 0x10, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x41, 0x50, 0x49, 0x12,
	0x14, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x31, 0x32, 0x37, 0x2e,
	0x30, 0x2e, 0x30, 0x2e, 0x31, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0xd8,
	0x01, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0xd0, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xbf, 0x01, 0x0a, 0x10, 0x58, 0x2d, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x49, 0x64, 0x12, 0xaa, 0x01, 0x0a,
	0x2b, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0x26, 0x22, 0x32, 0x34, 0x33,
	0x38, 0x61, 0x63, 0x33, 0x63, 0x2d, 0x33, 0x37, 0x65, 0x62, 0x2d, 0x34, 0x39, 0x30, 0x32, 0x2d,
	0x61, 0x64, 0x65, 0x66, 0x2d, 0x65, 0x64, 0x31, 0x36, 0x62, 0x34, 0x34, 0x33, 0x31, 0x30, 0x33,
	0x30, 0x22, 0x6a, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d,
	0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30,
	0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d,
	0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39,
	0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_product_proto_rawDescOnce sync.Once
	file_services_product_proto_rawDescData = file_services_product_proto_rawDesc
)

func file_services_product_proto_rawDescGZIP() []byte {
	file_services_product_proto_rawDescOnce.Do(func() {
		file_services_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_product_proto_rawDescData)
	})
	return file_services_product_proto_rawDescData
}

var file_services_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_product_proto_goTypes = []interface{}{
	(*SingleProductRequest)(nil),  // 0: go.tracking.services.SingleProductRequest
	(*SingleProductResponse)(nil), // 1: go.tracking.services.SingleProductResponse
	(*ProductRequest)(nil),        // 2: go.tracking.services.ProductRequest
	(*ProductResponse)(nil),       // 3: go.tracking.services.ProductResponse
	(*ProductCreateRequest)(nil),  // 4: go.tracking.services.ProductCreateRequest
	(*ProductCreateResponse)(nil), // 5: go.tracking.services.ProductCreateResponse
	(*entities.ProductInfo)(nil),  // 6: go.tracking.entities.ProductInfo
}
var file_services_product_proto_depIdxs = []int32{
	6, // 0: go.tracking.services.SingleProductResponse.data:type_name -> go.tracking.entities.ProductInfo
	6, // 1: go.tracking.services.ProductResponse.data:type_name -> go.tracking.entities.ProductInfo
	6, // 2: go.tracking.services.ProductCreateResponse.data:type_name -> go.tracking.entities.ProductInfo
	0, // 3: go.tracking.services.ProductService.GetSingle:input_type -> go.tracking.services.SingleProductRequest
	2, // 4: go.tracking.services.ProductService.Get:input_type -> go.tracking.services.ProductRequest
	4, // 5: go.tracking.services.ProductService.Create:input_type -> go.tracking.services.ProductCreateRequest
	1, // 6: go.tracking.services.ProductService.GetSingle:output_type -> go.tracking.services.SingleProductResponse
	3, // 7: go.tracking.services.ProductService.Get:output_type -> go.tracking.services.ProductResponse
	5, // 8: go.tracking.services.ProductService.Create:output_type -> go.tracking.services.ProductCreateResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_product_proto_init() }
func file_services_product_proto_init() {
	if File_services_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleProductRequest); i {
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
		file_services_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleProductResponse); i {
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
		file_services_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_services_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
		file_services_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreateRequest); i {
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
		file_services_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreateResponse); i {
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
			RawDescriptor: file_services_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_product_proto_goTypes,
		DependencyIndexes: file_services_product_proto_depIdxs,
		MessageInfos:      file_services_product_proto_msgTypes,
	}.Build()
	File_services_product_proto = out.File
	file_services_product_proto_rawDesc = nil
	file_services_product_proto_goTypes = nil
	file_services_product_proto_depIdxs = nil
}
