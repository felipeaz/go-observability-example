package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-observability-example/internal/app/service"
)

type ControllerDeps struct {
	Router  *gin.Engine
	Service service.Service
}

type controller struct {
	router  *gin.Engine
	service service.Service
}

func New(deps ControllerDeps) Controller {
	c := controller{
		router:  deps.Router,
		service: deps.Service,
	}
	return c.Initialize()
}

func (c *controller) CreateVehicle(ctx *gin.Context) {
	resp, err := c.service.CreateVehicle(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"plate": resp})
}

func (c *controller) GetVehicle(ctx *gin.Context) {
	resp, err := c.service.GetVehicle(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (c *controller) GetVehicleByPlate(ctx *gin.Context) {
	plate := ctx.Param("plate")
	if plate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "plate is empty"})
		return
	}

	resp, err := c.service.GetVehicleByPlate(ctx, plate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
