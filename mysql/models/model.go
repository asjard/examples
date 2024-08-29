package models

import (
	"context"

	"github.com/asjard/asjard/core/config"
	"github.com/asjard/asjard/core/logger"
	"github.com/asjard/asjard/pkg/stores/xgorm"
)

func Init() error {
	logger.Debug("create table")
	if !config.GetBool("examples.gorm.autoMigrate", false) {
		return nil
	}
	db, err := xgorm.DB(context.Background())
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&ExampleTable{}); err != nil {
		return err
	}
	return nil
}
