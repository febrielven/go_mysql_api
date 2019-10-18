package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=root dbname=go_gorm password=123456 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
