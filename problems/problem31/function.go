package problem31

import "fmt"

// Pair holds two strings to compare.
type Pair struct {
	A string
	B string
}

func (p Pair) String() string {
	return fmt.Sprintf("(%s, %s)", p.A, p.B)
}

func editDistance(p Pair) int {
	n := len(p.A)
	m := len(p.B)
	cnt := diff(n, m)
	for i := 0; i < min(n, m); i++ {
		if p.A[i] != p.B[i] {
			cnt++
		}
	}
	return cnt
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func diff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
