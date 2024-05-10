package main

import (
	"car-service1/app/internal/app"
	"car-service1/app/internal/config"
)

// @title car-service
// @version 1.0
// @description Тестовое задание на Go

// @host      localhost:8000
// @BasePath  /

func main() {
	app.StartApp(config.GetConfig())
}
