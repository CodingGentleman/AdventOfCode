package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    // "time"
)

func main() {
    fmt.Println("Day 10 - Part 1")
    file, _ := os.Open("input")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    start := make([]int, 2)

    for y := 0; scanner.Scan(); y++ {
        grid = append(grid, make([]string, len(scanner.Text())))
        for x, c := range scanner.Text() {
            grid[y][x] = string(c)
            if string(c) == "S" {
                start[0] = x
                start[1] = y
            }
        }
    }
    runner := make([]Runner,0)
    // find start points for the runners
    // check east of start if any west path exists
    if strings.Contains(ew+nw+sw, get(start[0]+1, start[1])) {
        newRunner := Runner{[]int{start[0]+1, start[1]}, 1}
        runner = append(runner, newRunner)
    }
    // check south of start if any north path exists
    if strings.Contains(ns+ne+nw, get(start[0], start[1]+1)) {
        newRunner := Runner{[]int{start[0], start[1]+1}, 1}
        runner = append(runner, newRunner)
    }
    // check west of start if any east path exists
    if strings.Contains(ew+ne+se, get(start[0]-1, start[1])) {
        newRunner := Runner{[]int{start[0]-1, start[1]}, 1}
        runner = append(runner, newRunner)
    }
    // check north of start if any south path exists
    if strings.Contains(ns+sw+se, get(start[0], start[1]-1)) {
        newRunner := Runner{[]int{start[0], start[1]-1}, 1}
        runner = append(runner, newRunner)
    }
    // set start to #
    set(start[0], start[1])
    fmt.Println("R:", runner)

    for runner[0].curr[0] != runner[1].curr[0] || runner[0].curr[1] != runner[1].curr[1] {
        // clearScreen()
        // fmt.Println("R1 (", runner[0].curr, ", ", runner[0].steps, ")", "R2 (", runner[1].curr, ", ", runner[1].steps, ")")
        // printGrid()
        for i := 0; i < len(runner); i++ {
            runner[i].move()
        }
        // wait for 500ms
        // time.Sleep(50 * time.Millisecond)
    }

    fmt.Println("R:", runner)
}

var grid = make([][]string, 0)
var burned = "#"

func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

func printGrid() {
    for _, row := range grid {
        for _, col := range row {
            fmt.Print(col)
        }
        fmt.Println()
    }
}

func get(x, y int) string {
    return grid[y][x]
}

func set(x, y int) {
    grid[y][x] = burned
}

type Runner struct {
    curr []int
    steps int
}

func (r *Runner) move() {
    currField := get(r.curr[0], r.curr[1])
    newX := r.curr[0]
    newY := r.curr[1]
    if currField == ns {
        if get(r.curr[0], r.curr[1]-1) != burned {
            newY = r.curr[1]-1
        } else {
            newY = r.curr[1]+1
        }
    }
    if currField == ew {
        if get(r.curr[0]-1, r.curr[1]) != burned {
            newX = r.curr[0]-1
        } else {
            newX = r.curr[0]+1
        }
    }
    if currField == ne {
        if get(r.curr[0], r.curr[1]-1) != burned {
            newY = r.curr[1]-1
        } else {
            newX = r.curr[0]+1
        }
    }
    if currField == nw {
        if get(r.curr[0], r.curr[1]-1) != burned {
            newY = r.curr[1]-1
        } else {
            newX = r.curr[0]-1
        }
    }
    if currField == sw {
        if get(r.curr[0], r.curr[1]+1) != burned {
            newY = r.curr[1]+1
        } else {
            newX = r.curr[0]-1
        }
    }
    if currField == se {
        if get(r.curr[0], r.curr[1]+1) != burned {
            newY = r.curr[1]+1
        } else {
            newX = r.curr[0]+1
        }
    }
    set(r.curr[0], r.curr[1])
    r.curr[0] = newX
    r.curr[1] = newY
    r.steps++
}

var ns = "|"
var ew = "-"
var ne = "L"
var nw = "J"
var sw = "7"
var se = "F"
