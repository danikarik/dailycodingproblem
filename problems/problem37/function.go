package problem37

func powerSet(set []int) [][]int {
	out := [][]int{[]int{}}
	if len(set) == 0 {
		return out
	}
	result := powerSet(set[1:])
	subset := [][]int{}
	for _, s := range result {
		subset = append(subset, append(s, set[0]))
	}
	return append(result, subset...)
}
