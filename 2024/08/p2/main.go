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
	tracking := make(map[string]struct{})
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			if string(c) != "." {
				p := Point{x, y}
				originMap[string(c)] = append(originMap[string(c)], p)
				tracking[p.ToString()] = struct{}{}
			}
			max.x = x
		}
		max.y = y
	}
	for key := range originMap {
		points := originMap[key]
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1 := points[i]
				p2 := points[j]
				pNext := Point{2*p1.x - p2.x, 2*p1.y - p2.y}
				for pNext.InBoundsOf(max) {
					tracking[pNext.ToString()] = struct{}{}
					p2 = p1
					p1 = pNext
					pNext = Point{2*p1.x - p2.x, 2*p1.y - p2.y}
				}
				p1 = points[i]
				p2 = points[j]
				pNext = Point{2*p2.x - p1.x, 2*p2.y - p1.y}
				for pNext.InBoundsOf(max) {
					tracking[pNext.ToString()] = struct{}{}
					p1 = p2
					p2 = pNext
					pNext = Point{2*p2.x - p1.x, 2*p2.y - p1.y}
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
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
func (p *Point) InBoundsOf(max Point) bool {
	return p.x >= 0 && p.x <= max.x && p.y >= 0 && p.y <= max.y
}
