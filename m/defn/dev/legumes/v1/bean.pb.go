// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: legumes/v1/bean.proto

// $schema: defn.dev.legumes.v1.Bean
// $title: Cool bean
// $description: So many cool beans
// $location: https://defn.dev/legumes/bean/v1

// cool

package legumesv1

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

// <!-- crd generation tags
// +cue-gen:Bean:groupName:legumes.defn.dev
// +cue-gen:Bean:version:v1
// +cue-gen:Bean:storageVersion
// +cue-gen:Bean:subresource:status
// +cue-gen:Bean:scope:Namespaced
// +cue-gen:Bean:preserveUnknownFields:pluginConfig
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=legumes.defn.dev/v1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// cool
type Bean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cool
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// cool
	Sha256 string `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *Bean) Reset() {
	*x = Bean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bean) ProtoMessage() {}

func (x *Bean) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bean.ProtoReflect.Descriptor instead.
func (*Bean) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{0}
}

func (x *Bean) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Bean) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

type GetBeanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Bean `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetBeanRequest) Reset() {
	*x = GetBeanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeanRequest) ProtoMessage() {}

func (x *GetBeanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeanRequest.ProtoReflect.Descriptor instead.
func (*GetBeanRequest) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{1}
}

func (x *GetBeanRequest) GetReq() *Bean {
	if x != nil {
		return x.Req
	}
	return nil
}

type GetBeanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Bean `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *GetBeanResponse) Reset() {
	*x = GetBeanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeanResponse) ProtoMessage() {}

func (x *GetBeanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeanResponse.ProtoReflect.Descriptor instead.
func (*GetBeanResponse) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{2}
}

func (x *GetBeanResponse) GetReq() *Bean {
	if x != nil {
		return x.Req
	}
	return nil
}

type PutBeanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Bean `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *PutBeanRequest) Reset() {
	*x = PutBeanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutBeanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutBeanRequest) ProtoMessage() {}

func (x *PutBeanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutBeanRequest.ProtoReflect.Descriptor instead.
func (*PutBeanRequest) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{3}
}

func (x *PutBeanRequest) GetReq() *Bean {
	if x != nil {
		return x.Req
	}
	return nil
}

type PutBeanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Bean `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *PutBeanResponse) Reset() {
	*x = PutBeanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutBeanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutBeanResponse) ProtoMessage() {}

func (x *PutBeanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutBeanResponse.ProtoReflect.Descriptor instead.
func (*PutBeanResponse) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{4}
}

func (x *PutBeanResponse) GetReq() *Bean {
	if x != nil {
		return x.Req
	}
	return nil
}

type DeleteBeanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Bean `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *DeleteBeanRequest) Reset() {
	*x = DeleteBeanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBeanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBeanRequest) ProtoMessage() {}

func (x *DeleteBeanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBeanRequest.ProtoReflect.Descriptor instead.
func (*DeleteBeanRequest) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBeanRequest) GetReq() *Bean {
	if x != nil {
		return x.Req
	}
	return nil
}

type DeleteBeanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *Bean `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *DeleteBeanResponse) Reset() {
	*x = DeleteBeanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legumes_v1_bean_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBeanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBeanResponse) ProtoMessage() {}

func (x *DeleteBeanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_legumes_v1_bean_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBeanResponse.ProtoReflect.Descriptor instead.
func (*DeleteBeanResponse) Descriptor() ([]byte, []int) {
	return file_legumes_v1_bean_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBeanResponse) GetReq() *Bean {
	if x != nil {
		return x.Req
	}
	return nil
}

var File_legumes_v1_bean_proto protoreflect.FileDescriptor

var file_legumes_v1_bean_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x04,
	0x42, 0x65, 0x61, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x22, 0x3d,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x3e, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x3d, 0x0a,
	0x0e, 0x50, 0x75, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64,
	0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x3e, 0x0a, 0x0f,
	0x50, 0x75, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64,
	0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x40, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x41,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67,
	0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x03, 0x72, 0x65,
	0x71, 0x32, 0xa3, 0x02, 0x0a, 0x10, 0x42, 0x65, 0x61, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x65, 0x61,
	0x6e, 0x12, 0x23, 0x2e, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67,
	0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x07, 0x50, 0x75, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x12, 0x23, 0x2e, 0x64, 0x65, 0x66, 0x6e,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x65, 0x61, 0x6e, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64,
	0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc8, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x65, 0x66, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x42, 0x65, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x66,
	0x6e, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2f, 0x64, 0x65, 0x66, 0x6e, 0x2f, 0x64, 0x65, 0x76,
	0x2f, 0x6c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x65, 0x67, 0x75,
	0x6d, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x44, 0x4c, 0xaa, 0x02, 0x13, 0x44, 0x65,
	0x66, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x4c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x13, 0x44, 0x65, 0x66, 0x6e, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x4c, 0x65, 0x67,
	0x75, 0x6d, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x44, 0x65, 0x66, 0x6e, 0x5c, 0x44,
	0x65, 0x76, 0x5c, 0x4c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x44, 0x65, 0x66, 0x6e,
	0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x4c, 0x65, 0x67, 0x75, 0x6d, 0x65, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_legumes_v1_bean_proto_rawDescOnce sync.Once
	file_legumes_v1_bean_proto_rawDescData = file_legumes_v1_bean_proto_rawDesc
)

func file_legumes_v1_bean_proto_rawDescGZIP() []byte {
	file_legumes_v1_bean_proto_rawDescOnce.Do(func() {
		file_legumes_v1_bean_proto_rawDescData = protoimpl.X.CompressGZIP(file_legumes_v1_bean_proto_rawDescData)
	})
	return file_legumes_v1_bean_proto_rawDescData
}

var file_legumes_v1_bean_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_legumes_v1_bean_proto_goTypes = []interface{}{
	(*Bean)(nil),               // 0: defn.dev.legumes.v1.Bean
	(*GetBeanRequest)(nil),     // 1: defn.dev.legumes.v1.GetBeanRequest
	(*GetBeanResponse)(nil),    // 2: defn.dev.legumes.v1.GetBeanResponse
	(*PutBeanRequest)(nil),     // 3: defn.dev.legumes.v1.PutBeanRequest
	(*PutBeanResponse)(nil),    // 4: defn.dev.legumes.v1.PutBeanResponse
	(*DeleteBeanRequest)(nil),  // 5: defn.dev.legumes.v1.DeleteBeanRequest
	(*DeleteBeanResponse)(nil), // 6: defn.dev.legumes.v1.DeleteBeanResponse
}
var file_legumes_v1_bean_proto_depIdxs = []int32{
	0, // 0: defn.dev.legumes.v1.GetBeanRequest.req:type_name -> defn.dev.legumes.v1.Bean
	0, // 1: defn.dev.legumes.v1.GetBeanResponse.req:type_name -> defn.dev.legumes.v1.Bean
	0, // 2: defn.dev.legumes.v1.PutBeanRequest.req:type_name -> defn.dev.legumes.v1.Bean
	0, // 3: defn.dev.legumes.v1.PutBeanResponse.req:type_name -> defn.dev.legumes.v1.Bean
	0, // 4: defn.dev.legumes.v1.DeleteBeanRequest.req:type_name -> defn.dev.legumes.v1.Bean
	0, // 5: defn.dev.legumes.v1.DeleteBeanResponse.req:type_name -> defn.dev.legumes.v1.Bean
	1, // 6: defn.dev.legumes.v1.BeanStoreService.GetBean:input_type -> defn.dev.legumes.v1.GetBeanRequest
	3, // 7: defn.dev.legumes.v1.BeanStoreService.PutBean:input_type -> defn.dev.legumes.v1.PutBeanRequest
	5, // 8: defn.dev.legumes.v1.BeanStoreService.DeleteBean:input_type -> defn.dev.legumes.v1.DeleteBeanRequest
	2, // 9: defn.dev.legumes.v1.BeanStoreService.GetBean:output_type -> defn.dev.legumes.v1.GetBeanResponse
	4, // 10: defn.dev.legumes.v1.BeanStoreService.PutBean:output_type -> defn.dev.legumes.v1.PutBeanResponse
	6, // 11: defn.dev.legumes.v1.BeanStoreService.DeleteBean:output_type -> defn.dev.legumes.v1.DeleteBeanResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_legumes_v1_bean_proto_init() }
func file_legumes_v1_bean_proto_init() {
	if File_legumes_v1_bean_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_legumes_v1_bean_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bean); i {
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
		file_legumes_v1_bean_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeanRequest); i {
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
		file_legumes_v1_bean_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeanResponse); i {
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
		file_legumes_v1_bean_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutBeanRequest); i {
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
		file_legumes_v1_bean_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutBeanResponse); i {
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
		file_legumes_v1_bean_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBeanRequest); i {
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
		file_legumes_v1_bean_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBeanResponse); i {
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
			RawDescriptor: file_legumes_v1_bean_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_legumes_v1_bean_proto_goTypes,
		DependencyIndexes: file_legumes_v1_bean_proto_depIdxs,
		MessageInfos:      file_legumes_v1_bean_proto_msgTypes,
	}.Build()
	File_legumes_v1_bean_proto = out.File
	file_legumes_v1_bean_proto_rawDesc = nil
	file_legumes_v1_bean_proto_goTypes = nil
	file_legumes_v1_bean_proto_depIdxs = nil
}
