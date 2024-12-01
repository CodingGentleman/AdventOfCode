package main

import (
	"fmt"
	"internal/util"
)

func main() {
	content := util.ReadAsArray("input")
	arr0 := []int{}
	arr1 := []int{}
	for _, line := range content {
		if len(line) == 0 {
			continue
		}
		var val0, val1 int
		fmt.Sscanf(line, "%d   %d", &val0, &val1)
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
