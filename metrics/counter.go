package metrics

import "github.com/prometheus/client_golang/prometheus"

// HTTP请求计数
var HTTPRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "path", "status"},
)

// 注册计数器指标
func RegisterCounters() {
	prometheus.MustRegister(HTTPRequestsTotal)
}