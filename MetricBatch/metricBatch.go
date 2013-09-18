package MetricBatch

import (
	"github.com/msiebuhr/GraphiteBase"
)

type MetricBatch struct {
	Name string
	Data []*GraphiteBase.MetricValues
}

// Return a new MetricBatch with the given name.
func NewMetricBatch(name string) *MetricBatch {
	n := &MetricBatch{}
	n.Name = name
	n.Data = make([]*GraphiteBase.MetricValues, 0)
	return n
}

// Roll an existing metric into the MetricBatch.
func (b *MetricBatch) AddMetric(metric *GraphiteBase.Metric) {
	b.Data = append(
		b.Data,
		&GraphiteBase.MetricValues{Time: metric.Time, Value: metric.Value},
	)
}

// Un-roll the bulk metric into a list of the given metrics
func (b *MetricBatch) GetMetrics() []*GraphiteBase.Metric {
	out := make([]*GraphiteBase.Metric, len(b.Data))
	for i, d := range b.Data {
		out[i] = GraphiteBase.NewMetric(b.Name, d.Time, d.Value)
	}
	return out
}
