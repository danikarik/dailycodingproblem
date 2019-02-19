package problem19

import (
	"math"
)

func brute(matrix [][]int) int {
	k := len(matrix[0])
	solnRow := make([]int, k)

	for _, row := range matrix {
		newRow := make([]int, 0)
		for c, val := range row {
			min := math.MaxInt32
			for i := 0; i < k; i++ {
				if i != c && solnRow[i] < min {
					min = solnRow[i]
				}
			}
			newRow = append(newRow, min+val)
		}
		solnRow = newRow
	}

	min := math.MaxInt32
	for _, cost := range solnRow {
		if cost < min {
			min = cost
		}
	}

	return min
}

func elegant(matrix [][]int) int {
	minCost, minIndex := 0, -1
	secondCost := 0

	for _, row := range matrix {
		tempCost, idx := math.MaxInt32, -1
		tempSecondCost := math.MaxInt32
		for i, val := range row {
			var prevMinCost int
			if i == minIndex {
				prevMinCost = secondCost
			} else {
				prevMinCost = minCost
			}
			cost := prevMinCost + val
			if cost < tempCost {
				tempSecondCost = tempCost
				tempCost, idx = cost, i
			} else if cost < tempSecondCost {
				tempSecondCost = cost
			}
		}
		minCost = tempCost
		minIndex = idx
		secondCost = tempSecondCost
	}

	return minCost
}
