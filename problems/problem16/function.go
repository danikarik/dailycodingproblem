package problem16

// SliceContainer holds logs in a simple slice.
type SliceContainer struct {
	n    int
	logs []int
	cur  int
}

// NewSliceContainer createa a new instance of SliceContainer.
func NewSliceContainer(n int) *SliceContainer {
	return &SliceContainer{
		n:    n,
		logs: make([]int, 0),
		cur:  0,
	}
}

// Record adds the order_id to the log.
func (sc *SliceContainer) Record(id int) {
	if len(sc.logs) == sc.n {
		sc.logs[sc.cur] = id
	} else {
		sc.logs = append(sc.logs, id)
	}
	sc.cur = (sc.cur + 1) % sc.n
}

// GetLast gets the ith last element from the log.
// i is guaranteed to be smaller than or equal to N.
func (sc *SliceContainer) GetLast(i int) int {
	idx := sc.cur - i
	if idx < 0 {
		return sc.logs[len(sc.logs)+sc.cur-i]
	}
	return sc.logs[sc.cur-i]
}
