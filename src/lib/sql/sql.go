package sql

import (
	"fmt"
	"go-clean/src/business/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

func Init(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic(err)
	}

	return db
}
