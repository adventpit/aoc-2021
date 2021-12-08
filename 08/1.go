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
	day := "08"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	noSegments := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	appearanceEasyNumbers := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
		inputSplit := strings.Split(inputLine, " | ")
		// uniqueSignalPattern := inputSplit[0]
		for _, fourDigitOutput := range strings.Split(inputSplit[1], " ") {
			if len(fourDigitOutput) == noSegments[1] || len(fourDigitOutput) == noSegments[4] ||
				len(fourDigitOutput) == noSegments[7] || len(fourDigitOutput) == noSegments[8] {
				appearanceEasyNumbers++
			}
		}
	}

	fmt.Println(appearanceEasyNumbers)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
