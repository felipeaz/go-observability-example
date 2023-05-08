package vehicle

import (
	"go-observability-example/internal/app/domain"
	"log"
)

type Vehicle struct {
	Plate    string
	Doors    int
	MaxSpeed float64
	Wheels   []domain.Wheel
	FuelTank domain.FuelTank
	Model    domain.VehicleModel
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

func (v *Vehicle) GetWheels() []domain.Wheel {
	return v.Wheels
}

func (v *Vehicle) GetFuel() domain.FuelTank {
	return v.FuelTank
}
