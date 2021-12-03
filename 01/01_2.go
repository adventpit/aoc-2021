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

	fmt.Println("Advent of Code challenge Day " + day + " Part " + challengePart)
	fmt.Println("Triple line increases ", multiIncreases)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
