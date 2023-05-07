package controller

func (c *controller) initializeRoutes() {
	vehicles := c.router.Group("/api")

	v1 := vehicles.Group("/v1")
	v1.POST("/vehicle", c.CreateVehicle)
	v1.GET("/vehicle", c.GetVehicle)
	v1.GET("/vehicle/:plate", c.GetVehicleByPlate)
}
