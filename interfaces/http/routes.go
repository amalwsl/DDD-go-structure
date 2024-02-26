package http

import (
	"cars-stucture/domain"
	"fmt"
	"net/http"
)

type handler struct {
	service domain.Service
}

func NewHandler() *handler {
	return &handler{}
}

func RegisterHandlers(service domain.Service, mux *http.ServeMux) {
	fmt.Println("here u r in routes ")
	h := &handler{service}

	mux.Handle("GET /cars", h.handleGetAllCars())
	mux.Handle("GET /cars/{registration}", h.handleGetCarByRegistration())
	mux.Handle("POST /cars", h.handleAddCar())
	mux.Handle("PUT /cars/{registration}", h.handleUpdateCar())
	mux.Handle("DELETE /cars/{registration}", h.handleDeleteCar())
}
