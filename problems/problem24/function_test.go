package problem24

import (
	"testing"
)

func TestIsLocked(t *testing.T) {
	testCases := []struct {
		Given    *Node
		Expected bool
	}{
		{
			Given: &Node{
				isLocked: true,
			},
			Expected: true,
		},
		{
			Given:    &Node{},
			Expected: false,
		},
	}
	for _, c := range testCases {
		res := c.Given.IsLocked()
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

func TestAddChild(t *testing.T) {
	testCases := []struct {
		Given         *Node
		Left          *Node
		Right         *Node
		ExpectedLeft  *Node
		ExpectedRight *Node
	}{
		{
			Given:         &Node{Value: 1},
			Left:          &Node{Value: 2, isLocked: true},
			Right:         (&Node{Value: 3}).AddChild(&Node{Value: 4}),
			ExpectedLeft:  &Node{Value: 2, isLocked: true},
			ExpectedRight: (&Node{Value: 3}).AddChild(&Node{Value: 4}),
		},
	}
	for _, c := range testCases {
		res := c.Given.AddChild(c.Left)
		res = res.AddChild(c.Right)
		if res.Left.Value != c.ExpectedLeft.Value {
			t.Errorf("failed left node: got %v, expected %v", res.Left.Value, c.ExpectedLeft.Value)
		}
		if res.Right.Value != c.ExpectedRight.Value {
			t.Errorf("failed right node: got %v, expected %v", res.Right.Value, c.ExpectedRight.Value)
		}
	}
}

func TestLock(t *testing.T) {
	var testCases = []struct {
		Given    *Node
		Expected bool
	}{
		{
			Given:    (&Node{Value: 1, lockedDescendants: 1}).AddChild(&Node{Value: 2}).AddChild(&Node{Value: 3, isLocked: true}),
			Expected: false,
		},
		{
			Given:    (&Node{Value: 1}).AddChild(&Node{Value: 2}).AddChild(&Node{Value: 3}),
			Expected: true,
		},
		{
			Given:    (&Node{Value: 1, lockedDescendants: 1}).AddChild((&Node{Value: 2, lockedDescendants: 1}).AddChild(&Node{Value: 4}).AddChild(&Node{Value: 5, isLocked: true})).AddChild(&Node{Value: 3}),
			Expected: false,
		},
	}
	for _, c := range testCases {
		res := c.Given.Lock()
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

func TestUnlock(t *testing.T) {
	var testCases = []struct {
		Given    *Node
		Expected bool
	}{
		{
			Given:    (&Node{Value: 1, isLocked: true, lockedDescendants: 1}).AddChild(&Node{Value: 2}).AddChild(&Node{Value: 3, isLocked: true}),
			Expected: false,
		},
		{
			Given:    (&Node{Value: 1, isLocked: true}).AddChild(&Node{Value: 2}).AddChild(&Node{Value: 3}),
			Expected: true,
		},
		{
			Given:    (&Node{Value: 1, isLocked: true, lockedDescendants: 1}).AddChild((&Node{Value: 2, lockedDescendants: 1}).AddChild(&Node{Value: 4}).AddChild(&Node{Value: 5, isLocked: true})).AddChild(&Node{Value: 3}),
			Expected: false,
		},
	}
	for _, c := range testCases {
		res := c.Given.Unlock()
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}
