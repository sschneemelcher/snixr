package main

import (
	"sschneemelcher/snixr/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func main() {
    app := fiber.New()

    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    app.Post("/api/links", handlers.CreateLink(rdb))
    app.Get("/:code", handlers.RedirectLink(rdb))
    app.Get("/api/links/:id", handlers.GetLink)
    app.Get("/api/links", handlers.ListLinks)


    app.Listen(":3000")
}
