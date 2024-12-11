package main

import (
	"fmt"
	"internal/util"
)

var grid = make([][]int, 0)
var count = 0

func main() {
	util.StreamFileWithIndex("input", func(y int, line string) {
		grid = append(grid, make([]int, len(line)))
		for x, c := range line {
			grid[y][x] = util.ToInt(string(c))
		}
	})
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				score(x, y)
			}
		}
	}
	fmt.Println(count)
}

func score(x, y int) {
	if grid[y][x] == 9 {
		count++
	}
	if y-1 >= 0 && grid[y][x]+1 == grid[y-1][x] {
		score(x, y-1)
	}
	if x+1 < len(grid[0]) && grid[y][x]+1 == grid[y][x+1] {
		score(x+1, y)
	}
	if y+1 < len(grid) && grid[y][x]+1 == grid[y+1][x] {
		score(x, y+1)
	}
	if x-1 >= 0 && grid[y][x]+1 == grid[y][x-1] {
		score(x-1, y)
	}
}
