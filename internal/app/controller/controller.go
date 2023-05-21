package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	errTest      = errors.New("test error")
	promRegistry = prometheus.NewRegistry()
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
	successCounter := promauto.With(promRegistry).NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_success_ops_total",
			Help: "total success operation",
		})
	successCounter.Inc()
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *controller) Error(ctx *gin.Context) {
	errorCounter := promauto.With(promRegistry).NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_error_ops_total",
			Help: "total operation errors",
		})
	errorCounter.Inc()
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": errTest.Error()})
}
