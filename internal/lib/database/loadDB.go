package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DevAthhh/DoZen/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PWD")

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable", user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("err with connecting to db: %v", err)
	}

	return db
}

func SyncDB(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Group{}, &models.Task{}); err != nil {
		log.Fatalf("err with sync database: %v", err)
	}
}
