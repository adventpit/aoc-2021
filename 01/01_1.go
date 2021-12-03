package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day := "01"
	challengePart := "1"
	file, err := os.Open(fmt.Sprintf("/home/pitm-geshdo/dev/training/aoc/aoc-2021/%s/input.txt", day))
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

	fmt.Println("Advent of Code challenge Day " + day + " Part " + challengePart)
	fmt.Println("Single line increases ", singleIncreases)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
