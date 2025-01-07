package main

func equalPairs(grid [][]int) int {
	n := len(grid)
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			match := true
			for k := 0; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					match = false
					break
				}
			}
			if match {
				result++
			}

		}
	}
	return result
}
