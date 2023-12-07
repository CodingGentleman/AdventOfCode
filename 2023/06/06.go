package main

import (
    "fmt"
)

func main() {
    fmt.Println("Day 6 - Part 1")
    vals := [][]int{
        {7, 9},
        {15, 40},
        {30, 200},
    }

    result := 1
    for _, v := range vals {
        result *= calc(v[0], v[1])
    }
    fmt.Println("result: ", result)

    fmt.Println("Day 6 - Part 2")
    fmt.Println("result: ", calc(71530, 940200))
    fmt.Println("result: ", calc(42686985, 284100511221341))
}

func calc(x int, y int) int {
    count := 0
    for v := 0; v <= x; v++ {
        distance := v * (x - v)
        if distance > y {
            count++
        }
    }
    return count
}
