package main

import (
	"fmt"
	"internal/util"
)

func main() {
    arr0 := []int{}
    dict1 := make(map[int]int)
    util.StreamFile("input", func(line string) {
		var val0, val1 int
		fmt.Sscanf(line, "%d   %d", &val0, &val1)
		arr0 = append(arr0, val0)
        dict1[val1] = dict1[val1] + 1
	})
	sumSimilarity := 0
	for _, val0 := range arr0 {
        sumSimilarity += dict1[val0] * val0
	}
	fmt.Println(sumSimilarity)
}
