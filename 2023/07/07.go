package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "slices"
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
    for scanner.Scan() {
        part1(scanner.Text())
    }
    
    factor := len(rankedHands)
    result := 0
    for _, v := range rankedHands {
        fmt.Println(string(v.cards), ": ", v.bid, "x", factor)
        result += v.bid * factor
        factor--
    }
    fmt.Println("Result:", result)
}

type Hand struct {
    cards []rune
    bid int
    rank int
}

var rankedCards = map[rune]int{
    'A': 14,
    'K': 13,
    'Q': 12,
    'J': 11,
    'T': 10,
    '9': 9,
    '8': 8,
    '7': 7,
    '6': 6,
    '5': 5,
    '4': 4,
    '3': 3,
    '2': 2,
}
var rankedHands = []Hand{}

func part1(line string) {
    parts := strings.Split(line, " ")
    hand := Hand{cards: []rune(parts[0])}
    hand.bid, _ = strconv.Atoi(parts[1])
    hand.rank = calcHandRank(hand.cards)
    insertHand(hand)

    for _, v := range rankedHands {
        fmt.Println(string(v.cards) )
    }
}

func insertHand(hand Hand) {
    for i := 0; i < len(rankedHands); i++ {
        fmt.Println(string(hand.cards), "-", string(rankedHands[i].cards))
        if hand.rank == rankedHands[i].rank {
            // if the rank is the same check the cards. higher cards win
            for j := 0; j < len(hand.cards); j++ {
                if hand.cards[j] == rankedHands[i].cards[j] {
                    continue
                }
                if rankedCards[hand.cards[j]] > rankedCards[rankedHands[i].cards[j]] {
                    rankedHands = slices.Insert(rankedHands, i, hand)
                    return
                }
                break
            }
        }
        // insert hand right before the first hand with a lower rank
        if hand.rank > rankedHands[i].rank {
            rankedHands = slices.Insert(rankedHands, i, hand)
            return
        }
    }
    // if we get here the hand is the lowest rank
    rankedHands = append(rankedHands, hand)
}

func calcHandRank(hand []rune) int {
    stats := make(map[rune]int)
    for i := 0; i < len(hand); i++ {
        stats[hand[i]]++
    }
    maxValue := 0
    for _, v := range stats {
        if v > maxValue {
            maxValue = v
        }
    }

    // 5 of a kind
    if len(stats) == 1 {
        return 6
    }
    // 4 of a kind
    if len(stats) == 2 && maxValue == 4 {
        return 5
    }
    // Full house
    if len(stats) == 2 && maxValue == 3 {
        return 4
    }
    // Three of a kind
    if len(stats) == 3 && maxValue == 3 {
        return 3
    }
    // Two pair
    if len(stats) == 3 && maxValue == 2 {
        return 2
    }
    // One pair
    if len(stats) == 4 {
        return 1
    }
    // High card
    return 0
}
