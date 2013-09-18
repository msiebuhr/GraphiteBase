package Hashing

import "testing"

var ringKnownPositions = []struct {
	in  string
	out uint16
}{
	{"a", 3265},
	{"b", 37611},
	{"c", 19082},
	{"d", 33399},
	{"e", 57703},
	{"f", 36769},
}

func TestComputeRingPosition(t *testing.T) {
	for _, tt := range ringKnownPositions {
		out := computeRingPosition(tt.in)
		if out != tt.out {
			t.Errorf("Expected '%v' to map to '%v', but got '%v'", tt.in, tt.out, out)
		}
	}
}

func BenchmarkComputeRingPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = computeRingPosition("a")
	}
}
