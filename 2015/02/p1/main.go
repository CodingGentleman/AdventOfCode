package main

import (
	"fmt"
	"internal/util"
	"slices"
)

func main() {
	sum := 0
	util.StreamFile("input", func(line string) {
		l, w, h := 0, 0, 0
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		sides := []int{l * w, w * h, h * l}
		sum += slices.Min(sides) + sides[0]*2 + sides[1]*2 + sides[2]*2
	})
	fmt.Println(sum)
}
