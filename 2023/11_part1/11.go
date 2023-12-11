package main

import (
    "fmt"
    "strings"
    _ "embed"
)

//go:embed input
var input string

func main() {
    fmt.Println("Day 11 - Part 1")

    points := make([]Point, 0)
    grid := strings.Split(input, "\n")
    expansionRows := []int{}
    expansionCols := []int{}
    maxCol := len(grid[0])
    for i := 0; i < maxCol; i++ {
        expansionCols = append(expansionCols, 1)
    }
    maxRow := len(grid)
    for i := 0; i < maxRow; i++ {
        expansionRows = append(expansionRows, 1)
    }

    for y, line := range grid {
        values := strings.Split(line, "")
        for x, value := range values {
            if value == "#" {
                expansionRows[y] = 0
                expansionCols[x] = 0
                points = append(points, Point{x, y})
            }
        }
    }
    
    for i := 1; i < len(expansionRows); i++ {
        expansionRows[i] = expansionRows[i] + expansionRows[i-1]
    }
    for i := 1; i < len(expansionCols); i++ {
        expansionCols[i] = expansionCols[i] + expansionCols[i-1]
    }
    for i := 0; i < len(points); i++ {
        points[i].x += expansionCols[points[i].x]
        points[i].y += expansionRows[points[i].y]
    }
    fmt.Println(points)

    for i := 0; i < len(points); i++ {
        from := points[i]
        for j := i+1; j < len(points); j++ {
            to := points[j]
            distance := calculateDistance(from, to)
            fmt.Println(from, "-", to, ":", distance)
            
        }
    }
}

type Point struct {
    x int
    y int
}

func calculateDistance(from Point, to Point) int {
    return 0
}
