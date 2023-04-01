package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sschneemelcher/snixr/internal/api"
	"github.com/sschneemelcher/snixr/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
    loadEnv()

    // Create a new Fiber instance
    app := fiber.New()
    
    // Setup redis connection
    rdb := db.SetupDB()
    
    // Set up Cross-Origin Resource Sharing (CORS) so that any client can access the API
    app.Use(cors.New())

    // Setup routes
    api.SetupRoutes(app, rdb)

    app.Listen(":" + os.Getenv("PORT"))
}
