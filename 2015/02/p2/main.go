package main

import (
	"fmt"
	"internal/util"
	"slices"
)

func main() {
	sum := 0
	util.StreamFile("input", func(line string) {
		d := make([]int, 3)
		fmt.Sscanf(line, "%dx%dx%d", &d[0], &d[1], &d[2])
		sum += d[0]*d[1]*d[2] + d[0]*2 + d[1]*2 + d[2]*2 - slices.Max(d)*2
	})
	fmt.Println(sum)
}
