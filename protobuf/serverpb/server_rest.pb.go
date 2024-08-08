// Code generated by protoc-gen-go-rest. DO NOT EDIT.
// versions:
// - protoc-gen-go-rest v1.3.0
// - protoc             v5.27.0
// source: server.proto

package serverpb

import (
	context "context"
	server "github.com/asjard/asjard/core/server"
	rest "github.com/asjard/asjard/pkg/server/rest"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func _Server_Say_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(HelloReq)
	if interceptor == nil {
		return srv.(ServerServer).Say(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: "api.v1.examples.Server.Say",
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(ServerServer).Say(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// rest请求
func _Server_Hello_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(emptypb.Empty)
	if interceptor == nil {
		return srv.(ServerServer).Hello(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: "api.v1.examples.Server.Hello",
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(ServerServer).Hello(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// sse请求
func _Server_Log_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(emptypb.Empty)
	if interceptor == nil {
		return srv.(ServerServer).Log(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: "api.v1.examples.Server.Log",
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(ServerServer).Log(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// ServerRestServiceDesc is the rest.ServiceDesc for Server service.
// It's only intended for direct use with rest.AddHandler,
// and not to be introspected or modified (even as a copy)
var ServerRestServiceDesc = rest.ServiceDesc{
	ServiceName: "api.v1.examples.Server",
	HandlerType: (*ServerServer)(nil),
	OpenAPI:     file_server_proto_openapi,
	Methods: []rest.MethodDesc{
		{
			MethodName: "Say",
			Desc:       ".",
			Method:     "POST",
			Path:       Server_Say_RestPath,
			Handler:    _Server_Say_RestHandler,
		},
		{
			MethodName: "Say",
			Desc:       ".",
			Method:     "GET",
			Path:       Server_Say_RestPath_1,
			Handler:    _Server_Say_RestHandler,
		},
		{
			MethodName: "Hello",
			Desc:       "rest请求.",
			Method:     "GET",
			Path:       Server_Hello_RestPath,
			Handler:    _Server_Hello_RestHandler,
		},
		{
			MethodName: "Hello",
			Desc:       "rest请求.",
			Method:     "GET",
			Path:       Server_Hello_RestPath_1,
			Handler:    _Server_Hello_RestHandler,
		},
		{
			MethodName: "Log",
			Desc:       "sse请求.",
			Method:     "GET",
			Path:       Server_Log_RestPath,
			Handler:    _Server_Log_RestHandler,
		},
	},
}

const (
	Server_Log_RestPath     = "/api/v1/examples/server/log"
	Server_Say_RestPath     = "/api/v1/examples/server/region/{region_id}/project/{project_id}/user/{user_id}"
	Server_Say_RestPath_1   = "/api/v1/examples/server/region/{region_id}/project/{project_id}/user/{user_id}"
	Server_Hello_RestPath   = "/api/v1/examples/server/hello"
	Server_Hello_RestPath_1 = "/hello"
)

var file_server_proto_openapi = []byte{
	0x0a, 0x05, 0x33, 0x2e, 0x30, 0x2e, 0x33, 0x12, 0x07, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31,
	0x22, 0x81, 0x17, 0x0a, 0x86, 0x02, 0x0a, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0xe4, 0x01, 0x22, 0xe1, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x1a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0xe8, 0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0x2a,
	0x1e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x30, 0x42,
	0xaa, 0x01, 0x12, 0x48, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x41, 0x0a, 0x3f, 0x0a, 0x02, 0x4f,
	0x4b, 0x1a, 0x39, 0x0a, 0x37, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x21, 0x12, 0x1f, 0x0a, 0x1d, 0x23,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x5e, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a,
	0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0xc7, 0x01, 0x0a,
	0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x12, 0xa7, 0x01, 0x22,
	0xa4, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x09, 0x73, 0x73, 0x65, 0xe8,
	0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0x2a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x5f, 0x30, 0x42, 0x71, 0x12, 0x0f, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x08, 0x0a, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a,
	0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0xb9, 0x11, 0x0a, 0x4e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xe6, 0x10, 0x22, 0xb7, 0x0d, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x53, 0x61, 0x79, 0x5f, 0x31, 0x32, 0x2c, 0x0a, 0x2a, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x12, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x08, 0xe5, 0x8c, 0xba, 0xe5,
	0x9f, 0x9f, 0x49, 0x44, 0x20, 0x01, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x2d, 0x0a, 0x2b, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x12, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x08, 0xe9, 0xa1, 0xb9, 0xe7, 0x9b,
	0xae, 0x49, 0x44, 0x20, 0x01, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x32, 0x36, 0x0a, 0x34, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x08, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x20,
	0x01, 0x52, 0x17, 0x0a, 0x15, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02,
	0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x40, 0x0a, 0x3e, 0x0a, 0x08,
	0x73, 0x74, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x0f, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4, 0xb8, 0xb2, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8,
	0x52, 0x1a, 0x0a, 0x18, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0xf2, 0x01, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x49, 0x0a, 0x47,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x0c, 0xe6, 0x95, 0xb0, 0xe5, 0xad, 0x97, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x52,
	0x26, 0x0a, 0x24, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0xf2, 0x01, 0x19, 0x0a, 0x17,
	0x0a, 0x15, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x2e, 0x0a, 0x2c, 0x0a, 0x0d, 0x6f, 0x62, 0x6a,
	0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a,
	0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x32, 0x25, 0x0a, 0x23, 0x0a, 0x0d, 0x6f, 0x62, 0x6a,
	0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x27,
	0x0a, 0x25, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x4a, 0x0a, 0x48, 0x0a, 0x32, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x69,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x52, 0x0a, 0x50, 0x0a, 0x3a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x2e, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69,
	0x6e, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x2d, 0x0a, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x06, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5,
	0x52, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a, 0x02,
	0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x32, 0x33, 0x0a, 0x31, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0c, 0xe6, 0xaf, 0x8f, 0xe9, 0xa1, 0xb5, 0xe5,
	0xa4, 0xa7, 0xe5, 0xb0, 0x8f, 0x52, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x32, 0x24, 0x0a, 0x22, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x06, 0xe6, 0x8e,
	0x92, 0xe5, 0xba, 0x8f, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x32, 0x29, 0x0a, 0x27, 0x0a, 0x02, 0x6f, 0x6b, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x0c, 0xe5, 0xb8, 0x83, 0xe5, 0xb0, 0x94, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0x52, 0x0c,
	0x0a, 0x0a, 0xca, 0x01, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x32, 0x47, 0x0a, 0x45,
	0x0a, 0x12, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x12, 0xe5, 0x8f, 0xaf,
	0xe9, 0x80, 0x89, 0xe6, 0x95, 0xb4, 0xe5, 0xbd, 0xa2, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x52,
	0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x9a, 0x02, 0x05,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x32, 0x44, 0x0a, 0x42, 0x0a, 0x15, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x15, 0xe5, 0x8f, 0xaf, 0xe9, 0x80, 0x89, 0xe5,
	0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4, 0xb8, 0xb2, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x52, 0x0b,
	0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x47, 0x0a, 0x45, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x12, 0xe5, 0x8f,
	0xaf, 0xe9, 0x80, 0x89, 0xe6, 0x9e, 0x9a, 0xe4, 0xb8, 0xbe, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0,
	0x52, 0x22, 0x0a, 0x20, 0xc2, 0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f, 0x41, 0xc2, 0x01, 0x05, 0x12,
	0x03, 0x4b, 0x5f, 0x42, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x9a, 0x02, 0x04,
	0x65, 0x6e, 0x75, 0x6d, 0x32, 0x51, 0x0a, 0x4f, 0x0a, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x12,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0c, 0xe6, 0x9e, 0x9a, 0xe4, 0xb8, 0xbe, 0xe5, 0x88,
	0x97, 0xe8, 0xa1, 0xa8, 0x52, 0x31, 0x0a, 0x2f, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79,
	0xf2, 0x01, 0x24, 0x0a, 0x22, 0x0a, 0x20, 0xc2, 0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f, 0x41, 0xc2,
	0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f, 0x42, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x9a, 0x02, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x32, 0x2b, 0x0a, 0x29, 0x0a, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x13, 0x0a, 0x11, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x9a, 0x02, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x32, 0x5e, 0x0a, 0x5c, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x2c, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x20, 0xe4, 0xbc, 0x9a, 0xe6, 0x8a, 0x8a, 0xe8, 0xbf, 0x99,
	0xe4, 0xb8, 0xaa, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe8, 0xa7, 0xa3, 0xe6, 0x9e, 0x90, 0xe4,
	0xb8, 0xba, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4, 0xb8, 0xb2, 0x52, 0x17, 0x0a, 0x15, 0xca,
	0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x32, 0x36, 0x0a, 0x34, 0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x17, 0x0a, 0x15, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x1b, 0x0a, 0x19,
	0x0a, 0x03, 0x61, 0x70, 0x70, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x1e, 0x0a, 0x1c, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x1a, 0x0a, 0x18, 0x0a, 0x02, 0x61,
	0x7a, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x23, 0x0a, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a,
	0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x25, 0x0a, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x32, 0x2c, 0x0a, 0x2a, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32,
	0x28, 0x0a, 0x26, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0xaa, 0x01, 0x12, 0x48, 0x0a, 0x03,
	0x32, 0x30, 0x30, 0x12, 0x41, 0x0a, 0x3f, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x39, 0x0a, 0x37, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x21, 0x12, 0x1f, 0x0a, 0x1d, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a,
	0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa9, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x79, 0x5f, 0x30, 0x32,
	0x2c, 0x0a, 0x2a, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x12, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x1a, 0x08, 0xe5, 0x8c, 0xba, 0xe5, 0x9f, 0x9f, 0x49, 0x44, 0x20, 0x01,
	0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x2d, 0x0a,
	0x2b, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x1a, 0x08, 0xe9, 0xa1, 0xb9, 0xe7, 0x9b, 0xae, 0x49, 0x44, 0x20, 0x01, 0x52,
	0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x36, 0x0a, 0x34,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a,
	0x08, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x20, 0x01, 0x52, 0x17, 0x0a, 0x15, 0xca,
	0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x3a, 0x3f, 0x0a, 0x3d, 0x12, 0x39, 0x0a, 0x37, 0x0a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x21, 0x12, 0x1f, 0x0a, 0x1d, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x18, 0x01, 0x42, 0xaa, 0x01, 0x12, 0x48, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12,
	0x41, 0x0a, 0x3f, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x39, 0x0a, 0x37, 0x0a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x21, 0x12, 0x1f, 0x0a, 0x1d, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a,
	0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x0a, 0xef, 0x01, 0x0a, 0x06, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0xe4, 0x01,
	0x22, 0xe1, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x0a, 0x72, 0x65, 0x73,
	0x74, 0xe8, 0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0x2a, 0x1e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x31, 0x42, 0xaa, 0x01, 0x12, 0x48, 0x0a, 0x03, 0x32, 0x30,
	0x30, 0x12, 0x41, 0x0a, 0x3f, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x39, 0x0a, 0x37, 0x0a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x21, 0x12, 0x1f, 0x0a, 0x1d, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x85, 0x16, 0x0a, 0x82, 0x16, 0x0a, 0xd7, 0x01, 0x0a, 0x11, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x41, 0x6e, 0x79,
	0x12, 0xc1, 0x01, 0x0a, 0xbe, 0x01, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa,
	0x01, 0x3c, 0x0a, 0x3a, 0x0a, 0x05, 0x40, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x2f, 0xca,
	0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x23, 0x54, 0x68, 0x65, 0x20, 0x74,
	0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x82, 0x02,
	0x02, 0x10, 0x01, 0x92, 0x02, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x61,
	0x6e, 0x20, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x20, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x61,
	0x6c, 0x6f, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x40, 0x74, 0x79, 0x70,
	0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x0a, 0xa2, 0x09, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x12, 0x95, 0x09, 0x0a, 0x92, 0x09, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0xfa, 0x01, 0x85, 0x09, 0x0a, 0x23, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x14, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02,
	0x08, 0xe5, 0x8c, 0xba, 0xe5, 0x9f, 0x9f, 0x49, 0x44, 0x0a, 0x24, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x14, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x08, 0xe9, 0xa1, 0xb9, 0xe7, 0x9b, 0xae, 0x49, 0x44, 0x0a,
	0x2d, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x20, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x92, 0x02, 0x08, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x49, 0x44, 0x0a, 0x38,
	0x0a, 0x08, 0x73, 0x74, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x2a, 0xca, 0x01,
	0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0xf2, 0x01, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x0f, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4,
	0xb8, 0xb2, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x0a, 0x41, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x33, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79,
	0xf2, 0x01, 0x19, 0x0a, 0x17, 0x0a, 0x15, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x92, 0x02, 0x0c, 0xe6,
	0x95, 0xb0, 0xe5, 0xad, 0x97, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x0a, 0x3a, 0x0a, 0x03, 0x6f,
	0x62, 0x6a, 0x12, 0x33, 0x0a, 0x31, 0xd2, 0x01, 0x25, 0x12, 0x23, 0x0a, 0x21, 0x23, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x5f, 0x4f, 0x62, 0x6a, 0x92, 0x02,
	0x06, 0xe5, 0xaf, 0xb9, 0xe8, 0xb1, 0xa1, 0x0a, 0x4b, 0x0a, 0x04, 0x6f, 0x62, 0x6a, 0x73, 0x12,
	0x43, 0x0a, 0x41, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0xf2, 0x01, 0x27, 0x0a, 0x25,
	0x12, 0x23, 0x0a, 0x21, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x5f, 0x4f, 0x62, 0x6a, 0x92, 0x02, 0x0c, 0xe5, 0xaf, 0xb9, 0xe8, 0xb1, 0xa1, 0xe5, 0x88,
	0x97, 0xe8, 0xa1, 0xa8, 0x0a, 0x42, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x37, 0x0a, 0x35, 0xd2, 0x01, 0x29, 0x12, 0x27, 0x0a, 0x25, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x5f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x92,
	0x02, 0x06, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0x0a, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x1b, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x92, 0x02,
	0x06, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a,
	0x2b, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x21, 0xca, 0x01, 0x07, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x92, 0x02, 0x0c, 0xe6, 0xaf, 0x8f, 0xe9, 0xa1, 0xb5, 0xe5, 0xa4,
	0xa7, 0xe5, 0xb0, 0x8f, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0x1c, 0x0a, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x92, 0x02, 0x06, 0xe6, 0x8e, 0x92, 0xe5, 0xba, 0x8f, 0x0a, 0x21, 0x0a, 0x02, 0x6f, 0x6b,
	0x12, 0x1b, 0x0a, 0x19, 0xca, 0x01, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x92, 0x02,
	0x0c, 0xe5, 0xb8, 0x83, 0xe5, 0xb0, 0x94, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0x0a, 0x3f, 0x0a,
	0x12, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x27, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x92, 0x02, 0x12, 0xe5, 0x8f, 0xaf, 0xe9, 0x80, 0x89, 0xe6, 0x95, 0xb4, 0xe5, 0xbd, 0xa2,
	0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0x3c,
	0x0a, 0x15, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x21, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x15, 0xe5, 0x8f, 0xaf, 0xe9, 0x80, 0x89, 0xe5, 0xad, 0x97,
	0xe7, 0xac, 0xa6, 0xe4, 0xb8, 0xb2, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x0a, 0x3f, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x35, 0xc2, 0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f, 0x41,
	0xc2, 0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f, 0x42, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x92, 0x02, 0x12, 0xe5, 0x8f, 0xaf, 0xe9, 0x80, 0x89, 0xe6, 0x9e, 0x9a, 0xe4, 0xb8, 0xbe,
	0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x9a, 0x02, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x0a, 0x49, 0x0a,
	0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x3e, 0xca, 0x01, 0x05, 0x61, 0x72, 0x72,
	0x61, 0x79, 0xf2, 0x01, 0x24, 0x0a, 0x22, 0x0a, 0x20, 0xc2, 0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f,
	0x41, 0xc2, 0x01, 0x05, 0x12, 0x03, 0x4b, 0x5f, 0x42, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x9a, 0x02, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x92, 0x02, 0x0c, 0xe6, 0x9e, 0x9a, 0xe4,
	0xb8, 0xbe, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x0a, 0x22, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x13, 0x0a, 0x11, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x9a, 0x02, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x0a, 0x56, 0x0a, 0x0c,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x44,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x92, 0x02, 0x2c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x20,
	0xe4, 0xbc, 0x9a, 0xe6, 0x8a, 0x8a, 0xe8, 0xbf, 0x99, 0xe4, 0xb8, 0xaa, 0xe5, 0xad, 0x97, 0xe6,
	0xae, 0xb5, 0xe8, 0xa7, 0xa3, 0xe6, 0x9e, 0x90, 0xe4, 0xb8, 0xba, 0xe5, 0xad, 0x97, 0xe7, 0xac,
	0xa6, 0xe4, 0xb8, 0xb2, 0x0a, 0x2d, 0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x15, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x0a, 0x12, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x11,
	0x0a, 0x02, 0x61, 0x7a, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x0a, 0x36, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x12,
	0x28, 0x0a, 0x26, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x5f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x0a, 0xbb, 0x01, 0x0a, 0x10, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x5f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0xa6,
	0x01, 0x0a, 0xa3, 0x01, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x96,
	0x01, 0x0a, 0x16, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x0b, 0x0a, 0x09,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x39, 0x0a, 0x2a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x0a, 0x41, 0x0a, 0x32, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x61,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0xab, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x5f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x95, 0x01,
	0x0a, 0x92, 0x01, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x85, 0x01,
	0x0a, 0x11, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x0a, 0x13, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x09, 0xca,
	0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x1a, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x0a, 0x16, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x27, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x19, 0xca, 0x01, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x82, 0x02, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x0a, 0x5b, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x5f, 0x4f, 0x62, 0x6a, 0x12, 0x4b, 0x0a, 0x49, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0xfa, 0x01, 0x3d, 0x0a, 0x21, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x12, 0xca, 0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0x18, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x73, 0x74, 0x72, 0x12, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x0a, 0xb7, 0x07, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0xac, 0x07,
	0x0a, 0xa9, 0x07, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0xef, 0x03,
	0x0a, 0x74, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x6c, 0x0a, 0x6a, 0xca, 0x01, 0x07, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x92, 0x02, 0x55, 0x54, 0x68, 0x65, 0x20, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20,
	0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x6e, 0x75,
	0x6d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x5b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x5d, 0x5b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x5d, 0x2e, 0x9a, 0x02,
	0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0xf5, 0x01, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0xe9, 0x01, 0x0a, 0xe6, 0x01, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x92, 0x02, 0xd9, 0x01, 0x41, 0x20, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72,
	0x2d, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x73, 0x68, 0x6f,
	0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x68, 0x2e, 0x20, 0x41, 0x6e, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x66, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5d,
	0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5d, 0x20, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2c, 0x20, 0x6f, 0x72, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x0a, 0x7f,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x77, 0x0a, 0x75, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x92, 0x02, 0x69, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x63, 0x61,
	0x72, 0x72, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x20, 0x20, 0x54, 0x68, 0x65, 0x72, 0x65, 0x20, 0x69, 0x73,
	0x20, 0x61, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65, 0x2e, 0x92,
	0x02, 0xa9, 0x03, 0x54, 0x68, 0x65, 0x20, 0x60, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x60, 0x20,
	0x74, 0x79, 0x70, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x69, 0x73, 0x20, 0x73, 0x75, 0x69, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x74, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2c, 0x20, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x52, 0x45, 0x53, 0x54, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x52, 0x50, 0x43, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2e, 0x20, 0x49, 0x74,
	0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x5b, 0x67, 0x52, 0x50,
	0x43, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x29, 0x2e, 0x20, 0x45, 0x61, 0x63,
	0x68, 0x20, 0x60, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x60, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x72, 0x65,
	0x65, 0x20, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x61, 0x74, 0x61,
	0x3a, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2c, 0x20, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2c, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x20,
	0x59, 0x6f, 0x75, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74,
	0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x68, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x69, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x5b, 0x41, 0x50, 0x49,
	0x20, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x20, 0x47, 0x75, 0x69, 0x64, 0x65, 0x5d, 0x28, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x29, 0x2e,
}
