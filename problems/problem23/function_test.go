package problem23

import "testing"

var testCases = []struct {
	Given    [][]bool
	Start    coords
	End      coords
	Expected int
}{
	{
		Given: [][]bool{
			[]bool{false, false, false, false},
			[]bool{true, true, false, true},
			[]bool{false, false, false, false},
			[]bool{false, false, false, false},
		},
		Start:    coords{row: 3, col: 0},
		End:      coords{row: 0, col: 0},
		Expected: 7,
	},
}

func TestSteps(t *testing.T) {
	for _, c := range testCases {
		res := steps(c.Given, c.Start, c.End)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkSteps(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = steps(testCases[0].Given, testCases[0].Start, testCases[0].End)
	}
	result = r
}
