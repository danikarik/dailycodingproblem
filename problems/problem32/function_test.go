package problem32

import "testing"

var testCases = []struct {
	Given    [][]float64
	Expected bool
}{
	{
		Given: [][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{1.0 / 2.0, 1.0, 3.0 / 2.0},
			[]float64{1.0 / 3.0, 2.0 / 3.0, 1.0},
		},
		Expected: true,
	},
}

func TestArbitrage(t *testing.T) {
	for _, c := range testCases {
		res := arbitrage(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

var result bool

func BenchmarkArbitrage(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = arbitrage(testCases[0].Given)
	}
	result = r
}
