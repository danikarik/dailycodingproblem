package problem05

type pair func() (int, int)

func cons(a, b int) pair {
	return func() (int, int) {
		return a, b
	}
}

func car(p pair) int {
	f, _ := p()
	return f
}

func cdr(p pair) int {
	_, l := p()
	return l
}
