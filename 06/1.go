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
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", challengeDay, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), challengeDay))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fish []int

	for scanner.Scan() {
		inputLine := scanner.Text()
		for _, fishReproduceTimerStr := range strings.Split(inputLine, ",") {
			if fishReproduceTimer, err := strconv.Atoi(fishReproduceTimerStr); err == nil {
				fish = append(fish, fishReproduceTimer)
			}
		}
	}

	days := 80

	for day := 1; day <= days; day++ {
		for index, fishReproduceTimer := range fish {
			if fishReproduceTimer == 0 {
				fish[index] = 6
				fish = append(fish, 8)
			} else {
				fish[index]--
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day %d Fish %d", days, len(fish)))

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
