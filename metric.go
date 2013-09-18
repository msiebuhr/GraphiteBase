package GraphiteBase

import (
	"fmt"
)

// Data structures
type MetricValues struct {
	Time  int64
	Value float64
}

type Metric struct {
	MetricValues
	Name string
}

func NewMetric(name string, timestamp int64, value float64) *Metric {
	m := &Metric{}
	m.Name = name
	m.Time = timestamp
	m.Value = value
	return m
}

// Plain stringification
func (m *Metric) String() string {
	return fmt.Sprintf("%s %d %v", m.Name, m.Time, m.Value)
}
