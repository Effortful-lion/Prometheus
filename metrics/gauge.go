package metrics

import "github.com/prometheus/client_golang/prometheus"

// 当前连接数
var CurrentConnections = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "current_connections",
		Help: "Current number of connections",
	},
)

// 注册仪表盘指标
func RegisterGauges() {
	prometheus.MustRegister(CurrentConnections)
}