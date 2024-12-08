package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	scanner := bufio.NewScanner(f)
	obstacles := make(map[string]struct{})
	movementLog := make(map[string]struct{})
	maxX, maxY, guardX, guardY := 0, 0, 0, 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			if string(c) == "#" {
				obstacles[coords(j, i)] = struct{}{}
			}
			if string(c) == "^" {
				movementLog[coords(j, i)] = struct{}{}
				guardX, guardY = j, i
			}
		}
		maxX = len(line)
		maxY = i
	}
	direction := NewLoopingIndex(3)
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
		movementLog[coords(guardX, guardY)] = struct{}{}
	}
	fmt.Println(len(movementLog) - 1)
}

func coords(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

type LoopingIndex struct {
	index int
	max   int
}

func NewLoopingIndex(max int) *LoopingIndex {
	return &LoopingIndex{0, max}
}

func (l *LoopingIndex) Current() int {
	return l.index
}

func (l *LoopingIndex) Next() int {
	l.index++
	if l.index > l.max {
		l.index = 0
	}
	return l.index
}
