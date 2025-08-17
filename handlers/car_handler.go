package handler

import (
	"cars/config"
	"cars/models"
	"encoding/json"
	"fmt"

	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"

	"net/http"
)

var mu sync.Mutex

// CarInventory godoc
// @Summary      Create a new car
// @Description  Add a new car to the inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        cars   body    models.Car  true  "Add car"
// @Success      200  {object}  models.Car
// @Failure      400  {object}  models.Errors
// @Router       /car [post]

func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.Errors{
			Error:   "Incorrect input body",
			Details: err.Error(),
		})
	}
	//car.Insert()

	//mongodb
	coll := config.MongoDB.Database("car-inventory").Collection("cars")
	_, err := coll.InsertOne(c.Context(), car)
	if err != nil {
		fmt.Printf("car not saved %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(&models.Errors{
			Error:   "unable to add a car",
			Details: err.Error(),
		})
	}

	fmt.Println("Car saved successfully with ID:", car.ID)
	return c.Status(fiber.StatusCreated).JSON(car)
}

func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	car_id := c.Params("car_id")
	if value, err := strconv.Atoi(car_id); err == nil {
		car.ID = int64(value)
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	if err := car.Get(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "car with the given id not found",
			"id":    car.ID,
		})
	}
	//fmt.Println("Car found with the id", car_id)
	return c.Status(fiber.StatusFound).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}
	car_id := c.Params("car_id")
	if value, err := strconv.Atoi(car_id); err == nil {
		car.ID = int64(value)
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	car.Delete()
	fmt.Println("car delete with the id:", car_id)
	return c.SendStatus(fiber.StatusNoContent) // No body, no content type needed
}

func ListCars(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	cars := make([]models.Car, 0, len(models.Cars))
	for _, car := range models.Cars {
		cars = append(cars, car)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
