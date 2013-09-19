package MetricBatch

import (
	"github.com/msiebuhr/GraphiteBase"
	"testing"
)

func TestMetricBatcher(t *testing.T) {
	// Create a new builk metric
	b := NewMetricBatcher()

	b.AddMetric(GraphiteBase.NewMetric("foo", 1, 2))
	b.AddMetric(GraphiteBase.NewMetric("foo", 3, 4))
	b.AddMetric(GraphiteBase.NewMetric("bar", 5, 6))

	// Check it's in there
	if len(b.data) != 2 {
		t.Errorf("Expected internal map to have 2 elements, got %v", len(b.data))
	}
	if b.Size() != 3 {
		t.Errorf("Expected Size() to return 3, got %v", b.Size())
	}

	/*
		if _, ok := b.data["foobar"]; ok == true {
			t.Errorf("Expected internal map to have foobar element. It doesn't")
		}
	*/

	// Fetch data in various ways
	batch := b.GetLargestBatch()

	if batch.Name != "foo" {
		t.Errorf("Expected 'foo'-batch, got '%v'", batch.Name)
	}

	if batch.Size() != 2 {
		t.Errorf("Expected returned batch to have Size() 2, got '%v'", batch.Size())
	}

	// Check batcher is now empty
	if len(b.data) != 1 {
		t.Errorf("Expected batch to have one element, but it has %v elements", len(b.data))
	}
	if b.Size() != 1 {
		t.Errorf("Expected Size() to return 1, got %v", b.Size())
	}
}

func BenchmarkMetricBatcherMax1000Elements(b *testing.B) {
	batcher := NewMetricBatcher()
	for i := 0; i < b.N; i++ {
		// Load a metric
		batcher.AddMetric(GraphiteBase.NewMetric("foo", 1, 1.0))

		// Check if batch has gotten too large
		if batcher.Size() >= 1000 {
			batch := batcher.GetLargestBatch()
			if batch.Size() > 1000 {
				b.Errorf("Expected batch of 1000 element tops, got %v", batch.Size())
			}
		}
	}
}
