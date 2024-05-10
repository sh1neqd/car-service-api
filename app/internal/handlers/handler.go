package handlers

import (
	_ "car-service1/app/internal/docs"
	"car-service1/app/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/gofiber/swagger"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Accept, Content-Type, Content-Length, Accept-Encoding",
	}))

	app.Get("/api/", h.getAllCars)
	app.Get("/api/:id", h.getCarById)
	app.Post("/api/", h.addCar)
	app.Delete("/api/:id", h.deleteCarById)
	app.Patch("/api/:id", h.updateCar)

	app.Get("/swagger/*", swagger.HandlerDefault)

}
