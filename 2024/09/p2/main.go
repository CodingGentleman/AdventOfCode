package main

import (
	"fmt"
	"internal/util"
	"strings"
)

func main() {
	values := make([]obj, 0)
	util.StreamFile("input", func(line string) {
		for i, s := range strings.Split(line, "") {
			obj := obj{len: util.ToInt(s), val: -1}
			if i%2 == 0 {
				obj.val = i / 2
			}
			values = append(values, obj)
		}
	})
	for j := len(values) - 1; j > 0; j-- {
		if values[j].val != -1 {
			for i := 0; i < j; i++ {
				if values[i].val == -1 && values[i].len >= values[j].len {
					values[i].len -= values[j].len
					values = append(values[:i], append([]obj{values[j]}, values[i:]...)...)
					values[j+1].val = -1
					break
				}
			}
		}
	}
	num := 0
	sum := 0
	for i := 0; i < len(values); i++ {
		for j := 0; j < values[i].len; j++ {
			if values[i].val != -1 {
				sum += values[i].val * num
			}
			num++
		}
	}
	fmt.Println(sum)
}

type obj struct {
	len int
	val int
}
