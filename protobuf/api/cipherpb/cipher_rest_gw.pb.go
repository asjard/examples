// Code generated by protoc-gen-go-rest2grpc-gw. DO NOT EDIT.
// versions:
// - protoc-gen-go-rest2grpc-gw v1.3.0
// - protoc             v5.27.0
// source: api/cipher.proto

package cipherpb

import (
	context "context"
	client "github.com/asjard/asjard/core/client"
	config "github.com/asjard/asjard/core/config"
	grpc "github.com/asjard/asjard/pkg/client/grpc"
	rest "github.com/asjard/asjard/pkg/server/rest"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type CipherAPI struct {
	UnimplementedCipherServer
	client CipherClient
}

func (api *CipherAPI) Start() error {
	conn, err := client.NewClient(grpc.Protocol, config.GetString("asjard.topology.services.cipher.name", "cipher")).Conn()
	if err != nil {
		return err
	}
	api.client = NewCipherClient(conn)
	return nil
}
func (api *CipherAPI) Stop() {
}

// 加密
func (api *CipherAPI) Encrypt(ctx context.Context, in *EncryptReq) (*EncryptResp, error) {
	return api.client.Encrypt(ctx, in)
}

// 解密
func (api *CipherAPI) Decrypt(ctx context.Context, in *emptypb.Empty) (*DecryptResp, error) {
	return api.client.Decrypt(ctx, in)
}
func (api *CipherAPI) RestServiceDesc() *rest.ServiceDesc {
	return &CipherRestServiceDesc
}
