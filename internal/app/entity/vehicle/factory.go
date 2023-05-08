package vehicle

import (
	"errors"
	"go-observability-example/internal/app/domain"
	"go-observability-example/internal/app/entity/vehicle/models/audi"
	"go-observability-example/internal/app/entity/vehicle/models/bmw"
)

var errorInvalidVehicleModel = errors.New("invalid vehicle model")

func Factory(model domain.VehicleModel) (*Vehicle, error) {
	switch model {
	case bmw.Model:
		return bmw.NewVehicle(), nil
	case audi.Model:
		return audi.NewVehicle(), nil
	default:
		return nil, errorInvalidVehicleModel
	}
}
