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

	var adminName string
	if adminName = os.Getenv("ADMIN_NAME"); adminName == "" {
		log.Fatal("ADMIN_NAME environment variable is not set")
	}

	var adminEmail string
	if adminEmail = os.Getenv("ADMIN_EMAIL"); adminEmail == "" {
		log.Fatal("ADMIN_EMAIL environment variable is not set")
	}

	if err := db.Exec("DELETE FROM users").Error; err != nil {
		log.Fatal("Failed to delete existing users:", err)
	}
	if err := db.Exec("DELETE FROM sqlite_sequence WHERE name = 'users'").Error; err != nil {
		log.Fatal("Failed to reset users ID sequence:", err)
	}

	password := adminPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	admin := model.User{
		Name:     adminName,
		Email:    adminEmail,
		Password: string(hashedPassword),
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("Failed to seed admin user:", err)
	}

	log.Println("Admin user seeded")
}
