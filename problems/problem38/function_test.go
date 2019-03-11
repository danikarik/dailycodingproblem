package problem38

import "testing"

var testCases = []struct {
	Given    int
	Expected int
}{
	{
		Given:    4,
		Expected: 4,
	},
}

func TestQueens(t *testing.T) {
	for _, c := range testCases {
		res := queens(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkQueens(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = queens(testCases[0].Given)
	}
	result = r
}
