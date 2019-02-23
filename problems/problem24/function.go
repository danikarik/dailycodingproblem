package problem24

// Node represents a locking binary tree node.
type Node struct {
	parent            *Node
	isLocked          bool
	lockedDescendants int
	Value             interface{}
	Left              *Node
	Right             *Node
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

func (n *Node) canLockOrUnlock() bool {
	if n.lockedDescendants > 0 {
		return false
	}
	cur := n.parent
	for cur != nil {
		if cur.IsLocked() {
			return false
		}
		cur = cur.parent
	}
	return true
}

// IsLocked returns whether the node is locked.
func (n *Node) IsLocked() bool {
	return n.isLocked
}

// Lock searches the node's children and parents for a true is_locked attribute.
// If it is set to true on any of them, then return false.
// Otherwise, set the current node's is_locked to true and return true.
func (n *Node) Lock() bool {
	if n.IsLocked() {
		return false
	}
	if !n.canLockOrUnlock() {
		return false
	}
	n.isLocked = true
	cur := n.parent
	for cur != nil {
		cur.lockedDescendants++
		cur = cur.parent
	}
	return true
}

// Unlock simply changes the node's attribute to false. If we want to be safe,
// then we should search the node's children and parents as in lock to make sure we can actually unlock the node,
// but that shouldn't ever happen.
func (n *Node) Unlock() bool {
	if !n.IsLocked() {
		return false
	}
	if !n.canLockOrUnlock() {
		return false
	}
	n.isLocked = false
	cur := n.parent
	for cur != nil {
		cur.lockedDescendants--
		cur = cur.parent
	}
	return true
}
