package app

import (
	"car-service1/app/internal/config"
	"car-service1/app/internal/handlers"
	"car-service1/app/internal/repositories"
	"car-service1/app/internal/services"
	"car-service1/app/pkg/client/postgresql"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func newApp() (*fiber.App, error) {
	return fiber.New(), nil
}

func StartApp(config *config.Config) {
	app, err := newApp()

	logrus.Println("postgresql initializing")
	db, err := postgresql.NewPostgresDB(config)
	if err != nil {
		logrus.Errorf("failed to initilize db: %v", err)
	}

	logrus.Println("initialize repos, services and handlers")
	repos := repositories.NewRepository(db)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)

	logrus.Println("initialize routes")
	handler.InitRoutes(app)
	if err != nil {
		logrus.Errorf("failed to initilize routes: %v", err)
	}

	err = app.Listen(":8000")
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
