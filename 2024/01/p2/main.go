package main

import (
	"fmt"
	"internal/util"
)

func main() {
	arr0 := []int{}
	arr1 := []int{}
	util.StreamFile("input", func(line string) {
		var val0, val1 int
		fmt.Sscanf(line, "%d   %d", &val0, &val1)
		arr0 = append(arr0, val0)
		arr1 = append(arr1, val1)
	})
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
