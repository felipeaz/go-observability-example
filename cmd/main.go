package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-observability-example/build/app"
	"go-observability-example/internal/app/controller"
)

func main() {
	ginEngine := gin.Default()
	appController := controller.New(controller.Params{
		Router: ginEngine,
	})

	serviceApp := app.Build(app.Params{
		Router:     ginEngine,
		Controller: appController,
	})

	err := serviceApp.Run(":8080")
	if err != nil {
		log.Fatal("failed to run app", err)
	}
}
