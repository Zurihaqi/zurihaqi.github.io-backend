package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"zurihaqi.github.io-backend/internal/database"
	"zurihaqi.github.io-backend/internal/route"
	seeder "zurihaqi.github.io-backend/internal/seeder"
)

func main() {
	_ = godotenv.Load()

	database.ConnectDatabase()

	runSeeders := os.Getenv("RUN_SEEDERS")
	if runSeeders == "true" {
		seeder.RunSeeders()
		log.Println("Seeding completed")
	}

	r := gin.Default()
	route.RegisterRoutes(r, database.DB)
	r.Run(":8080")
}
