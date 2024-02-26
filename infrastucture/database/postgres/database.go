package postgres

import (
	"cars-stucture/domain"
	"cars-stucture/repository"
	"log"

	"gorm.io/gorm"
)

type POSTGRES struct {
	db *gorm.DB
}

var _ repository.Cars = (*POSTGRES)(nil)

var DB *gorm.DB

func NewDatabase(db *gorm.DB) *POSTGRES {

	// autoMigrate the Car struct to create the "cars" table if it doesn't exist
	err := db.AutoMigrate(&domain.Car{})
	if err != nil {
		log.Fatalf("Error auto migrating table: %v", err)
	}

	return &POSTGRES{db: db}
}
