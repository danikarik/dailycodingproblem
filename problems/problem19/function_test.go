package problem19

import "testing"

var testCases = []struct {
	Given    [][]int
	Expected int
}{
	{
		Given: [][]int{
			[]int{4, 7, 2},
			[]int{10, 1, 14},
		},
		Expected: 3,
	},
	{
		Given: [][]int{
			[]int{5, 10, 15},
			[]int{2, 5, 11},
		},
		Expected: 10,
	},
}

func TestBrute(t *testing.T) {
	for _, c := range testCases {
		res := brute(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

func TestElegant(t *testing.T) {
	for _, c := range testCases {
		res := elegant(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var (
	resultBrute   int
	resultElegant int
)

func BenchmarkBrute(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = brute(testCases[0].Given)
	}
	resultBrute = r
}

func BenchmarkElegant(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = elegant(testCases[0].Given)
	}
	resultElegant = r
}
