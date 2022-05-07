package database

import (
	"gin-api-rest/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	connection := "host=localhost user=postgres password=? dbname=students port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic("Error connectiong to database")
	}
	DB.AutoMigrate(&model.Student{})
}
