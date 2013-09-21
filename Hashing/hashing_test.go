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

func TestAddRemoveNode(t *testing.T) {
	r := NewConsistentHashRing([]string{}, 10)

	// Add one node and confirm it only has those in the ring
	r.AddNode("node_1")

	for _, e := range r.ring {
		if e.key != "node_1" {
			t.Errorf("Expected node_1, got %v.", e.key)
		}
	}

	// Add another node and remove the first and check again
	r.AddNode("node_2")
	r.RemoveNode("node_1")

	for _, e := range r.ring {
		if e.key != "node_2" {
			t.Errorf("Expected node_2, got %v.", e.key)
		}
	}

	// Remove the second node and check it's empty
	r.RemoveNode("node_2")
	if len(r.ring) != 0 {
		t.Errorf("Expected ring to empty, got &%v.", r.ring)
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

	{"0", []string{"c", "b", "a"}},
	{"1", []string{"c", "b", "a"}},
	{"2", []string{"c", "b", "a"}},
	{"3", []string{"c", "b", "a"}},
	{"4", []string{"b", "a", "c"}},
	{"5", []string{"b", "c", "a"}},
	{"6", []string{"a", "c", "b"}},
	{"7", []string{"c", "a", "b"}},
	{"8", []string{"c", "b", "a"}},
	{"9", []string{"c", "b", "a"}},
	{"10", []string{"c", "b", "a"}},
	{"11", []string{"a", "c", "b"}},
	{"12", []string{"c", "b", "a"}},
	{"13", []string{"c", "b", "a"}},
	{"14", []string{"b", "a", "c"}},
	{"15", []string{"a", "b", "c"}},
	{"16", []string{"c", "b", "a"}},
	{"17", []string{"c", "b", "a"}},
	{"18", []string{"c", "b", "a"}},
	{"19", []string{"a", "c", "b"}},
	{"20", []string{"a", "b", "c"}},
	{"21", []string{"a", "b", "c"}},
	{"22", []string{"a", "c", "b"}},
	{"23", []string{"c", "a", "b"}},
	{"24", []string{"a", "c", "b"}},
	{"25", []string{"c", "a", "b"}},
	{"26", []string{"c", "b", "a"}},
	{"27", []string{"c", "a", "b"}},
	{"28", []string{"a", "c", "b"}},
	{"29", []string{"c", "b", "a"}},
	{"30", []string{"c", "a", "b"}},
	{"31", []string{"c", "b", "a"}},
	{"32", []string{"b", "a", "c"}},
	{"33", []string{"a", "c", "b"}},
	{"34", []string{"b", "c", "a"}},
	{"35", []string{"a", "c", "b"}},
	{"36", []string{"a", "c", "b"}},
	{"37", []string{"b", "a", "c"}},
	{"38", []string{"b", "a", "c"}},
	{"39", []string{"c", "b", "a"}},
	{"40", []string{"c", "b", "a"}},
	{"41", []string{"c", "a", "b"}},
	{"42", []string{"b", "a", "c"}},
	{"43", []string{"a", "c", "b"}},
	{"44", []string{"c", "b", "a"}},
	{"45", []string{"c", "b", "a"}},
	{"46", []string{"b", "c", "a"}},
	{"47", []string{"c", "b", "a"}},
	{"48", []string{"b", "a", "c"}},
	{"49", []string{"c", "b", "a"}},
	{"50", []string{"c", "b", "a"}},
	{"51", []string{"a", "c", "b"}},
	{"52", []string{"a", "b", "c"}},
	{"53", []string{"b", "c", "a"}},
	{"54", []string{"b", "a", "c"}},
	{"55", []string{"a", "c", "b"}},
	{"56", []string{"b", "a", "c"}},
	{"57", []string{"b", "c", "a"}},
	{"58", []string{"c", "b", "a"}},
	{"59", []string{"c", "a", "b"}},
	{"60", []string{"c", "a", "b"}},
	{"61", []string{"b", "c", "a"}},
	{"62", []string{"c", "b", "a"}},
	{"63", []string{"c", "a", "b"}},
	{"64", []string{"c", "b", "a"}},
	{"65", []string{"c", "a", "b"}},
	{"66", []string{"a", "c", "b"}},
	{"67", []string{"b", "c", "a"}},
	{"68", []string{"b", "a", "c"}},
	{"69", []string{"a", "c", "b"}},
	{"70", []string{"b", "c", "a"}},
	{"71", []string{"b", "c", "a"}},
	{"72", []string{"a", "c", "b"}},
	{"73", []string{"c", "b", "a"}},
	{"74", []string{"a", "b", "c"}},
	{"75", []string{"c", "b", "a"}},
	{"76", []string{"c", "a", "b"}},
	{"77", []string{"a", "c", "b"}},
	{"78", []string{"c", "a", "b"}},
	{"79", []string{"c", "b", "a"}},
	{"80", []string{"c", "b", "a"}},
	{"81", []string{"c", "b", "a"}},
	{"82", []string{"a", "b", "c"}},
	{"83", []string{"c", "a", "b"}},
	{"84", []string{"c", "b", "a"}},
	{"85", []string{"c", "b", "a"}},
	{"86", []string{"a", "b", "c"}},
	{"87", []string{"c", "b", "a"}},
	{"88", []string{"a", "c", "b"}},
	{"89", []string{"b", "c", "a"}},
	{"90", []string{"b", "c", "a"}},
	{"91", []string{"c", "b", "a"}},
	{"92", []string{"a", "b", "c"}},
	{"93", []string{"a", "b", "c"}},
	{"94", []string{"c", "b", "a"}},
	{"95", []string{"b", "c", "a"}},
	{"96", []string{"a", "c", "b"}},
	{"97", []string{"b", "c", "a"}},
	{"98", []string{"c", "b", "a"}},
	{"99", []string{"b", "a", "c"}},
	{"100", []string{"b", "c", "a"}},
	{"101", []string{"c", "a", "b"}},
	{"102", []string{"c", "b", "a"}},
	{"103", []string{"c", "b", "a"}},
	{"104", []string{"c", "b", "a"}},
	{"105", []string{"a", "c", "b"}},
	{"106", []string{"c", "b", "a"}},
	{"107", []string{"b", "a", "c"}},
	{"108", []string{"b", "a", "c"}},
	{"109", []string{"a", "c", "b"}},
	{"110", []string{"a", "b", "c"}},
	{"111", []string{"c", "b", "a"}},
	{"112", []string{"b", "c", "a"}},
	{"113", []string{"b", "c", "a"}},
	{"114", []string{"a", "b", "c"}},
	{"115", []string{"a", "c", "b"}},
	{"116", []string{"c", "b", "a"}},
	{"117", []string{"c", "b", "a"}},
	{"118", []string{"a", "b", "c"}},
	{"119", []string{"c", "a", "b"}},
	{"120", []string{"b", "c", "a"}},
	{"121", []string{"c", "b", "a"}},
	{"122", []string{"b", "a", "c"}},
	{"123", []string{"a", "c", "b"}},
	{"124", []string{"c", "b", "a"}},
	{"125", []string{"b", "c", "a"}},
	{"126", []string{"c", "a", "b"}},
	{"127", []string{"c", "b", "a"}},
	{"128", []string{"b", "c", "a"}},
	{"129", []string{"c", "b", "a"}},
	{"130", []string{"a", "b", "c"}},
	{"131", []string{"a", "c", "b"}},
	{"132", []string{"a", "c", "b"}},
	{"133", []string{"b", "a", "c"}},
	{"134", []string{"c", "a", "b"}},
	{"135", []string{"b", "c", "a"}},
	{"136", []string{"c", "b", "a"}},
	{"137", []string{"a", "b", "c"}},
	{"138", []string{"c", "a", "b"}},
	{"139", []string{"b", "c", "a"}},
	{"140", []string{"a", "c", "b"}},
	{"141", []string{"a", "c", "b"}},
	{"142", []string{"b", "a", "c"}},
	{"143", []string{"a", "b", "c"}},
	{"144", []string{"c", "a", "b"}},
	{"145", []string{"a", "c", "b"}},
	{"146", []string{"b", "a", "c"}},
	{"147", []string{"c", "a", "b"}},
	{"148", []string{"c", "b", "a"}},
	{"149", []string{"c", "b", "a"}},
	{"150", []string{"b", "c", "a"}},
	{"151", []string{"b", "a", "c"}},
	{"152", []string{"c", "a", "b"}},
	{"153", []string{"a", "c", "b"}},
	{"154", []string{"a", "c", "b"}},
	{"155", []string{"a", "c", "b"}},
	{"156", []string{"a", "c", "b"}},
	{"157", []string{"c", "b", "a"}},
	{"158", []string{"c", "a", "b"}},
	{"159", []string{"a", "c", "b"}},
	{"160", []string{"a", "c", "b"}},
	{"161", []string{"a", "c", "b"}},
	{"162", []string{"b", "c", "a"}},
	{"163", []string{"c", "a", "b"}},
	{"164", []string{"c", "a", "b"}},
	{"165", []string{"b", "a", "c"}},
	{"166", []string{"b", "c", "a"}},
	{"167", []string{"c", "b", "a"}},
	{"168", []string{"c", "a", "b"}},
	{"169", []string{"c", "a", "b"}},
	{"170", []string{"a", "c", "b"}},
	{"171", []string{"b", "a", "c"}},
	{"172", []string{"a", "c", "b"}},
	{"173", []string{"b", "c", "a"}},
	{"174", []string{"c", "b", "a"}},
	{"175", []string{"b", "c", "a"}},
	{"176", []string{"c", "a", "b"}},
	{"177", []string{"b", "a", "c"}},
	{"178", []string{"a", "b", "c"}},
	{"179", []string{"c", "a", "b"}},
	{"180", []string{"c", "a", "b"}},
	{"181", []string{"c", "a", "b"}},
	{"182", []string{"c", "b", "a"}},
	{"183", []string{"c", "b", "a"}},
	{"184", []string{"c", "b", "a"}},
	{"185", []string{"c", "b", "a"}},
	{"186", []string{"a", "b", "c"}},
	{"187", []string{"a", "c", "b"}},
	{"188", []string{"a", "b", "c"}},
	{"189", []string{"b", "a", "c"}},
	{"190", []string{"c", "b", "a"}},
	{"191", []string{"c", "a", "b"}},
	{"192", []string{"c", "b", "a"}},
	{"193", []string{"a", "c", "b"}},
	{"194", []string{"b", "a", "c"}},
	{"195", []string{"c", "a", "b"}},
	{"196", []string{"c", "a", "b"}},
	{"197", []string{"b", "c", "a"}},
	{"198", []string{"a", "c", "b"}},
	{"199", []string{"b", "c", "a"}},
	{"200", []string{"c", "a", "b"}},
	{"201", []string{"b", "c", "a"}},
	{"202", []string{"b", "c", "a"}},
	{"203", []string{"b", "c", "a"}},
	{"204", []string{"a", "c", "b"}},
	{"205", []string{"c", "b", "a"}},
	{"206", []string{"b", "c", "a"}},
	{"207", []string{"c", "b", "a"}},
	{"208", []string{"c", "a", "b"}},
	{"209", []string{"a", "c", "b"}},
	{"210", []string{"c", "b", "a"}},
	{"211", []string{"c", "b", "a"}},
	{"212", []string{"a", "c", "b"}},
	{"213", []string{"a", "b", "c"}},
	{"214", []string{"c", "b", "a"}},
	{"215", []string{"a", "b", "c"}},
	{"216", []string{"c", "b", "a"}},
	{"217", []string{"b", "a", "c"}},
	{"218", []string{"c", "b", "a"}},
	{"219", []string{"c", "b", "a"}},
	{"220", []string{"c", "b", "a"}},
	{"221", []string{"c", "a", "b"}},
	{"222", []string{"a", "c", "b"}},
	{"223", []string{"a", "c", "b"}},
	{"224", []string{"a", "c", "b"}},
	{"225", []string{"c", "b", "a"}},
	{"226", []string{"a", "b", "c"}},
	{"227", []string{"c", "b", "a"}},
	{"228", []string{"b", "c", "a"}},
	{"229", []string{"c", "b", "a"}},
	{"230", []string{"c", "b", "a"}},
	{"231", []string{"a", "b", "c"}},
	{"232", []string{"a", "c", "b"}},
	{"233", []string{"b", "c", "a"}},
	{"234", []string{"a", "c", "b"}},
	{"235", []string{"c", "b", "a"}},
	{"236", []string{"c", "a", "b"}},
	{"237", []string{"c", "b", "a"}},
	{"238", []string{"b", "a", "c"}},
	{"239", []string{"c", "b", "a"}},
	{"240", []string{"a", "c", "b"}},
	{"241", []string{"c", "b", "a"}},
	{"242", []string{"b", "c", "a"}},
	{"243", []string{"c", "b", "a"}},
	{"244", []string{"a", "b", "c"}},
	{"245", []string{"c", "a", "b"}},
	{"246", []string{"a", "b", "c"}},
	{"247", []string{"a", "b", "c"}},
	{"248", []string{"b", "a", "c"}},
	{"249", []string{"c", "a", "b"}},
	{"250", []string{"c", "b", "a"}},
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
