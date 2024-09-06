// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: api/cipherpb/cipher.proto

package cipherpb

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
	Cipher_Encrypt_FullMethodName = "/api.v1.cipher.Cipher/Encrypt"
	Cipher_Decrypt_FullMethodName = "/api.v1.cipher.Cipher/Decrypt"
)

// CipherClient is the client API for Cipher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CipherClient interface {
	// 加密
	Encrypt(ctx context.Context, in *EncryptReq, opts ...grpc.CallOption) (*EncryptResp, error)
	// 解密
	Decrypt(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DecryptResp, error)
}

type cipherClient struct {
	cc grpc.ClientConnInterface
}

func NewCipherClient(cc grpc.ClientConnInterface) CipherClient {
	return &cipherClient{cc}
}

func (c *cipherClient) Encrypt(ctx context.Context, in *EncryptReq, opts ...grpc.CallOption) (*EncryptResp, error) {
	out := new(EncryptResp)
	err := c.cc.Invoke(ctx, Cipher_Encrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipherClient) Decrypt(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DecryptResp, error) {
	out := new(DecryptResp)
	err := c.cc.Invoke(ctx, Cipher_Decrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CipherServer is the server API for Cipher service.
// All implementations must embed UnimplementedCipherServer
// for forward compatibility
type CipherServer interface {
	// 加密
	Encrypt(context.Context, *EncryptReq) (*EncryptResp, error)
	// 解密
	Decrypt(context.Context, *emptypb.Empty) (*DecryptResp, error)
	mustEmbedUnimplementedCipherServer()
}

// UnimplementedCipherServer must be embedded to have forward compatible implementations.
type UnimplementedCipherServer struct {
}

func (UnimplementedCipherServer) Encrypt(context.Context, *EncryptReq) (*EncryptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedCipherServer) Decrypt(context.Context, *emptypb.Empty) (*DecryptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedCipherServer) mustEmbedUnimplementedCipherServer() {}

// UnsafeCipherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CipherServer will
// result in compilation errors.
type UnsafeCipherServer interface {
	mustEmbedUnimplementedCipherServer()
}

func RegisterCipherServer(s grpc.ServiceRegistrar, srv CipherServer) {
	s.RegisterService(&Cipher_ServiceDesc, srv)
}

func _Cipher_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cipher_Encrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherServer).Encrypt(ctx, req.(*EncryptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cipher_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cipher_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherServer).Decrypt(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Cipher_ServiceDesc is the grpc.ServiceDesc for Cipher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cipher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.cipher.Cipher",
	HandlerType: (*CipherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _Cipher_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _Cipher_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cipherpb/cipher.proto",
}
