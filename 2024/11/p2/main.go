package main

import (
	"fmt"
	"internal/util"
	"strconv"
	"strings"
)

func main() {
	util.StreamFile("input", func(line string) {
		dict := make(map[string]int)
		for _, v := range strings.Split(line, " ") {
			dict[v]++
		}
		for i := 0; i < 75; i++ {
			copyDict := make(map[string]int)
			for k, v := range dict {
				copyDict[k] = v
			}
			dict = make(map[string]int)
			for k, v := range copyDict {
				if k == "0" {
					dict["1"] += v
				} else if len(k)%2 == 0 {
					dict[strconv.Itoa(util.ToInt(k[0:len(k)/2]))] += v
					dict[strconv.Itoa(util.ToInt(k[len(k)/2:]))] += v
				} else {
					dict[strconv.Itoa(util.ToInt(k)*2024)] += v
				}
			}
		}
		sum := 0
		for k := range dict {
			sum += dict[k]
		}
		fmt.Println(sum)
	})
}
