package metrics

// 注册所有指标
func RegisterAll() {
	RegisterCounters()
	RegisterGauges()
	RegisterHistograms()
	RegisterSummaries()
}