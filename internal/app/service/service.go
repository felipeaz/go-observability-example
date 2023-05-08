package service

import (
	"context"
	"encoding/json"
	"go-observability-example/internal/app/entity/vehicle"

	"go-observability-example/internal/app/domain"
	"go-observability-example/internal/app/storage"
)

type Params struct {
	Storage storage.Storage
}

type service struct {
	storage storage.Storage
}

func New(p Params) Service {
	return &service{
		storage: p.Storage,
	}
}

func (s *service) CreateVehicle(ctx context.Context, request CreateVehicleRequest) (plate string, err error) {
	v, err := vehicle.Factory(request.VehicleModel)
	if err != nil {
		return "", err
	}

	err = s.storage.Set(v.Plate, v)
	if err != nil {
		return "", err
	}

	return v.Plate, nil
}

func (s *service) GetVehicleByPlate(ctx context.Context, plate string) (domain.VehicleInterface, error) {
	var vobj *domain.Vehicle
	vbytes, err := s.storage.Get(plate)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(vbytes, &vobj)
	if err != nil {
		return nil, err
	}

	return vobj, nil
}
