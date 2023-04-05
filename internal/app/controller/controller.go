package controller

import (
	"github.com/gin-gonic/gin"
	"go-observability-example/internal/app/domain"
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
	return c
}

func (c controller) GetVehicle() domain.Vehicle {
	//TODO implement me
	panic("implement me")
}
