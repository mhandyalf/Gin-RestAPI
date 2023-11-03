package database

import (
	"os"
	"restapi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	url := os.Getenv("supabase")

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(models.User{})

	return db

}
