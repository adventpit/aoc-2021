package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var xLength int = 0
var yLength int = 0
var caveMap [][]int

func main() {
	day := "09"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/data/dev/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

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

	var threeMajorBasins [3]int

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
			var basin [][2]int
			basin = findBasinPositions(x, y, basin)
			basin = unique(basin)
			basinSize := len(basin)
			if basinSize > threeMajorBasins[0] {
				threeMajorBasins[0] = basinSize
			}
			sort.Ints(threeMajorBasins[:])
		}
	}

	multipliedBasins := threeMajorBasins[0] * threeMajorBasins[1] * threeMajorBasins[2]
	fmt.Println("Multiplied basins: ", multipliedBasins)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func findBasinPositions(x int, y int, basin [][2]int) [][2]int {
	value := caveMap[y][x]
	basin = append(basin, [2]int{x, y})
	if value == 8 {
		return basin
	}
	if y > 0 {
		if value+1 == caveMap[y-1][x] {
			basin = findBasinPositions(x, y-1, basin)
		}
	}
	if y < yLength-1 {
		if value+1 == caveMap[y+1][x] {
			basin = findBasinPositions(x, y+1, basin)
		}
	}
	if x > 0 {
		if value+1 == caveMap[y][x-1] {
			basin = findBasinPositions(x-1, y, basin)
		}
	}
	if x < xLength-1 {
		if value+1 == caveMap[y][x+1] {
			basin = findBasinPositions(x+1, y, basin)
		}
	}
	return basin
}

func unique(arr [][2]int) [][2]int {
	occurred := map[[2]int]bool{}
	result := [][2]int{}
	for e := range arr {
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true

			result = append(result, arr[e])
		}
	}

	return result
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
