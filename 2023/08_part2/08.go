package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    fmt.Println("Day 8 - Part 2")
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    startNodes := make([]string, 0)
    for scanner.Scan() {
       // AAA = (BBB, BBB)
       parts := strings.Split( scanner.Text(), " = " )
       nodes[parts[0]] = strings.Split(parts[1][1:9], ", ")
       if strings.Split(parts[0], "")[2] == "A" {
           startNodes = append(startNodes, parts[0])
       }
    }
    fmt.Println("Current nodes:", startNodes)
    navigation := []rune("LRRLRLRRRLLRLRRRLRLLRLRLRRLRLRRLRRLRLRLLRRRLRRLLRRRLRRLRRRLRRLRLRLLRRLRLRRLLRRRLLLRRRLLLRRLRLRRLRLLRRRLRRLRRRLRRLLRRRLRRRLRRRLRLRRLRLRRRLRRRLRRLRLRRLLRRRLRRLLRRLRRLRLRLRRRLRLLRRRLRRLRRRLLRRLLLLLRRRLRRLLLRRRLRRRLRRLRLLLLLRLRRRLRRRLRLRRLLLLRLRRRLLRRRLRRRLRLRLRRLRRLRRLRLRLLLRLRRLRRLRRRLRRRLLRRRR")

    navIndex := 0

    allCounter := make([]int, 0)
    for _, currentNode := range startNodes { 
        counter := 0
        for strings.Split(currentNode, "")[2] != "Z" {
            counter++
            step := navigation[navIndex]
            if step == 'L' {
                currentNode = nodes[currentNode][0]
            } else {
                currentNode = nodes[currentNode][1]
            }
            navIndex++
            if navIndex == len(navigation) {
                navIndex = 0
            }
        }
        allCounter = append(allCounter, counter)
    }

    result := LCM(allCounter[0], allCounter[1], allCounter[2:]...)
    fmt.Println("Result:", result)
}

var nodes = make(map[string][]string)

func GCD(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func LCM(a, b int, integers ...int) int {
    result := a * b / GCD(a, b)
    for i := 0; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }
    return result
}
