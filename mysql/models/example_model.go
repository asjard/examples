package models

import (
	"context"
	"fmt"

	"github.com/asjard/asjard/core/bootstrap"
	"github.com/asjard/asjard/core/status"
	"github.com/asjard/asjard/pkg/cache"
	"github.com/asjard/asjard/pkg/stores"
	pb "github.com/asjard/examples/protobuf/api/mysqlpb"
	"google.golang.org/grpc/codes"
)

type ExampleModel struct {
	stores.Model
	*ExampleTable
	kvCache *cache.CacheRedis
}

// NewExampleModel 模型初始化
// 用XXXModel去复写XXXTable的方法
func NewExampleModel() *ExampleModel {
	exampleModel := &ExampleModel{
		ExampleTable: &ExampleTable{},
	}
	bootstrap.AddBootstrap(exampleModel)
	return exampleModel
}

// Start 缓存初始化
func (model *ExampleModel) Start() (err error) {
	localCache, err := cache.NewLocalCache(model.ExampleTable)
	if err != nil {
		return err
	}
	model.kvCache, err = cache.NewRedisKeyValueCache(model.ExampleTable,
		cache.WithLocalCache(localCache))
	if err != nil {
		return err
	}
	return nil
}

// Stop .
func (model *ExampleModel) Stop() {}

func (model *ExampleModel) Create(ctx context.Context, in *pb.CreateOrUpdateReq) error {
	// 需要删除删除缓存
	// 创建需要提前分配缓存Key
	// 如果是按照主键ID缓存的，提前生成主键ID，创建记录使用提前生成的主键ID创建记录
	if err := model.SetData(ctx,
		model.kvCache.WithGroup(model.searchGroup()).WithKey(model.getCacheKey(in.Name)),
		func() error {
			return model.ExampleTable.Create(ctx, in)
		}); err != nil {
		return err
	}
	return nil
}

func (model *ExampleModel) Update(ctx context.Context, in *pb.CreateOrUpdateReq) (*pb.ExampleInfo, error) {
	if err := model.SetData(ctx,
		model.kvCache.WithGroup(model.searchGroup()).WithKey(model.getCacheKey(in.Name)),
		func() error {
			if _, err := model.ExampleTable.Update(ctx, in); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return nil, err
	}
	return model.Get(ctx, &pb.ReqWithName{Name: in.Name})
}

func (model *ExampleModel) Get(ctx context.Context, in *pb.ReqWithName) (*pb.ExampleInfo, error) {
	var resp pb.ExampleInfo
	if err := model.GetData(ctx,
		&resp,
		model.kvCache.WithKey(model.getCacheKey(in.Name)),
		func() (any, error) {
			return model.ExampleTable.Get(ctx, in)
		}); err != nil || resp.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "record not found")
	}
	return &resp, nil
}

func (model *ExampleModel) Search(ctx context.Context, in *pb.SearchReq) (*pb.ExampleList, error) {
	var result pb.ExampleList
	if err := model.GetData(ctx, &result,
		model.kvCache.WithKey(model.searchCacheKey(in)).WithGroup(model.searchGroup()),
		func() (any, error) {
			return model.ExampleTable.Search(ctx, in)
		}); err != nil {
		return nil, err
	}
	return &result, nil
}

func (model *ExampleModel) Del(ctx context.Context, in *pb.ReqWithName) error {
	return model.SetData(ctx,
		model.kvCache.WithKey(model.getCacheKey(in.Name)).WithGroup(model.searchGroup()),
		func() error {
			return model.ExampleTable.Del(ctx, in)
		})
}

func (model *ExampleModel) searchGroup() string {
	return "search"
}

func (model *ExampleModel) searchCacheKey(in *pb.SearchReq) string {
	return fmt.Sprintf("search:%d:%d:%s", in.Page, in.Size, in.Sort)
}

func (model *ExampleModel) getCacheKey(name string) string {
	return fmt.Sprintf("info:%s", name)
}
