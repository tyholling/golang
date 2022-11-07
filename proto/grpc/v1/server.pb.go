// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/grpc/v1/server.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_v1_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_proto_grpc_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *Heartbeat) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu           float64 `protobuf:"fixed64,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory        float64 `protobuf:"fixed64,2,opt,name=memory,proto3" json:"memory,omitempty"`
	BytesSent     uint64  `protobuf:"varint,3,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	BytesReceived uint64  `protobuf:"varint,4,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	ErrorsIn      uint64  `protobuf:"varint,5,opt,name=errors_in,json=errorsIn,proto3" json:"errors_in,omitempty"`
	ErrorsOut     uint64  `protobuf:"varint,6,opt,name=errors_out,json=errorsOut,proto3" json:"errors_out,omitempty"`
	DiscardsIn    uint64  `protobuf:"varint,7,opt,name=discards_in,json=discardsIn,proto3" json:"discards_in,omitempty"`
	DiscardsOut   uint64  `protobuf:"varint,8,opt,name=discards_out,json=discardsOut,proto3" json:"discards_out,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_proto_grpc_v1_server_proto_rawDescGZIP(), []int{1}
}

func (x *Metrics) GetCpu() float64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Metrics) GetMemory() float64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Metrics) GetBytesSent() uint64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *Metrics) GetBytesReceived() uint64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

func (x *Metrics) GetErrorsIn() uint64 {
	if x != nil {
		return x.ErrorsIn
	}
	return 0
}

func (x *Metrics) GetErrorsOut() uint64 {
	if x != nil {
		return x.ErrorsOut
	}
	return 0
}

func (x *Metrics) GetDiscardsIn() uint64 {
	if x != nil {
		return x.DiscardsIn
	}
	return 0
}

func (x *Metrics) GetDiscardsOut() uint64 {
	if x != nil {
		return x.DiscardsOut
	}
	return 0
}

var File_proto_grpc_v1_server_proto protoreflect.FileDescriptor

var file_proto_grpc_v1_server_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x09,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0xf9, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x63, 0x70,
	0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79,
	0x68, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_grpc_v1_server_proto_rawDescOnce sync.Once
	file_proto_grpc_v1_server_proto_rawDescData = file_proto_grpc_v1_server_proto_rawDesc
)

func file_proto_grpc_v1_server_proto_rawDescGZIP() []byte {
	file_proto_grpc_v1_server_proto_rawDescOnce.Do(func() {
		file_proto_grpc_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_v1_server_proto_rawDescData)
	})
	return file_proto_grpc_v1_server_proto_rawDescData
}

var file_proto_grpc_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_grpc_v1_server_proto_goTypes = []interface{}{
	(*Heartbeat)(nil),             // 0: proto.grpc.v1.Heartbeat
	(*Metrics)(nil),               // 1: proto.grpc.v1.Metrics
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_grpc_v1_server_proto_depIdxs = []int32{
	2, // 0: proto.grpc.v1.Heartbeat.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_grpc_v1_server_proto_init() }
func file_proto_grpc_v1_server_proto_init() {
	if File_proto_grpc_v1_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
		file_proto_grpc_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
			RawDescriptor: file_proto_grpc_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_grpc_v1_server_proto_goTypes,
		DependencyIndexes: file_proto_grpc_v1_server_proto_depIdxs,
		MessageInfos:      file_proto_grpc_v1_server_proto_msgTypes,
	}.Build()
	File_proto_grpc_v1_server_proto = out.File
	file_proto_grpc_v1_server_proto_rawDesc = nil
	file_proto_grpc_v1_server_proto_goTypes = nil
	file_proto_grpc_v1_server_proto_depIdxs = nil
}
