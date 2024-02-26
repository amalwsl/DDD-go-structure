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

// Repository represents the data access layer for managing cars.
type Repository struct {
	storage map[string]domain.Car
}

// NewRepository creates a new instance of Repository.
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string]domain.Car),
	}
}

// AddCar adds a new car to the repository.
func (r *Repository) AddCar(car domain.Car) error {
	if _, exists := r.storage[car.Registration]; exists {
		return errors.New("car already exists")
	}
	r.storage[car.Registration] = car
	return nil
}

// GetCarByRegistration retrieves a car by its registration number from the repository.
func (r *Repository) GetCarByRegistration(registration string) (*domain.Car, error) {
	car, exists := r.storage[registration]
	if !exists {
		return nil, errors.New("car not found")
	}
	return &car, nil
}

// UpdateCar updates the details of a car in the repository.
func (r *Repository) UpdateCar(registration string, mileage int, rented bool) error {
	car, exists := r.storage[registration]
	if !exists {
		return errors.New("car not found")
	}
	car.Mileage = mileage
	car.Rented = rented
	r.storage[registration] = car
	return nil
}

// DeleteCar deletes a car by its registration number from the repository.
func (r *Repository) DeleteCar(registration string) error {
	if _, exists := r.storage[registration]; !exists {
		return errors.New("car not found")
	}
	delete(r.storage, registration)
	return nil
}

// GetAllCars retrieves all cars from the repository.
func (r *Repository) GetAllCars() ([]domain.Car, error) {
	var cars []domain.Car
	for _, car := range r.storage {
		cars = append(cars, car)
	}
	return cars, nil
}
