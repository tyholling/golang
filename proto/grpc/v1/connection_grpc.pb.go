// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/grpc/v1/connection.proto

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

const (
	ConnectionService_Connect_FullMethodName = "/proto.grpc.v1.ConnectionService/Connect"
)

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (ConnectionService_ConnectClient, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (ConnectionService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectionService_ServiceDesc.Streams[0], ConnectionService_Connect_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionServiceConnectClient{stream}
	return x, nil
}

type ConnectionService_ConnectClient interface {
	Send(*ConnectRequest) error
	Recv() (*ConnectResponse, error)
	grpc.ClientStream
}

type connectionServiceConnectClient struct {
	grpc.ClientStream
}

func (x *connectionServiceConnectClient) Send(m *ConnectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionServiceConnectClient) Recv() (*ConnectResponse, error) {
	m := new(ConnectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	Connect(ConnectionService_ConnectServer) error
	mustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) Connect(ConnectionService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedConnectionServiceServer) mustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	mustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionServiceServer).Connect(&connectionServiceConnectServer{stream})
}

type ConnectionService_ConnectServer interface {
	Send(*ConnectResponse) error
	Recv() (*ConnectRequest, error)
	grpc.ServerStream
}

type connectionServiceConnectServer struct {
	grpc.ServerStream
}

func (x *connectionServiceConnectServer) Send(m *ConnectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionServiceConnectServer) Recv() (*ConnectRequest, error) {
	m := new(ConnectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.grpc.v1.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ConnectionService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/grpc/v1/connection.proto",
}
