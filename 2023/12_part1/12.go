package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "regexp"
)

func main() {
    fmt.Println("Day 12 - Part 1")

    file, _ := os.Open("input")
    defer file.Close()

    entries := make([]Entry, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")
        record := split[0]
        groups := make([]string, 0)
        for _, group := range strings.Split(split[1], ",") {
            groups = append(groups, group)
        }
        unknowns := make([]int, 0)
        for i, char := range record {
            if char == '?' {
                unknowns = append(unknowns, i)
            }
        }
        entries = append(entries, Entry{record, groups, unknowns})
    }
    fmt.Println(entries)

    sum := 0
    for _, entry := range entries {
        fmt.Println(entry.record, entry.groups, entry.getRegex())
        sum += entry.calcArrangements()
    }
    fmt.Println(sum)
}

type Entry struct {
    record string
    groups []string
    unknowns []int
}

func (e Entry) getRegex() string {
    regex := "^[^#]*#{" // start with anything but #
    regex += strings.Join(e.groups, "}[^#]{1,}#{")
    regex += "}[^#]*$" // end with anything but #
    return regex
}

// ???.### 1,1,3
// .??..??...?##. 1,1,3
// ?#?#?#?#?#?#?#? 1,3,1,6
// ????.#...#... 4,1,1
// ????.######..#####. 1,6,5
// ?###???????? 3,2,1
func (e Entry) calcArrangements() int {
    arrangements := 1 << len(e.unknowns)
    regex, err := regexp.Compile(e.getRegex())
    if err != nil {
        fmt.Println(err)
        return 0
    }
    matches := 0
    for i := 0; i < arrangements; i++ {
        current := i
        result := e.record
        
        for _, index := range e.unknowns {
            if current%2 == 0 {
                result = replaceAtIndex(result, index, '.')
            } else {
                result = replaceAtIndex(result, index, '#')
            }
            current /= 2
        }

        match := regex.MatchString(result)
        if match {
            fmt.Println(result, match)
            matches++
        }
    }

    return matches
}

func replaceAtIndex(str string, index int, replacement rune) string {
    return str[:index] + string(replacement) + str[index+1:]
}
