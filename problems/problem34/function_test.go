package problem34

import "testing"

var testCases = []struct {
	Given    string
	Expected string
}{
	{
		Given:    "race",
		Expected: "ecarace",
	},
	{
		Given:    "google",
		Expected: "elgoogle",
	},
}

func TestPalindrome(t *testing.T) {
	for _, c := range testCases {
		res := palindrome(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %s, expected %s, values %s", res, c.Expected, c.Given)
		}
	}
}

var result string

func BenchmarkPalindrome(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = palindrome(testCases[0].Given)
	}
	result = r
}
