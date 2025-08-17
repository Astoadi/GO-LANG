package handler

import (
	"cars/config"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCarAdd(t *testing.T) {
	config.ConnectDB()

	app := fiber.New()

	app.Post("/cars", CreateCar)

	body := `
	{
		"name": "rounda rousy",
		"model": "v7 inline",
		"brand": "buggati",
		"year": 2024,
		"price": 9000000
	}`
	req, _ := http.NewRequest("POST", "/cars", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("request failed :%v", err)
	}

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
