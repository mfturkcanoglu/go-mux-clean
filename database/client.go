package database

import (
	"log"
	"standart-lib-rest-api/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to db", err)
		return
	}
	log.Println("Connected to database.")
}
func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Db migration completed.")
}
