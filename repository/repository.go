package repository

import "cars-stucture/domain"

type Cars interface {
	AddCar(car domain.Car) error
	GetCarByRegistration(registration string) (*domain.Car, error)
	UpdateCar(registration string, mileage int, rented bool) error
	DeleteCar(registration string) error
	GetAllCars() ([]domain.Car, error)
}
