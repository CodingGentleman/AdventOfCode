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
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] != "A" {
				continue
			}
			southEast := grid[i+1][j+1]
			northWest := grid[i-1][j-1]
			southWest := grid[i+1][j-1]
			northEast := grid[i-1][j+1]
			if (southEast == "M" && northWest == "S" || southEast == "S" && northWest == "M") &&
				(southWest == "M" && northEast == "S" || southWest == "S" && northEast == "M") {
				sum += 1
			}
		}
	}
	fmt.Println("sum: ", sum)
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
