package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	day := "10"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/data/dev/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var syntaxCost = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	syntaxErrorScore := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
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
					syntaxErrorScore += syntaxCost[char]
					break
				} else {
					delimStack[stackLastElement] = ""
					delimStack = delimStack[:stackLastElement]
				}
			}
		}
	}

	fmt.Println("Syntax Error Score:", syntaxErrorScore)

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
