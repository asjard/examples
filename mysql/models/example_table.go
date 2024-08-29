package models

import (
	"context"
	"errors"
	"time"

	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/core/status"
	"github.com/asjard/asjard/pkg/protobuf/requestpb"
	"github.com/asjard/asjard/pkg/stores/xgorm"
	pb "github.com/asjard/examples/protobuf/api/mysqlpb"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type ExampleTable struct {
	Id        int64  `gorm:"column:id;type:INT(20);primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;type:VARCHAR(20);uniqueIndex"`
	Age       uint32 `gorm:"column:age;type:INT"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 数据库表名
func (ExampleTable) TableName() string {
	return "example_table"
}

// ModelName 全局唯一的表名
func (ExampleTable) ModelName() string {
	return "example_database_example_table"
}

func (t ExampleTable) Create(ctx context.Context, in *pb.CreateOrUpdateReq) error {
	db, err := xgorm.DB(ctx)
	if err != nil {
		return err
	}
	if err := db.Where("name=?", in.Name).First(&ExampleTable{}).Error; err == nil {
		return status.Errorf(codes.AlreadyExists, "record %s already exist", in.Name)
	}

	if err := db.Create(&ExampleTable{
		Name: in.Name,
		Age:  in.Age,
	}).Error; err != nil {
		logger.Error("create fail", "name", in.Name, "err", err)
		return status.InternalServerError()
	}
	return nil
}

func (t ExampleTable) Update(ctx context.Context, in *pb.CreateOrUpdateReq) (*pb.ExampleInfo, error) {
	db, err := xgorm.DB(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := t.Get(ctx, &pb.ReqWithName{Name: in.Name}); err != nil {
		return nil, err
	}
	if err := db.Model(&ExampleTable{}).Where("name=?", in.Name).Updates(&ExampleTable{
		Name: in.Name,
		Age:  in.Age,
	}).Error; err != nil {
		logger.Error("update record fail", "name", in.Name, "err", err)
		return nil, status.InternalServerError()
	}
	return t.Get(ctx, &pb.ReqWithName{Name: in.Name})
}

func (t ExampleTable) Get(ctx context.Context, in *pb.ReqWithName) (*pb.ExampleInfo, error) {
	db, err := xgorm.DB(ctx)
	if err != nil {
		return nil, err
	}
	var record ExampleTable
	if err := db.Where("name=?", in.Name).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "record not found")
		}
		logger.Error("get record fail", "name", in.Name, "err", err)
		return nil, status.InternalServerError()
	}
	return record.info(), nil
}

func (t ExampleTable) Search(ctx context.Context, in *pb.SearchReq) (*pb.ExampleList, error) {
	db, err := xgorm.DB(ctx)
	if err != nil {
		return nil, err
	}
	var total int64
	db.Model(&ExampleTable{}).Count(&total)
	records := make([]ExampleTable, 0, in.Size)
	db.Scopes(requestpb.ReqWithPageGormScope(in.Page, in.Size, in.Sort)).Find(&records)
	exampleInfos := make([]*pb.ExampleInfo, 0, in.Size)
	for _, record := range records {
		exampleInfos = append(exampleInfos, record.info())
	}
	return &pb.ExampleList{
		Total: int32(total),
		List:  exampleInfos,
	}, nil
}

func (t ExampleTable) Del(ctx context.Context, in *pb.ReqWithName) error {
	db, err := xgorm.DB(ctx)
	if err != nil {
		return err
	}
	if err := db.Where("name=?", in.Name).Delete(&ExampleTable{}).Error; err != nil {
		logger.Error("delete record fail", "name", in.Name, "err", err)
		return status.InternalServerError()
	}
	return nil
}

func (t ExampleTable) info() *pb.ExampleInfo {
	return &pb.ExampleInfo{
		Id:        t.Id,
		Name:      t.Name,
		Age:       t.Age,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
		UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
	}
}
