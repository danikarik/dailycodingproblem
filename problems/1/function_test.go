package function

import "testing"

func TestSolver(t *testing.T) {
	testCases := []struct {
		Given    []int
		Number   int
		Expected bool
	}{
		{
			Given:    []int{10, 15, 3, 7},
			Number:   10,
			Expected: true,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   13,
			Expected: true,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   17,
			Expected: true,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   18,
			Expected: true,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   22,
			Expected: true,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   25,
			Expected: true,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   16,
			Expected: false,
		},
		{
			Given:    []int{10, 15, 3, 7},
			Number:   11,
			Expected: false,
		},
	}
	for _, c := range testCases {
		res := solver(c.Given, c.Number)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v - %v", res, c.Expected, c.Given, c.Number)
		}
	}
}

var result bool

func BenchmarkSolver(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = solver([]int{10, 15, 3, 7}, 17)
		b.Log(r)
	}
	result = r
}
