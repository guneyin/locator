package database

import (
	"fmt"
	"github.com/guneyin/locator/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Error)}

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	return gorm.Open(mysql.Open(dsn), gormConfig)
}

func NewTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}
