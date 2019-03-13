package problem42

import (
	"sort"
	"testing"
)

var isEqual = func(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	left, right := make([]int, len(x)), make([]int, len(y))
	copy(left, x)
	copy(right, y)
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))
	for i, n := range left {
		if n != right[i] {
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
		Given:    []int{12, 1, 61, 5, 9, 2},
		K:        24,
		Expected: []int{12, 9, 2, 1},
	},
	{
		Given:    []int{3, 34, 4, 12, 5, 2},
		K:        9,
		Expected: []int{4, 5},
	},
}

func TestSubsetSum(t *testing.T) {
	for _, c := range testCases {
		res := subsetSum(c.Given, c.K)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v %v", res, c.Expected, c.Given, c.K)
		}
	}
}
