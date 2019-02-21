package problem22

func findSentence(s string, dict map[string]struct{}) []string {
	starts := map[int]string{0: ""}
	for i := 0; i < len(s)+1; i++ {
		newStarts := starts
		for k := range starts {
			word := s[k:i]
			if _, ok := dict[word]; ok {
				newStarts[i] = word
			}
		}
		starts = newStarts
	}
	result := []string{}
	currLen := len(s)
	if _, ok := starts[currLen]; !ok {
		return nil
	}
	for currLen > 0 {
		word := starts[currLen]
		currLen -= len(word)
		result = append(result, word)
	}
	return reverse(result)
}

func reverse(arr []string) []string {
	out := make([]string, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		out = append(out, arr[i])
	}
	return out
}
