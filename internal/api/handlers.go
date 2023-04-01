package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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
            log.Printf("createlink error: Failed to parse body: %s\n", err)
		    return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse body"})
	    }

        // Generate short code for CreateLink
        shortCode, err := utils.GenerateCode(body.URL, rdb)
        if err != nil {
            log.Printf("createlink error: Failed to generate shortCode: %s\n", err)
		    return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
        }

        log.Printf("created new link: {shortCode: %s, url: %s}", shortCode, body.URL)
        
        // Return new link as JSON response
        return c.JSON(fiber.Map{"url": body.URL, "shortUrl": fmt.Sprintf("%s%s", os.Getenv("BASE_URL"), shortCode)})
    }

}

func RedirectLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
        // Look up link by code in database
        val, err := rdb.Get(context.Background(), fmt.Sprintf("shortcode:%s", c.Params("code"))).Result()

        if err != nil {
            return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Link not found"})
        }


        // Update link click count
        


        // Redirect user to original URL
        return c.Redirect(val, http.StatusMovedPermanently)
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

