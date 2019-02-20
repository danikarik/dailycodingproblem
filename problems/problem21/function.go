package problem21

import (
	"fmt"
	"sort"
)

type lecture struct {
	start int
	end   int
}

func (l lecture) String() string {
	return fmt.Sprintf("(%v, %v)", l.start, l.end)
}

func rooms(lectures []lecture) int {
	n := len(lectures)
	starts := make([]lecture, n)
	ends := make([]lecture, n)
	copy(starts, lectures)
	copy(ends, lectures)
	sort.SliceStable(starts, func(i, j int) bool {
		return starts[i].start < starts[j].start
	})
	sort.SliceStable(ends, func(i, j int) bool {
		return ends[i].end < ends[j].end
	})
	maxRooms, overlap := 0, 0
	i, j := 0, 0
	for i < n && j < n {
		if starts[i].start < ends[j].end {
			overlap++
			maxRooms = max(maxRooms, overlap)
			i++
		} else {
			overlap--
			j++
		}
	}
	return maxRooms
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
