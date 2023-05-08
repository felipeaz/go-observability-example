package audi

import (
	"fmt"

	"go-observability-example/internal/app/domain"
	"go-observability-example/internal/app/entity/vehicle/models/plate"
)

const (
	Model                domain.VehicleModel = "AUDI"
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
				State: domain.FlatWheel,
				Size:  _wheelSize,
			},
			{
				State: domain.FlatWheel,
				Size:  _wheelSize,
			},
			{
				State: domain.FlatWheel,
				Size:  _wheelSize,
			},
			{
				State: domain.FlatWheel,
				Size:  _wheelSize,
			},
		},
		FuelTank: domain.FuelTank{
			AcceptsGasoline: _fuelAcceptsGasoline,
			AcceptsAlcohol:  _fuelAcceptsAlcohol,
			State:           domain.FuelFull,
		},
		Model: Model,
	}
}
