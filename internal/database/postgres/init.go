package postgres

import (
	"log"

	//"./models"
	"github.com/ricardo95fraga/goplication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://go:password@localhost:5432/godatabase"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
