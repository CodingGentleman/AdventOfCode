package main

import (
	"fmt"
	"internal/util"
)

func main() {
	grid := util.ReadAsGrid("input", func(r rune) string { return string(r) })
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// east
			if len(grid[i]) >= j+4 && "XMAS" == grid[i][j]+grid[i][j+1]+grid[i][j+2]+grid[i][j+3] {
				sum += 1
			}
			// southeast
			if len(grid) >= i+4 && len(grid[i]) >= j+4 && "XMAS" == grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2]+grid[i+3][j+3] {
				sum += 1
			}
			// south
			if len(grid) >= i+4 && "XMAS" == grid[i][j]+grid[i+1][j]+grid[i+2][j]+grid[i+3][j] {
				sum += 1
			}
			// southwest
			if len(grid) >= i+4 && j-4 >= -1 && "XMAS" == grid[i][j]+grid[i+1][j-1]+grid[i+2][j-2]+grid[i+3][j-3] {
				sum += 1
			}
			// west
			if j-4 >= -1 && "XMAS" == grid[i][j]+grid[i][j-1]+grid[i][j-2]+grid[i][j-3] {
				sum += 1
			}
			// northwest
			if i-4 >= -1 && j-3 >= 0 && "XMAS" == grid[i][j]+grid[i-1][j-1]+grid[i-2][j-2]+grid[i-3][j-3] {
				sum += 1
			}
			// north
			if i-4 >= -1 && "XMAS" == grid[i][j]+grid[i-1][j]+grid[i-2][j]+grid[i-3][j] {
				sum += 1
			}
			// northeast
			if i-4 >= -1 && len(grid[i]) >= j+4 && "XMAS" == grid[i][j]+grid[i-1][j+1]+grid[i-2][j+2]+grid[i-3][j+3] {
				sum += 1
			}
		}
	}
	fmt.Println("sum: ", sum)
}
