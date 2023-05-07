package main

import (
	"go-observability-example/infra/redis"
	"log"

	"github.com/gin-gonic/gin"
	"go-observability-example/build/app"
	"go-observability-example/internal/app/controller"
	"go-observability-example/internal/app/service"
)

func main() {
	ginEngine := gin.Default()

	redisServer := redis.New("localhost", "6379")

	appController := controller.New(controller.Params{
		Router: ginEngine,
		Service: service.New(
			service.Params{
				Storage: redisServer,
			},
		),
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
