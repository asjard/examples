// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: api/serverpb/server.proto

package serverpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Server_Say_FullMethodName  = "/api.v1.server.Server/Say"
	Server_Log_FullMethodName  = "/api.v1.server.Server/Log"
	Server_Call_FullMethodName = "/api.v1.server.Server/Call"
)

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	// 注释，描述这个接口的作用
	Say(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error)
	// sse请求
	Log(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// grpc请求
	Call(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Say(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := c.cc.Invoke(ctx, Server_Say_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Log(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Server_Log_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Call(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := c.cc.Invoke(ctx, Server_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	// 注释，描述这个接口的作用
	Say(context.Context, *HelloReq) (*HelloReq, error)
	// sse请求
	Log(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// grpc请求
	Call(context.Context, *HelloReq) (*HelloReq, error)
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) Say(context.Context, *HelloReq) (*HelloReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}
func (UnimplementedServerServer) Log(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedServerServer) Call(context.Context, *HelloReq) (*HelloReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_Say_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Say(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_Log_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Log(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Call(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Server_Say_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _Server_Log_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _Server_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/serverpb/server.proto",
}
