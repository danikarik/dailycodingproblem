package problem26

// Node represents a list's item structure.
type Node struct {
	Value interface{}
	Next  *Node
}

// List is singly linked list.
type List struct {
	len   int
	items []*Node
}

// NewList creates a new instance of list.
func NewList() *List {
	return &List{
		len:   0,
		items: make([]*Node, 0),
	}
}

// Len returns list's current size.
func (l *List) Len() int {
	return l.len
}

// Add inserts new item.
func (l *List) Add(val interface{}) {
	node := &Node{
		Value: val,
		Next:  nil,
	}
	l.items = append(l.items, node)
	l.len++
	if l.len > 0 {
		l.items[l.len-1].Next = node
	}
}

// Del removes the kth last element from the list
func (l *List) Del(k int) {
	l.items = l.items[:l.len-k]
	l.len = l.len - k
}
