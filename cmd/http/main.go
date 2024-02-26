package main

import (
	"cars-stucture/application"
	"cars-stucture/infrastucture/database/postgres"
	"cars-stucture/interfaces/http"
	"log"
	ghttp "net/http"

	postgres_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	var err error
	dsn := "host=surus.db.elephantsql.com user=hrjdepqm password=FoNrkYIQipb_YSXG3Is6wKDoYTb8ig9F dbname=hrjdepqm port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres_.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	repo := postgres.NewDatabase(DB)
	service := application.NewCarService(repo)

	mux := ghttp.NewServeMux()

	http.RegisterHandlers(service, mux)

	if err := ghttp.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
