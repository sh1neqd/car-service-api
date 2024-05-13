package handlers

import (
	"car-service1/app/internal/domain/car"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary GetAllCars
// @Description Getting all cars
// @Tags cars
// @ID get-cars
// @Produce json
// @Success 200 {object} []car.Car
// @Failure 400
// @Failure 500
// @Router /api/ [get]
func (h *Handler) getAllCars(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)
	filter := c.Query("filter")

	cars, err := h.services.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	filteredCars := filterCars(cars, filter)
	paginatedCars := paginateCars(filteredCars, limit, offset)
	return c.Status(fiber.StatusOK).JSON(paginatedCars)
}

// @Summary GetCar
// @Description Getting car by id
// @Tags cars
// @ID get-car
// @Produce json
// @Success 200 {object} car.Car
// @Failure 400
// @Failure 500
// @Router /api/{id} [get]
func (h *Handler) getCarById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	carById, err := h.services.Car.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(carById)

}

// @Summary Add cars
// @Description Creating cars
// @Tags cars
// @ID get-car
// @Accept json
// @Produce json
// @Param input body car.CreateCarDTO true "car dto"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/ [post]
func (h *Handler) addCar(c *fiber.Ctx) error {
	var cars car.CreateCarDTO
	if err := c.BodyParser(&cars); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	//response := `{
	//	"regNum": "AC23",
	//	"mark": "Porsche",
	//	"model": "911 GT3",
	//	"owner": {
	//		"name":"Alex",
	//		"surname":"Bogdan"
	//	}
	//}` // for testing

	var counter int
	for _, regNum := range cars.RegNums {
		url := fmt.Sprintf("https://vasheApi.net/api?regNum=%s", regNum) // НУЖНО ВСТАВИТЬ СВОЕ API

		res, err := http.Get(url)
		if err != nil {
			fmt.Println("request execution err: ", err)
		}
		defer res.Body.Close()

		var dto car.CarDTO

		err = json.NewDecoder(res.Body).Decode(&dto)
		//err := json.Unmarshal([]byte(response), &dto) // for testing
		if err != nil {
			fmt.Println("failed to decode: ", err)
		} else {
			err := h.services.Car.Create(dto)
			if err != nil {
				logrus.Printf("failed to create car with err: %f", err)
			} else {
				counter++
			}
		}
	}

	resp := fmt.Sprintf("created cars counter: %v. others already exist in our database or do not exist in the api database", counter)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": resp,
	})
}

// @Summary DeleteCar
// @Description Deleting car by id
// @Tags cars
// @ID delete-car
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/{id} [delete]
func (h *Handler) deleteCarById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.services.Car.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("car with id %d deleted", id),
	})
}

// @Summary Update car
// @Description Updating car by id
// @Tags cars
// @ID update-car
// @Accept json
// @Produce json
// @Param input body car.UpdateCarDTO true "update car"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/ [patch]
func (h *Handler) updateCar(c *fiber.Ctx) error {
	var err error
	var updatedCar car.UpdateCarDTO
	if err := c.BodyParser(&updatedCar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.services.Update(id, updatedCar)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("car with id %d updated", id),
	})
}
