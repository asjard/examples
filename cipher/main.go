package main

import (
	"context"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/config"
	"github.com/asjard/asjard/core/security"
	cs "github.com/asjard/asjard/pkg/security"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/cipherpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	AESCipherName = "aesCBCPkcs5paddingCipherExample"
)

func init() {
	// 添加一个自定义的aes加解密组件
	security.AddCipher(AESCipherName, cs.NewAESCipher)
}

// CipherAPI 加解密示例
type CipherAPI struct {
	cipherpb.UnimplementedCipherServer
}

func (api *CipherAPI) Encrypt(ctx context.Context, in *cipherpb.EncryptReq) (*cipherpb.EncryptResp, error) {
	return &cipherpb.EncryptResp{}, nil
}

func (api *CipherAPI) Decrypt(ctx context.Context, in *emptypb.Empty) (*cipherpb.DecryptResp, error) {
	return &cipherpb.DecryptResp{
		AesEncryptValueInPlainFile:         config.GetString("testAESEncrptValue", "", config.WithCipher(cs.AESCipherName)),
		Base64EncryptValueInPlainFile:      config.GetString("testBase64EncryptValue", "", config.WithCipher(security.Base64CipherName)),
		PlainValueInAesEncryptFile:         config.GetString("testPlainValueInAESEncryptFile", ""),
		AesEncryptValueInAesEncryptFile:    config.GetString("testAESEncryptValueInAESEncryptFile", "", config.WithCipher(AESCipherName)),
		Base64EncryptValueInAesEncryptFile: config.GetString("testBase64EncryptValueInAESEncryptFile", "", config.WithCipher(security.Base64CipherName)),
	}, nil
}

func (CipherAPI) RestServiceDesc() *rest.ServiceDesc {
	return &cipherpb.CipherRestServiceDesc
}

func main() {
	server := asjard.New()
	server.AddHandler(&CipherAPI{}, rest.Protocol)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
