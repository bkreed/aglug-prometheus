type Counter interface {
	Metric
	Collector
	Set(float64)
	Inc()
	Add(float64)
}

type Gauge interface {
	Metric
	Collector
	Set(float64)
	Inc()
	Dec()
	Add(float64)
	Sub(float64)
}

type Histogram interface {
	Metric
	Collector
	// Observe adds a single observation to the histogram.
	Observe(float64)
}

type Summary interface {
	Metric
	Collector
	// Observe adds a single observation to the summary.
	Observe(float64)
}
