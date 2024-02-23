package http

import (
	"cars-stucture/domain"
	"net/http"
)

type handler struct {
	service domain.Service
}

func NewHandler() *handler {
	return &handler{}
}

func RegisterHandlers(service domain.Service) {
	h := &handler{service}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cars", h.getAllCars)
	mux.HandleFunc("GET /cars/{registration}", h.getCarByRegistration)
	mux.HandleFunc("POST /cars", h.addCar)
	mux.HandleFunc("PUT /cars/{registration}", h.updateCar)
	mux.HandleFunc("DELETE /cars/{registration}", h.deleteCar)
}
