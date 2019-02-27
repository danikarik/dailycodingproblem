package problem28

import "testing"

var isEqual = func(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, n := range x {
		if n != y[i] {
			return false
		}
	}
	return true
}

var testCases = []struct {
	Given    []string
	K        int
	Expected []string
}{
	{
		Given: []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
		K:     16,
		Expected: []string{
			"the  quick brown",
			"fox  jumps  over",
			"the   lazy   dog",
		},
	},
	{
		Given: []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
		K:     6,
		Expected: []string{
			"the   ",
			"quick ",
			"brown ",
			"fox   ",
			"jumps ",
			"over  ",
			"the   ",
			"lazy  ",
			"dog   ",
		},
	},
	{
		Given: []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
		K:     10,
		Expected: []string{
			"the  quick",
			"brown  fox",
			"jumps over",
			"the   lazy",
			"dog       ",
		},
	},
}

func TestJustifyText(t *testing.T) {
	for _, c := range testCases {
		res := justifyText(c.Given, c.K)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

func TestJustifyText2(t *testing.T) {
	for _, c := range testCases {
		res := justifyText2(c.Given, c.K)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var (
	result  []string
	result2 []string
)

func BenchmarkJustifyText(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = justifyText(testCases[0].Given, testCases[0].K)
	}
	result = r
}

func BenchmarkJustifyText2(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = justifyText2(testCases[0].Given, testCases[0].K)
	}
	result2 = r
}
