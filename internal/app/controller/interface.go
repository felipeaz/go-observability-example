package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateVehicle(ctx *gin.Context)
	GetVehicleByPlate(ctx *gin.Context)
}
