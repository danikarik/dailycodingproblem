package problem27

func balance(s string) bool {
	stack := make([]string, 0)
	for _, ch := range s {
		char := string(ch)
		if char == "(" || char == "[" || char == "{" {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if (char == ")" && stack[len(stack)-1] != "(") ||
				(char == "]" && stack[len(stack)-1] != "[") ||
				(char == "}" && stack[len(stack)-1] != "{") {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
