package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-observability-example/build/app"
	"go-observability-example/internal/app/controller"
	"go-observability-example/internal/app/service"
)

func main() {
	ginEngine := gin.Default()

	appController := controller.New(controller.ControllerDeps{
		Router: ginEngine,
		Service: service.New(
			service.ServiceDeps{},
		),
	})

	serviceApp := app.Build(app.AppDeps{
		Router:     ginEngine,
		Controller: appController,
	})

	err := serviceApp.Run(":8080")
	if err != nil {
		log.Fatal("failed to run app", err)
	}
}
