package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	day := "10"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/data/dev/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var syntaxScore = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var missingSyntaxScores []int

	for scanner.Scan() {
		inputLine := scanner.Text()
		isCorrupt := false
		var delimStack []string
		for _, char := range strings.Split(inputLine, "") {
			if char == "(" {
				delimStack = append(delimStack, char)
			} else if char == "[" {
				delimStack = append(delimStack, char)
			} else if char == "{" {
				delimStack = append(delimStack, char)
			} else if char == "<" {
				delimStack = append(delimStack, char)
			} else {
				stackLastElement := len(delimStack) - 1
				topDelim := delimStack[stackLastElement]
				expectedDelimiter := matchingDelimiter(topDelim)
				if expectedDelimiter != char {
					isCorrupt = true
					break
				} else {
					delimStack[stackLastElement] = ""
					delimStack = delimStack[:stackLastElement]
				}
			}
		}
		if !isCorrupt {
			lineScore := 0

			for i := range delimStack {
				missingDelimiter := matchingDelimiter(delimStack[len(delimStack)-i-1])
				lineScore = lineScore*5 + syntaxScore[missingDelimiter]
			}
			missingSyntaxScores = append(missingSyntaxScores, lineScore)
		}
	}

	sort.Ints(missingSyntaxScores[:])
	middleScore := missingSyntaxScores[len(missingSyntaxScores)/2]
	fmt.Println("Middle Score:", middleScore)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func matchingDelimiter(char string) string {
	if char == "(" {
		return ")"
	} else if char == "[" {
		return "]"
	} else if char == "{" {
		return "}"
	} else if char == "<" {
		return ">"
	} else {
		return ""
	}
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
