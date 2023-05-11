package service

import (
	"context"

	"go-observability-example/internal/app/domain"
)

type Service interface {
	CreateVehicle(ctx context.Context, request CreateVehicleRequest) (plate string, err error)
	GetVehicleByPlate(ctx context.Context, plate string) (domain.Vehicle, error)
}
