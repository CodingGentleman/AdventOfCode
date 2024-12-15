package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadAsString(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return string(data)
}

func ReadAsArray(filePath string) []string {
	data := ReadAsString(filePath)
	return strings.Split(data, "\n")
}

func ReadAsGrid[T any](filePath string, convert func(int, int, rune)T) [][]T {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]T, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		grid = append(grid, make([]T, len(line)))
		for j, c := range line {
            grid[i][j] = convert(i,j,c)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}
	return grid
}

func StreamFile(filePath string, callback func(string)) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		callback(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func StreamFileWithIndex(filePath string, callback func(int, string)) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for i := 0; sc.Scan(); i++ {
		callback(i, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func PrintGrid(grid [][]string) {
    for _, row := range grid {
        for _, cell := range row {
            fmt.Print(cell)
        }
        fmt.Println()
    }
}

func Abs(value int) int {
    if value < 0 {
        return -value
    }
    return value
}

func ToInt(value string) int {
    i, _ := strconv.Atoi(value)
    return i
}

func ToIntArray(values []string) []int {
    result := make([]int, len(values))
    for i, v := range values {
        result[i] = ToInt(v)
    }
    return result
}
