package GraphiteBase

import (
	"fmt"
)

// The time and measured value of a metric.
type MetricValues struct {
	Time  int64
	Value float64
}

// A whole metric, whith name, time and value.
type Metric struct {
	MetricValues
	Name string
}

// Create a new metric from the given parameters.
func NewMetric(name string, value float64, timestamp int64) *Metric {
	return &Metric{
		Name: name,
		MetricValues: MetricValues{
			Value: value,
			Time:  timestamp,
		},
	}
}

// Stringifies metric with "<Name> <Value> <Time>", to match Graphites Text
// protocol.
func (m *Metric) String() string {
	return fmt.Sprintf("%s %v %d", m.Name, m.Value, m.Time)
}
