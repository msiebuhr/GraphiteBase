package MetricBatch

import (
	"github.com/msiebuhr/GraphiteBase"
	"testing"
)

func TestMetricBatch(t *testing.T) {
	// Create a new builk metric
	bm := NewMetricBatch("foobar")

	bm.AddMetric(GraphiteBase.NewMetric(
		"foobar",
		1,
		2,
	))

	// Check size
	if bm.Len() != 1 {
		t.Errorf("Len() to return 1, got %v", bm.Len())
	}

	// Gets the same metrics out?
	metricList := bm.GetMetrics()
	if len(metricList) != 1 {
		t.Errorf("GetMetrics() to return 1 element, got %v", len(metricList))
	}
}
