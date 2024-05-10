package services

import (
	"car-service1/app/internal/domain/car"
	"car-service1/app/internal/repositories"
)

type Car interface {
	GetAll() ([]car.Car, error)
	GetById(id int) (car.Car, error)
	Create(dto car.CarDTO) error
	Delete(id int) error
	Update(id int, dto car.UpdateCarDTO) error
}

type Service struct {
	Car
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Car: NewCarService(repos.Car),
	}
}
