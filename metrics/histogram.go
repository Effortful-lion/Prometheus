package metrics

import "github.com/prometheus/client_golang/prometheus"

// 响应时间直方图
var ResponseTime = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "response_time_seconds",
		Help:    "Response time in seconds",
		Buckets: prometheus.DefBuckets,
	},
)

// 注册直方图指标
func RegisterHistograms() {
	prometheus.MustRegister(ResponseTime)
}