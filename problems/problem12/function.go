package problem12

func climb(n int, x []int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	for i := 1; i < n+1; i++ {
		sum := 0
		for _, n := range x {
			if i-n >= 0 {
				sum += cache[i-n]
			}
		}
		cache[i] += sum
	}
	return cache[n]
}
