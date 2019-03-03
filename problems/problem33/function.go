package problem33

import (
	"fmt"
	"math"
	"sort"
)

func median(arr []int) []float64 {
	n := len(arr)
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		out[i] = find(arr[:i+1])
	}
	return out
}

func find(arr []int) float64 {
	n := len(arr)
	x := make([]int, n)
	copy(x, arr)
	sort.Sort(sort.IntSlice(x))
	i := n / 2
	if n%2 == 0 {
		return float64(x[i-1]+x[i]) / 2.0
	}
	return float64(x[i])
}

type heap struct {
	items []int
}

func newHeap() *heap {
	return &heap{
		items: make([]int, 0),
	}
}

func (h *heap) len() int {
	return len(h.items)
}

func (h *heap) findMin() (int, int) {
	temp := math.MaxInt64
	idx := -1
	for i := 0; i < len(h.items); i++ {
		if h.items[i] < temp {
			temp = h.items[i]
			idx = i
		}
	}
	return idx, h.items[idx]
}

func (h *heap) extractMin() int {
	idx, _ := h.findMin()
	defer func() { h.items = append(h.items[:idx], h.items[idx+1:]...) }()
	return h.items[idx]
}

func (h *heap) findMax() (int, int) {
	temp := math.MinInt64
	idx := -1
	for i := 0; i < len(h.items); i++ {
		if h.items[i] > temp {
			temp = h.items[i]
			idx = i
		}
	}
	return idx, h.items[idx]
}

func (h *heap) extractMax() int {
	idx, _ := h.findMax()
	defer func() { h.items = append(h.items[:idx], h.items[idx+1:]...) }()
	return h.items[idx]
}

func (h *heap) insert(num int) {
	h.items = append(h.items, num)
}

func (h *heap) print() {
	fmt.Printf("%v\n", h.items)
}

func getMedian(minHeap, maxHeap *heap) float64 {
	if minHeap.len() > maxHeap.len() {
		_, min := minHeap.findMin()
		return float64(min)
	} else if minHeap.len() < maxHeap.len() {
		_, max := maxHeap.findMax()
		return float64(max)
	}
	_, minRoot := minHeap.findMin()
	_, maxRoot := maxHeap.findMax()
	return float64(minRoot+maxRoot) / 2.0
}

func add(num int, minHeap, maxHeap *heap) {
	if minHeap.len()+maxHeap.len() <= 1 {
		maxHeap.insert(num)
		return
	}
	med := getMedian(minHeap, maxHeap)
	if float64(num) > med {
		minHeap.insert(num)
		return
	}
	maxHeap.insert(num)
	return
}

func rebalance(minHeap, maxHeap *heap) {
	if minHeap.len() > maxHeap.len()+1 {
		root := minHeap.extractMin()
		maxHeap.insert(root)
	} else if maxHeap.len() > minHeap.len()+1 {
		root := maxHeap.extractMax()
		minHeap.insert(root)
	}
}

func printMedian(minHeap, maxHeap *heap) float64 {
	return getMedian(minHeap, maxHeap)
}

func medianWithHeap(arr []int) []float64 {
	out := make([]float64, 0)
	minHeap := newHeap()
	maxHeap := newHeap()
	for _, num := range arr {
		add(num, minHeap, maxHeap)
		rebalance(minHeap, maxHeap)
		med := printMedian(minHeap, maxHeap)
		out = append(out, med)
	}
	return out
}
