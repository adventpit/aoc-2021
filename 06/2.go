package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	challengeDay := "06"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", challengeDay, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), challengeDay))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var durations [9]int

	for scanner.Scan() {
		inputLine := scanner.Text()
		for _, fishReproduceTimerStr := range strings.Split(inputLine, ",") {
			if fishReproduceTimer, err := strconv.Atoi(fishReproduceTimerStr); err == nil {
				durations[fishReproduceTimer]++
			}
		}
	}

	days := 256
	fish := 0

	for day := 1; day <= days; day++ {
		newFish := 0
		for i := 0; i < 9; i++ {
			if i == 0 {
				newFish = durations[0]
			} else {
				durations[i-1] = durations[i]
			}
		}
		durations[8] = newFish
		durations[6] += newFish
	}
	for _, v := range durations {
		fish += v
	}
	fmt.Println(fmt.Sprintf("Day %d Fish %d", days, fish))

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
