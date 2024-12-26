// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.27.1
// source: commonpb/commonpb.proto

package commonpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	mi := &file_commonpb_commonpb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonpb_commonpb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_commonpb_commonpb_proto_rawDescGZIP(), []int{0}
}

func (x *HealthResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HealthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_commonpb_commonpb_proto protoreflect.FileDescriptor

var file_commonpb_commonpb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e,
	0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5f,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x55, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x12, 0x0c, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x2b, 0x5a, 0x29, 0x73, 0x65, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x70, 0x6f, 0x6f, 0x70, 0x2f, 0x67,
	0x72, 0x65, 0x65, 0x64, 0x79, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonpb_commonpb_proto_rawDescOnce sync.Once
	file_commonpb_commonpb_proto_rawDescData = file_commonpb_commonpb_proto_rawDesc
)

func file_commonpb_commonpb_proto_rawDescGZIP() []byte {
	file_commonpb_commonpb_proto_rawDescOnce.Do(func() {
		file_commonpb_commonpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonpb_commonpb_proto_rawDescData)
	})
	return file_commonpb_commonpb_proto_rawDescData
}

var file_commonpb_commonpb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commonpb_commonpb_proto_goTypes = []any{
	(*HealthResponse)(nil), // 0: commonpb.HealthResponse
	(*emptypb.Empty)(nil),  // 1: google.protobuf.Empty
}
var file_commonpb_commonpb_proto_depIdxs = []int32{
	1, // 0: commonpb.Health.CheckHealth:input_type -> google.protobuf.Empty
	0, // 1: commonpb.Health.CheckHealth:output_type -> commonpb.HealthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonpb_commonpb_proto_init() }
func file_commonpb_commonpb_proto_init() {
	if File_commonpb_commonpb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonpb_commonpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonpb_commonpb_proto_goTypes,
		DependencyIndexes: file_commonpb_commonpb_proto_depIdxs,
		MessageInfos:      file_commonpb_commonpb_proto_msgTypes,
	}.Build()
	File_commonpb_commonpb_proto = out.File
	file_commonpb_commonpb_proto_rawDesc = nil
	file_commonpb_commonpb_proto_goTypes = nil
	file_commonpb_commonpb_proto_depIdxs = nil
}
