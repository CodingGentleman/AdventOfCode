package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Println("Day 10 - Part 2")
    file, _ := os.Open("input")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    startPos := []int{0, 0}
    for i := 0; scanner.Scan(); i++ {
        // expand grid by 3 rows
        grid = append(grid, make([]string, 0))
        grid = append(grid, make([]string, 0))
        grid = append(grid, make([]string, 0))
       
        line := scanner.Text()
        for _, c := range line {
            if c == 'S' {
                startPos[0] = i*3+1
                startPos[1] = len(grid[i*3])+1
            }
            value := explosion[string(c)]
            grid[i*3] = append(grid[i*3], value[0]...)
            grid[i*3+1] = append(grid[i*3+1], value[1]...)
            grid[i*3+2] = append(grid[i*3+2], value[2]...)
        }
    }
    // fix start position
    grid[startPos[0]][startPos[1]] = "#"

    if grid[startPos[0]-2][startPos[1]] == "#" {
        grid[startPos[0]-1][startPos[1]] = "#"
    }
    if grid[startPos[0]+2][startPos[1]] == "#" {
        grid[startPos[0]+1][startPos[1]] = "#"
    }
    if grid[startPos[0]][startPos[1]-2] == "#" {
        grid[startPos[0]][startPos[1]-1] = "#"
    }
    if grid[startPos[0]][startPos[1]+2] == "#" {
        grid[startPos[0]][startPos[1]+1] = "#"
    }

    // recursively fill the non-enclosed area with spaces
    fill(0, 0)

    // now lets walk through the grid by 3x3 steps and if a step does not contain a space, count it
    count := 0
    for i := 0; i < len(grid); i += 3 {
        for j := 0; j < len(grid[i]); j += 3 {
            if grid[i][j] != " " && grid[i+1][j] != " " && grid[i+2][j] != " " && grid[i][j+1] != " " && grid[i+1][j+1] != " " && grid[i+2][j+1] != " " && grid[i][j+2] != " " && grid[i+1][j+2] != " " && grid[i+2][j+2] != " " {
                count++
            }
        }
    }
    fmt.Println("Count: ", count)
}

// recursively fill the non-enclosed area with spaces
// step in every direction and if the point is a dot, fill it with a space and continue
func fill(x int, y int) {
    if grid[x][y] == "." {
        grid[x][y] = " "
        
        if isPoint(x-1, y) {
            fill(x-1, y)
        }
        if isPoint(x+1, y) {
            fill(x+1, y)
        }
        if isPoint(x, y-1) {
            fill(x, y-1)
        }
        if isPoint(x, y+1) {
            fill(x, y+1)
        }
    }
}

// check if the point is a dot and not out of bounds
func isPoint(x int, y int) bool {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
        return false
    }
    if grid[x][y] == "." {
        return true
    }
    return false
}

var grid = make([][]string, 0)

var explosion = map[string][][]string{
    "|": {
        {".", "#", "."},
        {".", "#", "."},
        {".", "#", "."},
    },
    "-": {
        {".", ".", "."},
        {"#", "#", "#"},
        {".", ".", "."},
    },
    "L": {
        {".", "#", "."},
        {".", "#", "#"},
        {".", ".", "."},
    },
    "J": {
        {".", "#", "."},
        {"#", "#", "."},
        {".", ".", "."},
    },
    "7": {
        {".", ".", "."},
        {"#", "#", "."},
        {".", "#", "."},
    },
    "F": {
        {".", ".", "."},
        {".", "#", "#"},
        {".", "#", "."},
    },
    ".": {
        {".", ".", "."},
        {".", ".", "."},
        {".", ".", "."},
    },
    "S": {
        {".", ".", "."},
        {".", "S", "."},
        {".", ".", "."},
    },
}

