package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input")
	result := 0
	for i, s := range strings.Split(string(data), "") {
		if s == "(" {
			result += 1
		} else {
			result -= 1
		}
		if result == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
