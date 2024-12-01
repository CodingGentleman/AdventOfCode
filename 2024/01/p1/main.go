package main

import (
	"fmt"
    "strconv"
	"strings"
    "sort"
	"internal/util"
)

func main() {
	content := util.ReadAsArray("input")
    arr0 := []int{}
    arr1 := []int{}
    for _, line := range content {
        fmt.Println(line)
        if len(line) == 0 {
            continue
        }
        vals := strings.Split(line, "   ")
        val0, _ := strconv.Atoi(vals[0])
        val1, _ := strconv.Atoi(vals[1])
        arr0 = append(arr0, val0)
        arr1 = append(arr1, val1)
    }
    sort.Ints(arr0)
    sort.Ints(arr1)
    sumDistance := 0
    for index, val := range arr0 {
        fmt.Println(val, arr1[index])
        distance := arr1[index] - val
        if distance < 0 {
            distance = -distance
        }
        sumDistance += distance
    }
    fmt.Println(sumDistance)
}

