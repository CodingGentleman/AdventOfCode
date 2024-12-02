package util

import (
	"bufio"
	"log"
	"os"
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

func Abs(value int) int {
    if value < 0 {
        return -value
    }
    return value
}
