package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
    fmt.Println("Day 11 - Part 2")

    points := make([]Point, 0)
    grid := strings.Split(input, "\n")
    expansionRows := []int{}
    expansionCols := []int{}
    maxCol := len(grid[0])
    for i := 0; i < maxCol; i++ {
        expansionCols = append(expansionCols, 999999)
    }
    maxRow := len(grid)
    for i := 0; i < maxRow; i++ {
        expansionRows = append(expansionRows, 999999)
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
    fmt.Println(points) 
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
    fmt.Println(expansionCols)
    fmt.Println(expansionRows)
    fmt.Println(points)

    sum := 0
    for i := 0; i < len(points); i++ {
        from := points[i]
        for j := i+1; j < len(points); j++ {
            to := points[j]
            distance := calculateDistance(from, to)
            fmt.Println(from, "-", to, ":", distance)
            sum += distance
        }
    }
    fmt.Println("Sum:", sum)
}

type Point struct {
    x int
    y int
}

func calculateDistance(from Point, to Point) int {
    x := abs(from.x - to.x)
    y := abs(from.y - to.y)
    return x + y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
