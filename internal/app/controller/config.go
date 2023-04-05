package controller

func (c *controller) Initialize() Controller {
	c.InitializeRoutes()
	return c
}
