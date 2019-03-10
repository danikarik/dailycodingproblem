package problem36

import "testing"

var testCases = []struct {
	Given    *Node
	Expected int
}{
	{
		Given: &Node{
			Value: 10,
			Left: &Node{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Expected: 5,
	},
	{
		Given: &Node{
			Value: 10,
			Left: &Node{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Value: 20,
				Left:  nil,
				Right: &Node{
					Value: 30,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Expected: 20,
	},
	{
		Given: &Node{
			Value: 50,
			Left: &Node{
				Value: 30,
				Left: &Node{
					Value: 30,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 40,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: 70,
				Left: &Node{
					Value: 60,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 80,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Expected: 70,
	},
}

func TestSecondLargest(t *testing.T) {
	for _, c := range testCases {
		res := secondLargest(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkSecondLargest(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = secondLargest(testCases[0].Given)
	}
	result = r
}
