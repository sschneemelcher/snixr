package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// SetupRoutes sets up all the routes for the API
func SetupRoutes(app *fiber.App, rdb *redis.Client) {
	// Landing page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
	app.Get("/assets/logo.png", func(c *fiber.Ctx) error {
		return c.SendFile("./public/img/logo.png")
	})
	// Shorten a link with a random code
	app.Post("/api/shorten", CreateLink(rdb))
	// Shorten a link with a user defined custom code
	app.Post("/api/custom", CreateCustomLink(rdb))
	// Redirect a link
	app.Get("/:code", RedirectLink(rdb))
}
