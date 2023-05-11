package domain

import "log"

type VehicleModel string

type VehicleInterface interface {
	Drive()
	GetDoors() int
	GetPlate() string
	GetMaxSpeed() float64
	GetWheels() []Wheel
	GetFuel() FuelTank
}

type Vehicle struct {
	Plate    string       `json:"plate"`
	Doors    int          `json:"doors"`
	MaxSpeed float64      `json:"maxSpeed"`
	Wheels   []Wheel      `json:"wheels"`
	FuelTank FuelTank     `json:"fuelTank"`
	Model    VehicleModel `json:"model"`
}

func (v *Vehicle) Drive() {
	log.Print()
}

func (v *Vehicle) GetDoors() int {
	return v.Doors
}

func (v *Vehicle) GetPlate() string {
	return v.GetPlate()
}

func (v *Vehicle) GetMaxSpeed() float64 {
	return v.MaxSpeed
}

func (v *Vehicle) GetWheels() []Wheel {
	return v.Wheels
}

func (v *Vehicle) GetFuel() FuelTank {
	return v.FuelTank
}
