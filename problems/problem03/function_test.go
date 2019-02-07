package problem03

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

func TestNodeJSON(t *testing.T) {
	node := deserializeJSON(serializeJSON(testNode))
	actual := node.Left.Left.Val
	valid := "left.left"
	if actual != valid {
		t.Errorf("failed: got %s, expected %s", actual, valid)
	}
}

var (
	result     *Node
	resultJSON *Node
)

func BenchmarkNode(b *testing.B) {
	var node *Node
	for i := 0; i < b.N; i++ {
		node = deserialize(serialize(testNode))
	}
	result = node
}

func BenchmarkNodeJSON(b *testing.B) {
	var node *Node
	for i := 0; i < b.N; i++ {
		node = deserializeJSON(serializeJSON(testNode))
	}
	resultJSON = node
}
