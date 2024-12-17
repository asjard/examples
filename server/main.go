package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/bootstrap"
	"github.com/asjard/asjard/core/client"
	"github.com/asjard/asjard/core/config"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/status"
	"github.com/hibiken/asynq"

	// 加载etcd配置源
	_ "github.com/asjard/asjard/pkg/config/etcd"
	// 从etcd发现服务, 并把当前服务注册到etcd
	_ "github.com/asjard/asjard/pkg/registry/etcd"
	// 加载grpc服务
	"github.com/asjard/asjard/pkg/server/grpc"
	"github.com/asjard/asjard/pkg/server/xasynq"

	// 加载rest服务
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/api/serverpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	serverpb.UnimplementedServerServer
	exit        <-chan struct{}
	client      serverpb.ServerClient
	asynqClient *asynq.Client
}

var _ bootstrap.Initiator = &ServerAPI{}
var _ xasynq.Handler = &ServerAPI{}

// Bootstrap 服务启动前会自动调用这个方法
// 当前这个方法内初始化了grpc客户端
func (api *ServerAPI) Start() error {
	conn, err := client.NewClient(grpc.Protocol, config.GetString("asjard.topology.services.examples.name", "server")).Conn()
	if err != nil {
		return err
	}
	api.client = serverpb.NewServerClient(conn)
	redisConn, err := xasynq.NewRedisConn("")
	if err != nil {
		return err
	}
	api.asynqClient = asynq.NewClient(redisConn)
	return nil
}

// Shutdown 服务停止会调用这里
func (api *ServerAPI) Stop() {}

// Say 接受rest请求然后去请求grpc请求
func (api *ServerAPI) Say(ctx context.Context, in *serverpb.HelloReq) (*serverpb.HelloReq, error) {
	return api.client.Call(ctx, in)
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
	payload, err := json.Marshal(in)
	if err != nil {
		logger.Error("marshal fail", "err", err)
		return nil, status.InternalServerError()
	}
	// 添加异步任务
	taskInfo, err := api.asynqClient.EnqueueContext(ctx,
		asynq.NewTask(xasynq.Pattern(serverpb.Server_Asynq_FullMethodName), payload))
	if err != nil {
		logger.Error("asynq enqueue fail", "err", err)
		return nil, status.InternalServerError()
	}
	logger.Debug("add asynq task info", "task_info", taskInfo)
	return in, nil
}

func (api *ServerAPI) Asynq(ctx context.Context, in *serverpb.HelloReq) (*emptypb.Empty, error) {
	logger.Info("----------asynq message recived----", "ctx", fmt.Sprintf("%T", ctx), "in", in.String())
	return nil, nil
}

// GrpcServiceDesc 提供grpc服务,需要实现这个方法
func (api *ServerAPI) GrpcServiceDesc() *grpc.ServiceDesc {
	return &serverpb.Server_ServiceDesc
}

// RestServiceDesc 提供rest服务,需要实现这个方法
func (api *ServerAPI) RestServiceDesc() *rest.ServiceDesc {
	return &serverpb.ServerRestServiceDesc
}

// AsynqServiceDesc asynq消费需要实现的方法
func (api *ServerAPI) AsynqServiceDesc() *xasynq.ServiceDesc {
	return &serverpb.ServerAsynqServiceDesc
}

func main() {
	server := asjard.New()
	// 添加grpc和rest服务
	server.AddHandler(&ServerAPI{
		exit: server.Exit(),
	}, rest.Protocol, grpc.Protocol, xasynq.Protocol)
	// 启动服务
	if err := server.Start(); err != nil {
		panic(err)
	}
}
