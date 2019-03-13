package problem41

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
	Given    []Flight
	Start    string
	Expected []string
}{
	{
		Given: []Flight{
			{"SFO", "HKO"},
			{"YYZ", "SFO"},
			{"YUL", "YYZ"},
			{"HKO", "ORD"},
		},
		Start:    "YUL",
		Expected: []string{"YUL", "YYZ", "SFO", "HKO", "ORD"},
	},
	{
		Given: []Flight{
			{"SFO", "COM"},
			{"COM", "YYZ"},
		},
		Start:    "COM",
		Expected: []string{},
	},
	// {
	// 	Given: []Flight{
	// 		{"A", "B"},
	// 		{"A", "C"},
	// 		{"B", "C"},
	// 		{"C", "A"},
	// 	},
	// 	Start:    "A",
	// 	Expected: []string{"A", "B", "C", "A", "C"},
	// },
}

func TestItinerary(t *testing.T) {
	for _, c := range testCases {
		res := itinerary(c.Given, c.Start)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v, values %v, '%s'", res, c.Expected, c.Given, c.Start)
		}
	}
}
