package application

import (
	"cars-stucture/domain"
	"errors"
)

// CarService represents the service layer for managing cars.
type CarService struct {
	carRepo domain.Service
}

// NewCarService creates a new instance of CarService.
func NewCarService(carRepo domain.Service) *CarService {
	return &CarService{carRepo: carRepo}
}

// AddCar adds a new car.
func (cs *CarService) AddCar(car domain.Car) error {
	// Validate input parameters
	if car.Model == "" || car.Registration == "" {
		return errors.New("model and registration are required")
	}

	// Call repository method to add car
	return cs.carRepo.AddCar(car)
}

// GetCarByRegistration retrieves a car by its registration number.
func (cs *CarService) GetCarByRegistration(registration string) (*domain.Car, error) {
	// Call repository method to get car
	return cs.carRepo.GetCarByRegistration(registration)
}

// UpdateCar updates the details of a car.
func (cs *CarService) UpdateCar(registration string, mileage int, rented bool) error {
	// Call repository method to update car
	return cs.carRepo.UpdateCar(registration, mileage, rented)
}

// DeleteCar deletes a car by its registration number.
func (cs *CarService) DeleteCar(registration string) error {
	// Call repository method to delete car
	return cs.carRepo.DeleteCar(registration)
}

// GetAllCars retrieves all cars.
func (cs *CarService) GetAllCars() ([]domain.Car, error) {
	// Call repository method to get all cars
	return cs.carRepo.GetAllCars()
}

// GetAvailableCars retrieves all available cars.
func (cs *CarService) GetAvailableCars() ([]domain.Car, error) {
	// Call repository method to get available cars
	cars, err := cs.carRepo.GetAllCars()
	if err != nil {
		return nil, err
	}

	var availableCars []domain.Car
	for _, car := range cars {
		if !car.Rented {
			availableCars = append(availableCars, car)
		}
	}
	return availableCars, nil
}
