package main

import (
	"fmt"
	"internal/util"
	"strings"
)

func main() {
	values := make([]int, 0)
	index := 0
	util.StreamFile("input", func(line string) {
		for i, s := range strings.Split(line, "") {
			val := -1
			if i%2 == 0 {
				val = i / 2
			}
			for limit := index + util.ToInt(s); index < limit; index++ {
				values = append(values, val)
			}
		}
	})
	index--
	sum := 0
	for i := 0; i < index; i++ {
		if values[i] == -1 {
			for ; i < index; index-- {
				if values[index] != -1 {
					break
				}
			}
			values[i] = values[index]
			values[index] = -1
		}
		sum += i * values[i]
	}
	fmt.Println(sum)
}
