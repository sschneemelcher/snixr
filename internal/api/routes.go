package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)


func SetupRoutes(app *fiber.App, rdb *redis.Client) {
    // Define a route for the root path
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile("./public/index.html");
    })
    app.Post("/api/shorten", CreateLink(rdb))
    app.Get("/:code", RedirectLink(rdb))
    //app.Get("/api/links/:id", GetLink)
    //app.Get("/api/links", ListLinks)
}

