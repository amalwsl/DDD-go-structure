package http

import (
	"cars-stucture/domain"
	"encoding/json"
	"net/http"
)

func (h *handler) getAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.service.GetAllCars()
	if err != nil {
		http.Error(w, "Failed to retrieve cars", http.StatusInternalServerError)
		return
	}

	// Serialize cars to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func (h *handler) getCarByRegistration(w http.ResponseWriter, r *http.Request) {
	registration := r.URL.Query().Get("registration")
	if registration == "" {
		http.Error(w, "Registration parameter is required", http.StatusBadRequest)
		return
	}

	car, err := h.service.GetCarByRegistration(registration)
	if err != nil {
		http.Error(w, "Failed to retrieve car", http.StatusInternalServerError)
		return
	}

	// Serialize car to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func (h *handler) addCar(w http.ResponseWriter, r *http.Request) {
	var car domain.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.AddCar(car); err != nil {
		http.Error(w, "Failed to add car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) updateCar(w http.ResponseWriter, r *http.Request) {
	registration := r.URL.Query().Get("registration")
	if registration == "" {
		http.Error(w, "Registration parameter is required", http.StatusBadRequest)
		return
	}

	var updatedCar domain.Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCar(registration, updatedCar.Mileage, updatedCar.Rented); err != nil {
		http.Error(w, "Failed to update car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) deleteCar(w http.ResponseWriter, r *http.Request) {
	registration := r.URL.Query().Get("registration")
	if registration == "" {
		http.Error(w, "Registration parameter is required", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCar(registration); err != nil {
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
