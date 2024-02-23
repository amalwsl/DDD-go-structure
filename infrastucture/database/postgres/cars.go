package postgres

import (
	"cars-stucture/domain"
	"log"
)

func (cdb *POSTGRES) AddCar(car domain.Car) error {
	_, err := cdb.db.Exec("INSERT INTO cars (model, registration, mileage, rented) VALUES (?, ?, ?, ?)", car.Model, car.Registration, car.Mileage, car.Rented)
	if err != nil {
		log.Printf("Error inserting car: %v", err)
		return err
	}
	return nil
}

func (cdb *POSTGRES) GetCarByRegistration(registration string) (*domain.Car, error) {
	var car domain.Car
	row := cdb.db.QueryRow("SELECT model, registration, mileage, rented FROM cars WHERE registration = ?", registration)
	err := row.Scan(&car.Model, &car.Registration, &car.Mileage, &car.Rented)
	if err != nil {
		log.Printf("Error fetching car: %v", err)
		return nil, err
	}
	return &car, nil
}

func (cdb *POSTGRES) UpdateCar(registration string, mileage int, rented bool) error {
	_, err := cdb.db.Exec("UPDATE cars SET mileage = ?, rented = ? WHERE registration = ?", mileage, rented, registration)
	if err != nil {
		log.Printf("Error updating car: %v", err)
		return err
	}
	return nil
}

func (cdb *POSTGRES) DeleteCar(registration string) error {
	_, err := cdb.db.Exec("DELETE FROM cars WHERE registration = ?", registration)
	if err != nil {
		log.Printf("Error deleting car: %v", err)
		return err
	}
	return nil
}

func (cdb *POSTGRES) GetAllCars() ([]domain.Car, error) {
	var cars []domain.Car
	rows, err := cdb.db.Query("SELECT model, registration, mileage, rented FROM cars")
	if err != nil {
		log.Printf("Error querying cars: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car domain.Car
		err := rows.Scan(&car.Model, &car.Registration, &car.Mileage, &car.Rented)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
