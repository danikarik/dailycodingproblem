package problem28

import (
	"strings"
)

func justifyText(words []string, k int) []string {
	size := 0
	curr := make([]string, 0)
	lines := make([]string, 0)
	for _, w := range words {
		if size+len(w) > k {
			size = len(w) + 1
			lines = append(lines, fillSpaces(curr, k))
			curr = []string{w}
			continue
		}
		size += len(w) + 1
		curr = append(curr, w)
	}
	lines = append(lines, fillSpaces(curr, k))
	return lines
}

func fillSpaces(words []string, k int) string {
	n := len(words)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return words[0] + strings.Repeat(" ", k-len(words[0]))
	}
	line := strings.Join(words, " ")
	if len(line) < k {
		less := true
		i := 0
		for less {
			if i+1 == n {
				i = 0
				continue
			}
			words[i] = words[i] + " "
			i++
			line = strings.Join(words, " ")
			if len(line) == k {
				less = false
				continue
			}
		}
	}
	return line
}
