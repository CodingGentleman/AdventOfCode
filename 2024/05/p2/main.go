package main

import (
	"fmt"
	"internal/util"
	"slices"
	"strings"
)

func main() {
	rules := make(map[string]map[string]struct{})
	sortFunc := func(a string, b string) int {
		if _, ok := rules[a][b]; ok {
			return -1
		}
		return 0
	}

	sum := 0
	util.StreamFile("input", func(line string) {
		if strs := strings.Split(line, "|"); len(strs) > 1 {
			if rules[strs[0]] == nil {
				rules[strs[0]] = make(map[string]struct{})
			}
			rules[strs[0]][strs[1]] = struct{}{}
		}
		if strs := strings.Split(line, ","); len(strs) > 1 && !slices.IsSortedFunc(strs, sortFunc) {
			slices.SortFunc(strs, sortFunc)
			sum += util.ToInt(strs[len(strs)/2])
		}
	})
	fmt.Println(sum)
}
