package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalRequests = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "myapp_http_requests_total",
		Help: "Number of get requests.",
	},
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == prometheusMetricPath {
			c.Next()
			return
		}
		totalRequests.Inc()
	}
}
