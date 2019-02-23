package problem25

import "testing"

var testCases = []struct {
	Given    string
	Pattern  string
	Expected bool
}{
	{
		Given:    "ray",
		Pattern:  "ra.",
		Expected: true,
	},
	{
		Given:    "raymond",
		Pattern:  "ra.",
		Expected: false,
	},
	{
		Given:    "chat",
		Pattern:  ".*at",
		Expected: true,
	},
	{
		Given:    "chats",
		Pattern:  ".*at",
		Expected: false,
	},
}

func TestMatch(t *testing.T) {
	for _, c := range testCases {
		res := matches(c.Given, c.Pattern)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %q", res, c.Expected, c.Given)
		}
	}
}

var result bool

func BenchmarkMatch(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = matches(testCases[0].Given, testCases[0].Pattern)
	}
	result = r
}
