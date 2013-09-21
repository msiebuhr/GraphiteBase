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

// Add a metric to the batcher
func (mb *MetricBatcher) AddMetric(m *GraphiteBase.Metric) {
	// If it isn't there, add a new MetricBatch
	if _, ok := mb.data[m.Name]; ok == false {
		mb.data[m.Name] = NewMetricBatch(m.Name)
	}

	mb.data[m.Name].AddMetric(m)
}

// Return the largest batch available.
//
// Loops over all internal MetricBatch'es and returns the largest one it finds.
func (mb *MetricBatcher) GetLargestBatch() *MetricBatch {
	largest := &MetricBatch{}
	for _, batch := range mb.data {
		if batch.Len() > largest.Len() {
			largest = batch
		}
	}

	// Drop the batch
	delete(mb.data, largest.Name)

	return largest
}

// Return the total number of metrics currently cached inside the batcher.
func (mb *MetricBatcher) Len() (size int) {
	size = 0
	for _, batch := range mb.data {
		size += batch.Len()
	}
	return size
}
