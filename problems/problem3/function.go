package problem3

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Node is a tree node.
type Node struct {
	Val   string `json:"val"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}

// Iter holds serialized node data.
type Iter struct {
	Base []string
}

// Next returns next val.
func (i *Iter) Next() string {
	l := len(i.Base)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return i.Base[0]
	}
	v := i.Base[0]
	i.Base = i.Base[1:]
	return v
}

func serialize(root *Node) string {
	if root == nil {
		return "#"
	}
	return fmt.Sprintf("%s %s %s", root.Val, serialize(root.Left), serialize(root.Right))
}

func deserialize(data string) *Node {
	iter := &Iter{Base: strings.Split(data, " ")}
	return helper(iter)
}

func helper(iter *Iter) *Node {
	val := iter.Next()
	if val == "#" {
		return nil
	}
	node := &Node{Val: val}
	node.Left = helper(iter)
	node.Right = helper(iter)
	return node
}

func serializeJSON(root *Node) string {
	data, err := json.Marshal(root)
	if err != nil {
		return ""
	}
	return string(data)
}

func deserializeJSON(data string) *Node {
	var node Node
	err := json.Unmarshal([]byte(data), &node)
	if err != nil {
		return nil
	}
	return &node
}
