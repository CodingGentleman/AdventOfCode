package main

import (
	"fmt"
	"internal/util"
	"regexp"
)

func main() {
	sum := 0
	util.StreamFile("input", func(line string) {
		pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sum += util.ToInt(match[1]) * util.ToInt(match[2])
		}
	})
	fmt.Println(sum)
}
