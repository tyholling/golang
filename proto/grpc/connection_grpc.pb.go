// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: connection.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnectionClient is the client API for Connection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Connection_ConnectClient, error)
}

type connectionClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionClient(cc grpc.ClientConnInterface) ConnectionClient {
	return &connectionClient{cc}
}

func (c *connectionClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Connection_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Connection_ServiceDesc.Streams[0], "/grpc.Connection/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionConnectClient{stream}
	return x, nil
}

type Connection_ConnectClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type connectionConnectClient struct {
	grpc.ClientStream
}

func (x *connectionConnectClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionServer is the server API for Connection service.
// All implementations must embed UnimplementedConnectionServer
// for forward compatibility
type ConnectionServer interface {
	Connect(Connection_ConnectServer) error
	mustEmbedUnimplementedConnectionServer()
}

// UnimplementedConnectionServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServer struct {
}

func (UnimplementedConnectionServer) Connect(Connection_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedConnectionServer) mustEmbedUnimplementedConnectionServer() {}

// UnsafeConnectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServer will
// result in compilation errors.
type UnsafeConnectionServer interface {
	mustEmbedUnimplementedConnectionServer()
}

func RegisterConnectionServer(s grpc.ServiceRegistrar, srv ConnectionServer) {
	s.RegisterService(&Connection_ServiceDesc, srv)
}

func _Connection_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionServer).Connect(&connectionConnectServer{stream})
}

type Connection_ConnectServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type connectionConnectServer struct {
	grpc.ServerStream
}

func (x *connectionConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionConnectServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Connection_ServiceDesc is the grpc.ServiceDesc for Connection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Connection",
	HandlerType: (*ConnectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Connection_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "connection.proto",
}