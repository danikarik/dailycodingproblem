package problem40

import "testing"

var testCases = []struct {
	Given    []int
	Expected int
}{
	{
		Given:    []int{3, 3, 2, 3},
		Expected: 2,
	},
	{
		Given:    []int{6, 1, 3, 3, 3, 6, 6},
		Expected: 1,
	},
	{
		Given:    []int{13, 19, 13, 13},
		Expected: 19,
	},
}

func TestNonDuplicated(t *testing.T) {
	for _, c := range testCases {
		res := nonDuplicated(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkNonDuplicated(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = nonDuplicated(testCases[0].Given)
	}
	result = r
}
