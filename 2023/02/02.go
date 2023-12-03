package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
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
    for scanner.Scan() {
        // sum1 += day02_part1(scanner.Text())
        sum2 += day02_part2(scanner.Text())
    }
    fmt.Println("day2 part1: ", sum1)
    fmt.Println("day2 part2: ", sum2)
}

var cubeMap = map[string]int {
    "red": 12,
    "green": 13,
    "blue": 14,
}

func day02_part1(line string) int {
    // first substring "Game " from line
    line = strings.TrimPrefix(line, "Game ")
    // second split line by ": "
    split := strings.Split(line, ": ")
    gameNumber, _ := strconv.Atoi(split[0])
    line = split[1]
    // third split line by "; "
    rounds := strings.Split(line, "; ")
    for _, round := range rounds {
        // fourth split round by ", "
        sets := strings.Split(round, ", ")
        for _, set := range sets {
            // fifth split set by " "
            values := strings.Split(set, " ")
            count, _ := strconv.Atoi(values[0])
            color := values[1]
            // sixth check if count exceeds value in cubeMap
            if count > cubeMap[color] {
                fmt.Println("Game not valid: ", line)
                return 0
            }
        }
    }
    return gameNumber
}

func day02_part2(line string) int {
    // remember the max value for each color
    maxValueMap := map[string]int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    // first substring "Game " from line
    line = strings.TrimPrefix(line, "Game ")
    // second split line by ": "
    split := strings.Split(line, ": ")
    line = split[1]
    // third split line by "; "
    rounds := strings.Split(line, "; ")
    for _, round := range rounds {
        // fourth split round by ", "
        sets := strings.Split(round, ", ")
        for _, set := range sets {
            // fifth split set by " "
            values := strings.Split(set, " ")
            count, _ := strconv.Atoi(values[0])
            color := values[1]
            // sixth if count exceeds the current value in maxValueMap, update it
            if count > maxValueMap[color] {
                maxValueMap[color] = count
            }
        }
    }
    // lastly multiply the values in maxValueMap
    return maxValueMap["red"] * maxValueMap["green"] * maxValueMap["blue"]
}
