package domain

var (
	ErrGetCars            = "Failed to retrieve available cars"
	ErrProcessingCars     = "Failed to process car data"
	ErrEncodeJSON         = "Failed to encode JSON response"
	ErrInvalidRequestBody = "Invalid request body"
	ErrAddingCars         = "Failed to add car"
	ErrCarUnavailable     = "Car is already rented"
	ErrUpdatingStatus     = "Failed to update car rental status"

	ErrCarNotFound = "Car not found"
	ErrReturnCar   = "Car was not rented"

	ErrInvalidMileage = "Invalid mileage"

	ErrUpdatingCar = "Failed to update car data"
)
