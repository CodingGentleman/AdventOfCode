package main

import (
	"fmt"
    "strconv"
	"strings"
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
    sumSimilarity := 0
    for _, val0 := range arr0 {
        valSimilarity := 0
        for _, val1 := range arr1 {
            if val0 == val1 {
                valSimilarity += 1
            }
        }
        sumSimilarity += valSimilarity * val0
    }
    fmt.Println(sumSimilarity)
}

