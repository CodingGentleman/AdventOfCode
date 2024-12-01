package util

import (
    "os"
	"log"
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
