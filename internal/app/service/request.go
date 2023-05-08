package service

import "go-observability-example/internal/app/domain"

type CreateVehicleRequest struct {
	VehicleModel domain.VehicleModel `json:"vehicleModel"`
}
