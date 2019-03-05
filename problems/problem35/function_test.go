package problem35

import "testing"

var isEqual = func(x, y []string) bool {
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
	Given    []string
	Expected []string
}{
	{
		Given:    []string{"G", "B", "R", "R", "B", "R", "G"},
		Expected: []string{"R", "R", "R", "G", "G", "B", "B"},
	},
	{
		Given:    []string{"R", "B", "R", "R", "B", "B", "G"},
		Expected: []string{"R", "R", "R", "G", "B", "B", "B"},
	},
}

func TestSort(t *testing.T) {
	for _, c := range testCases {
		res := sort(c.Given)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result []string

func BenchmarkSort(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = sort(testCases[0].Given)
	}
	result = r
}
