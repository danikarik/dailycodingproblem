package problem36

// Node represents a binary tree node.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func helper(root *Node, c, r *int) {
	if root == nil || *c >= 2 {
		return
	}
	helper(root.Right, c, r)
	*c++
	if *c == 2 {
		*r = root.Value
	}
	helper(root.Left, c, r)
}

func secondLargest(root *Node) int {
	c := 0
	r := 0
	helper(root, &c, &r)
	return r
}
