package domain

type Vehicle interface {
	Drive()
	Doors() int
	Plate() string
	MaxSpeed() float64
	Wheels() []Wheel
	Fuel() FuelTank
}
