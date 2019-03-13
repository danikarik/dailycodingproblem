package problem42

func subsetSum(arr []int, k int) []int {

	n := len(arr)
	table := makeTable(n+1, k+1)

	for i := 0; i < n+1; i++ {
		table[i][0] = []int{}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < k+1; j++ {
			last := arr[i-1]
			if last > j {
				table[i][j] = table[i-1][j]
				continue
			}
			if table[i-1][j] != nil {
				table[i][j] = table[i-1][j]
			} else if table[i-1][j-last] != nil {
				l := table[i-1][j-last].([]int)
				table[i][j] = append(l, last)
			} else {
				table[i][j] = nil
			}
		}
	}

	return table[n][k].([]int)
}

func makeTable(n, k int) [][]interface{} {
	table := make([][]interface{}, n)

	for i := 0; i < n; i++ {
		r := make([]interface{}, k)
		for j := 0; j < k; j++ {
			r[j] = nil
		}
		table[i] = r
	}

	return table
}
