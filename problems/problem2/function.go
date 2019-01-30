package problem2

func solver(in []int) []int {
	out := make([]int, len(in))
	n := 1
	for i := 0; i < len(in); i++ {
		out[i] = n
		n *= in[i]
	}
	n = 1
	for i := len(in) - 1; i >= 0; i-- {
		out[i] *= n
		n *= in[i]
	}
	return out
}
