package database

import (
	"fmt"
	"github.com/guneyin/locator/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Error)}

//func NewDB(cfg *config.Config) (*gorm.DB, error) {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/gorm?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.Port)
//	return gorm.Open(mysql.Open(dsn), gormConfig)
//}

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func NewTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}
