package problem27

import "testing"

var testCases = []struct {
	Given    string
	Expected bool
}{
	{
		Given:    "([])[]({})",
		Expected: true,
	},
	{
		Given:    "([)]",
		Expected: false,
	},
	{
		Given:    "((()",
		Expected: false,
	},
}

func TestBalance(t *testing.T) {
	for _, c := range testCases {
		res := balance(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %q", res, c.Expected, c.Given)
		}
	}
}

var result bool

func BenchmarkBalance(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = balance(testCases[0].Given)
	}
	result = r
}
