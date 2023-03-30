package api

import (
	"context"
	"fmt"
	"net/http"

    "github.com/sschneemelcher/snixr/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func CreateLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
        // Parse request body and create new link 
    	type Body struct {
		    URL string `json:"url" xml:"url" form:"url"`
	    }

	    body := new(Body)
	    if err := c.BodyParser(body); err != nil {
		    return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse body"})
	    }

        // Generate short code for CreateLink
        shortCode, err := utils.GenerateCode(body.URL, rdb)
        if err != nil {
		    return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate short url"})
        }
        
        // Return new link as JSON response
        return c.SendString(fmt.Sprintf("New short url for %s: %s", body.URL, shortCode))
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

