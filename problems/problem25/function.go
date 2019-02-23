package problem25

func mathesFirstChar(s, r string) bool {
	if len(s) > 0 {
		return s[0] == r[0] || string(r[0]) == "."
	}
	return false
}

func matches(s, r string) bool {
	if len(r) == 0 {
		return len(s) == 0
	}
	if len(r) == 1 || string(r[1]) != "*" {
		// The first character in the regex is not proceeded by a *.
		if mathesFirstChar(s, r) {
			return matches(s[1:], r[1:])
		}
		return false
	}
	// The first character is proceeded by a *.
	// First, try zero length.
	if matches(s, r[2:]) {
		return true
	}
	i := 0
	for mathesFirstChar(s[i:], r) {
		if matches(s[i+1:], r[2:]) {
			return true
		}
		i++
	}
	return false
}
