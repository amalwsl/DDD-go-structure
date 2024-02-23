package domain

type Service interface {
	AddCar(Car) error
	GetCarByRegistration(registration string) (*Car, error)
	UpdateCar(registration string, mileage int, rented bool) error
	DeleteCar(registration string) error
	GetAllCars() ([]Car, error)
}
