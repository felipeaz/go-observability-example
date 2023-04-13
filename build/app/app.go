package app

import (
	"github.com/gin-gonic/gin"

	"go-observability-example/api/prometheus"
	"go-observability-example/internal/app/controller"
)

type AppDeps struct {
	Router     *gin.Engine
	Controller controller.Controller
}

type App interface {
	Run(addr string) error
}

type app struct {
	deps AppDeps
}

func Build(deps AppDeps) App {
	prometheus.RegisterMetrics(deps.Router)

	return &app{
		deps: deps,
	}
}

func (a *app) Run(addr string) error {
	return a.deps.Router.Run(addr)
}
