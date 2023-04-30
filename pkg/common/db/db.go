package db

import (
	"log"

	"github.com/nexentra/inteligpt/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Comment{},&models.User{})

	return db
}
