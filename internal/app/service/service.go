package service

import (
	"context"

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
	return "", nil
}

func (s *service) GetVehicle(ctx context.Context) ([]domain.Vehicle, error) {
	return nil, nil
}

func (s *service) GetVehicleByPlate(ctx context.Context, plate string) (domain.Vehicle, error) {
	return nil, nil
}
