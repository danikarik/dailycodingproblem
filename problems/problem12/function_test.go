package problem12

import "testing"

var testCases = []struct {
	Given    int
	Steps    []int
	Expected int
}{
	{
		Given:    2,
		Steps:    []int{1, 2},
		Expected: 2,
	},
	{
		Given:    1,
		Steps:    []int{1, 2},
		Expected: 1,
	},
	{
		Given:    4,
		Steps:    []int{1, 2},
		Expected: 5,
	},
}

func TestClimb(t *testing.T) {
	for _, c := range testCases {
		res := climb(c.Given, c.Steps)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkClimb(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = climb(4, []int{1, 2})
	}
	result = r
}
