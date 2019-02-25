package problem26

import (
	"testing"
)

var testCases = []struct {
	Given    *List
	N        int
	K        int
	Expected int
}{
	{
		Given:    NewList(),
		N:        10,
		K:        7,
		Expected: 3,
	},
	{
		Given:    NewList(),
		N:        7,
		K:        6,
		Expected: 1,
	},
}

func prepList(l *List, n int) {
	for i := 0; i < n; i++ {
		l.Add(i + 1)
	}
}

func TestDel(t *testing.T) {
	for _, c := range testCases {
		prepList(c.Given, c.N)
		c.Given.Del(c.K)
		if c.Given.Len() != c.Expected {
			t.Errorf("failed: got %v, expected %v", c.Given.Len(), c.Expected)
		}
	}
}
