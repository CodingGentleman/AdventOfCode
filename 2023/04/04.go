package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "math"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum1 := 0
    sum2 := 0
    index := 0
    for scanner.Scan() {
        // sum1 += day04_part1(scanner.Text())
        day04_part2(index, scanner.Text())
        // print cards
        fmt.Println("cards: ", cards)
        index++
    }
    // sum cards
    for i := 0; i < len(cards); i++ {
        sum2 += cards[i]
    }
    fmt.Println("day4 part1: ", sum1)
    fmt.Println("day4 part2: ", sum2)
}

var cards = make([]int, 200)

func day04_part2(index int, line string) {
    cards[index]++
    fmt.Println("line: ", line)
    card := strings.Split(line, ": ")[1]
    parts := strings.Split(card, " | ")
    winningPart := parts[0]
    playerPart := parts[1]
    winningNumbers := strings.Split(winningPart, " ")
    playerNumbers := strings.Split(playerPart, " ")
    winningIndex := index+1
    for i := 0; i < len(winningNumbers); i++ {
        if winningNumbers[i] == "" {
            continue
        }
        for j := 0; j < len(playerNumbers); j++ {
            if winningNumbers[i] == playerNumbers[j] {
                for k := 0; k < cards[index]; k++ {
                    cards[winningIndex]++
                }
                winningIndex++
            }
        }
    }
}

func day04_part1(line string) int {
    card := strings.Split(line, ": ")[1]
    parts := strings.Split(card, " | ")
    winningPart := parts[0]
    playerPart := parts[1]
    winningNumbers := strings.Split(winningPart, " ")
    playerNumbers := strings.Split(playerPart, " ")
    countOfCorrect := 0
    for i := 0; i < len(winningNumbers); i++ {
        if winningNumbers[i] == "" {
            continue
        }
        for j := 0; j < len(playerNumbers); j++ {
            if winningNumbers[i] == playerNumbers[j] {
                countOfCorrect++
            }
        }
    }

    result := int(math.Pow(2, float64(countOfCorrect-1)))
    return result
}
