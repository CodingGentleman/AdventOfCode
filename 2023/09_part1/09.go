package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    fmt.Println("Day 7 - Part 1")

    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        sum += part1(scanner.Text())
    }
    fmt.Println("Result:", sum)
} 

func part1(input string) int {
    parts := strings.Split(input, " ")
    numbers := make([]int, len(parts))
    for i, v := range parts {
        numbers[i], _ = strconv.Atoi(v)
    }
    return rec(numbers)
}

func allSame(numbers []int) bool {
    for _, v := range numbers {
        if v != numbers[0] {
            return false
        }
    }
    return true
}

func reduceArray(numbers []int) []int {
    newNumbers := make([]int, len(numbers)-1)
    for i := 0; i < len(numbers)-1; i++ {
        newNumbers[i] = numbers[i+1] - numbers[i]
    }
    return newNumbers
}

func rec(numbers []int) int {
    if allSame(numbers) {
        return numbers[0]
    }
    reducedNumbers := reduceArray(numbers)
    // go down recursion
    appendedNumber := rec(reducedNumbers)
    // get the last number from input
    return numbers[len(numbers)-1] + appendedNumber
}
