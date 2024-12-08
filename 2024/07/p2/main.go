package main

import (
	"fmt"
	"internal/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	util.StreamFile("input", func(line string) {
		parts := strings.Split(line, ": ")
		result := util.ToInt(parts[0])
		values := strings.Split(parts[1], " ")
		for i := 0; i < int(math.Pow(float64(3), float64(len(values)-1))); i++ {
			if result == calc(values, i) {
				sum += result
				break
			}
		}
	})
	fmt.Println(sum)
}

func calc(values []string, operator int) int {
	result := util.ToInt(values[0])
	opcode := fmt.Sprintf("%0"+strconv.Itoa(len(values)-1)+"s", strconv.FormatInt(int64(operator), 3))
	ops := strings.Split(opcode, "")
	for i := 1; i < len(values); i++ {
		v := util.ToInt(values[i])
		if ops[i-1] == "2" {
			result = util.ToInt(fmt.Sprintf("%d%d", result, v))
		} else if ops[i-1] == "1" {
			result *= v
		} else {
			result += v
		}
	}
	return result
}
