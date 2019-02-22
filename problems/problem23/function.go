package problem23

import (
	"fmt"
)

type coords struct {
	row int
	col int
}

func (c coords) String() string {
	return fmt.Sprintf("(%v, %v)", c.row, c.col)
}

func (c coords) Key() string {
	return fmt.Sprintf("%v-%v", c.row, c.col)
}

type node struct {
	coords coords
	dist   int
}

type queue struct {
	items []node
}

func (q *queue) pushBack(n node) {
	q.items = append(q.items, n)
}

func (q *queue) front() node {
	return q.items[0]
}

func (q *queue) back() node {
	return q.items[len(q.items)-1]
}

func (q *queue) popFront() node {
	i := len(q.items) - 1
	defer func() {
		q.items = append(q.items[:i], q.items[i+1:]...)
	}()
	return q.items[i]
}

func (q *queue) popBack() node {
	defer func() {
		q.items = q.items[1:]
	}()
	return q.items[0]
}

func (q *queue) nonEmpty() bool {
	return len(q.items) > 0
}

func walkable(board [][]bool, row, col int) bool {
	if row < 0 || row >= len(board) {
		return false
	}
	if col < 0 || col >= len(board[0]) {
		return false
	}
	return !board[row][col]
}

func walkableNeighbours(board [][]bool, row, col int) []coords {
	out := make([]coords, 0)
	neighbours := []coords{
		coords{row, col - 1},
		coords{row - 1, col},
		coords{row + 1, col},
		coords{row, col + 1},
	}
	for _, n := range neighbours {
		if walkable(board, n.row, n.col) {
			out = append(out, n)
		}
	}
	return out
}

func steps(board [][]bool, start, end coords) int {
	seen := make(map[string]struct{})
	q := &queue{}
	q.pushBack(node{start, 0})
	for q.nonEmpty() {
		temp := q.popFront()
		curr, dist := temp.coords, temp.dist
		if curr.row == end.row && curr.col == end.col {
			return dist
		}
		neighbours := walkableNeighbours(board, curr.row, curr.col)
		for _, neighbour := range neighbours {
			key := neighbour.Key()
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				q.pushBack(node{neighbour, dist + 1})
			}
		}
	}
	return -1
}
