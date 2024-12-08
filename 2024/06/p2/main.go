package main

import (
	"bufio"
	"fmt"
	"internal/util"
	"os"
)

func main() {
	f, _ := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	scanner := bufio.NewScanner(f)
	obstacles := make(map[string]struct{})
	movementLog := make(map[string]int)
	direction := util.NewLoopingIndex(3)
	maxX, maxY, guardX, guardY := 0, 0, 0, 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			if string(c) == "#" {
				obstacles[coords(j, i)] = struct{}{}
			}
			if string(c) == "^" {
				movementLog[coords(j, i)] = direction.Current()
				guardX, guardY = j, i
			}
		}
		maxX = len(line)
		maxY = i
	}
	// run once to get the path
	runGuard(movementLog, obstacles, direction, maxX, maxY, guardX, guardY)
	loopsFound := 0
	for k := range movementLog {
		direction.Reset()
		obstacles[k] = struct{}{}
		tmpMovementLog := make(map[string]int)
		tmpMovementLog[coords(guardX, guardY)] = direction.Current()
		loopsFound += runGuard(tmpMovementLog, obstacles, direction, maxX, maxY, guardX, guardY)
		delete(obstacles, k)
	}
	fmt.Println(loopsFound)
	fmt.Println(len(movementLog) - 1)
}

func runGuard(movementLog map[string]int, obstacles map[string]struct{}, direction *util.LoopingIndex, maxX, maxY, guardX, guardY int) int {
	for guardX > -1 && guardX <= maxX && guardY > -1 && guardY <= maxY {
		nextX, nextY := guardX, guardY
		if direction.Current() == 3 { //west
			nextX--
		} else if direction.Current() == 2 { //south
			nextY++
		} else if direction.Current() == 1 { //east
			nextX++
		} else { // north
			nextY--
		}
		if _, ok := obstacles[coords(nextX, nextY)]; ok {
			direction.Next()
			continue
		}
		guardX, guardY = nextX, nextY
		if d, ok := movementLog[coords(guardX, guardY)]; ok && d == direction.Current() {
			return 1
		}
		movementLog[coords(guardX, guardY)] = direction.Current()
	}
	return 0
}

func coords(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
