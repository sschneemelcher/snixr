package main

import ("github.com/gofiber/fiber/v2"
    "sschneemelcher/snixr/handlers"
)

func main() {
    app := fiber.New()

    app.Post("/api/links", handlers.CreateLink)
    app.Get("/:code", handlers.RedirectLink)
    app.Get("/api/links/:id", handlers.GetLink)
    app.Get("/api/links", handlers.ListLinks)

    app.Listen(":3000")
}
