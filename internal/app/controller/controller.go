package controller

import (
	"github.com/gin-gonic/gin"

	"go-observability-example/internal/app/service"
)

type Args struct {
	Router  *gin.Engine
	Service service.Service
}

type controller struct {
	router  *gin.Engine
	service service.Service
}

func New(args Args) Controller {
	c := controller{
		router:  args.Router,
		service: args.Service,
	}
	return c.Initialize()
}

func (c *controller) CreateVehicle(ctx *gin.Context) {

}

func (c *controller) GetVehicle(ctx *gin.Context) {

}

func (c *controller) GetVehicleByPlate(ctx *gin.Context) {

}
