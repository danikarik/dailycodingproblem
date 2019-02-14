package problem16

import "testing"

var testCases = []struct {
	Given    []int
	N        int
	I        int
	Expected int
}{
	{
		Given:    []int{1, 2, 3, 4, 5, 6, 7, 8},
		N:        7,
		I:        4,
		Expected: 5,
	},
	{
		Given:    []int{1, 2, 3, 4, 5, 6, 7},
		N:        7,
		I:        7,
		Expected: 1,
	},
	{
		Given:    []int{1, 2, 3, 4, 5, 6, 7},
		N:        7,
		I:        3,
		Expected: 5,
	},
}

func TestSliceContainer(t *testing.T) {
	for _, c := range testCases {
		container := NewSliceContainer(c.N)
		for _, rec := range c.Given {
			container.Record(rec)
		}
		lastRec := container.GetLast(c.I)
		if lastRec != c.Expected {
			t.Errorf("failed: got %v, expected %v, values: %v, %v", lastRec, c.Expected, c.Given, c.I)
		}
		t.Logf("last ith %v record from %v: %v", c.I, c.Given, lastRec)
	}
}

var (
	resultSlice int
)

func BenchmarkSliceContainer(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		container := NewSliceContainer(4)
		container.Record(1)
		container.Record(2)
		container.Record(3)
		r = container.GetLast(3)
	}
	resultSlice = r
}
