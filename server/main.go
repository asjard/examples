package main

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/client"
	"github.com/asjard/asjard/core/config"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/status"

	// 加载consul配置源
	_ "github.com/asjard/asjard/pkg/config/consul"
	// 从consul发现服务,并把服务注册到consul
	_ "github.com/asjard/asjard/pkg/registry/consul"
	// 加载etcd配置源
	_ "github.com/asjard/asjard/pkg/config/etcd"
	// 从etcd发现服务, 并把当前服务注册到etcd
	_ "github.com/asjard/asjard/pkg/registry/etcd"
	// 加载grpc服务
	"github.com/asjard/asjard/pkg/server/grpc"
	// 加载rest服务
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/serverpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	serverpb.UnimplementedServerServer
	exit   <-chan struct{}
	client serverpb.ServerClient
}

// Bootstrap 服务启动前会自动调用这个方法
// 当前这个方法内初始化了grpc客户端
func (api *ServerAPI) Bootstrap() error {
	conn, err := client.NewClient(grpc.Protocol, config.GetString("asjard.topology.services.examples.name", "server")).Conn()
	if err != nil {
		return err
	}
	api.client = serverpb.NewServerClient(conn)
	return nil
}

// Shutdown 服务停止会调用这里
func (api *ServerAPI) Shutdown() {}

// Say 接受rest请求然后去请求grpc请求
func (api *ServerAPI) Say(ctx context.Context, in *serverpb.HelloReq) (*serverpb.HelloReq, error) {
	return api.client.Call(ctx, in)
}

// Hello 直接处理逻辑, 请求参数为emptypb.Empty说明没有参数
func (api *ServerAPI) Hello(ctx context.Context, in *emptypb.Empty) (*serverpb.HelloReq, error) {
	return &serverpb.HelloReq{
		RegionId: "hello",
	}, nil
}

// Log SSE请求
func (api *ServerAPI) Log(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	rtx, ok := ctx.(*rest.Context)
	if !ok {
		return nil, status.UnsupportProtocol()
	}
	rtx.SetContentType("text/event-stream")
	rtx.SetBodyStreamWriter(func(w *bufio.Writer) {
		for {
			select {
			case <-api.exit:
				return
			default:
				w.Write([]byte(fmt.Sprintf("data: %s\n\n", time.Now())))

				if err := w.Flush(); err != nil {
					logger.Debug("client disconnected", "err", err)
					return
				}

				time.Sleep(time.Second)
			}
		}
	})
	return nil, nil
}

// Call 实时获取配置并返回
func (api *ServerAPI) Call(ctx context.Context, in *serverpb.HelloReq) (*serverpb.HelloReq, error) {
	in.Configs = &serverpb.HelloReq_Configs{
		KeyInDifferentSourcer: config.GetString("test_key", ""),
	}
	return in, nil
}

// GrpcServiceDesc 提供grpc服务,需要实现这个方法
func (api *ServerAPI) GrpcServiceDesc() *grpc.ServiceDesc {
	return &serverpb.Server_ServiceDesc
}

// RestServiceDesc 提供rest服务,需要实现这个方法
func (api *ServerAPI) RestServiceDesc() *rest.ServiceDesc {
	return &serverpb.ServerRestServiceDesc
}

func main() {
	server := asjard.New()
	// 添加grpc和rest服务
	server.AddHandler(&ServerAPI{
		exit: server.Exit(),
	}, rest.Protocol, grpc.Protocol)
	// 启动服务
	if err := server.Start(); err != nil {
		panic(err)
	}
}
