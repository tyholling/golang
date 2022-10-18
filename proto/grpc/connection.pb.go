// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: connection.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RequestResponse:
	//
	//	*Message_Request
	//	*Message_Response
	RequestResponse isMessage_RequestResponse `protobuf_oneof:"request_response"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_connection_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetRequestResponse() isMessage_RequestResponse {
	if m != nil {
		return m.RequestResponse
	}
	return nil
}

func (x *Message) GetRequest() *anypb.Any {
	if x, ok := x.GetRequestResponse().(*Message_Request); ok {
		return x.Request
	}
	return nil
}

func (x *Message) GetResponse() *anypb.Any {
	if x, ok := x.GetRequestResponse().(*Message_Response); ok {
		return x.Response
	}
	return nil
}

type isMessage_RequestResponse interface {
	isMessage_RequestResponse()
}

type Message_Request struct {
	Request *anypb.Any `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type Message_Response struct {
	Response *anypb.Any `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

func (*Message_Request) isMessage_RequestResponse() {}

func (*Message_Response) isMessage_RequestResponse() {}

var File_connection_proto protoreflect.FileDescriptor

var file_connection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x30, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x39, 0x0a, 0x0a, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x68, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connection_proto_rawDescOnce sync.Once
	file_connection_proto_rawDescData = file_connection_proto_rawDesc
)

func file_connection_proto_rawDescGZIP() []byte {
	file_connection_proto_rawDescOnce.Do(func() {
		file_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_connection_proto_rawDescData)
	})
	return file_connection_proto_rawDescData
}

var file_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_connection_proto_goTypes = []interface{}{
	(*Message)(nil),   // 0: grpc.Message
	(*anypb.Any)(nil), // 1: google.protobuf.Any
}
var file_connection_proto_depIdxs = []int32{
	1, // 0: grpc.Message.request:type_name -> google.protobuf.Any
	1, // 1: grpc.Message.response:type_name -> google.protobuf.Any
	0, // 2: grpc.Connection.Connect:input_type -> grpc.Message
	0, // 3: grpc.Connection.Connect:output_type -> grpc.Message
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_connection_proto_init() }
func file_connection_proto_init() {
	if File_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_connection_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Request)(nil),
		(*Message_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connection_proto_goTypes,
		DependencyIndexes: file_connection_proto_depIdxs,
		MessageInfos:      file_connection_proto_msgTypes,
	}.Build()
	File_connection_proto = out.File
	file_connection_proto_rawDesc = nil
	file_connection_proto_goTypes = nil
	file_connection_proto_depIdxs = nil
}
