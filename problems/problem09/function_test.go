package problem09

import "testing"

func TestSolver(t *testing.T) {
	testCases := []struct {
		Given    []int
		Expected int
	}{
		{
			Given:    []int{2, 4, 6, 2, 5},
			Expected: 13,
		},
		{
			Given:    []int{5, 1, 1, 5},
			Expected: 10,
		},
	}
	for _, c := range testCases {
		res := solver(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkSolver(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = solver([]int{2, 4, 6, 2, 5})
	}
	result = r
}
