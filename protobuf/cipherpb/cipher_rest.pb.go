// Code generated by protoc-gen-go-rest. DO NOT EDIT.
// versions:
// - protoc-gen-go-rest v1.3.0
// - protoc             v5.27.0
// source: cipher.proto

package cipherpb

import (
	context "context"
	server "github.com/asjard/asjard/core/server"
	rest "github.com/asjard/asjard/pkg/server/rest"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// 加密
func _Cipher_Encrypt_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(EncryptReq)
	if interceptor == nil {
		return srv.(CipherServer).Encrypt(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cipher_Encrypt_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(CipherServer).Encrypt(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// 解密
func _Cipher_Decrypt_RestHandler(ctx *rest.Context, srv any, interceptor server.UnaryServerInterceptor) (any, error) {
	in := new(emptypb.Empty)
	if interceptor == nil {
		return srv.(CipherServer).Decrypt(ctx, in)
	}
	info := &server.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cipher_Decrypt_FullMethodName,
		Protocol:   rest.Protocol,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(CipherServer).Decrypt(ctx, in)
	}
	return interceptor(ctx, in, info, handler)
}

// CipherRestServiceDesc is the rest.ServiceDesc for Cipher service.
// It's only intended for direct use with rest.AddHandler,
// and not to be introspected or modified (even as a copy)
//
// Cipher 加解密相关接口
var CipherRestServiceDesc = rest.ServiceDesc{
	ServiceName: "api.v1.cipher.Cipher",
	HandlerType: (*CipherServer)(nil),
	OpenAPI:     file_cipher_proto_openapi,
	Methods: []rest.MethodDesc{
		{
			MethodName: "Encrypt",
			Desc:       "加密.",
			Method:     "GET",
			Path:       Cipher_Encrypt_RestPath,
			Handler:    _Cipher_Encrypt_RestHandler,
		},
		{
			MethodName: "Decrypt",
			Desc:       "解密.",
			Method:     "GET",
			Path:       Cipher_Decrypt_RestPath,
			Handler:    _Cipher_Decrypt_RestHandler,
		},
	},
}

const (
	Cipher_Encrypt_RestPath = "/api/v1/examples/cipher/encrypt"
	Cipher_Decrypt_RestPath = "/api/v1/examples/cipher/decrypt"
)

var file_cipher_proto_openapi = []byte{
	0x0a, 0x05, 0x33, 0x2e, 0x30, 0x2e, 0x33, 0x12, 0x07, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31,
	0x22, 0xe5, 0x04, 0x0a, 0x87, 0x02, 0x0a, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x2f,
	0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0xe3, 0x01, 0x22, 0xe0, 0x01, 0x0a, 0x06, 0x43,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x1a, 0x06, 0xe8, 0xa7, 0xa3, 0xe5, 0xaf, 0x86, 0x2a, 0x1e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x30, 0x42, 0xad, 0x01,
	0x12, 0x4b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x44, 0x0a, 0x42, 0x0a, 0x02, 0x4f, 0x4b, 0x1a,
	0x3c, 0x0a, 0x3a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x24, 0x12, 0x22, 0x0a, 0x20, 0x23, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5e, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d,
	0x0a, 0x1b, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0xd8, 0x02,
	0x0a, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x12, 0xb4, 0x02, 0x22, 0xb1, 0x02, 0x0a, 0x06, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x1a,
	0x06, 0xe5, 0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0x2a, 0x1e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x30, 0x32, 0x2a, 0x0a, 0x28, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x06,
	0xe6, 0x98, 0x8e, 0xe6, 0x96, 0x87, 0x52, 0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x23, 0x0a, 0x21, 0x0a, 0x0b, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x0a, 0x09, 0xca,
	0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0xad, 0x01, 0x12, 0x4b, 0x0a, 0x03, 0x32,
	0x30, 0x30, 0x12, 0x44, 0x0a, 0x42, 0x0a, 0x02, 0x4f, 0x4b, 0x1a, 0x3c, 0x0a, 0x3a, 0x0a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x24, 0x12, 0x22, 0x0a, 0x20, 0x23, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x51, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a,
	0x37, 0x0a, 0x35, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1f, 0x12, 0x1d, 0x0a, 0x1b, 0x23, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xf4, 0x0d, 0x0a, 0xf1, 0x0d, 0x0a, 0xeb,
	0x03, 0x0a, 0x0b, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0xdb,
	0x03, 0x0a, 0xd8, 0x03, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0xb3,
	0x03, 0x0a, 0x4f, 0x0a, 0x1f, 0x61, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x2a, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x92, 0x02, 0x1e, 0xe6, 0x98, 0x8e, 0xe6, 0x96, 0x87, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6,
	0xe4, 0xb8, 0xad, 0xe7, 0x9a, 0x84, 0x61, 0x65, 0x73, 0xe5, 0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0xe5,
	0x80, 0xbc, 0x0a, 0x55, 0x0a, 0x22, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x2d, 0xca, 0x01, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x21, 0xe6, 0x98, 0x8e, 0xe6, 0x96, 0x87, 0xe6, 0x96,
	0x87, 0xe4, 0xbb, 0xb6, 0xe4, 0xb8, 0xad, 0xe7, 0x9a, 0x84, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34,
	0xe5, 0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0xe5, 0x80, 0xbc, 0x0a, 0x4f, 0x0a, 0x1f, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x61, 0x65, 0x73, 0x5f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x2a,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x1e, 0x61, 0x65, 0x73, 0xe5,
	0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe4, 0xb8, 0xad, 0xe7, 0x9a,
	0x84, 0xe6, 0x98, 0x8e, 0xe6, 0x96, 0x87, 0xe5, 0x80, 0xbc, 0x0a, 0x58, 0x0a, 0x25, 0x61, 0x65,
	0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x69, 0x6e, 0x5f, 0x61, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x2d, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x92, 0x02, 0x21, 0x61, 0x65, 0x73, 0xe5, 0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0xe6, 0x96, 0x87, 0xe4,
	0xbb, 0xb6, 0xe4, 0xb8, 0xad, 0xe7, 0x9a, 0x84, 0x61, 0x65, 0x73, 0xe5, 0x8a, 0xa0, 0xe5, 0xaf,
	0x86, 0xe5, 0x80, 0xbc, 0x0a, 0x5e, 0x0a, 0x28, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x5f,
	0x61, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x32, 0x0a, 0x30, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x24,
	0x61, 0x65, 0x73, 0xe5, 0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe4,
	0xb8, 0xad, 0xe7, 0x9a, 0x84, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0xe5, 0x8a, 0xa0, 0xe5, 0xaf,
	0x86, 0xe5, 0x80, 0xbc, 0x92, 0x02, 0x15, 0xe5, 0x8a, 0xa0, 0xe8, 0xa7, 0xa3, 0xe5, 0xaf, 0x86,
	0xe7, 0xa4, 0xba, 0xe4, 0xbe, 0x8b, 0xe8, 0xbf, 0x94, 0xe5, 0x9b, 0x9e, 0x0a, 0x6d, 0x0a, 0x0b,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5e, 0x0a, 0x5c, 0xca,
	0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa, 0x01, 0x41, 0x0a, 0x23, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x12, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x06, 0xe5, 0xaf, 0x86, 0xe6, 0x96, 0x87,
	0x0a, 0x1a, 0x0a, 0x0b, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0b, 0x0a, 0x09, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x0c, 0xe5,
	0x8a, 0xa0, 0xe5, 0xaf, 0x86, 0xe8, 0xbf, 0x94, 0xe5, 0x9b, 0x9e, 0x0a, 0xd7, 0x01, 0x0a, 0x11,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x41, 0x6e,
	0x79, 0x12, 0xc1, 0x01, 0x0a, 0xbe, 0x01, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0xfa, 0x01, 0x3c, 0x0a, 0x3a, 0x0a, 0x05, 0x40, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x2f,
	0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x23, 0x54, 0x68, 0x65, 0x20,
	0x74, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x82,
	0x02, 0x02, 0x10, 0x01, 0x92, 0x02, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20,
	0x61, 0x6e, 0x20, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x20, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
	0x61, 0x6c, 0x6f, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x40, 0x74, 0x79,
	0x70, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x0a, 0xb7, 0x07, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0xac, 0x07, 0x0a, 0xa9, 0x07, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0xfa,
	0x01, 0xef, 0x03, 0x0a, 0x74, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x6c, 0x0a, 0x6a, 0xca,
	0x01, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x92, 0x02, 0x55, 0x54, 0x68, 0x65, 0x20,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2c, 0x20, 0x77, 0x68, 0x69,
	0x63, 0x68, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x6e, 0x75, 0x6d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x5b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x5d, 0x5b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x5d,
	0x2e, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0xf5, 0x01, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xe9, 0x01, 0x0a, 0xe6, 0x01, 0xca, 0x01, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0xd9, 0x01, 0x41, 0x20, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x2d, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20,
	0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x45, 0x6e, 0x67,
	0x6c, 0x69, 0x73, 0x68, 0x2e, 0x20, 0x41, 0x6e, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x66,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x6e, 0x74,
	0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x5d, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5d, 0x20, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2c, 0x20, 0x6f, 0x72, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x0a, 0x7f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x77, 0x0a, 0x75, 0xca, 0x01, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x92, 0x02, 0x69, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74,
	0x20, 0x63, 0x61, 0x72, 0x72, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x20, 0x20, 0x54, 0x68, 0x65, 0x72, 0x65,
	0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x41, 0x50, 0x49, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73,
	0x65, 0x2e, 0x92, 0x02, 0xa9, 0x03, 0x54, 0x68, 0x65, 0x20, 0x60, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x60, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x20,
	0x61, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x69, 0x73, 0x20, 0x73, 0x75,
	0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67,
	0x20, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2c, 0x20, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x52, 0x45, 0x53, 0x54, 0x20, 0x41, 0x50,
	0x49, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x52, 0x50, 0x43, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2e,
	0x20, 0x49, 0x74, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x5b,
	0x67, 0x52, 0x50, 0x43, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x29, 0x2e, 0x20,
	0x45, 0x61, 0x63, 0x68, 0x20, 0x60, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x60, 0x20, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x74,
	0x68, 0x72, 0x65, 0x65, 0x20, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x3a, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2c,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2c, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x20, 0x59, 0x6f, 0x75, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x20,
	0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x68, 0x6f, 0x77, 0x20, 0x74, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x69, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x5b,
	0x41, 0x50, 0x49, 0x20, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x20, 0x47, 0x75, 0x69, 0x64, 0x65,
	0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x29, 0x2e,
}
