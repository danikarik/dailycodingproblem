package problem18

import "testing"

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

var testCases = []struct {
	Given    []int
	K        int
	Expected []int
}{
	{
		Given:    []int{10, 5, 2, 7, 8, 7},
		K:        3,
		Expected: []int{10, 7, 8, 8},
	},
}

func TestSolver(t *testing.T) {
	for _, c := range testCases {
		res := solver(c.Given, c.K)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v, %v", res, c.Expected, c.Given, c.K)
		}
	}
}

func TestDeque(t *testing.T) {
	for _, c := range testCases {
		res := withDeque(c.Given, c.K)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v, %v", res, c.Expected, c.Given, c.K)
		}
	}
}

var (
	result      []int
	resultDeque []int
)

func BenchmarkSolver(b *testing.B) {
	var r []int
	for i := 0; i < b.N; i++ {
		r = solver([]int{10, 5, 2, 7, 8, 7}, 3)
	}
	result = r
}

func BenchmarkDeque(b *testing.B) {
	var r []int
	for i := 0; i < b.N; i++ {
		r = withDeque([]int{10, 5, 2, 7, 8, 7}, 3)
	}
	resultDeque = r
}
