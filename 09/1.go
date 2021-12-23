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
	day := "09"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/data/dev/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var caveMap [][]int
	xLength := 0
	yLength := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputLine := scanner.Text()
		var lineMap []int
		for _, valuePos := range strings.Split(inputLine, "") {
			if intPos, err := strconv.Atoi(valuePos); err == nil {
				lineMap = append(lineMap, intPos)
			}
			if yLength == 0 {
				xLength++
			}
		}
		caveMap = append(caveMap, lineMap)
		yLength++
	}

	var riskLevels []int

	for y, line := range caveMap {
		for x, value := range line {
			if y > 0 {
				if value >= caveMap[y-1][x] {
					continue
				}
			}
			if y < yLength-1 {
				if value >= caveMap[y+1][x] {
					continue
				}
			}
			if x > 0 {
				if value >= caveMap[y][x-1] {
					continue
				}
			}
			if x < xLength-1 {
				if value >= caveMap[y][x+1] {
					continue
				}
			}
			riskLevels = append(riskLevels, value+1)
		}
	}

	fmt.Println(riskLevels)
	summedRiskLevels := 0
	for _, v := range riskLevels {
		summedRiskLevels += v
	}

	fmt.Printf("Sum of Risk levels: %d \n", summedRiskLevels)

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
