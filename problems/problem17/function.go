package problem17

import (
	"path/filepath"
	"strings"
)

const (
	lineSep = "\\n"
	tabSep  = "\\t"
)

type stat struct {
	level int
	path  string
}

func longestPath(s string) int {
	stats := make([]stat, 0)
	path, fname := "", ""
	maxDepth, maxIndex := -1, 0
	for _, line := range strings.Split(s, lineSep) {
		indent := strings.Count(line, tabSep)
		if indent == 0 {
			path = line
		} else {
			filename := strings.TrimLeft(line, tabSep)
			stats = append(stats, stat{
				level: indent,
				path:  filename,
			})
			if isFile(filename) && indent > maxDepth {
				maxDepth = indent
				maxIndex = len(stats) - 1
				fname = filename
			}
		}
	}
	if maxIndex > 0 {
		for _, s := range stats[maxIndex-maxDepth+1 : maxIndex] {
			path = filepath.Join(path, s.path)
		}
		path = filepath.Join(path, fname)
		return len(path)
	}
	return 0
}

func isFile(filename string) bool {
	return strings.Count(filename, ".") > 0
}
