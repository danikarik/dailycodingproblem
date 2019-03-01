package problem30

func capacity(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}

	total := 0
	maxIdx := maxIndex(arr)

	leftMax := arr[0]
	for i := 1; i < maxIdx; i++ {
		total += leftMax - arr[i]
		leftMax = max(leftMax, arr[i])
	}

	rightMax := arr[n-1]
	for i := n - 2; i >= maxIdx; i-- {
		total += rightMax - arr[i]
		rightMax = max(rightMax, arr[i])
	}

	return total
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxIndex(arr []int) int {
	temp := 0
	idx := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] >= temp {
			temp = arr[i]
			idx = i
		}
	}
	return idx
}
