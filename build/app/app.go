package app

import (
	"github.com/gin-gonic/gin"

	"go-observability-example/api/prometheus"
	"go-observability-example/internal/app/controller"
)

type Params struct {
	Router     *gin.Engine
	Controller controller.Controller
}

type App interface {
	Run(addr string) error
}

type app struct {
	router     *gin.Engine
	controller controller.Controller
}

func Build(p Params) App {
	prometheus.RegisterMetrics(p.Router)
	return &app{
		router:     p.Router,
		controller: p.Controller,
	}
}

func (a *app) Run(addr string) error {
	return a.router.Run(addr)
}
