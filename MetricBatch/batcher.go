package MetricBatch

import (
	"github.com/msiebuhr/GraphiteBase"
)

type MetricBatcher struct {
	data map[string]*MetricBatch
}

func NewMetricBatcher() *MetricBatcher {
	return &MetricBatcher{
		data: make(map[string]*MetricBatch),
	}
}

func (mb *MetricBatcher) AddMetric(m *GraphiteBase.Metric) {
	// If it isn't there, add a new MetricBatch
	if _, ok := mb.data[m.Name]; ok == false {
		mb.data[m.Name] = NewMetricBatch(m.Name)
	}

	mb.data[m.Name].AddMetric(m)
}
