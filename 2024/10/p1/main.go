package main

import (
	"fmt"
	"image"
	"internal/util"
)

func main() {
	grid := make([][]int, 0)
	util.StreamFileWithIndex("input", func(y int, line string) {
		grid = append(grid, make([]int, len(line)))
		for x, c := range line {
			grid[y][x] = util.ToInt(string(c))
		}
	})
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				set := make(map[image.Point]struct{})
				score(set, grid, x, y)
				sum += len(set) - 1
			}
		}
	}
	fmt.Println(sum)
}

func score(set map[image.Point]struct{}, grid [][]int, x, y int) image.Point {
	if grid[y][x] == 9 {
		return image.Point{x, y}
	}
	if y-1 >= 0 && grid[y][x]+1 == grid[y-1][x] {
		set[score(set, grid, x, y-1)] = struct{}{}
	}
	if x+1 < len(grid[0]) && grid[y][x]+1 == grid[y][x+1] {
		set[score(set, grid, x+1, y)] = struct{}{}
	}
	if y+1 < len(grid) && grid[y][x]+1 == grid[y+1][x] {
		set[score(set, grid, x, y+1)] = struct{}{}
	}
	if x-1 >= 0 && grid[y][x]+1 == grid[y][x-1] {
		set[score(set, grid, x-1, y)] = struct{}{}
	}
	return image.Point{X: -1, Y: -1}
}
