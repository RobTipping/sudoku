package solver


func ValidGrid(grid [9][9]int) bool{
	for i:= range grid {
		for j:= range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			// check rows
			for k:= range grid[i] {
				if j==k {
					continue
				}
				if grid[i][j] == grid[i][k] {
					return false
				}
				if i==k {
					continue
				}
				if grid[i][j] == grid[k][j] {
					return false
				}
			}
			// check columns
			/*
			for k:= range grid {
				if i==k {
					continue
				}
				if grid[i][j] == grid[k][j] {
					return false
				}
			}
		*/
		}
	}
	return true
}
