package main

import (
	"cars/config"
	handler "cars/handlers"
	"cars/middleware"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.ConnectDB()
	config.ConnectMongo()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.SecurityHeader)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin":   "12345",
			"manager": "qwert",
			"john":    "doe",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user is not authorized",
			})
		},
	}))

	app.Use(etag.New())

	app.Post("/cars", handler.CreateCar)
	app.Get("/cars/:car_id", handler.GetCar)
	app.Delete("/cars/:car_id", handler.DeleteCar)

	fmt.Println("Fiber HTTP server listening...")
	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("Couldn't listen on port 3015,error:%v", err)
	}
}
