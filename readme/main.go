package main

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/bootstrap"
	"github.com/asjard/asjard/core/client"
	"github.com/asjard/asjard/core/config"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/runtime"
	"github.com/asjard/asjard/core/status"
	_ "github.com/asjard/asjard/pkg/config/etcd"
	"github.com/asjard/asjard/pkg/protobuf/requestpb"
	_ "github.com/asjard/asjard/pkg/registry/etcd"
	"github.com/asjard/asjard/pkg/server/grpc"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/api/readmepb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	readmepb.UnimplementedExamplesServer
	exit   <-chan struct{}
	client readmepb.ExamplesClient
}

var _ bootstrap.Initiator = &ServerAPI{}

// Bootstrap 服务启动前会自动调用这个方法
// 当前这个方法内初始化了grpc客户端
func (api *ServerAPI) Start() error {
	conn, err := client.NewClient(grpc.Protocol, config.GetString("asjard.topology.services.readme.name", "readme")).Conn()
	if err != nil {
		return err
	}
	api.client = readmepb.NewExamplesClient(conn)
	return nil
}

// Shutdown 服务停止会调用这里
func (api *ServerAPI) Stop() {}

// Say 接受rest请求然后去请求grpc请求
func (api *ServerAPI) Say(ctx context.Context, in *readmepb.HelloReq) (*readmepb.HelloReq, error) {
	return api.client.Call(ctx, in)
}

// Hello 直接处理逻辑, 请求参数为emptypb.Empty说明没有参数
func (api *ServerAPI) Hello(ctx context.Context, in *emptypb.Empty) (*readmepb.HelloReq, error) {
	return &readmepb.HelloReq{
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
func (api *ServerAPI) Call(ctx context.Context, in *readmepb.HelloReq) (*readmepb.HelloReq, error) {
	in.Configs = &readmepb.HelloReq_Configs{
		Timeout:               config.GetDuration("timeout", time.Second).String(),
		KeyInDifferentSourcer: config.GetString("test_key", ""),
	}
	app := runtime.GetAPP()
	in.Instance = &readmepb.HelloReq_Instance{
		Id:         app.Instance.ID,
		Name:       app.Instance.Name,
		Version:    app.Instance.Version,
		SystemCode: app.Instance.SystemCode,
		Metadata:   app.Instance.MetaData,
	}
	reqWithPage := &requestpb.ReqWithPage{
		Page: in.Page,
		Size: in.Size,
		Sort: in.Sort,
	}
	if err := reqWithPage.IsValid(20, []string{"created_at", "updated_at"}); err != nil {
		return nil, err
	}
	in.Page = reqWithPage.Page
	in.Size = reqWithPage.Size
	in.Sort = reqWithPage.Sort
	return in, nil
}

// GrpcServiceDesc 提供grpc服务,需要实现这个方法
func (api *ServerAPI) GrpcServiceDesc() *grpc.ServiceDesc {
	return &readmepb.Examples_ServiceDesc
}

// RestServiceDesc 提供rest服务,需要实现这个方法
func (api *ServerAPI) RestServiceDesc() *rest.ServiceDesc {
	return &readmepb.ExamplesRestServiceDesc
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
