package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(conn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(conn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
}

func NewTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
}
