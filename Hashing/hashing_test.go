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

var hashRingOutput_a_b_c_10 = []struct {
	in  string
	out []string
}{
	{"foo", []string{"b", "a", "c"}},
	{"bar", []string{"c", "a", "b"}},
	{"baz", []string{"b", "c", "a"}},
	{"foobar", []string{"c", "a", "b"}},
	{"a", []string{"c", "a", "b"}},
	{"b", []string{"a", "b", "c"}},
	{"c", []string{"c", "b", "a"}},
	{"1", []string{"c", "b", "a"}},
	{"2", []string{"c", "b", "a"}},
	{"3", []string{"c", "b", "a"}},
}

func TestBasicHashRing10Replicas(t *testing.T) {
	// Create new ring
	r := NewConsistentHashRing([]string{"a", "b", "c"}, 10)

	for _, tt := range hashRingOutput_a_b_c_10 {
		// GetNode(key) should return the first node it should hit
		node := r.GetNode(tt.in)
		if node != tt.out[0] {
			t.Errorf("Expected GetNode(%v) to return '%v', got '%v'", tt.in, tt.out[0], node)
		}

		// GetNodes(key) should return all the nodes in order of appearence in the hash
		nodes := r.GetNodes(tt.in)
		if len(nodes) != 3 {
			t.Errorf("Expected GetNodes(%v) to return 3 elements, got %v", tt.in, len(nodes))
		} else if nodes[0] != tt.out[0] || nodes[1] != tt.out[1] || nodes[2] != tt.out[2] {
			t.Errorf("Expected GetNodes(%v) to return '%v', got '%v'", tt.in, tt.out, nodes)
		}
	}
}

func BenchmarkGetNode_3_10(b *testing.B) {
	r := NewConsistentHashRing([]string{"a", "b", "c"}, 10)
	keys := []string{"foo", "bar", "baz", "1", "2", "3"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.GetNode(keys[i%len(keys)])
	}
}

func BenchmarkGetNode_12_100(b *testing.B) {
	r := NewConsistentHashRing([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}, 100)
	keys := []string{"foo", "bar", "baz", "1", "2", "3"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.GetNode(keys[i%len(keys)])
	}
}
