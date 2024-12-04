package main

import (
	"fmt"
	"internal/util"
	"strings"
)

func main() {
	grid := util.ReadAsGrid("input")
	sum := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			val := grid[i+1][j+1] + grid[i-1][j-1] + grid[i+1][j-1] + grid[i-1][j+1]
			if grid[i][j] == "A" && strings.Contains("MSMS;MSSM;SMMS;SMSM", val) {
				sum += 1
			}
		}
	}
	fmt.Println("sum: ", sum)
}
