package problem34

import (
	"strings"
)

func palindrome(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	table := make([][]string, n+1)
	for i := 0; i < n+1; i++ {
		row := make([]string, n+1)
		table[i] = row
	}

	for i := 0; i < n; i++ {
		table[i][1] = string(s[i])
	}

	for j := 2; j < n+1; j++ {
		for i := 0; i < n-j+1; i++ {
			term := s[i : i+j]
			first, last := string(term[0]), string(term[len(term)-1])
			if first == last {
				table[i][j] = first + table[i+1][j-2] + last
				continue
			}
			one := first + table[i+1][j-1] + first
			two := last + table[i][j-1] + last
			if len(one) < len(two) {
				table[i][j] = one
			} else if len(one) > len(two) {
				table[i][j] = two
			} else {
				table[i][j] = min(one, two)
			}
		}
	}

	return table[0][n]
}

func min(a, b string) string {
	if strings.Compare(a, b) < 0 {
		return a
	}
	return b
}
