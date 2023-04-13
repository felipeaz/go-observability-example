package service

import (
	"context"

	"go-observability-example/internal/app/domain"
)

type ServiceDeps struct {
}

type service struct {
}

func New(deps ServiceDeps) Service {
	return &service{}
}

func (s *service) CreateVehicle(ctx context.Context) (plate string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetVehicle(ctx context.Context) ([]domain.Vehicle, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetVehicleByPlate(ctx context.Context, plate string) (domain.Vehicle, error) {
	//TODO implement me
	panic("implement me")
}
