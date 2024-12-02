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
		arr := toIntArray(line)
		sumSafe += IsSafe(arr, true)
	})
	fmt.Println(sumSafe)
}

func IsSafe(arr []int, withFaultTolerance bool) int {
	asc := arr[0] - arr[1]
	for i := 0; i < len(arr)-1; i++ {
		if util.Abs(asc) < 4 && util.Abs(arr[i]-arr[i+1]) < 4 {
			if asc < 0 && arr[i] < arr[i+1] {
				continue
			}
			if asc > 0 && arr[i] > arr[i+1] {
				continue
			}
		}
		if withFaultTolerance {
			if (i > 0 && IsSafe(removeIndex(arr, i-1), false) > 0) ||
				IsSafe(removeIndex(arr, i), false) > 0 ||
				IsSafe(removeIndex(arr, i+1), false) > 0 {
				return 1
			}
		}
		return 0
	}
	return 1
}

func removeIndex(s []int, index int) []int {
	result := make([]int, 0)
	result = append(result, s[:index]...)
	return append(result, s[index+1:]...)
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
