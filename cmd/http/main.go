package main

import (
	"cars-stucture/application"
	"cars-stucture/infrastucture/database/postgres"
	"cars-stucture/interfaces/http"
	"log"
	ghttp "net/http"
	"os"

	postgres_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no settings file provided")
	}

	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres_.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	repo := postgres.NewDatabase(DB)
	service := application.NewCarService(repo)

	http.RegisterHandlers(service)

	if err := ghttp.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
