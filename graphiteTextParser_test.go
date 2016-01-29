package GraphiteBase

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestParseMetricLineOK(t *testing.T) {
	m, err := parseGraphiteLine("foo.bar 42.1 10")

	if err != nil {
		t.Errorf("Didn't expect an error got %v", err)
	}

	if m.Name != "foo.bar" {
		t.Errorf("Expected name to be foo.bar, got %v", m.Name)
	}

	if m.Time != 10 {
		t.Errorf("Expected time to be 10, got %v", m.Time)
	}

	if m.Value != 42.1 {
		t.Errorf("Expected value to be 42.1, got %v", m.Value)
	}
}

func BenchmarkParseMetricLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = parseGraphiteLine("foo.bar 123 456")
		b.SetBytes(15)
	}
}

func TestGraphiteProtocolReaderOK(t *testing.T) {
	out := make(chan *Metric)
	in := ioutil.NopCloser(
		bytes.NewReader([]byte("foo 1 1\nbar 2 2\nbaz 3 3\n")),
	)

	go func() {
		GraphiteProtocolReader(in, out)
		close(out)
	}()

	// Read output
	data := []*Metric{}
	for m := range out {
		data = append(data, m)
	}

	// Basic checks on output
	if len(data) != 3 {
		t.Errorf("Expected to get 3 metrics, got %v", len(data))
	}

	// We rely on output being in-order for testing purposes. It shouldn't have
	// to be that way in real life.
	for i, m := range data {
		// Test datapoint
		if m.Time != int64(i+1) {
			t.Errorf("Expected time to be %v, got %v", i+1, m.Time)
		}

		if m.Value != float64(i+1) {
			t.Errorf("Expected value to be %v, got %v", i+1, m.Value)
		}
	}
}
