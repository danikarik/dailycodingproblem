package problem08

// Node is a tree node.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func solver(root *Node) int {
	cnt, _ := helper(root)
	return cnt
}

func helper(root *Node) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftCount, isLeftUnival := helper(root.Left)
	rightCount, isRightUnival := helper(root.Right)
	totalCount := leftCount + rightCount
	if isLeftUnival && isRightUnival {
		if root.Left != nil && root.Value != root.Left.Value {
			return totalCount, false
		}
		if root.Right != nil && root.Value != root.Right.Value {
			return totalCount, false
		}
		return totalCount + 1, true
	}
	return totalCount, false
}
