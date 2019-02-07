package problem09

func solver(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	incl := nums[0]
	excl, temp := 0, 0
	for i := 1; i < len(nums); i++ {
		temp = max(incl, excl)
		incl = excl + nums[i]
		excl = temp
	}
	return max(incl, excl)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
