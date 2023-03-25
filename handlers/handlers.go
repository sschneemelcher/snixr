package handlers

import (
	"github.com/gofiber/fiber/v2"
    "fmt"
)

func CreateLink(c *fiber.Ctx) error {
    // Parse request body and create new link
    // Save link to database
    // Return new link as JSON response
    return c.SendString("I'm a POST request at `create new link`!")

}

func RedirectLink(c *fiber.Ctx) error {
    // Look up link by code in database
    // Update link click count
    // Redirect user to original URL
    return c.SendString(fmt.Sprintf("Redirect Link %s!", c.Params("code")))
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

