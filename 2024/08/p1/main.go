package main

import (
	"fmt"
	"internal/util"
)

func main() {
	originMap := make(map[string][]Point)
	max := Point{0, 0}
	util.StreamFileWithIndex("input", func(y int, line string) {
		for x, c := range line {
			if string(c) != "." {
				originMap[string(c)] = append(originMap[string(c)], Point{x, y})
			}
			max.Set(x, y)
		}
	})
	tracking := make(map[string]struct{})
	for key := range originMap {
		points := originMap[key]
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p := Point{2*points[i].x - points[j].x, 2*points[i].y - points[j].y}
				if p.InBoundsOf(max) {
					tracking[p.ToString()] = struct{}{}
				}
				p.Set(2*points[j].x-points[i].x, 2*points[j].y-points[i].y)
				if p.InBoundsOf(max) {
					tracking[p.ToString()] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(tracking))
}

type Point struct {
	x, y int
}

func (p *Point) Set(x, y int) {
	p.x = x
	p.y = y
}
func (p *Point) ToString() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
func (p *Point) InBoundsOf(max Point) bool {
	return p.x >= 0 && p.x <= max.x && p.y >= 0 && p.y <= max.y
}
