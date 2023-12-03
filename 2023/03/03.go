package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0
    dotRow := []rune(strings.Repeat(".", 142))

    // read the whole file into 2d rune array
    content := [][]rune{}
    content = append(content, dotRow)
    for scanner.Scan() {
        line := []rune("."+scanner.Text()+".")
        content = append(content, line)
    }
    content = append(content, dotRow)
    
    // iterate through content     
    for i := 1; i < len(content)-1; i++ {
        // sum += day03_part1(content[i-1], content[i], content[i+1])
        sum += day03_part2(content[i-1], content[i], content[i+1])
    }
    fmt.Println("result: ", sum)
}

var digits = "0123456789"

func day03_part1(prev []rune, curr []rune, next []rune) int {
    sum := 0
    // we go through the second line char by char
    for i := 0; i < len(curr); i++ {
        // if the char is a dot or digit, we skip it
        if curr[i] == '.' || strings.Contains(digits, string(curr[i])) {
            continue
        }
        // else, we found a symbol. we need to check around it to find a digit
        // first we check east of it
        if strings.Contains(digits, string(curr[i-1])) {
            sum = sum + getNumberAndOverwrite(curr, i-1)
        }
        // then we check north-east of it
        if strings.Contains(digits, string(prev[i-1])) {
            sum = sum + getNumberAndOverwrite(prev, i-1)
        }
        // then we check north of it
        if strings.Contains(digits, string(prev[i])) {
            sum = sum + getNumberAndOverwrite(prev, i)
        }
        // then we check north-west of it
        if strings.Contains(digits, string(prev[i+1])) {
            sum = sum + getNumberAndOverwrite(prev, i+1)
        }
        // then we check west of it
        if strings.Contains(digits, string(curr[i+1])) {
            sum = sum + getNumberAndOverwrite(curr, i+1)
        }
        // then we check south-west of it
        if strings.Contains(digits, string(next[i+1])) {
            sum = sum + getNumberAndOverwrite(next, i+1)
        }
        // then we check south of it
        if strings.Contains(digits, string(next[i])) {
            sum = sum + getNumberAndOverwrite(next, i)
        }
        // then we check south-east of it
        if strings.Contains(digits, string(next[i-1])) {
            sum = sum + getNumberAndOverwrite(next, i-1)
        }
    }
    return sum
}

type Pair struct {
    first, second int
}
func (p *Pair) add(value int) bool {
    if p.first == 0 {
        p.first = value
        return true
    } else if p.second == 0 {
        p.second = value
        return true
    } 
    return false
}

func day03_part2(prev []rune, curr []rune, next []rune) int {
    sum := 0
    // we go through the second line char by char
    for i := 0; i < len(curr); i++ {
        result := Pair{0, 0}
        // if the char is not a asterisk, we skip it
        if curr[i] != '*' {
            continue
        }
        // else, we found a gear. we need to check around it to find exactly two numbers
        // first we check east of it
        if strings.Contains(digits, string(curr[i-1])) {
            success := result.add(getNumberAndOverwrite(curr, i-1))
            if !success {
                continue
            }
        }
        // then we check north-east of it
        if strings.Contains(digits, string(prev[i-1])) {
            success := result.add(getNumberAndOverwrite(prev, i-1))
            if !success {
                continue
            }
        }
        // then we check north of it
        if strings.Contains(digits, string(prev[i])) {
            success := result.add(getNumberAndOverwrite(prev, i))
            if !success {
                continue
            }
        }
        // then we check north-west of it
        if strings.Contains(digits, string(prev[i+1])) {
            success := result.add(getNumberAndOverwrite(prev, i+1))
            if !success {
                continue
            }
        }
        // then we check west of it
        if strings.Contains(digits, string(curr[i+1])) {
            success := result.add(getNumberAndOverwrite(curr, i+1))
            if !success {
                continue
            }
        }
        // then we check south-west of it
        if strings.Contains(digits, string(next[i+1])) {
            success := result.add(getNumberAndOverwrite(next, i+1))
            if !success {
                continue
            }
        }
        // then we check south of it
        if strings.Contains(digits, string(next[i])) {
            success := result.add(getNumberAndOverwrite(next, i))
            if !success {
                continue
            }
        }
        // then we check south-east of it
        if strings.Contains(digits, string(next[i-1])) {
            success := result.add(getNumberAndOverwrite(next, i-1))
            if !success {
                continue
            }
        }
        sum = sum + result.first * result.second
    }
    return sum
}

func getNumberAndOverwrite(line []rune, index int) int {
    foundNumber := string((line)[index])
    line[index] = '.'

    for j:= index-1; j > 0; j-- {
        if strings.Contains(digits, string((line)[j])) {
            foundNumber = string((line)[j]) + foundNumber
            line[j] = '.'
        } else {
            break
        }
    }

    for j:= index+1; j < len(line); j++ {
        if strings.Contains(digits, string((line)[j])) {
            foundNumber = foundNumber + string((line)[j])
            line[j] = '.'
        } else {
            break
        }
    }
    result, _ := strconv.Atoi(foundNumber)
    return result
}

