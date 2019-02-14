package problem14

import (
	"testing"
)

var testCases = []struct {
	Given int
}{
	{Given: 10000000},
	// {Given: 100000000},
	// {Given: 1000000000},
}

func TestEstimate(t *testing.T) {
	for _, c := range testCases {
		pi := estimate(c.Given)
		t.Logf("π: %v, n: %v\n", pi, c.Given)
	}
}

var result float64

func BenchmarkEstimate(b *testing.B) {
	var r float64
	for i := 0; i < b.N; i++ {
		r = estimate(1000000)
	}
	result = r
}
