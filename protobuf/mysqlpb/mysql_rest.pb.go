// Code generated by protoc-gen-go-rest. DO NOT EDIT.
// versions:
// - protoc-gen-go-rest v1.3.0
// - protoc             v5.27.0
// source: mysql.proto

package mysqlpb

import (
	context "context"
	server "github.com/asjard/asjard/core/server"
	rest "github.com/asjard/asjard/pkg/server/rest"
)

// 创建
func _Mysql_Create_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(CreateOrUpdateReq)
	if interceptor == nil {
		return srv.(MysqlServer).Create(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Create_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MysqlServer).Create(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// 更新
func _Mysql_Update_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(CreateOrUpdateReq)
	if interceptor == nil {
		return srv.(MysqlServer).Update(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Update_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MysqlServer).Update(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// 获取详情
func _Mysql_Get_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(ReqWithName)
	if interceptor == nil {
		return srv.(MysqlServer).Get(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Get_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MysqlServer).Get(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// 查询
func _Mysql_Search_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(SearchReq)
	if interceptor == nil {
		return srv.(MysqlServer).Search(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Search_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MysqlServer).Search(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// 删除
func _Mysql_Del_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(ReqWithName)
	if interceptor == nil {
		return srv.(MysqlServer).Del(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mysql_Del_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(MysqlServer).Del(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// MysqlRestServiceDesc is the rest.ServiceDesc for Mysql service.
// It's only intended for direct use with rest.AddHandler,
// and not to be introspected or modified (even as a copy)
var MysqlRestServiceDesc = rest.ServiceDesc{
	ServiceName: "api.v1.mysql.Mysql",
	HandlerType: (*MysqlServer)(nil),
	OpenAPI:     file_mysql_proto_openapi,
	Methods: []rest.MethodDesc{
		{
			MethodName: "Create",
			Desc:       "创建.",
			Method:     "POST",
			Path:       Mysql_Create_RestPath,
			Handler:    _Mysql_Create_RestHandler,
		},
		{
			MethodName: "Update",
			Desc:       "更新.",
			Method:     "PUT",
			Path:       Mysql_Update_RestPath,
			Handler:    _Mysql_Update_RestHandler,
		},
		{
			MethodName: "Get",
			Desc:       "获取详情.",
			Method:     "GET",
			Path:       Mysql_Get_RestPath,
			Handler:    _Mysql_Get_RestHandler,
		},
		{
			MethodName: "Search",
			Desc:       "查询.",
			Method:     "GET",
			Path:       Mysql_Search_RestPath,
			Handler:    _Mysql_Search_RestHandler,
		},
		{
			MethodName: "Del",
			Desc:       "删除.",
			Method:     "DELETE",
			Path:       Mysql_Del_RestPath,
			Handler:    _Mysql_Del_RestHandler,
		},
	},
}

const (
	Mysql_Create_RestPath = "/api/v1/examples/mysql"
	Mysql_Update_RestPath = "/api/v1/examples/mysql/{name}"
	Mysql_Get_RestPath    = "/api/v1/examples/mysql/{name}"
	Mysql_Search_RestPath = "/api/v1/examples/mysql"
	Mysql_Del_RestPath    = "/api/v1/examples/mysql/mysql/{name}"
)

var file_mysql_proto_openapi = []byte{
	0x0a, 0x05, 0x33, 0x2e, 0x30, 0x2e, 0x33, 0x12, 0x07, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31,
	0x22, 0xac, 0x0b, 0x0a, 0xd2, 0x04, 0x0a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x12, 0xb7,
	0x04, 0x22, 0xc8, 0x02, 0x0a, 0x05, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x1a, 0x06, 0xe6, 0x9f, 0xa5,
	0xe8, 0xaf, 0xa2, 0x2a, 0x1b, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x79, 0x73, 0x71,
	0x6c, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x30,
	0x32, 0x25, 0x0a, 0x23, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a,
	0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x32, 0x25, 0x0a, 0x23, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x32, 0x1c,
	0x0a, 0x1a, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0xad, 0x01, 0x12,
	0x4b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x44, 0x0a, 0x42, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x3c,
	0x0a, 0x3a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x24, 0x12, 0x22, 0x0a, 0x20, 0x23, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a,
	0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xe9, 0x01, 0x0a,
	0x05, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x1a, 0x06, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0x2a, 0x1b,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4d, 0x79, 0x73,
	0x71, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x30, 0x3a, 0x48, 0x0a, 0x46, 0x12,
	0x42, 0x0a, 0x40, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x2a, 0x12, 0x28, 0x0a, 0x26, 0x23, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x18, 0x01, 0x42, 0x71, 0x12, 0x0f, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x08,
	0x0a, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a,
	0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0xe6, 0x01, 0x0a, 0x23, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x12, 0xbe, 0x01, 0x3a, 0xbb, 0x01, 0x0a, 0x05, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x1a, 0x06, 0xe5,
	0x88, 0xa0, 0xe9, 0x99, 0xa4, 0x2a, 0x18, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x5f, 0x30, 0x32,
	0x1d, 0x0a, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x70, 0x61, 0x74, 0x68, 0x20,
	0x01, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x71,
	0x12, 0x0f, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x08, 0x0a, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x1a,
	0x00, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51,
	0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x0a, 0xeb, 0x04, 0x0a, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0xc9, 0x04, 0x22, 0xfe, 0x01, 0x0a, 0x05, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x1a, 0x0c, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe8, 0xaf, 0xa6, 0xe6, 0x83, 0x85, 0x2a, 0x18,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4d, 0x79, 0x73,
	0x71, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x5f, 0x30, 0x32, 0x1d, 0x0a, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x04, 0x70, 0x61, 0x74, 0x68, 0x20, 0x01, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0xad, 0x01, 0x12, 0x4b, 0x0a, 0x03, 0x32, 0x30,
	0x30, 0x12, 0x44, 0x0a, 0x42, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x3c, 0x0a, 0x3a, 0x0a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x24, 0x12, 0x22, 0x0a, 0x20, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37,
	0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xc5, 0x02, 0x0a, 0x05, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0x1a, 0x06, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0x2a, 0x1b, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x30, 0x32, 0x1d, 0x0a, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x04, 0x70, 0x61, 0x74, 0x68, 0x20, 0x01, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x48, 0x0a, 0x46, 0x12, 0x42, 0x0a, 0x40, 0x0a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x2a, 0x12, 0x28, 0x0a, 0x26, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x18, 0x01, 0x42,
	0xad, 0x01, 0x12, 0x4b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x44, 0x0a, 0x42, 0x0a, 0x02, 0x4f,
	0x4b, 0x1a, 0x3c, 0x0a, 0x3a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x24, 0x12, 0x22, 0x0a, 0x20, 0x23,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x2f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f,
	0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a,
	0x95, 0x0c, 0x0a, 0x92, 0x0c, 0x0a, 0x56, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x41, 0x0a, 0x3f, 0xca, 0x01,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x33, 0x0a, 0x13, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a,
	0x1c, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x13, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x9a, 0x02, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0xa8, 0x01,
	0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x98, 0x01,
	0x0a, 0x95, 0x01, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x88, 0x01,
	0x0a, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x15, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a,
	0x13, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x0a, 0x1c, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x13, 0xca,
	0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a, 0x02, 0x06, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x0a, 0x19, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x19, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x0b, 0x0a, 0x09, 0xca,
	0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x79, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6a, 0x0a, 0x68, 0xca, 0x01, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x5c, 0x0a, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a, 0x02,
	0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x31, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0xf2, 0x01, 0x26, 0x0a, 0x24, 0x12,
	0x22, 0x0a, 0x20, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x0a, 0xd7, 0x01, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x41, 0x6e, 0x79, 0x12, 0xc1, 0x01, 0x0a, 0xbe, 0x01, 0xca,
	0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x3c, 0x0a, 0x3a, 0x0a, 0x05, 0x40,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x2f, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x92, 0x02, 0x23, 0x54, 0x68, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x82, 0x02, 0x02, 0x10, 0x01, 0x92, 0x02, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x72, 0x62, 0x69, 0x74,
	0x72, 0x61, 0x72, 0x79, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x61, 0x6c, 0x6f, 0x6e, 0x67, 0x20, 0x77, 0x69,
	0x74, 0x68, 0x20, 0x61, 0x20, 0x40, 0x74, 0x79, 0x70, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x79,
	0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x0a, 0xb7, 0x07,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0xac, 0x07, 0x0a, 0xa9, 0x07, 0xca, 0x01,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0xef, 0x03, 0x0a, 0x74, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x6c, 0x0a, 0x6a, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x92, 0x02, 0x55, 0x54, 0x68, 0x65, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x63,
	0x6f, 0x64, 0x65, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x6e, 0x75, 0x6d, 0x20, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x5d, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x5d, 0x2e, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x0a, 0xf5, 0x01, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xe9, 0x01,
	0x0a, 0xe6, 0x01, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0xd9, 0x01,
	0x41, 0x20, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2d, 0x66, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62,
	0x65, 0x20, 0x69, 0x6e, 0x20, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x20, 0x41, 0x6e,
	0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5d, 0x5b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x5d, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2c, 0x20, 0x6f, 0x72,
	0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x0a, 0x7f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x77, 0x0a, 0x75, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x92, 0x02,
	0x69, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x63, 0x61, 0x72, 0x72, 0x79, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x20, 0x20, 0x54, 0x68, 0x65, 0x72, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x41, 0x50,
	0x49, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65, 0x2e, 0x92, 0x02, 0xa9, 0x03, 0x54, 0x68,
	0x65, 0x20, 0x60, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x60, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x20, 0x74, 0x68,
	0x61, 0x74, 0x20, 0x69, 0x73, 0x20, 0x73, 0x75, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2c, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67,
	0x20, 0x52, 0x45, 0x53, 0x54, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x52,
	0x50, 0x43, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2e, 0x20, 0x49, 0x74, 0x20, 0x69, 0x73, 0x20, 0x75,
	0x73, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x5b, 0x67, 0x52, 0x50, 0x43, 0x5d, 0x28, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x29, 0x2e, 0x20, 0x45, 0x61, 0x63, 0x68, 0x20, 0x60, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x60, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x72, 0x65, 0x65, 0x20, 0x70, 0x69, 0x65,
	0x63, 0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x20, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x20, 0x59, 0x6f, 0x75, 0x20, 0x63,
	0x61, 0x6e, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65,
	0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x6f, 0x77, 0x20,
	0x74, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x69, 0x74, 0x20,
	0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x5b, 0x41, 0x50, 0x49, 0x20, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x20, 0x47, 0x75, 0x69, 0x64, 0x65, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x29, 0x2e,
}
