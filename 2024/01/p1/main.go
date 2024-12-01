package main

import (
	"fmt"
	"internal/util"
	"sort"
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
	sort.Ints(arr0)
	sort.Ints(arr1)
	sumDistance := 0
	for index, val := range arr0 {
		distance := arr1[index] - val
		if distance < 0 {
			distance = -distance
		}
		sumDistance += distance
	}
	fmt.Println(sumDistance)
}
