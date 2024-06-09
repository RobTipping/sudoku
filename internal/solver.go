package solver

import ()

func ValidGrid(grid [9][9]int) bool {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			if checkSquare(grid, i, j) == false {
				return false
			}
		}
	}
	return true
}

func checkSquare(grid [9][9]int, i, j int) bool {
	// find 3x3
	var a int = (i / 3) * 3
	var b int = (j / 3) * 3
	for k := range grid[i] {
		// Check Row
		if j != k && grid[i][j] == grid[i][k] {
			return false
		}
		// Check Column
		if i != k && grid[i][j] == grid[k][j] {
			return false
		}
		// Check 3x3 grid
		x := a + k/3
		y := b + k%3
		if i == x && j == y {
			continue
		}
		if grid[i][j] == grid[x][y] {
			return false
		}
	}
	return true
}
