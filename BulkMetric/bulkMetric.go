package BulkMetric

import (
	"github.com/msiebuhr/GraphiteBase"
)

type BulkMetric struct {
	Name string
	Data []*GraphiteBase.MetricValues
}

// Return a new BulkMetric with the given name.
func NewBulkMetric(name string) *BulkMetric {
	n := &BulkMetric{}
	n.Name = name
	n.Data = make([]*GraphiteBase.MetricValues, 0)
	return n
}

// Roll an existing metric into the BulkMetric.
func (b *BulkMetric) AddMetric(metric *GraphiteBase.Metric) {
	b.Data = append(
		b.Data,
		&GraphiteBase.MetricValues{Time: metric.Time, Value: metric.Value},
	)
}

// Un-roll the bulk metric into a list of the given metrics
func (b *BulkMetric) GetMetrics() []*GraphiteBase.Metric {
	out := make([]*GraphiteBase.Metric, len(b.Data))
	for i, d := range b.Data {
		out[i] = GraphiteBase.NewMetric(b.Name, d.Time, d.Value)
	}
	return out
}
