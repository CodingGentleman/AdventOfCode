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
	moved := true
	for moved {
		moved = rearrange(values)
	}
	fmt.Println(values)
}

func rearrange(values []obj) bool {
	for j := len(values) - 1; j > 0; j-- {
		if values[j].val != -1 {
			for i := 0; i < j; i++ {
				if values[i].val == -1 && values[i].len >= values[j].len {
					values[i].len -= values[j].len
					// move element at j to i-1
					values = append(values[:i], append([]obj{values[j]}, values[i:]...)...)
					return true
				}
			}
		}
	}
	return false
}

type obj struct {
	len int
	val int
}
