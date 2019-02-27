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

func minLine(words []string) string {
	return strings.Join(words, " ")
}

func groupLines(words []string, k int) [][]string {
	groups := make([][]string, 0)
	cur := make([]string, 0)
	for _, word := range words {
		if len(minLine(append(cur, word))) > k {
			groups = append(groups, cur)
			cur = make([]string, 0)
		}
		cur = append(cur, word)
	}
	groups = append(groups, cur)
	return groups
}

func justify(words []string, k int) string {
	if len(words) == 1 {
		word := words[0]
		numSpaces := k - len(word)
		spaces := strings.Repeat(" ", numSpaces)
		return word + spaces
	}
	sum := 0
	for _, word := range words {
		sum += len(word)
	}
	spacesToDistribute := k - sum
	numberOfSpaces := len(words) - 1
	smallestSpace := spacesToDistribute / numberOfSpaces
	leftoverSpaces := spacesToDistribute - (numberOfSpaces * smallestSpace)
	justifiedWords := make([]string, 0)
	for _, word := range words {
		justifiedWords = append(justifiedWords, word)
		currentSpace := strings.Repeat(" ", smallestSpace)
		if leftoverSpaces > 0 {
			currentSpace += " "
			leftoverSpaces--
		}
		justifiedWords = append(justifiedWords, currentSpace)
	}
	return strings.TrimRight(strings.Join(justifiedWords, ""), " ")
}

func justifyText2(words []string, k int) []string {
	out := make([]string, 0)
	for _, group := range groupLines(words, k) {
		out = append(out, justify(group, k))
	}
	return out
}
