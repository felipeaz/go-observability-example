package service

import (
	"context"

	"go-observability-example/internal/app/domain"
)

type Params struct {
}

type service struct {
}

func New(p Params) Service {
	return &service{}
}

func (s *service) CreateVehicle(ctx context.Context) (plate string, err error) {
	return "", nil
}

func (s *service) GetVehicle(ctx context.Context) ([]domain.Vehicle, error) {
	return nil, nil
}

func (s *service) GetVehicleByPlate(ctx context.Context, plate string) (domain.Vehicle, error) {
	return nil, nil
}
