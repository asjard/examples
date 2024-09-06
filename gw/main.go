package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/status"
	_ "github.com/asjard/asjard/pkg/config/etcd"
	_ "github.com/asjard/asjard/pkg/registry/etcd"

	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/protobuf/api/cipherpb"
	"github.com/asjard/examples/protobuf/api/mysqlpb"
	"github.com/asjard/examples/protobuf/api/readmepb"
	pb "github.com/asjard/examples/protobuf/api/serverpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb.ServerServer = &pb.ServerAPI{}
var _ rest.Handler = &pb.ServerAPI{}

type Rewrite struct {
	*pb.ServerAPI
	exit <-chan struct{}
}

// Response 自定义rest输出
type Response struct {
	Err  error `json:"err"`
	Data any   `json:"data"`
}

func init() {
	// 添加自定义输出
	rest.AddWriter("custome_writer", CustomeWriter)
}

// CustomeWriter 自定义输出
func CustomeWriter(c *rest.Context, data any, err error) {
	b, err := json.Marshal(&Response{
		Err:  err,
		Data: data,
	})
	if err != nil {
		logger.Error("custome writer fail", "err", err)
	}
	c.Write(b)
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
		&cipherpb.CipherAPI{},
		&readmepb.ExamplesAPI{})
	if err := server.Start(); err != nil {
		panic(err)
	}
}
