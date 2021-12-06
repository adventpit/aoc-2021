package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	day := "01"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	singleIncreases := 0
	prevDepth := 0
	index := 0

	for scanner.Scan() {
		depthString := scanner.Text()
		if depth, err := strconv.Atoi(depthString); err == nil {
			if prevDepth == 0 {
				prevDepth = depth
			} else if depth > prevDepth {
				singleIncreases++
			}
			prevDepth = depth
		}
		index++
	}

	fmt.Println("Single line increases ", singleIncreases)

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
