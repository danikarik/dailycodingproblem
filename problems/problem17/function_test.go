package problem17

import "testing"

var testCases = []struct {
	Given    string
	Expected int
}{
	{
		Given:    `dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext`,
		Expected: 32,
	},
	{
		Given:    `dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext`,
		Expected: 20,
	},
	{
		Given:    `dir\n\tsubdir1\n\tsubdir2`,
		Expected: 0,
	},
	{
		Given:    `dir\n\tsubdir1\n\t\tfile1.ext`,
		Expected: 21,
	},
	{
		Given:    `dir`,
		Expected: 0,
	},
	{
		Given:    `dir\n\tsubdir1`,
		Expected: 0,
	},
}

func TestLongestPath(t *testing.T) {
	for _, c := range testCases {
		res := longestPath(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values: %q", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkLongestPath(b *testing.B) {
	var r int
	s := `dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext`
	for i := 0; i < b.N; i++ {
		r = longestPath(s)
	}
	result = r
}
