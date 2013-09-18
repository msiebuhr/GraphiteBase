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
func NewMetric(name string, timestamp int64, value float64) *Metric {
	m := &Metric{}
	m.Name = name
	m.Time = timestamp
	m.Value = value
	return m
}

// Stringifies metric with "<Name> <Time> <Value>", to match Graphites Text
// protocol.
func (m *Metric) String() string {
	return fmt.Sprintf("%s %d %v", m.Name, m.Time, m.Value)
}
