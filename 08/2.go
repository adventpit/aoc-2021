package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	day := "08"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// segments
	//   aaaa
	// b      c
	// b      c
	//   dddd
	// e      f
	// e      f
	//   gggg

	scanner := bufio.NewScanner(file)
	noSegments := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	structSegments := [10][]string{
		[]string{"a", "b", "c", "e", "f", "g"},
		[]string{"c", "f"},
		[]string{"a", "c", "d", "e", "g"},
		[]string{"a", "c", "d", "f", "g"},
		[]string{"b", "c", "d", "f"},
		[]string{"a", "b", "d", "f", "g"},
		[]string{"a", "b", "d", "e", "f", "g"},
		[]string{"a", "c", "f"},
		[]string{"a", "b", "c", "d", "e", "f", "g"},
		[]string{"a", "b", "c", "d", "f", "g"},
	}
	// var occSegments = map[string]int{
	// 	"a": 8,
	// 	"b": 6,
	// 	"c": 8,
	// 	"d": 7,
	// 	"e": 4,
	// 	"f": 9,
	// 	"g": 7,
	// }
	sumOutput := 0

	for scanner.Scan() {
		inputLine := scanner.Text()
		inputSplit := strings.Split(inputLine, " | ")
		var occStructSegment = map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
			"d": 0,
			"e": 0,
			"f": 0,
			"g": 0,
		}
		var mapStructSegment = make(map[string][]string)
		var numberSegments [10][]string
		for _, uniqueSignalPattern := range strings.Split(inputSplit[0], " ") {
			splitSignalPattern := strings.Split(uniqueSignalPattern, "")
			for _, v := range splitSignalPattern {
				occStructSegment[v]++
			}
			if len(uniqueSignalPattern) == noSegments[1] {
				mapStructSegment["c"] = splitSignalPattern
			} else if len(uniqueSignalPattern) == noSegments[4] {
				mapStructSegment["d"] = splitSignalPattern
			} else if len(uniqueSignalPattern) == noSegments[7] {
				mapStructSegment["a"] = splitSignalPattern
			} else if len(uniqueSignalPattern) == noSegments[8] {
				mapStructSegment["g"] = splitSignalPattern
			}
		}
		for k, occSegment := range occStructSegment {
			if occSegment == 4 {
				mapStructSegment["e"] = []string{k}
			} else if occSegment == 9 {
				mapStructSegment["f"] = []string{k}
			} else if occSegment == 6 {
				mapStructSegment["b"] = []string{k}
			}
		}
		for _, v := range mapStructSegment["c"] {
			if v != mapStructSegment["f"][0] {
				mapStructSegment["c"] = []string{v}
			}
		}
		for _, v := range mapStructSegment["a"] {
			if v != mapStructSegment["f"][0] && v != mapStructSegment["c"][0] {
				mapStructSegment["a"] = []string{v}
			}
		}
		for _, v := range mapStructSegment["d"] {
			if v != mapStructSegment["f"][0] && v != mapStructSegment["c"][0] && v != mapStructSegment["b"][0] {
				mapStructSegment["d"] = []string{v}
			}
		}
		for _, v := range mapStructSegment["g"] {
			if v != mapStructSegment["f"][0] && v != mapStructSegment["e"][0] && v != mapStructSegment["d"][0] && v != mapStructSegment["c"][0] && v != mapStructSegment["b"][0] && v != mapStructSegment["a"][0] {
				mapStructSegment["g"] = []string{v}
			}
		}
		for number, segments := range structSegments {
			var newSegments []string
			for _, v := range segments {
				newSegments = append(newSegments, mapStructSegment[v][0])
			}
			numberSegments[number] = newSegments
		}
		rowOutput := 0
		for digit, fourDigitOutput := range strings.Split(inputSplit[1], " ") {
			for number, v := range numberSegments {
				if stringSlicesEqualDisregardOrder(v, strings.Split(fourDigitOutput, "")) {
					rowOutput += number * int(math.Pow(float64(10), float64(3-digit)))
					break
				}
			}
		}
		sumOutput += rowOutput
	}
	fmt.Println(sumOutput)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func stringSlicesEqualDisregardOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		appears := false
		for _, v2 := range b {
			if v == v2 {
				appears = true
				break
			}
		}
		if !appears {
			return false
		}
	}
	return true
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
