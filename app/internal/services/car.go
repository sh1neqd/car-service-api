package services

import (
	"car-service1/app/internal/domain/car"
	"car-service1/app/internal/repositories"
	"fmt"
	"github.com/sirupsen/logrus"
)

type CarService struct {
	repo repositories.Car
}

func (s CarService) Update(id int, dto car.UpdateCarDTO) error {
	return s.repo.Update(id, dto)
}

func (s CarService) GetAll() ([]car.Car, error) {
	return s.repo.GetAll()
}

func (s CarService) GetById(id int) (car.Car, error) {
	return s.repo.GetById(id)
}

func (s CarService) Create(dto car.CarDTO) error {
	fmt.Println(dto.RegNum)
	err := s.repo.Create(dto)
	if err != nil {
		logrus.Errorf("failed to add car: %v", err)
	}
	return err
}

func (s CarService) Delete(id int) error {
	return s.repo.Delete(id)
}

func NewCarService(repo repositories.Car) *CarService {
	return &CarService{repo: repo}
}
