package problem13

import "testing"

var testCases = []struct {
	Given    string
	K        int
	Expected int
}{
	{
		Given:    "abcba",
		K:        2,
		Expected: 3,
	},
	{
		Given:    "aabbcc",
		K:        1,
		Expected: 2,
	},
	{
		Given:    "aabbcc",
		K:        2,
		Expected: 4,
	},
	{
		Given:    "aabbcc",
		K:        3,
		Expected: 6,
	},
	{
		Given:    "aaabbb",
		K:        3,
		Expected: -1,
	},
}

func TestSolver(t *testing.T) {
	for _, c := range testCases {
		res := solver(c.K, c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkSolver(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = solver(2, "abcba")
	}
	result = r
}
