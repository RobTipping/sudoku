package solver

import ()

var originalGrid [9][9]int
var solvedGrid [9][9]int

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

func SolveGrid(grid [9][9]int) ([9][9]int, bool) {
	originalGrid = grid
	if findGrid(grid, 0) == 1 {
		return solvedGrid, true
	}
	return grid, false
}

func findGrid(grid [9][9]int, pos int) int {
	x := int(pos / 9)
	y := int(pos % 9)
	if originalGrid[x][y] != 0 {
		if pos < 80 {
			return findGrid(grid, pos+1)
		}
		if checkSquare(grid, x, y) == true {
			solvedGrid = grid
			return 1
		}
		return 0
	}
	for i := 1; i < 10; i++ {
		grid[x][y] = i
		if checkSquare(grid, x, y) == true {
			if pos == 80 {
				solvedGrid = grid
				return 1
			}
			if findGrid(grid, pos+1) == 1 {
				return 1
			}
		}
	}
	return 0
}
