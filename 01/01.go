package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	singleIncreases := 0
	prevDepth := 0
	index := 0
	var depthGroup [3]int
	depthSum := 0
	prevDepthSum := 0
	multiIncreases := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depthString := scanner.Text()
		if depth, err := strconv.Atoi(depthString); err == nil {
			if prevDepth == 0 {
				prevDepth = depth
			} else if depth > prevDepth {
				singleIncreases++
			}
			prevDepth = depth

			depthGroup[index%3] = depth
			if index > 1 {
				depthSum = depthGroup[0] + depthGroup[1] + depthGroup[2]
			}
			if prevDepthSum == 0 {
				prevDepthSum = depthSum
			} else if depthSum > prevDepthSum {
				fmt.Println(depthSum, ">", prevDepthSum)
				multiIncreases++
			}
			prevDepthSum = depthSum
		}
		index++
	}

	fmt.Println("Single line increases ", singleIncreases)
	fmt.Println("Triple line increases ", multiIncreases)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
