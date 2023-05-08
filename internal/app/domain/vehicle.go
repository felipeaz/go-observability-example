package domain

type VehicleModel string

var (
	BMW  = "BMW"
	Audi = "AUDI"
)

type Vehicle interface {
	Drive()
	Doors() int
	Plate() string
	MaxSpeed() float64
	Wheels() []Wheel
	Fuel() FuelTank
}
