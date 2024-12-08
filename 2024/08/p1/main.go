package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	scanner := bufio.NewScanner(f)
	originMap := make(map[string][]Point)
	max := Point{0, 0}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			if string(c) != "." {
				originMap[string(c)] = append(originMap[string(c)], Point{x, y})
			}
			max.x = x
		}
		max.y = y
	}
	tracking := make(map[string]struct{})
	for key := range originMap {
		points := originMap[key]
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1 := Point{2*points[i].x - points[j].x, 2*points[i].y - points[j].y}
				if p1.InBoundsOf(max) {
					tracking[p1.ToString()] = struct{}{}
				}
				p2 := Point{2*points[j].x - points[i].x, 2*points[j].y - points[i].y}
				if p2.InBoundsOf(max) {
					tracking[p2.ToString()] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(tracking))
}

type Point struct {
	x, y int
}

func (p *Point) ToString() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
func (p *Point) InBoundsOf(max Point) bool {
	return p.x >= 0 && p.x <= max.x && p.y >= 0 && p.y <= max.y
}
