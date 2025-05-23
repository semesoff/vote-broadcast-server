// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/websocket.proto

package websocket

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WebSocketService_GetPolls_FullMethodName = "/websocket.WebSocketService/GetPolls"
	WebSocketService_GetVotes_FullMethodName = "/websocket.WebSocketService/GetVotes"
)

// WebSocketServiceClient is the client API for WebSocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebSocketServiceClient interface {
	GetPolls(ctx context.Context, in *PollsRequest, opts ...grpc.CallOption) (*PollsResponse, error)
	GetVotes(ctx context.Context, in *VotesRequest, opts ...grpc.CallOption) (*VotesResponse, error)
}

type webSocketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebSocketServiceClient(cc grpc.ClientConnInterface) WebSocketServiceClient {
	return &webSocketServiceClient{cc}
}

func (c *webSocketServiceClient) GetPolls(ctx context.Context, in *PollsRequest, opts ...grpc.CallOption) (*PollsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PollsResponse)
	err := c.cc.Invoke(ctx, WebSocketService_GetPolls_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webSocketServiceClient) GetVotes(ctx context.Context, in *VotesRequest, opts ...grpc.CallOption) (*VotesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VotesResponse)
	err := c.cc.Invoke(ctx, WebSocketService_GetVotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebSocketServiceServer is the server API for WebSocketService service.
// All implementations must embed UnimplementedWebSocketServiceServer
// for forward compatibility.
type WebSocketServiceServer interface {
	GetPolls(context.Context, *PollsRequest) (*PollsResponse, error)
	GetVotes(context.Context, *VotesRequest) (*VotesResponse, error)
	mustEmbedUnimplementedWebSocketServiceServer()
}

// UnimplementedWebSocketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebSocketServiceServer struct{}

func (UnimplementedWebSocketServiceServer) GetPolls(context.Context, *PollsRequest) (*PollsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolls not implemented")
}
func (UnimplementedWebSocketServiceServer) GetVotes(context.Context, *VotesRequest) (*VotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotes not implemented")
}
func (UnimplementedWebSocketServiceServer) mustEmbedUnimplementedWebSocketServiceServer() {}
func (UnimplementedWebSocketServiceServer) testEmbeddedByValue()                          {}

// UnsafeWebSocketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebSocketServiceServer will
// result in compilation errors.
type UnsafeWebSocketServiceServer interface {
	mustEmbedUnimplementedWebSocketServiceServer()
}

func RegisterWebSocketServiceServer(s grpc.ServiceRegistrar, srv WebSocketServiceServer) {
	// If the following call pancis, it indicates UnimplementedWebSocketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WebSocketService_ServiceDesc, srv)
}

func _WebSocketService_GetPolls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketServiceServer).GetPolls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebSocketService_GetPolls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketServiceServer).GetPolls(ctx, req.(*PollsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebSocketService_GetVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketServiceServer).GetVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebSocketService_GetVotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketServiceServer).GetVotes(ctx, req.(*VotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebSocketService_ServiceDesc is the grpc.ServiceDesc for WebSocketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebSocketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "websocket.WebSocketService",
	HandlerType: (*WebSocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolls",
			Handler:    _WebSocketService_GetPolls_Handler,
		},
		{
			MethodName: "GetVotes",
			Handler:    _WebSocketService_GetVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/websocket.proto",
}
