package main

import (
	"fmt"
	"internal/util"
	"sort"
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
	sort.Ints(arr0)
	sort.Ints(arr1)
	sumDistance := 0
	for index, val := range arr0 {
		sumDistance += util.Abs(arr1[index] - val)
	}
	fmt.Println(sumDistance)
}
