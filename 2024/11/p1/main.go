package main

import (
	"fmt"
	"internal/util"
	"strconv"
	"strings"
)

func main() {
	util.StreamFile("input", func(line string) {
		a := strings.Split(line, " ")
		for i := 0; i < 25; i++ {
			b := make([]string, 0)
			for _, v := range a {
				if v == "0" {
					b = append(b, "1")
				} else if len(v)%2 == 0 {
					b = append(b, strconv.Itoa(util.ToInt(v[0:len(v)/2])))
					b = append(b, strconv.Itoa(util.ToInt(v[len(v)/2:])))
				} else {
					b = append(b, strconv.Itoa(util.ToInt(v)*2024))
				}
			}
			a = make([]string, len(b))
			copy(a, b)
		}
		fmt.Println(len(a))
	})
}
