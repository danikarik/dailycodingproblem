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
	{
		Given:    Pair{"geek", "gesek"},
		Expected: 1,
	},
	{
		Given:    Pair{"cat", "cut"},
		Expected: 1,
	},
	{
		Given:    Pair{"sunday", "saturday"},
		Expected: 3,
	},
}

func TestMin(t *testing.T) {
	cases := []struct {
		X int
		Y int
		Z int
		R int
	}{
		{
			3, 1, 5, 1,
		},
		{
			1, 2, 5, 1,
		},
		{
			5, 2, 1, 1,
		},
	}
	for _, c := range cases {
		if res := min(c.X, c.Y, c.Z); res != c.R {
			t.Errorf("failed: got %v, expected: %v, values %v %v %v", res, c.R, c.X, c.Y, c.Z)
		}
	}
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
