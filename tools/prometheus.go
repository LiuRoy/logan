package tools

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	historyBuckets = [...]float64{
		10., 20., 30., 50., 80., 100., 200., 300., 500., 1000., 2000., 3000.}
	DefaultMetricPath string = "/kingkong/metrics"

	ResponseCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kingkong_request_total",
		Help: "Total request counts"}, []string{"method", "endpoint"})
	ErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kingkong_error_total",
		Help: "Total Error counts"}, []string{"method", "endpoint"})
	ResponseLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "kingkong_response_latency_millisecond",
		Help: "Response latency (millisecond)",
		Buckets: historyBuckets[:]}, []string{"method", "endpoint"})
)

func init() {
	prometheus.MustRegister(ResponseCounter)
	prometheus.MustRegister(ErrorCounter)
	prometheus.MustRegister(ResponseLatency)
}

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		endPoint := c.Request.URL.String()
		if endPoint != DefaultMetricPath {
			c.Next()
		} else {
			start := time.Now()
			method := c.Request.Method

			c.Next()

			elapsed := float64(time.Since(start)) / float64(time.Millisecond)
			ResponseCounter.WithLabelValues(method, endPoint).Inc()
			ResponseLatency.WithLabelValues(method, endPoint).Observe(elapsed)
		}

	}
}

func LatestMetrics(c *gin.Context) {
	handler := promhttp.Handler()
	handler.ServeHTTP(c.Writer, c.Request)
}
