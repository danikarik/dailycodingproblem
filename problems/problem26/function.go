package problem26

import (
	"fmt"
)

// Node represents a list's item structure.
type Node struct {
	Value interface{}
	Next  *Node
}

func (n *Node) String() string {
	cur := n
	result := []interface{}{}
	for cur != nil {
		result = append(result, cur.Value)
		cur = cur.Next
	}
	return fmt.Sprintf("%v", result)
}

// IsEqual compares given node with current.
func (n *Node) IsEqual(x *Node) bool {
	cur := n
	for cur != nil {
		if cur.Value != x.Value {
			return false
		}
		cur = cur.Next
		x = x.Next
	}
	return x == nil
}

// RemoveKth removes the kth last element from the list.
func RemoveKth(head *Node, k int) {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	prev := &Node{}
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = slow.Next
}
