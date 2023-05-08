package domain

type VehicleModel string

type Vehicle interface {
	Drive()
	GetDoors() int
	GetPlate() string
	GetMaxSpeed() float64
	GetWheels() []Wheel
	GetFuel() FuelTank
}
