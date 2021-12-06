package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	day := "03"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const bitLength = 12
	occ := 0
	var bitSubcols = make(map[string][]string)
	binaryEntries := 0

	for scanner.Scan() {
		binaryInput := scanner.Text()
		bit := binaryInput[0:1]
		if bit == "1" {
			occ += 1
		}
		subcol, _ := bitSubcols[bit]
		bitSubcols[bit] = append(subcol, binaryInput)
		binaryEntries++
	}

	var bitOxygen []string
	var bitCO2 []string

	if determineBitValue(occ, binaryEntries, true) {
		bitOxygen = bitSubcols["1"]
		bitCO2 = bitSubcols["0"]
	} else {
		bitOxygen = bitSubcols["0"]
		bitCO2 = bitSubcols["1"]
	}

	for i := 1; i < bitLength; i++ {
		if len(bitOxygen) > 1 {
			bitOxygen = checkBitSubcol(bitOxygen, i, true)
		}
		if len(bitCO2) > 1 {
			bitCO2 = checkBitSubcol(bitCO2, i, false)
		}
	}

	oxygenRate := convertBitsToDecimal(bitOxygen[0])
	co2Rate := convertBitsToDecimal(bitCO2[0])

	fmt.Println(fmt.Sprintf("Decimal oxygen rate multiplied with decimal CO2 rate %d", oxygenRate*co2Rate))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func checkBitSubcol(bitArray []string, bitPos int, oxygen bool) []string {
	occ := 0
	entries := 0
	var subcols = make(map[string][]string)

	for _, bits := range bitArray {
		bit := bits[bitPos : bitPos+1]
		if bit == "1" {
			occ += 1
		}
		subcol, _ := subcols[bit]
		subcols[bit] = append(subcol, bits)
		entries++
	}

	if determineBitValue(occ, entries, oxygen) {
		return subcols["1"]
	} else {
		return subcols["0"]
	}
}

func determineBitValue(bitOccurences int, binaryEntries int, oxygen bool) bool {
	entries := float32(binaryEntries) / float32(2)
	occ := float32(bitOccurences)
	if (oxygen && occ >= entries) || (!oxygen && occ < entries) {
		return true
	} else {
		return false
	}
}

func convertBitsToDecimal(bits string) int64 {
	decimal, _ := strconv.ParseInt(bits, 2, 64)
	return decimal
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
