package problem4

import "testing"

var testCases = []struct {
	Given    []int
	Expected int
}{
	{
		Given:    []int{3, 4, -1, 1},
		Expected: 2,
	},
	{
		Given:    []int{1, 2, 0},
		Expected: 3,
	},
	{
		Given:    []int{-1, 0, 2},
		Expected: 1,
	},
	{
		Given:    []int{-4, -2, -1},
		Expected: 1,
	},
	{
		Given:    []int{1, 2, 5},
		Expected: 3,
	},
	{
		Given:    []int{-1, 0, 3},
		Expected: 1,
	},
	{
		Given:    []int{2, -1, 1, 4},
		Expected: 3,
	},
}

func TestSolver(t *testing.T) {
	for _, c := range testCases {
		res := solver(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

func TestSolverSort(t *testing.T) {
	for _, c := range testCases {
		res := solverSort(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var (
	result     int
	resultSort int
)

func BenchmarkSolver(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = solver([]int{1, 2, 3, 4, 5})
		b.Log(r)
	}
	result = r
}
func BenchmarkSolverSort(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = solverSort([]int{1, 2, 3, 4, 5})
		b.Log(r)
	}
	resultSort = r
}
