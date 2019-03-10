package problem37

import (
	"testing"
)

var isEqual = func(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}
	found := make([]bool, len(x))
	for i := 0; i < len(x); i++ {
		ok := false
		for j := 0; j < len(y); j++ {
			if compare(x[i], y[j]) {
				ok = true
				break
			}
		}
		found[i] = ok
	}
	return true
}

var compare = func(x, y []int) bool {
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
	Expected [][]int
}{
	{
		Given: []int{1, 2, 3},
		Expected: [][]int{
			[]int{},
			[]int{1},
			[]int{2},
			[]int{3},
			[]int{1, 2},
			[]int{1, 3},
			[]int{2, 3},
			[]int{1, 2, 3},
		},
	},
	{
		Given: []int{},
		Expected: [][]int{
			[]int{},
		},
	},
	{
		Given: []int{1},
		Expected: [][]int{
			[]int{},
			[]int{1},
		},
	},
}

func TestPowerSet(t *testing.T) {
	for _, c := range testCases {
		res := powerSet(c.Given)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

var result [][]int

func BenchmarkPowerSet(b *testing.B) {
	var r [][]int
	for i := 0; i < b.N; i++ {
		r = powerSet(testCases[0].Given)
	}
	result = r
}
