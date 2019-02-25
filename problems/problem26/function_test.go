package problem26

import (
	"testing"
)

var testCases = []struct {
	Given    *Node
	K        int
	Expected *Node
}{
	{
		Given:    &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
		K:        3,
		Expected: &Node{1, &Node{2, &Node{4, &Node{5, nil}}}},
	},
	{
		Given:    &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
		K:        2,
		Expected: &Node{1, &Node{2, &Node{3, &Node{5, nil}}}},
	},
}

func TestRemoveKth(t *testing.T) {
	for _, c := range testCases {
		RemoveKth(c.Given, c.K)
		if !c.Given.IsEqual(c.Expected) {
			t.Errorf("failed: got %v, expected %v", c.Given, c.Expected)
		}
	}
}

var result *Node

func BenchmarkRemoveKth(b *testing.B) {
	var r *Node
	for i := 0; i < b.N; i++ {
		RemoveKth(testCases[0].Given, testCases[0].K)
		r = testCases[0].Given
	}
	result = r
}
