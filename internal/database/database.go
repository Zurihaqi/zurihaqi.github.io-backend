package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"

	"zurihaqi.github.io-backend/internal/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	var dbName string
	if dbName = os.Getenv("DATABASE_NAME"); dbName == "" {
		log.Fatal("DATABASE_NAME environment variable is not set")
	}

	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	if err = DB.AutoMigrate(
		&model.Project{},
		&model.Category{},
		&model.Technology{},
		&model.Image{},
		&model.ImgAlt{},
		&model.User{},
	); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database connected and migrated.")
}
