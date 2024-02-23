package postgres

import (
	"cars-stucture/repository"

	"gorm.io/gorm"
)

type POSTGRES struct {
	db *gorm.DB
}

var _ repository.Cars = (*POSTGRES)(nil)

var DB *gorm.DB

func NewDatabase(db *gorm.DB) *POSTGRES {
	if db == nil {
		panic("postgres: database is nil")
	}
	return &POSTGRES{db}
}
