// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: TokenServiceServer.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Input         string                 `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	mi := &file_TokenServiceServer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_TokenServiceServer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_TokenServiceServer_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type UsernameResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Output        string                 `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsernameResponse) Reset() {
	*x = UsernameResponse{}
	mi := &file_TokenServiceServer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameResponse) ProtoMessage() {}

func (x *UsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_TokenServiceServer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameResponse.ProtoReflect.Descriptor instead.
func (*UsernameResponse) Descriptor() ([]byte, []int) {
	return file_TokenServiceServer_proto_rawDescGZIP(), []int{1}
}

func (x *UsernameResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_TokenServiceServer_proto protoreflect.FileDescriptor

var file_TokenServiceServer_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x2a,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x59, 0x0a, 0x0c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_TokenServiceServer_proto_rawDescOnce sync.Once
	file_TokenServiceServer_proto_rawDescData []byte
)

func file_TokenServiceServer_proto_rawDescGZIP() []byte {
	file_TokenServiceServer_proto_rawDescOnce.Do(func() {
		file_TokenServiceServer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_TokenServiceServer_proto_rawDesc), len(file_TokenServiceServer_proto_rawDesc)))
	})
	return file_TokenServiceServer_proto_rawDescData
}

var file_TokenServiceServer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TokenServiceServer_proto_goTypes = []any{
	(*TokenRequest)(nil),     // 0: tokenservice.TokenRequest
	(*UsernameResponse)(nil), // 1: tokenservice.UsernameResponse
}
var file_TokenServiceServer_proto_depIdxs = []int32{
	0, // 0: tokenservice.TokenService.GetUsername:input_type -> tokenservice.TokenRequest
	1, // 1: tokenservice.TokenService.GetUsername:output_type -> tokenservice.UsernameResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TokenServiceServer_proto_init() }
func file_TokenServiceServer_proto_init() {
	if File_TokenServiceServer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_TokenServiceServer_proto_rawDesc), len(file_TokenServiceServer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TokenServiceServer_proto_goTypes,
		DependencyIndexes: file_TokenServiceServer_proto_depIdxs,
		MessageInfos:      file_TokenServiceServer_proto_msgTypes,
	}.Build()
	File_TokenServiceServer_proto = out.File
	file_TokenServiceServer_proto_goTypes = nil
	file_TokenServiceServer_proto_depIdxs = nil
}
