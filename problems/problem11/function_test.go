package problem11

import (
	"strings"
	"testing"
)

var (
	isEqual = func(x, y []string) bool {
		if len(x) != len(y) {
			return false
		}
		for i, n := range x {
			if strings.ToLower(n) != strings.ToLower(y[i]) {
				return false
			}
		}
		return true
	}
	testCases = []struct {
		Query    string
		Set      []string
		Expected []string
	}{
		{
			Query:    "de",
			Set:      []string{"dog", "deer", "deal"},
			Expected: []string{"deer", "deal"},
		},
		{
			Query:    "dE",
			Set:      []string{"Dog", "Deer", "DEal"},
			Expected: []string{"deer", "deal"},
		},
	}
)

func TestAutocomplete(t *testing.T) {
	for _, c := range testCases {
		res := autocomplete(c.Query, c.Set)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %s %v", res, c.Expected, c.Query, c.Set)
		}
	}
}

func TestAutocompleteRadix(t *testing.T) {
	for _, c := range testCases {
		res := autocompleteRadix(c.Query, c.Set)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %s %v", res, c.Expected, c.Query, c.Set)
		}
	}
}

var (
	result      []string
	resultRadix []string
)

func BenchmarkAutocomplete(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = autocomplete("de", []string{"dog", "deer", "deal"})
	}
	result = r
}

func BenchmarkAutocompleteRadix(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = autocompleteRadix("de", []string{"dog", "deer", "deal"})
	}
	resultRadix = r
}
