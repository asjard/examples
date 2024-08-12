package model

import (
	"context"

	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/pkg/stores/xgorm"
)

func Init() error {
	logger.Debug("create table")
	db, err := xgorm.DB(context.Background())
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&ExampleTable{}); err != nil {
		return err
	}
	return nil
}
