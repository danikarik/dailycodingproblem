package problem31

import "testing"

var testCases = []struct {
	Given    Pair
	Expected int
}{
	{
		Given:    Pair{"kitten", "sitting"},
		Expected: 3,
	},
	{
		Given:    Pair{"dog", "cat"},
		Expected: 3,
	},
	{
		Given:    Pair{"formula", "for"},
		Expected: 4,
	},
	{
		Given:    Pair{"bobby", "bobsleigh"},
		Expected: 6,
	},
}

func TestEditDistance(t *testing.T) {
	for _, c := range testCases {
		res := editDistance(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %s", res, c.Expected, c.Given.String())
		}
	}
}

var result int

func BenchmarkEditDistance(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = editDistance(testCases[0].Given)
	}
	result = r
}
