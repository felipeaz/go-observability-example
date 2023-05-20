package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Router *gin.Engine
}

type controller struct {
	router *gin.Engine
}

func New(p Params) Controller {
	ctrl := &controller{
		router: p.Router,
	}
	defer ctrl.initializeRoutes()

	return ctrl
}

func (c *controller) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *controller) Error(ctx *gin.Context) {
	err := errors.New("test error")
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
}
