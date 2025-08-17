package handler

import (
	"cars/config"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func BenchmarkCarGet(b *testing.B) {
	//connect to db
	config.ConnectDB()
	//getting app from fiber
	app := fiber.New()

	app.Get("/cars", GetCar)

	req, _ := http.NewRequest("GET", "/cars/19", nil)
	req.Header.Set("Content-Type", "application/json")

	for b.Loop() {
		_, _ = app.Test(req, 5000)
	}
}
