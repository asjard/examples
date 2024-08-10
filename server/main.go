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
	"github.com/asjard/asjard/pkg/server/grpc"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/serverpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type traceContextKeyType int

const (
	currentSpanKey traceContextKeyType = iota
)

type ServerAPI struct {
	serverpb.UnimplementedServerServer
	exit   <-chan struct{}
	client serverpb.ServerClient
}

func (api *ServerAPI) Bootstrap() error {
	conn, err := client.NewClient(grpc.Protocol, config.GetString("asjard.topology.services.examples.name", "server")).Conn()
	if err != nil {
		return err
	}
	api.client = serverpb.NewServerClient(conn)
	return nil
}

func (ServerAPI) Shutdown() {}

func (api *ServerAPI) Say(ctx context.Context, in *serverpb.HelloReq) (*serverpb.HelloReq, error) {
	return in, nil
}

func (api *ServerAPI) Hello(ctx context.Context, in *emptypb.Empty) (*serverpb.HelloReq, error) {
	return api.client.Say(ctx, &serverpb.HelloReq{})
}

// Log SSE请求
func (api *ServerAPI) Log(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	rtx, ok := ctx.(*rest.Context)
	if !ok {
		return nil, status.UnsupportProtocol
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

func (api *ServerAPI) GrpcServiceDesc() *grpc.ServiceDesc {
	return &serverpb.Server_ServiceDesc
}

func (api *ServerAPI) RestServiceDesc() *rest.ServiceDesc {
	return &serverpb.ServerRestServiceDesc
}

func main() {
	server := asjard.New()
	server.AddHandler(&ServerAPI{
		exit: server.Exit(),
	}, rest.Protocol, grpc.Protocol)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
