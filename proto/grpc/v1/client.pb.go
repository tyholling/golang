// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/grpc/v1/client.proto

package grpc

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

type SubscriptionType int32

const (
	SubscriptionType_HEARTBEAT SubscriptionType = 0
	SubscriptionType_METRICS   SubscriptionType = 1
)

// Enum value maps for SubscriptionType.
var (
	SubscriptionType_name = map[int32]string{
		0: "HEARTBEAT",
		1: "METRICS",
	}
	SubscriptionType_value = map[string]int32{
		"HEARTBEAT": 0,
		"METRICS":   1,
	}
)

func (x SubscriptionType) Enum() *SubscriptionType {
	p := new(SubscriptionType)
	*p = x
	return p
}

func (x SubscriptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_grpc_v1_client_proto_enumTypes[0].Descriptor()
}

func (SubscriptionType) Type() protoreflect.EnumType {
	return &file_proto_grpc_v1_client_proto_enumTypes[0]
}

func (x SubscriptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriptionType.Descriptor instead.
func (SubscriptionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_grpc_v1_client_proto_rawDescGZIP(), []int{0}
}

type Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type SubscriptionType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.grpc.v1.SubscriptionType" json:"type,omitempty"`
}

func (x *Subscribe) Reset() {
	*x = Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_v1_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscribe) ProtoMessage() {}

func (x *Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_v1_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscribe.ProtoReflect.Descriptor instead.
func (*Subscribe) Descriptor() ([]byte, []int) {
	return file_proto_grpc_v1_client_proto_rawDescGZIP(), []int{0}
}

func (x *Subscribe) GetType() SubscriptionType {
	if x != nil {
		return x.Type
	}
	return SubscriptionType_HEARTBEAT
}

type Unsubscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type SubscriptionType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.grpc.v1.SubscriptionType" json:"type,omitempty"`
}

func (x *Unsubscribe) Reset() {
	*x = Unsubscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_v1_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unsubscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unsubscribe) ProtoMessage() {}

func (x *Unsubscribe) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_v1_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unsubscribe.ProtoReflect.Descriptor instead.
func (*Unsubscribe) Descriptor() ([]byte, []int) {
	return file_proto_grpc_v1_client_proto_rawDescGZIP(), []int{1}
}

func (x *Unsubscribe) GetType() SubscriptionType {
	if x != nil {
		return x.Type
	}
	return SubscriptionType_HEARTBEAT
}

var File_proto_grpc_v1_client_proto protoreflect.FileDescriptor

var file_proto_grpc_v1_client_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x22, 0x40, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x42, 0x0a,
	0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x2a, 0x2e, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45,
	0x41, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10,
	0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x79, 0x68, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_v1_client_proto_rawDescOnce sync.Once
	file_proto_grpc_v1_client_proto_rawDescData = file_proto_grpc_v1_client_proto_rawDesc
)

func file_proto_grpc_v1_client_proto_rawDescGZIP() []byte {
	file_proto_grpc_v1_client_proto_rawDescOnce.Do(func() {
		file_proto_grpc_v1_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_v1_client_proto_rawDescData)
	})
	return file_proto_grpc_v1_client_proto_rawDescData
}

var file_proto_grpc_v1_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_grpc_v1_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_grpc_v1_client_proto_goTypes = []interface{}{
	(SubscriptionType)(0), // 0: proto.grpc.v1.SubscriptionType
	(*Subscribe)(nil),     // 1: proto.grpc.v1.Subscribe
	(*Unsubscribe)(nil),   // 2: proto.grpc.v1.Unsubscribe
}
var file_proto_grpc_v1_client_proto_depIdxs = []int32{
	0, // 0: proto.grpc.v1.Subscribe.type:type_name -> proto.grpc.v1.SubscriptionType
	0, // 1: proto.grpc.v1.Unsubscribe.type:type_name -> proto.grpc.v1.SubscriptionType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_grpc_v1_client_proto_init() }
func file_proto_grpc_v1_client_proto_init() {
	if File_proto_grpc_v1_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_v1_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscribe); i {
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
		file_proto_grpc_v1_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unsubscribe); i {
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
			RawDescriptor: file_proto_grpc_v1_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_grpc_v1_client_proto_goTypes,
		DependencyIndexes: file_proto_grpc_v1_client_proto_depIdxs,
		EnumInfos:         file_proto_grpc_v1_client_proto_enumTypes,
		MessageInfos:      file_proto_grpc_v1_client_proto_msgTypes,
	}.Build()
	File_proto_grpc_v1_client_proto = out.File
	file_proto_grpc_v1_client_proto_rawDesc = nil
	file_proto_grpc_v1_client_proto_goTypes = nil
	file_proto_grpc_v1_client_proto_depIdxs = nil
}
