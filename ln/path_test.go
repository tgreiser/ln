package ln

import "testing"

func TestFind(t *testing.T) {
	var p = Paths{
		{{0, 0, 0}, {0, 0, 1}},
		{{0, 0, 0}, {0, 1, 0}},
		{{0, 0, 0}, {1, 0, 0}},
	}
	var findTests = []struct {
		paths     Paths
		vec       Vector
		expected1 int
		expected2 int
	}{
		{p, Vector{0, 0, 0}, 0, 0},
		{p, Vector{0, 1, 0}, 1, 1},
		{p[1:], Vector{0, 0, 0}, 0, 0},
		{p, Vector{0, 10, 0}, -1, -1},
	}

	for _, tt := range findTests {
		i1, i2 := tt.paths.find(tt.vec)
		if i1 != tt.expected1 {
			t.Errorf("index1 %d != %d\n", i1, tt.expected1)
		}
		if i2 != tt.expected2 {
			t.Errorf("index2 %d != %d\n", i2, tt.expected2)
		}
	}
}

func TestReverse(t *testing.T) {
	var revTests = []struct {
		given    Path
		expected Path
	}{
		{Path{Vector{0, 0, 0}, Vector{1, 1, 1}}, Path{Vector{1, 1, 1}, Vector{0, 0, 0}}},
		{Path{Vector{0, 0, 0}, Vector{0, 0, 0}}, Path{Vector{0, 0, 0}, Vector{0, 0, 0}}},
	}
	for _, tt := range revTests {
		tt.given.Reverse()
		if tt.given[0].Distance(tt.expected[0]) != 0 {
			t.Errorf("%v != %v\n", tt.expected[0], tt.given[0])
		}
		if tt.given[1].Distance(tt.expected[1]) != 0 {
			t.Errorf("%v != %v\n", tt.expected[1], tt.given[1])
		}
	}
}
