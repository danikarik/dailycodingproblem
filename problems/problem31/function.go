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
	x := len(p.A) + 1
	y := len(p.B) + 1

	cache := make([][]int, y)
	for j := 0; j < y; j++ {
		r := make([]int, x)
		for i := 0; i < x; i++ {
			r[i] = -1
		}
		cache[j] = r
	}

	for i := 0; i < x; i++ {
		cache[0][i] = i
	}

	for j := 0; j < y; j++ {
		cache[j][0] = j
	}

	for i := 1; i < y; i++ {
		for j := 1; j < x; j++ {
			if p.A[j-1] == p.B[i-1] {
				cache[i][j] = cache[i-1][j-1]
				continue
			}
			cache[i][j] = min(
				cache[i-1][j]+1,
				cache[i][j-1]+1,
				cache[i-1][j-1]+1,
			)
		}
	}

	return cache[y-1][x-1]
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	} else if y < z {
		return y
	}
	return z
}
