package main

import (
	"fmt"
	"internal/util"
	"regexp"
)

func main() {
	sum := 0
	enabled := true
	util.StreamFile("input", func(line string) {
		pattern := regexp.MustCompile(`(?:(?:do(?:n't)?\(\))|(?:mul\((\d+),(\d+)\)))`)
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == `don't()` {
				enabled = false
			} else if enabled {
				sum += util.ToInt(match[1]) * util.ToInt(match[2])
			}
		}
	})
	fmt.Println(sum)
}
