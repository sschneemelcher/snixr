//go:build integration

package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sschneemelcher/snixr/internal/db"
	"github.com/sschneemelcher/snixr/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
	// Load environment variables
	utils.LoadEnv("../../.env")

	// Create a new Fiber app for the test
	app := fiber.New()

	// Setup redis client
	rdb := db.SetupDB()

	// Set up the handler being tested
	app.Post("/api/shorten", CreateLink(rdb))

	// Create a new HTTP request for the test
	reqBody := `{"url":"https://www.example.org/very/long/url"}`
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Execute the request and get the response
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Check the response headers
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	// Check the response body
	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, result["shortCode"])
}

func TestGetClicks(t *testing.T) {
	// Load environment variables
	utils.LoadEnv("../../.env")

	// Create a new Fiber app for the test
	app := fiber.New()

	// Setup redis client
	rdb := db.SetupDB()
	rdb.HSetNX(context.Background(), "shortcode:123", "url", "https://example.org")
	rdb.HSetNX(context.Background(), "shortcode:123", "clicks", 0)

	// Set up the neccessary handlers
	app.Get("/:code/clicks", GetClicks(rdb))

	// Create a new HTTP request for the test
	req := httptest.NewRequest(http.MethodGet, "/123/clicks", nil)

	// Execute the request and get the response
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check the response headers
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	// Check the response body
	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0", result["clicks"])
}
