package GraphiteBase

import "testing"

// String()
func TestMetricString(t *testing.T) {
	m := NewMetric("foo", 42, 10.1)
	s := m.String()

	if s != "foo 42 10.1" {
		t.Errorf("Expected 'foo 42 10.1', got '%v'", s)
	}
}

func BenchmarkMetricString(b *testing.B) {
	m := NewMetric("foo.bar", 123, 1.23)
	for i := 0; i < b.N; i++ {
		s := m.String()
		b.SetBytes(int64(len(s)))
	}
}
