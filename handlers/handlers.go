package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func CreateLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
        // Parse request body and create new link
        // Save link to database
        // Return new link as JSON response

        // Store the link in Redis
        err := rdb.Set(context.Background(), "hello", "world", 0).Err()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to store link in Redis",
            })
        }


        return c.SendString("I'm a POST request at `create new link`!")
    }

}

func RedirectLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
        // Look up link by code in database
        // Update link click count
        // Redirect user to original URL

        val, _ := rdb.Get(context.Background(), c.Params("code")).Result()

        return c.SendString(fmt.Sprintf("Redirect Link %s : %s!", c.Params("code"), val))
    }
}

func GetLink(c *fiber.Ctx) error {
    // Look up link by ID in database
    // Return link as JSON response
    return c.SendString("Get Link!")
}

func ListLinks(c *fiber.Ctx) error {
    // Query database for all links
    // Return list of links as JSON response
    return c.SendString("List Links!")
}

