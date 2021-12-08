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

const matrixSize = 1000

func main() {
	day := "05"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [matrixSize][matrixSize]int

	for scanner.Scan() {
		inputLine := scanner.Text()
		points := strings.Split(inputLine, " -> ")
		pointStart := strings.Split(points[0], ",")
		pointEnd := strings.Split(points[1], ",")
		if xStart, err := strconv.Atoi(pointStart[0]); err == nil {
			if xEnd, err := strconv.Atoi(pointEnd[0]); err == nil {
				if yStart, err := strconv.Atoi(pointStart[1]); err == nil {
					if yEnd, err := strconv.Atoi(pointEnd[1]); err == nil {
						matrix = increasePointsOnMatrix(xStart, xEnd, yStart, yEnd, matrix)
					}
				}
			}
		}
	}

	counterDangerZone := 0

	for x := 0; x < matrixSize; x++ {
		for y := 0; y < matrixSize; y++ {
			if matrix[x][y] > 1 {
				counterDangerZone++
			}
		}
	}

	fmt.Println("Danger zones: ", counterDangerZone)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func increasePointsOnMatrix(xStart int, xEnd int, yStart int, yEnd int, matrix [matrixSize][matrixSize]int) [matrixSize][matrixSize]int {
	if xStart == xEnd { // Only horizontal lines
		if yStart > yEnd {
			for y := yStart; y >= yEnd; y-- {
				matrix[xStart][y]++
			}
		} else {
			for y := yStart; y <= yEnd; y++ {
				matrix[xStart][y]++
			}
		}
	} else if yStart == yEnd { // Only vertical lines
		if xStart > xEnd {
			for x := xStart; x >= xEnd; x-- {
				matrix[x][yStart]++
			}
		} else {
			for x := xStart; x <= xEnd; x++ {
				matrix[x][yStart]++
			}
		}
	} else { // diagonals
		x := xStart
		y := yStart
		if xStart > xEnd && yStart > yEnd {
			for ; x >= xEnd; x-- {
				matrix[x][y]++
				y--
			}
		} else if xStart > xEnd && yStart < yEnd {
			for ; x >= xEnd; x-- {
				matrix[x][y]++
				y++
			}
		} else if xStart < xEnd && yStart > yEnd {
			for ; x <= xEnd; x++ {
				matrix[x][y]++
				y--
			}
		} else if xStart < xEnd && yStart < yEnd {
			for ; x <= xEnd; x++ {
				matrix[x][y]++
				y++
			}
		}
	}
	return matrix
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
