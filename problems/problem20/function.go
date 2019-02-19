package problem20

type list struct {
	root *node
}

type node struct {
	value int
	next  *node
}

func (n *node) isLast() bool {
	return n.next == nil
}

func (n *node) length() int {
	if n.next == nil {
		return 0
	}
	return 1 + n.next.length()
}

func intersect(a, b *list) int {
	m, n := a.root.length(), b.root.length()
	curA, curB := a.root, b.root

	if m > n {
		i := 0
		for i < m-n {
			curA = curA.next
			i++
		}
	} else {
		i := 0
		for i < n-m {
			curB = curB.next
		}
	}

	return walk(curA, curB)
}

func walk(rootA, rootB *node) int {
	if val := rootA.value; val == rootB.value {
		return val
	}
	if rootA.isLast() || rootB.isLast() {
		return -1
	}
	return walk(rootA.next, rootB.next)
}
