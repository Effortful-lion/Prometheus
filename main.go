package main

import (
	"fmt"
	"log"
	"net/http"

	"Prometheus/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 注册指标
func init() {
	metrics.RegisterAll()
}

func main() {
	// 设置指标暴露端点
	http.Handle("/metrics", promhttp.Handler())

	RegisterHandlers()

	// 启动服务器
	port := 8080
	log.Printf("Server starting on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func RegisterHandlers() {
	RegisterExampleHandler()
}

func RegisterExampleHandler() {
	// 添加一个示例端点
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 增加请求计数
		metrics.HTTPRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()

		// 模拟响应时间
		metrics.ResponseTime.Observe(0.1)
		metrics.ResponseTimeSummary.Observe(0.1)

		// 更新当前连接数
		metrics.CurrentConnections.Inc()
		defer metrics.CurrentConnections.Dec()

		fmt.Fprintf(w, "Hello, Prometheus!\n")
	})
}
