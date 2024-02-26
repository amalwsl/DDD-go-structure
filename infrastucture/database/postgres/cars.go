package postgres

import (
	"cars-stucture/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func (cdb *POSTGRES) AddCar(car domain.Car) error {
	result := cdb.db.Create(&car)
	if result.Error != nil {
		log.Printf("Error inserting car: %v", result.Error)
		return result.Error
	}
	return nil
}

func (cdb *POSTGRES) GetCarByRegistration(registration string) (*domain.Car, error) {
	var car domain.Car
	result := cdb.db.Where("registration = ?", registration).First(&car)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("car with registration %s not found", registration)
		}
		log.Printf("Error fetching car: %v", result.Error)
		return nil, result.Error
	}
	return &car, nil
}

func (cdb *POSTGRES) UpdateCar(registration string, mileage int, rented bool) error {
	result := cdb.db.Model(&domain.Car{}).Where("registration = ?", registration).Updates(map[string]interface{}{"mileage": mileage, "rented": rented})
	if result.Error != nil {
		log.Printf("Error updating car: %v", result.Error)
		return result.Error
	}
	return nil
}

func (cdb *POSTGRES) DeleteCar(registration string) error {
	result := cdb.db.Where("registration = ?", registration).Delete(&domain.Car{})
	if result.Error != nil {
		log.Printf("Error deleting car: %v", result.Error)
		return result.Error
	}
	return nil
}

func (cdb *POSTGRES) GetAllCars() ([]domain.Car, error) {
	var cars []domain.Car
	result := cdb.db.Find(&cars)
	if result.Error != nil {
		log.Printf("Error querying cars: %v", result.Error)
		return nil, result.Error
	}
	return cars, nil
}
