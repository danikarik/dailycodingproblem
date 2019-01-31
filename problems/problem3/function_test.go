package problem3

import (
	"testing"
)

var testNode = &Node{
	Val: "root",
	Left: &Node{
		Val: "left",
		Left: &Node{
			Val:   "left.left",
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	},
	Right: &Node{
		Val:   "right",
		Left:  nil,
		Right: nil,
	},
}

func TestNode(t *testing.T) {
	node := deserialize(serialize(testNode))
	actual := node.Left.Left.Val
	valid := "left.left"
	if actual != valid {
		t.Errorf("failed: got %s, expected %s", actual, valid)
	}
}

var result *Node

func BenchmarkNode(b *testing.B) {
	var n *Node
	for i := 0; i < b.N; i++ {
		n = deserialize(serialize(testNode))
	}
	result = n
}
