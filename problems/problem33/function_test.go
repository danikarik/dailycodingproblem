package problem33

import "testing"

var isEqual = func(x, y []float64) bool {
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
	Expected []float64
}{
	{
		Given:    []int{2, 1, 5, 7, 2, 0, 5},
		Expected: []float64{2, 1.5, 2, 3.5, 2, 2, 2},
	},
}

func TestMedian(t *testing.T) {
	for _, c := range testCases {
		res := median(c.Given)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

func TestMedianWithHeap(t *testing.T) {
	for _, c := range testCases {
		res := medianWithHeap(c.Given)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var (
	result         []float64
	resultWithHeap []float64
)

func BenchmarkMedian(b *testing.B) {
	var r []float64
	for i := 0; i < b.N; i++ {
		r = median(testCases[0].Given)
	}
	result = r
}

func BenchmarkMedianWithHeap(b *testing.B) {
	var r []float64
	for i := 0; i < b.N; i++ {
		r = medianWithHeap(testCases[0].Given)
	}
	resultWithHeap = r
}
