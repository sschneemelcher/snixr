package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)


func SetupRoutes(app *fiber.App, rdb *redis.Client) {
    app.Post("/api/shorten", CreateLink(rdb))
    app.Get("/:code", RedirectLink(rdb))
    //app.Get("/api/links/:id", GetLink)
    //app.Get("/api/links", ListLinks)
}

