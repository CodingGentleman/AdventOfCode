package main

import (
	"fmt"
	"internal/util"
	"strconv"
	"strings"
)

func main() {
	sumSafe := 0
	util.StreamFile("input", func(line string) {
		sumSafe += IsSafe(line)
	})
	fmt.Println(sumSafe)
}

func IsSafe(line string) int {
	arr := toIntArray(line)
	asc := arr[0] - arr[1]
	for i := 1; i < len(arr)-1; i++ {
		if util.Abs(asc) < 4 && util.Abs(arr[i]-arr[i+1]) < 4 {
			if asc < 0 && arr[i] < arr[i+1] {
				continue
			}
			if asc > 0 && arr[i] > arr[i+1] {
				continue
			}
		}
		return 0
	}
	return 1
}

func toIntArray(s string) []int {
	strs := strings.Fields(s)
	arr := make([]int, len(strs))
	for index, strval := range strs {
		ival, _ := strconv.Atoi(strval)
		arr[index] = ival
	}
	return arr
}
