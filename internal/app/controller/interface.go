package controller

import "go-observability-example/internal/app/domain"

type Controller interface {
	GetVehicle() domain.Vehicle
}
