// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: mysql.proto

package mysqlpb

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
	Mysql_Create_FullMethodName = "/api.v1.examples.Mysql/Create"
	Mysql_Update_FullMethodName = "/api.v1.examples.Mysql/Update"
	Mysql_Get_FullMethodName    = "/api.v1.examples.Mysql/Get"
	Mysql_Search_FullMethodName = "/api.v1.examples.Mysql/Search"
	Mysql_Del_FullMethodName    = "/api.v1.examples.Mysql/Del"
)

// MysqlClient is the client API for Mysql service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MysqlClient interface {
	// 创建
	Create(ctx context.Context, in *CreateOrUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新
	Update(ctx context.Context, in *CreateOrUpdateReq, opts ...grpc.CallOption) (*ExampleInfo, error)
	// 获取详情
	Get(ctx context.Context, in *ReqWithName, opts ...grpc.CallOption) (*ExampleInfo, error)
	// 查询
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*ExampleList, error)
	// 删除
	Del(ctx context.Context, in *ReqWithName, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mysqlClient struct {
	cc grpc.ClientConnInterface
}

func NewMysqlClient(cc grpc.ClientConnInterface) MysqlClient {
	return &mysqlClient{cc}
}

func (c *mysqlClient) Create(ctx context.Context, in *CreateOrUpdateReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Mysql_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlClient) Update(ctx context.Context, in *CreateOrUpdateReq, opts ...grpc.CallOption) (*ExampleInfo, error) {
	out := new(ExampleInfo)
	err := c.cc.Invoke(ctx, Mysql_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlClient) Get(ctx context.Context, in *ReqWithName, opts ...grpc.CallOption) (*ExampleInfo, error) {
	out := new(ExampleInfo)
	err := c.cc.Invoke(ctx, Mysql_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*ExampleList, error) {
	out := new(ExampleList)
	err := c.cc.Invoke(ctx, Mysql_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlClient) Del(ctx context.Context, in *ReqWithName, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Mysql_Del_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MysqlServer is the server API for Mysql service.
// All implementations must embed UnimplementedMysqlServer
// for forward compatibility
type MysqlServer interface {
	// 创建
	Create(context.Context, *CreateOrUpdateReq) (*emptypb.Empty, error)
	// 更新
	Update(context.Context, *CreateOrUpdateReq) (*ExampleInfo, error)
	// 获取详情
	Get(context.Context, *ReqWithName) (*ExampleInfo, error)
	// 查询
	Search(context.Context, *SearchReq) (*ExampleList, error)
	// 删除
	Del(context.Context, *ReqWithName) (*emptypb.Empty, error)
	mustEmbedUnimplementedMysqlServer()
}

// UnimplementedMysqlServer must be embedded to have forward compatible implementations.
type UnimplementedMysqlServer struct {
}

func (UnimplementedMysqlServer) Create(context.Context, *CreateOrUpdateReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMysqlServer) Update(context.Context, *CreateOrUpdateReq) (*ExampleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMysqlServer) Get(context.Context, *ReqWithName) (*ExampleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMysqlServer) Search(context.Context, *SearchReq) (*ExampleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedMysqlServer) Del(context.Context, *ReqWithName) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedMysqlServer) mustEmbedUnimplementedMysqlServer() {}

// UnsafeMysqlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MysqlServer will
// result in compilation errors.
type UnsafeMysqlServer interface {
	mustEmbedUnimplementedMysqlServer()
}

func RegisterMysqlServer(s grpc.ServiceRegistrar, srv MysqlServer) {
	s.RegisterService(&Mysql_ServiceDesc, srv)
}

func _Mysql_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlServer).Create(ctx, req.(*CreateOrUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mysql_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlServer).Update(ctx, req.(*CreateOrUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mysql_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqWithName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlServer).Get(ctx, req.(*ReqWithName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mysql_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mysql_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqWithName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Del_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlServer).Del(ctx, req.(*ReqWithName))
	}
	return interceptor(ctx, in, info, handler)
}

// Mysql_ServiceDesc is the grpc.ServiceDesc for Mysql service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mysql_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.examples.Mysql",
	HandlerType: (*MysqlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Mysql_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Mysql_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Mysql_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Mysql_Search_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Mysql_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mysql.proto",
}
