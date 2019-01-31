package problem3

import (
	"encoding/json"
)

// Node is a tree node.
type Node struct {
	Val   string `json:"val"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}

func serialize(n *Node) string {
	data, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(data)
}

func deserialize(s string) *Node {
	var n Node
	err := json.Unmarshal([]byte(s), &n)
	if err != nil {
		return nil
	}
	return &n
}
