package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=postgres password=password dbname=bookstore port=5432 sslmode=disable TimeZone=UTC"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	db = d
}

func GetDB() *gorm.DB {
	return db
}