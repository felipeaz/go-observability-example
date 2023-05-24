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
	router        *gin.Engine
	successMetric prometheus.Counter
	errorMetric   prometheus.Counter
}

func New(p Params) Controller {
	ctrl := &controller{
		router: p.Router,
	}
	defer ctrl.initializeRoutes()

	ctrl.successMetric = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_success_ops_total",
			Help: "total success operation",
		})

	ctrl.errorMetric = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_error_ops_total",
			Help: "total operation errors",
		})

	return ctrl
}

func (c *controller) Success(ctx *gin.Context) {
	c.successMetric.Inc()
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *controller) Error(ctx *gin.Context) {
	c.errorMetric.Inc()
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": errTest.Error()})
}
