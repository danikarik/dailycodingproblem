package problem32

import (
	"math"
)

func arbitrage(table [][]float64) bool {
	n := len(table)
	m := len(table[0])

	transformed := make([][]float64, n)
	for i := 0; i < n; i++ {
		row := make([]float64, m)
		for j := 0; j < m; j++ {
			row[j] = math.Log(table[i][j]) * -1.0
		}
		transformed[i] = row
	}

	source := 0
	minDist := make([]float64, n)
	for i := 0; i < n; i++ {
		minDist[i] = math.Inf(0)
	}
	minDist[source] = 0.0

	for i := 0; i < n-1; i++ {
		for v := 0; v < n; v++ {
			for w := 0; w < n; w++ {
				sum := round(minDist[v]+transformed[v][w], 0.001)
				if minDist[w] > sum {
					minDist[w] = sum
				}
			}
		}
	}

	for v := 0; v < n; v++ {
		for w := 0; w < n; w++ {
			if minDist[w] > minDist[v]+transformed[v][w] {
				return true
			}
		}
	}

	return false
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
