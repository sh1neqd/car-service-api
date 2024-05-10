package handlers

import (
	"car-service1/app/internal/domain/car"
)

func filterCars(cars []car.Car, filter string) []car.Car {
	if filter == "" {
		return cars
	}

	filteredCars := make([]car.Car, 0)
	for _, c := range cars {
		if c.RegNum == filter || c.Mark == filter || c.Model == filter {
			filteredCars = append(filteredCars, c)
		}
	}

	return filteredCars
}

func paginateCars(cars []car.Car, limit, offset int) []car.Car {
	if offset >= len(cars) {
		return []car.Car{}
	}

	end := offset + limit
	if end > len(cars) {
		end = len(cars)
	}

	return cars[offset:end]
}
