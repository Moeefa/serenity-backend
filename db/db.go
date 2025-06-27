package db

import (
	"errors"
	"os"

	"github.com/moeefa/serenity/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	if dsn, exists := os.LookupEnv("DB_URL"); !exists || dsn == "" {
		return errors.New("DB_URL not set")
	}
 
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Recommendation{})

	DB = db
	return nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("DB is nil — did you forget to call db.Init()?")
	}
	return DB
}

