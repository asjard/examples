package main

import (
	"context"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/status"
	"github.com/asjard/asjard/pkg/database/redis"
	"github.com/asjard/asjard/pkg/server/rest"
	pb "github.com/asjard/examples/protobuf/serverpb"
)

type HelloAPI struct {
	pb.UnimplementedHelloServer
}

func (HelloAPI) Say(ctx context.Context, in *pb.SayReq) (*pb.SayReq, error) {
	rds, err := redis.Client()
	if err != nil {
		return nil, err
	}
	if result := rds.Set(ctx, "hello", "world", 0); result.Err() != nil {
		logger.Error("set redis data fail", "err", err)
		return nil, status.InternalServerError
	}
	result := rds.Get(ctx, "hello")
	if result.Err() != nil {
		return nil, status.InternalServerError
	}
	in.App = result.Val()
	return in, nil
}

func (HelloAPI) RestServiceDesc() *rest.ServiceDesc {
	return &pb.HelloRestServiceDesc
}

func main() {
	server := asjard.New()
	server.AddHandler(&HelloAPI{}, rest.Protocol)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
