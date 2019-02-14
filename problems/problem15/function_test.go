package problem15

import "testing"

func TestPick(t *testing.T) {
	testCases := []struct {
		Given []int
	}{
		{
			Given: []int{1, 2, 3, 4},
		},
	}
	for _, c := range testCases {
		res := pick(c.Given)
		t.Logf("res: %v\n", res)
	}
}

var result int

func BenchmarkPick(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = pick([]int{1, 2, 3, 4})
	}
	result = r
}
