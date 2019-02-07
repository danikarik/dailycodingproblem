package problem02

import "testing"

func TestSolver(t *testing.T) {
	var isEqual = func(x, y []int) bool {
		if len(x) != len(y) {
			return false
		}
		for i, n := range x {
			if n != y[i] {
				return false
			}
		}
		return true
	}
	testCases := []struct {
		Given    []int
		Expected []int
	}{
		{
			Given:    []int{1, 2, 3, 4, 5},
			Expected: []int{120, 60, 40, 30, 24},
		},
		{
			Given:    []int{3, 2, 1},
			Expected: []int{2, 3, 6},
		},
	}
	for _, c := range testCases {
		res := solver(c.Given)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result []int

func BenchmarkSolver(b *testing.B) {
	var r []int
	for i := 0; i < b.N; i++ {
		r = solver([]int{1, 2, 3, 4, 5})
		b.Log(r)
	}
	result = r
}
