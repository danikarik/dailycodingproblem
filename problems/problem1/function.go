package problem1

func solver(nums []int, k int) bool {
	d := make(map[int]struct{})
	for _, num := range nums {
		d[num] = struct{}{}
		if _, ok := d[k-num]; ok {
			return true
		}
	}
	return false
}
