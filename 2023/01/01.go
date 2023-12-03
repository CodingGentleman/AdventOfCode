package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
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
        // sum1 += day01_part1(scanner.Text())
        sum2 += day01_part2(scanner.Text())
    }
    fmt.Println("day1 part1: ", sum1)
    fmt.Println("day1 part2: ", sum2)

}

func day01_part1(line string) int {
    // remove all non-digits
    line = regexp.MustCompile(`\D`).ReplaceAllString(line, "")
    // get the first digit
    num := line[0:1] + line[len(line)-1:]
    // parse to int
    var i int
    fmt.Sscanf(num, "%d", &i)
    return i
}

var digitStrMap = map[string]string{
    "1": "one",
    "2": "two",
    "3": "three",
    "4": "four",
    "5": "five",
    "6": "six",
    "7": "seven",
    "8": "eight",
    "9": "nine",
}

var strDigitMap = map[string]string{
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

func day01_part2(line string) int {
    // replace all digits with words
    fmt.Println(line)
    for k, v := range digitStrMap {
        line = regexp.MustCompile(k).ReplaceAllString(line, v)
    }
    fmt.Println(line)
    // find the first occurrence of any value of the map
    var currentFirstIdx int = len(line)
    var firstDigit string
    var currentLastIdx int = -1
    var lastDigit string
    for _, v := range digitStrMap {
        firstIdx := strings.Index(line, v)
        lastIdx := strings.LastIndex(line, v)
        if firstIdx != -1 && firstIdx < currentFirstIdx {
            currentFirstIdx = firstIdx
            firstDigit = v
        }
        if lastIdx != -1 && lastIdx > currentLastIdx {
            currentLastIdx = lastIdx
            lastDigit = v
        }
    }
    resultStr := strDigitMap[firstDigit] + strDigitMap[lastDigit]
    result, _ := strconv.Atoi(resultStr)
    fmt.Println(result)
    return result
}
