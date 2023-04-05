package domain

type Vehicle interface {
	Drive()
	Doors() int
	MaxSpeed() float64
	Wheels() []Wheel
	Fuel() FuelTank
}
