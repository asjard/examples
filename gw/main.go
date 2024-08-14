package main

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/status"
	_ "github.com/asjard/asjard/pkg/config/etcd"
	_ "github.com/asjard/asjard/pkg/registry/consul"

	// _ "github.com/asjard/asjard/pkg/registry/etcd"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/cipherpb"
	"github.com/asjard/examples/protobuf/mysqlpb"
	pb "github.com/asjard/examples/protobuf/serverpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb.ServerServer = &pb.ServerAPI{}
var _ rest.Handler = &pb.ServerAPI{}

type Rewrite struct {
	*pb.ServerAPI
	exit <-chan struct{}
}

func (api Rewrite) Log(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	logger.Debug("------is stream request-------")
	rtx, ok := ctx.(*rest.Context)
	if !ok {
		return nil, status.Error(codes.Unimplemented, "unsupport protocol")
	}
	rtx.SetContentType("text/event-stream")
	rtx.SetBodyStreamWriter(func(w *bufio.Writer) {
		for {
			select {
			case <-api.exit:
				logger.Debug("------server done----------")
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

func main() {
	server := asjard.New()
	server.AddHandlers(rest.Protocol, &Rewrite{
		ServerAPI: &pb.ServerAPI{},
		exit:      server.Exit(),
	},
		&mysqlpb.MysqlAPI{},
		&cipherpb.CipherAPI{})
	if err := server.Start(); err != nil {
		panic(err)
	}
}
