package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "myapp_http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == prometheusMetricPath {
			c.Next()
			return
		}
		totalRequests.WithLabelValues("path").Inc()
	}
}
