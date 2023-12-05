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
    // first line contains seeds
    scanner.Scan()
    seedsString := strings.Split(strings.TrimPrefix(scanner.Text(), "seeds: "), " ")
    seeds := make([]int, len(seedsString))
    for i, seed := range seedsString {
        seeds[i], _ = strconv.Atoi(seed)
    } 
    almanac := Almanac{seeds: seeds}

    // every upcoming line is either empty, starts a converter section or contains a convertion
    for scanner.Scan() {
        // empty line                
        if scanner.Text() == "" {
            continue
        }
        // converter start with a letter
        if scanner.Text()[0] >= 'a' && scanner.Text()[0] <= 'z' {
            almanac.converters = append(almanac.converters, Converter{name: scanner.Text()})
            continue
        }
        // everything else is a convertion that gets added to the last converter
        convertionString := strings.Split(scanner.Text(), " ")
        src, _ := strconv.Atoi(convertionString[1])
        dst, _ := strconv.Atoi(convertionString[0])
        rng, _ := strconv.Atoi(convertionString[2])
        almanac.converters[len(almanac.converters) - 1].convertions = append(almanac.converters[len(almanac.converters) - 1].convertions, Convertion{src: src, dst: dst, rng: rng})
    }
    // part1(almanac)
    part2(almanac)
}

func part2(almanac Almanac) {
    smallest := 0
    results := make(chan int)
    // iterate through all seeds in pairs of two
    for i := 0; i < len(almanac.seeds) - 1; i+=2 {
        src := almanac.seeds[i];
        rng := almanac.seeds[i+1];
        for seed := src; seed < src + rng; seed++ {
            // fmt.Println("seed: ", seed)
            go func() { results <- convert(seed, &almanac) }()
            res := <- results
            if res < smallest || smallest == 0 {
                smallest = res
            }
        }
    }
    fmt.Println(smallest)
}

func part1(almanac Almanac) {
    // calculate smallest end convertion
    smallest := 0
    for _, seed := range almanac.seeds {
        fmt.Println("seed: ", seed)
        res := almanac.convert(seed)
        if res < smallest || smallest == 0 {
            smallest = res
        }
    }
    fmt.Println(smallest)
}

type Almanac struct {
    seeds []int
    converters []Converter
}

func convert(in int, a *Almanac) int {
    tmp := in
    for _, converter := range a.converters {
        tmp = converter.convert(tmp)
    }
    return tmp
}

func (a *Almanac) convert(in int) int {
    tmp := in
    for _, converter := range a.converters {
        tmp = converter.convert(tmp)
    }
    return tmp
}
type Converter struct {
    name string
    convertions []Convertion
}

func (c *Converter) convert(in int) int {
    for _, convertion := range c.convertions {
        res := convertion.calc(in)
        if res != in {
            // fmt.Println(c.name, "[", convertion.src,"]:", in, "->", res)
            return res
        }
    }
    return in
}

type Convertion struct {
    src int
    dst int
    rng int
}

func (c *Convertion) calc(in int) int {
    if in >= c.src && in < c.src+ c.rng {
        return in - c.src + c.dst
    }
    return in
}
