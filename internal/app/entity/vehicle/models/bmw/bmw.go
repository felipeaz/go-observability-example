package bmw

import (
	"fmt"

	"go-observability-example/internal/app/domain"
	"go-observability-example/internal/app/entity/vehicle/models/plate"
)

const (
	Model                domain.VehicleModel = "BMW"
	_plateFmt                                = "%s-%d"
	_doors                                   = 4
	_maxSpeed                                = 300
	_wheelSize                               = 21
	_fuelAcceptsGasoline                     = true
	_fuelAcceptsAlcohol                      = false
)

func NewVehicle() *domain.Vehicle {
	return &domain.Vehicle{
		Plate:    fmt.Sprintf(_plateFmt, Model, plate.New()),
		Doors:    _doors,
		MaxSpeed: _maxSpeed,
		Wheels: []domain.Wheel{
			{
				State: domain.FullWheel,
				Size:  _wheelSize,
			},
			{
				State: domain.FullWheel,
				Size:  _wheelSize,
			},
			{
				State: domain.FullWheel,
				Size:  _wheelSize,
			},
			{
				State: domain.FullWheel,
				Size:  _wheelSize,
			},
		},
		FuelTank: domain.FuelTank{
			AcceptsGasoline: _fuelAcceptsGasoline,
			AcceptsAlcohol:  _fuelAcceptsAlcohol,
			State:           domain.FuelEmpty,
		},
		Model: Model,
	}
}
