package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	prometheusMetricPath = "/metrics"
)

func RegisterMetrics(router *gin.Engine) {
	router.GET(prometheusMetricPath, gin.WrapH(promhttp.Handler()))
}
