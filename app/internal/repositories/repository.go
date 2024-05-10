package repositories

import (
	"car-service1/app/internal/domain/car"
	"github.com/jmoiron/sqlx"
)

type Car interface {
	GetAll() ([]car.Car, error)
	GetById(id int) (car.Car, error)
	Create(car car.CarDTO) error
	Delete(id int) error
	Update(id int, dto car.UpdateCarDTO) error
}

type Repository struct {
	Car
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Car: NewCarRepo(db),
	}
}
