package problem04

import (
	"sort"
)

// O(n log n)
func solverSort(in []int) int {
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

// O(n)
func solver(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		for i+1 != nums[i] && nums[i] > 0 && nums[i] <= len(nums) {
			v := nums[i]
			nums[i] = nums[v-1]
			nums[v-1] = v
			if nums[i] == nums[v-1] {
				break
			}
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
