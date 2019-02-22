package problem24

// Node represents a binary tree node.
type Node struct {
	parent *Node
	locked bool
	Value  int
	Left   *Node
	Right  *Node
}

// AddChild appends child nodes to root.
func (n *Node) AddChild(child *Node) *Node {
	if n.Left != nil && n.Right != nil {
		return n
	}
	if n.Left == nil {
		child.parent = n
		n.Left = child
		return n
	}
	if n.Right == nil {
		child.parent = n
		n.Right = child
		return n
	}
	return n
}

// IsLocked returns whether the node is locked.
func (n *Node) IsLocked() bool {
	return n.locked
}

// Lock attempts to lock the node.
// If it cannot be locked, then it should return false.
// Otherwise, it should lock it and return true.
func (n *Node) Lock() bool {
	if n.parent != nil && n.parent.locked {
		return false
	}
	if ok := isValid(n); !ok {
		return false
	}
	n.locked = true
	return true
}

// Unlock unlocks the node.
// If it cannot be unlocked, then it should return false.
// Otherwise, it should unlock it and return true.
func (n *Node) Unlock() bool {
	if n.parent != nil && n.parent.locked {
		return false
	}
	if ok := isValid(n); !ok {
		return false
	}
	n.locked = false
	return true
}

func isValid(root *Node) bool {
	leftOk, rightOk := true, true
	if root.Left != nil {
		if root.Left.IsLocked() {
			return false
		}
		leftOk = isValid(root.Left)
	}
	if root.Right != nil {
		if root.Right.IsLocked() {
			return false
		}
		rightOk = isValid(root.Right)
	}
	if leftOk && rightOk {
		return true
	}
	return false
}
