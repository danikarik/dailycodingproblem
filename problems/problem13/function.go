package problem13

import (
	"math"
	"strings"
)

func solver(k int, s string) int {
	if k == 0 {
		return 0
	}
	bounds := [2]int{0, 0}
	h := map[string]int{}
	maxLength := 0
	for i, r := range strings.ToLower(s) {
		h[string(r)] = i
		newLowerBound := bounds[0]
		if len(h) > k {
			key := min(h)
			newLowerBound = h[key] + 1
			delete(h, key)
		}
		bounds[0], bounds[1] = newLowerBound, bounds[1]+1
		maxLength = max(maxLength, bounds[1]-bounds[0])
	}
	if len(h) < k {
		return -1
	}
	return maxLength
}

func min(h map[string]int) string {
	var r string
	temp := math.MaxInt32
	for k, v := range h {
		if v < temp {
			temp = v
			r = k
		}
	}
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
