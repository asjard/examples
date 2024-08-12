package main

import (
	"log"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/core/bootstrap"
	_ "github.com/asjard/asjard/pkg/registry/etcd"
	"github.com/asjard/asjard/pkg/server/grpc"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/examples/mysql/apis"
	"github.com/asjard/examples/mysql/model"
)

type System struct{}

// Bootstrap 系统初始化
func (System) Bootstrap() error {
	if err := model.Init(); err != nil {
		return err
	}
	return nil
}

// Shutdown 系统停止
func (System) Shutdown() {}

func init() {
	bootstrap.AddBootstrap(&System{})
}

func main() {
	server := asjard.New()
	server.AddHandler(apis.NewExampleAPI(), rest.Protocol, grpc.Protocol)
	if err := server.Start(); err != nil {
		log.Println(err.Error())
	}
}
