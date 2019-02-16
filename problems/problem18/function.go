package problem18

func solver(arr []int, k int) []int {
	maxes := make([]int, 0)
	n := len(arr)
	for i := 0; i < n; i++ {
		max := -1
		if i+k-1 < n {
			for j := i; j < i+k; j++ {
				if arr[j] > max {
					max = arr[j]
				}
			}
			maxes = append(maxes, max)
		}
	}
	return maxes
}

type queue struct {
	items []int
}

func (q *queue) pushBack(n int) {
	q.items = append(q.items, n)
}

func (q *queue) front() int {
	return q.items[0]
}

func (q *queue) back() int {
	return q.items[len(q.items)-1]
}

func (q *queue) popFront() int {
	i := len(q.items) - 1
	defer func() {
		q.items = append(q.items[:i], q.items[i+1:]...)
	}()
	return q.items[i]
}

func (q *queue) popBack() int {
	defer func() {
		q.items = q.items[1:]
	}()
	return q.items[0]
}

func (q *queue) nonEmpty() bool {
	return len(q.items) > 0
}

func withDeque(arr []int, k int) []int {
	maxes := make([]int, 0)
	q := &queue{}
	for i := 0; i < k; i++ {
		for q.nonEmpty() && arr[i] >= arr[q.back()] {
			q.popBack()
		}
		q.pushBack(i)
	}
	for i := k; i < len(arr); i++ {
		maxes = append(maxes, arr[q.front()])
		for q.nonEmpty() && q.front() <= i-k {
			q.popFront()
		}
		for q.nonEmpty() && arr[i] >= arr[q.back()] {
			q.popBack()
		}
		q.pushBack(i)
	}
	maxes = append(maxes, arr[q.front()])
	return maxes
}
