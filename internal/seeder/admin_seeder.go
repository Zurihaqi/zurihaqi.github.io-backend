package seeder

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"zurihaqi.github.io-backend/internal/database"
	"zurihaqi.github.io-backend/internal/model"
)

func SeedAdmin() {
	db := database.DB

	var adminPassword string
	if adminPassword = os.Getenv("ADMIN_PASSWORD"); adminPassword == "" {
		log.Fatal("ADMIN_PASSWORD environment variable is not set")
	}

	var adminUsername string
	if adminUsername = os.Getenv("ADMIN_USERNAME"); adminUsername == "" {
		log.Fatal("ADMIN_USERNAME environment variable is not set")
	}

	var adminEmail string
	if adminEmail = os.Getenv("ADMIN_EMAIL"); adminEmail == "" {
		log.Fatal("ADMIN_EMAIL environment variable is not set")
	}

	password := adminPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	admin := model.User{
		Name:     adminUsername,
		Email:    adminPassword,
		Password: string(hashedPassword),
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("Failed to seed admin user:", err)
	}

	log.Println("Admin user seeded")
}
