package problem29

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return fmt.Sprintf("%v%s", 1, s)
	}
	encoded := ""
	cnt := 1
	seen := rune(s[0])
	for _, c := range s[1:] {
		if c == seen {
			cnt++
			continue
		}
		encoded += strconv.Itoa(cnt) + string(seen)
		seen = c
		cnt = 1
	}
	encoded += strconv.Itoa(cnt) + string(seen)
	return encoded
}

func decode(s string) string {
	n := len(s)
	if n < 2 {
		return ""
	}
	decoded := ""
	cnt := ""
	for _, c := range s {
		char := string(c)
		if isDigit(c) {
			cnt += char
			continue
		}
		n, _ := strconv.Atoi(cnt)
		decoded += strings.Repeat(string(c), n)
		cnt = ""
	}
	return decoded
}

func isDigit(r rune) bool {
	return r >= 49 && r <= 57
}
