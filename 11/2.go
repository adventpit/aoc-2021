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

const gridSize = 10

var octopusMap [gridSize][gridSize]int

func main() {
	day := "11"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/data/dev/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
		for x, octopusEnergyStr := range strings.Split(inputLine, "") {
			if octopusEnergy, err := strconv.Atoi(octopusEnergyStr); err == nil {
				octopusMap[y][x] = octopusEnergy
			}
		}
		y++
	}

	flashes := 0
	steps := 1000
	for i := 1; i <= steps; i++ {
		var flashedPos [][2]int
		for y, octoLine := range octopusMap {
			for x := range octoLine {
				flashedPos = increaseOctopusEnergy(x, y, flashedPos)
			}
		}
		if len(flashedPos) == 100 {
			fmt.Println("Step for simultaneous flashes:", i)
			break
		}
		flashes += len(flashedPos)
		for y, octoLine := range octopusMap {
			for x, octopusEnergy := range octoLine {
				if octopusEnergy > 9 {
					octopusMap[y][x] = 0
				}
			}
		}
	}
	fmt.Println(octopusMap)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func increaseOctopusEnergy(x int, y int, flashedPos [][2]int) [][2]int {
	octopusMap[y][x]++
	if octopusMap[y][x] > 9 && !contains(flashedPos, [2]int{x, y}) {
		flashedPos = append(flashedPos, [2]int{x, y})
		if y > 0 {
			flashedPos = increaseOctopusEnergy(x, y-1, flashedPos)
			if x > 0 {
				flashedPos = increaseOctopusEnergy(x-1, y-1, flashedPos)
			}
			if x < gridSize-1 {
				flashedPos = increaseOctopusEnergy(x+1, y-1, flashedPos)
			}
		}
		if y < gridSize-1 {
			flashedPos = increaseOctopusEnergy(x, y+1, flashedPos)
			if x > 0 {
				flashedPos = increaseOctopusEnergy(x-1, y+1, flashedPos)
			}
			if x < gridSize-1 {
				flashedPos = increaseOctopusEnergy(x+1, y+1, flashedPos)
			}
		}
		if x > 0 {
			flashedPos = increaseOctopusEnergy(x-1, y, flashedPos)
		}
		if x < gridSize-1 {
			flashedPos = increaseOctopusEnergy(x+1, y, flashedPos)
		}
	}
	return flashedPos
}

func contains(elems [][2]int, v [2]int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
