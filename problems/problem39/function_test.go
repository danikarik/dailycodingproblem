package problem39

import "testing"

var isEqual = func(x, y [][]string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, n := range x {
		if !compare(n, y[i]) {
			return false
		}
	}
	return true
}

var compare = func(x, y []string) bool {
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
	Given    [][]int
	N        int
	Expected [][]string
}{
	{
		Given: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		N: 1,
		Expected: [][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", "*", "*", ".", ".", ".", ".", "."},
			{".", ".", ".", "*", "*", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", "*", "*", "*", ".", ".", ".", ".", "."},
			{".", ".", "*", "*", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", "*", "*", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		},
	},
}

func TestGameOfLife(t *testing.T) {
	for _, c := range testCases {
		res := playGameOfLife(c.Given, c.N)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

var result [][]string

func BenchmarkGameOfLife(b *testing.B) {
	var r [][]string
	for i := 0; i < b.N; i++ {
		r = playGameOfLife(testCases[0].Given, testCases[0].N)
	}
	result = r
}
