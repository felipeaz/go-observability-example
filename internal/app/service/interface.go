package service

import (
	"context"

	"go-observability-example/internal/app/domain"
)

type Service interface {
	CreateVehicle(ctx context.Context) (plate string, err error)
	GetVehicle(ctx context.Context) ([]domain.Vehicle, error)
	GetVehicleByPlate(ctx context.Context, plate string) (domain.Vehicle, error)
}
