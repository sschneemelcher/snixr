package main

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/sschneemelcher/snixr/internal/api"
	"github.com/sschneemelcher/snixr/internal/db"
	"github.com/sschneemelcher/snixr/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.LoadEnv()

	// Create a new Fiber instance
	app := fiber.New()

	// Setup redis connection
	rdb := db.SetupDB()

	// Set up Cross-Origin Resource Sharing (CORS) so that any client can access the API
	app.Use(cors.New())

	// Setup routes
	api.SetupRoutes(app, rdb)

	if os.Getenv("ENV") == "prod" {
		// Create tls certificate
		cer, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/key.pem")
		if err != nil {
			log.Fatal(err)
		}

		config := &tls.Config{Certificates: []tls.Certificate{cer}}

		// Create custom listener
		ln, err := tls.Listen("tcp", ":443", config)
		if err != nil {
			panic(err)
		}

		// Start server with https/ssl enabled on http://localhost:443
		log.Fatal(app.Listener(ln))
	} else {
		app.Listen(":" + os.Getenv("PORT"))
	}
}
