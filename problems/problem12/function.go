package problem12

func climb(n int, x []int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	for stair := 1; stair < n+1; stair++ {
		sum := 0
		for _, step := range x {
			if stair-step >= 0 {
				sum += nums[stair-step]
			}
		}
		nums[stair] += sum
	}
	return nums[n]
}
