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

// Handler for creating custom Links 
// parses request body and creates new link with custom name
// if the name is still available
//
// path /api/custom
func CreateCustomLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
        // Parse request body and create new link 
    	type Body struct {
		    URL string `json:"url" xml:"url" form:"url"`
		    CUSTOM_NAME string `json:"custom_name" xml:"custom_name" form:"custom_name"`
	    }

	    body := new(Body)
	    if err := c.BodyParser(body); err != nil {
            log.Printf("createlink error: Failed to parse body: %s\n", err)
		    return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse body"})
	    }

        // Check if custom name is available 
        _, err := rdb.Get(context.Background(), fmt.Sprintf("shortcode:%s", body.CUSTOM_NAME)).Result()
        if err != nil {
            err := rdb.Set(context.Background(), fmt.Sprintf("shortcode:%s", body.CUSTOM_NAME), body.URL, 0).Err()
            if err != nil {
		        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "error setting custom url"})
            }
        } else {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "URL already in use"})
        }

        // Return new link as JSON response
        return c.JSON(fiber.Map{"url": body.URL, "shortUrl": fmt.Sprintf("%s%s", os.Getenv("BASE_URL"), body.CUSTOM_NAME)})
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
        // TODO

        // Redirect user to original URL
        return c.Redirect(val, http.StatusMovedPermanently)
    }
}
