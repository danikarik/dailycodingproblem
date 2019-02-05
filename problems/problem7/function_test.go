package problem7

import "testing"

func TestSolver(t *testing.T) {
	testCases := []struct {
		Given    string
		Expected int
	}{
		{
			Given:    "111",
			Expected: 3,
		},
		{
			Given:    "121",
			Expected: 3,
		},
		{
			Given:    "1234",
			Expected: 3,
		},
	}
	for _, c := range testCases {
		res := solver(c.Given, len(c.Given))
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkSolver(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = solver("111", 3)
	}
	result = r
}
