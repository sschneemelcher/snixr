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

// Handler for creating Links
// parses request body and creates new link
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
		return c.Status(http.StatusCreated).JSON(fiber.Map{"url": body.URL, "shortUrl": fmt.Sprintf("%s%s", os.Getenv("BASE_URL"), shortCode)})
	}

}

// Handler for getting Link information
// looks up link by code and returns shortcoe, url and click count
func GetClicks(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Look up link by code in database
		res, err := rdb.HGetAll(context.Background(), "shortcode:"+c.Params("code")).Result()
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Link not found"})
		}

		log.Printf("clicks: {shortcode: %s, url: %s, clicks: %s}", c.Params("code"), res["url"], res["clicks"])

		// Return clicks as JSON response
		return c.JSON(fiber.Map{"shortcode": c.Params("code"), "url": res["url"], "clicks": res["clicks"]})
	}
}

// Handler for creating custom Links
// parses request body and creates new link with custom name
// if the name is still available
func CreateCustomLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body and create new link
		type Body struct {
			URL         string `json:"url" xml:"url" form:"url"`
			CUSTOM_NAME string `json:"custom_name" xml:"custom_name" form:"custom_name"`
		}

		body := new(Body)
		if err := c.BodyParser(body); err != nil {
			log.Printf("createlink error: Failed to parse body: %s\n", err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse body"})
		}

		// Check if custom name is available
		_, err := rdb.HGet(context.Background(), "shortcode:"+body.CUSTOM_NAME, "url").Result()
		if err != nil {
			err := rdb.HSet(context.Background(), "shortcode:"+body.CUSTOM_NAME, "url", body.URL, "clicks", 0).Err()
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

// Handler for redirecting to original URL
// looks up link by code and redirects user to original URL
func RedirectLink(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Look up link by code in database
		val, err := rdb.HGet(context.Background(), "shortcode:"+c.Params("code"), "url").Result()

		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Link not found"})
		}

		// Update link click count
		err = rdb.HIncrBy(context.Background(), "shortcode:"+c.Params("code"), "clicks", 1).Err()
		if err != nil {
			log.Printf("Failed to update click count: %s", err)
		}

		log.Printf("redirection: {shortcode: %s, url %s", c.Params("code"), val)

		// Redirect user to original URL
		return c.Redirect(val, http.StatusTemporaryRedirect)
	}
}
