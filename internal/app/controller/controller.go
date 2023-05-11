package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-observability-example/internal/app/service"
)

const _plateParam = "plate"

type Params struct {
	Router  *gin.Engine
	Service service.Service
}

type controller struct {
	router  *gin.Engine
	service service.Service
}

func New(p Params) Controller {
	ctrl := &controller{
		router:  p.Router,
		service: p.Service,
	}
	defer ctrl.initializeRoutes()

	return ctrl
}

func (c *controller) CreateVehicle(ctx *gin.Context) {
	var req service.CreateVehicleRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.service.CreateVehicle(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"plate": resp})
}

func (c *controller) GetVehicleByPlate(ctx *gin.Context) {
	plate := ctx.Param(_plateParam)
	if plate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "plate is empty"})
		return
	}

	resp, err := c.service.GetVehicleByPlate(ctx, plate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
