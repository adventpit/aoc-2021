package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	day := "07"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var crabs []int
	maxCrabPos := 0
	minCrabPos := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
		for _, crabPositionStr := range strings.Split(inputLine, ",") {
			if crabPosition, err := strconv.Atoi(crabPositionStr); err == nil {
				crabs = append(crabs, crabPosition)
				if crabPosition > maxCrabPos {
					maxCrabPos = crabPosition
				}
				if crabPosition < minCrabPos {
					minCrabPos = crabPosition
				}
			}
		}
	}

	minFuelCost := 0

	for optCrabPos := minCrabPos; optCrabPos <= maxCrabPos; optCrabPos++ {
		fuelCost := 0
		for _, crabPos := range crabs {
			fuelSteps := int(math.Abs(float64(optCrabPos - crabPos)))
			fuelCost += convertFuelStepsIntoCost(fuelSteps)
		}
		if minFuelCost == 0 || fuelCost < minFuelCost {
			minFuelCost = fuelCost
		}
	}

	fmt.Println(minFuelCost)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func convertFuelStepsIntoCost(fuelSteps int) int {
	// Formula N * (N+1) / 2
	return fuelSteps * (fuelSteps + 1) / 2
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
