package MetricBatch

import (
	"github.com/msiebuhr/GraphiteBase"
	"testing"
)

func TestMetricBatcher(t *testing.T) {
	// Create a new builk metric
	b := NewMetricBatcher()

	b.AddMetric(GraphiteBase.NewMetric("foobar", 1, 2))
	b.AddMetric(GraphiteBase.NewMetric("foobar", 3, 4))

	// Check it's in there
	if len(b.data) != 1 {
		t.Errorf("Expected internal map to have 1 element, got %v", len(b.data))
	}

	/*
		if _, ok := b.data["foobar"]; ok == true {
			t.Errorf("Expected internal map to have foobar element. It doesn't")
		}
	*/

	// TODO: Fetch data in various ways
}
