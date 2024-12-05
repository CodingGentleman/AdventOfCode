package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input")
	santa, robo := []int{0, 0}, []int{0, 0}
	positionLog := make(map[string]struct{})
	positionLog["0:0"] = struct{}{}
	for i, s := range strings.Split(string(data), "") {
		current := santa
		if i%2 == 0 {
			current = robo
		}
		switch s {
		case ">":
			current[0] += 1
		case "v":
			current[1] += 1
		case "<":
			current[0] -= 1
		case "^":
			current[1] -= 1
		}
		positionLog[strconv.Itoa(current[0])+":"+strconv.Itoa(current[1])] = struct{}{}
	}
	fmt.Println(len(positionLog))
}
