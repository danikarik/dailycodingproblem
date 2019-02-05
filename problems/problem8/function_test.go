package problem8

import "testing"

func TestSolver(t *testing.T) {
	testCases := []struct {
		Given    *Node
		Expected int
	}{
		{
			Given: &Node{
				Value: 0,
				Left: &Node{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 0,
					Left: &Node{
						Value: 1,
						Left: &Node{
							Value: 1,
							Left:  nil,
							Right: nil,
						},
						Right: &Node{
							Value: 1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &Node{
						Value: 0,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Expected: 5,
		},
		{
			Given: &Node{
				Value: 1,
				Left: &Node{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Value: 1,
						Left:  nil,
						Right: &Node{
							Value: 0,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			Expected: 3,
		},
		{
			Given: &Node{
				Value: 1,
				Left: &Node{
					Value: 3,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 2,
					Left: &Node{
						Value: 2,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Value: 2,
						Left:  nil,
						Right: &Node{
							Value: 2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			Expected: 5,
		},
	}
	for _, c := range testCases {
		res := solver(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

var result int

func BenchmarkSolver(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = solver(&Node{
			Value: 0,
			Left: &Node{
				Value: 1,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Value: 0,
				Left: &Node{
					Value: 1,
					Left: &Node{
						Value: 1,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Value: 1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Value: 0,
					Left:  nil,
					Right: nil,
				},
			},
		})
	}
	result = r
}
