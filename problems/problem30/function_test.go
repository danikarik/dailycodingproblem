package problem30

import "testing"

var testCases = []struct {
	Given    []int
	Expected int
}{
	{
		Given:    []int{2, 1, 2},
		Expected: 1,
	},
	{
		Given:    []int{3, 0, 1, 3, 0, 5},
		Expected: 8,
	},
}

func TestCapacity(t *testing.T) {
	for _, c := range testCases {
		res := capacity(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkCapacity(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = capacity(testCases[0].Given)
	}
	result = r
}
