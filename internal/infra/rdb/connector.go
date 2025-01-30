package rdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RDB struct {
	DB *gorm.DB
}

func New() *RDB {
	dsn := "host=47.245.40.222 user=admin password=thF@AHgy3SUR dbname=drone_sphere port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &RDB{
		DB: db,
	}
}
