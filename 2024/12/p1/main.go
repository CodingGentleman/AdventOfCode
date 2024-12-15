package main

import (
	"fmt"
	"internal/util"
)

type Region struct {
	id     string
	fields []Field
}

type Field struct {
	char    string
	x, y    int
	visited bool
	region  *Region
}

var grid [][]Field
var regions []Region

func main() {
	grid = util.ReadAsGrid("input", func(y, x int, r rune) Field { return Field{char: string(r), x: x, y: y, visited: false} })
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			FillRegion(grid[y][x].char, y, x)
		}
	}

	sum := 0
	for _, region := range regions {
		fieldCount := len(region.fields)
		fenceCount := 0
		for _, field := range region.fields {
			if field.y == 0 || grid[field.y-1][field.x].region.id != region.id {
				fenceCount++
			}
			if field.y == len(grid)-1 || grid[field.y+1][field.x].region.id != region.id {
				fenceCount++
			}
			if field.x == 0 || grid[field.y][field.x-1].region.id != region.id {
				fenceCount++
			}
			if field.x == len(grid[field.y])-1 || grid[field.y][field.x+1].region.id != region.id {
				fenceCount++
			}
		}
		sum += fieldCount * fenceCount
	}
	fmt.Println(sum)
}

func FillRegion(char string, y, x int) {
	if y < 0 || x < 0 || y >= len(grid) || x >= len(grid[y]) || grid[y][x].visited || grid[y][x].char != char {
		return
	}
	var id string
	var region *Region

	if y > 0 && grid[y-1][x].char == grid[y][x].char && grid[y-1][x].visited {
		region = grid[y-1][x].region
	} else if x > 0 && grid[y][x-1].char == grid[y][x].char && grid[y][x-1].visited {
		region = grid[y][x-1].region
	} else if x < len(grid[y])-1 && grid[y][x+1].char == grid[y][x].char && grid[y][x+1].visited {
		region = grid[y][x+1].region
	} else if y < len(grid)-1 && grid[y+1][x].char == grid[y][x].char && grid[y+1][x].visited {
		region = grid[y+1][x].region
	} else {
		id = fmt.Sprintf("%s#%d-%d", grid[y][x].char, x, y)
		regions = append(regions, Region{id: id, fields: make([]Field, 0)})
		region = &regions[len(regions)-1]
	}
	grid[y][x] = Field{char: grid[y][x].char, x: x, y: y, visited: true, region: region}
	region.fields = append(region.fields, grid[y][x])

	FillRegion(char, y-1, x)
	FillRegion(char, y+1, x)
	FillRegion(char, y, x-1)
	FillRegion(char, y, x+1)
}
