// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: base.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int64             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message  string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current int32 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	Size    int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *PageRequest) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *PageRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current int32 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	Size    int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Total   int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PageReply) Reset() {
	*x = PageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageReply) ProtoMessage() {}

func (x *PageReply) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageReply.ProtoReflect.Descriptor instead.
func (*PageReply) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *PageReply) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *PageReply) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PageReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Asc   bool   `protobuf:"varint,2,opt,name=asc,proto3" json:"asc,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *Sort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sort) GetAsc() bool {
	if x != nil {
		return x.Asc
	}
	return false
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

func (x *Field) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Field) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type ListQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      *PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Sort      []*Sort      `protobuf:"bytes,2,rep,name=sort,proto3" json:"sort,omitempty"`
	Fields    []string     `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	Keyword   string       `protobuf:"bytes,5,opt,name=keyword,proto3" json:"keyword,omitempty"`
	StartAt   int64        `protobuf:"varint,6,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt     int64        `protobuf:"varint,7,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	TimeField string       `protobuf:"bytes,8,opt,name=time_field,json=timeField,proto3" json:"time_field,omitempty"`
}

func (x *ListQuery) Reset() {
	*x = ListQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuery) ProtoMessage() {}

func (x *ListQuery) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuery.ProtoReflect.Descriptor instead.
func (*ListQuery) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{5}
}

func (x *ListQuery) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListQuery) GetSort() []*Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListQuery) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ListQuery) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListQuery) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *ListQuery) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *ListQuery) GetTimeField() string {
	if x != nil {
		return x.TimeField
	}
	return ""
}

type ListQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *PageReply `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Fields []*Field   `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *ListQueryResult) Reset() {
	*x = ListQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQueryResult) ProtoMessage() {}

func (x *ListQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQueryResult.ProtoReflect.Descriptor instead.
func (*ListQueryResult) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{6}
}

func (x *ListQueryResult) GetPage() *PageReply {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListQueryResult) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x0b, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x20, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x09, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2e, 0x0a, 0x04, 0x53,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x73, 0x63, 0x22, 0x33, 0x0a, 0x05, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x22, 0xdd, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0x59, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x23, 0x0a, 0x03, 0x61,
	0x70, 0x69, 0x50, 0x01, 0x5a, 0x1a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_base_proto_goTypes = []interface{}{
	(*Response)(nil),        // 0: api.Response
	(*PageRequest)(nil),     // 1: api.PageRequest
	(*PageReply)(nil),       // 2: api.PageReply
	(*Sort)(nil),            // 3: api.Sort
	(*Field)(nil),           // 4: api.Field
	(*ListQuery)(nil),       // 5: api.ListQuery
	(*ListQueryResult)(nil), // 6: api.ListQueryResult
	nil,                     // 7: api.Response.MetadataEntry
}
var file_base_proto_depIdxs = []int32{
	7, // 0: api.Response.metadata:type_name -> api.Response.MetadataEntry
	1, // 1: api.ListQuery.page:type_name -> api.PageRequest
	3, // 2: api.ListQuery.sort:type_name -> api.Sort
	2, // 3: api.ListQueryResult.page:type_name -> api.PageReply
	4, // 4: api.ListQueryResult.fields:type_name -> api.Field
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageReply); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
		file_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuery); i {
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
		file_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQueryResult); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
