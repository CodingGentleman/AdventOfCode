package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    fmt.Println("Day 8 - Part 1")
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
       // AAA = (BBB, BBB)
       parts := strings.Split( scanner.Text(), " = " )
       nodes[parts[0]] = strings.Split(parts[1][1:9], ", ")
    }
    navigation := []rune("LRRLRLRRRLLRLRRRLRLLRLRLRRLRLRRLRRLRLRLLRRRLRRLLRRRLRRLRRRLRRLRLRLLRRLRLRRLLRRRLLLRRRLLLRRLRLRRLRLLRRRLRRLRRRLRRLLRRRLRRRLRRRLRLRRLRLRRRLRRRLRRLRLRRLLRRRLRRLLRRLRRLRLRLRRRLRLLRRRLRRLRRRLLRRLLLLLRRRLRRLLLRRRLRRRLRRLRLLLLLRLRRRLRRRLRLRRLLLLRLRRRLLRRRLRRRLRLRLRRLRRLRRLRLRLLLRLRRLRRLRRRLRRRLLRRRR")

    counter := 0
    navIndex := 0
    currentNode := "AAA"
    for currentNode != "ZZZ" {
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
    fmt.Println("Result:", counter)
}

var nodes = make(map[string][]string)
