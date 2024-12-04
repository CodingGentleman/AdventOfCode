package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid := readFile()
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum += walkEast(i, j, grid)
			sum += walkSouthEast(i, j, grid)
			sum += walkSouth(i, j, grid)
			sum += walkSouthWest(i, j, grid)
			sum += walkWest(i, j, grid)
			sum += walkNorthWest(i, j, grid)
			sum += walkNorth(i, j, grid)
			sum += walkNorthEast(i, j, grid)
		}
	}
	fmt.Println("sum: ", sum)
}

func walkEast(i int, j int, grid [][]string) int {
	if len(grid[i]) < j+4 || "XMAS" != grid[i][j]+grid[i][j+1]+grid[i][j+2]+grid[i][j+3] {
		return 0
	}
	return 1
}
func walkSouthEast(i int, j int, grid [][]string) int {
	if len(grid) < i+4 || len(grid[i]) < j+4 || "XMAS" != grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2]+grid[i+3][j+3] {
		return 0
	}
	return 1
}
func walkSouth(i int, j int, grid [][]string) int {
	if len(grid) < i+4 || "XMAS" != grid[i][j]+grid[i+1][j]+grid[i+2][j]+grid[i+3][j] {
		return 0
	}
	return 1
}
func walkSouthWest(i int, j int, grid [][]string) int {
	if len(grid) < i+4 || j-4 < -1 || "XMAS" != grid[i][j]+grid[i+1][j-1]+grid[i+2][j-2]+grid[i+3][j-3] {
		return 0
	}
	return 1
}
func walkWest(i int, j int, grid [][]string) int {
	if j-4 < -1 || "XMAS" != grid[i][j]+grid[i][j-1]+grid[i][j-2]+grid[i][j-3] {
		return 0
	}
	return 1
}
func walkNorthWest(i int, j int, grid [][]string) int {
	if i-4 < -1 || j-3 < 0 || "XMAS" != grid[i][j]+grid[i-1][j-1]+grid[i-2][j-2]+grid[i-3][j-3] {
		return 0
	}
	return 1
}
func walkNorth(i int, j int, grid [][]string) int {
	if i-4 < -1 || "XMAS" != grid[i][j]+grid[i-1][j]+grid[i-2][j]+grid[i-3][j] {
		return 0
	}
	return 1
}
func walkNorthEast(i int, j int, grid [][]string) int {
	if i-4 < -1 || len(grid[i]) < j+4 || "XMAS" != grid[i][j]+grid[i-1][j+1]+grid[i-2][j+2]+grid[i-3][j+3] {
		return 0
	}
	return 1
}

func readFile() [][]string {
	f, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]string, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		grid = append(grid, make([]string, len(line)))
		for j, c := range line {
			grid[i][j] = string(c)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}
	return grid
}
