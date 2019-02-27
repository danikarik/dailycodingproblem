package problem29

import "testing"

var testCases = []struct {
	Given    string
	Expected string
}{
	{
		Given:    "AAAABBBCCDAA",
		Expected: "4A3B2C1D2A",
	},
	{
		Given:    "AAABBBBBCDEFF",
		Expected: "3A5B1C1D1E2F",
	},
	{
		Given:    "QWERTYY",
		Expected: "1Q1W1E1R1T2Y",
	},
	{
		Given:    "AAAAAAAAA",
		Expected: "9A",
	},
}

func TestEncode(t *testing.T) {
	for _, c := range testCases {
		res := encode(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %q, expected %q, values %q", res, c.Expected, c.Given)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, c := range testCases {
		res := decode(c.Expected)
		if res != c.Given {
			t.Errorf("failed: got %q, expected %q, values %q", res, c.Given, c.Expected)
		}
	}
}

var (
	resultEncode string
	resultDecode string
)

func BenchmarkEncode(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = encode(testCases[0].Given)
	}
	resultEncode = r
}

func BenchmarkDecode(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = encode(testCases[0].Expected)
	}
	resultDecode = r
}
