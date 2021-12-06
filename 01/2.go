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
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index := 0
	var depthGroup [3]int
	depthSum := 0
	prevDepthSum := 0
	multiIncreases := 0

	for scanner.Scan() {
		depthString := scanner.Text()
		if depth, err := strconv.Atoi(depthString); err == nil {
			depthGroup[index%3] = depth
			if index > 1 {
				depthSum = depthGroup[0] + depthGroup[1] + depthGroup[2]
			}
			if prevDepthSum == 0 {
				prevDepthSum = depthSum
			} else if depthSum > prevDepthSum {
				multiIncreases++
			}
			prevDepthSum = depthSum
		}
		index++
	}

	fmt.Println("Triple line increases ", multiIncreases)

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
