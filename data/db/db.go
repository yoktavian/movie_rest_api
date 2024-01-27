package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBname   string
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/movie-dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
