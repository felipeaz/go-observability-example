package controller

func (c *controller) initializeRoutes() {
	vehicles := c.router.Group("/api")

	v1 := vehicles.Group("/v1")
	v1.GET("/success", c.Success)
	v1.GET("/error", c.Error)
}
