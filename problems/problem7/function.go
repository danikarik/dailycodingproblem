package problem7

func solver(s string, n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	cnt := 0
	s1 := string(s[n-1])
	s2 := string(s[n-2])
	if s1 > "0" {
		cnt = solver(s, n-1)
	}
	if s2 == "1" || (s2 == "2" && s1 < "7") {
		cnt = cnt + solver(s, n-2)
	}
	return cnt
}
