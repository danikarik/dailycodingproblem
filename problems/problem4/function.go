package problem4

import (
	"sort"
)

func solver(in []int) int {
	sort.Sort(sort.IntSlice(in))
	next := 1
	for _, n := range in {
		if n > 0 && n == next {
			next++
			continue
		}
	}
	return next
}
