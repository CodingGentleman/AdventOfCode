package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input")
	x, y := 0, 0
	positionLog := make(map[string]struct{})
	positionLog["0:0"] = struct{}{}
	for _, s := range strings.Split(string(data), "") {
		switch s {
		case ">":
			x += 1
		case "v":
			y += 1
		case "<":
			x -= 1
		case "^":
			y -= 1
		}
		positionLog[strconv.Itoa(x)+":"+strconv.Itoa(y)] = struct{}{}
	}
	fmt.Println(len(positionLog))
}
