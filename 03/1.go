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
	day := "03"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const bitLength = 12
	var bits = make(map[int]int)
	binaryEntries := 0

	for scanner.Scan() {
		binaryInput := scanner.Text()
		bitsInput := strings.Split(binaryInput, "")
		for i, bit := range bitsInput {
			occ, exists := bits[i]
			if exists {
				if bit == "1" {
					occ += 1
				}
			} else {
				if bit == "1" {
					occ = 1
				}
			}
			bits[i] = occ
		}
		binaryEntries++
	}

	var bitGamma [bitLength]string
	var bitEpsilon [bitLength]string

	for i := 0; i < bitLength; i++ {
		if determineBitValue(bits[i], binaryEntries) {
			bitGamma[i] = "1"
			bitEpsilon[i] = "0"
		} else {
			bitGamma[i] = "0"
			bitEpsilon[i] = "1"
		}
	}

	gammaRate := convertBitsToDecimal(strings.Join(bitGamma[:], ""))
	epsilonRate := convertBitsToDecimal(strings.Join(bitEpsilon[:], ""))

	fmt.Println(fmt.Sprintf("Decimal Gamma rate multiplied with decimal epsilon rate %d", gammaRate*epsilonRate))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func determineBitValue(bitOccurences int, binaryEntries int) bool {
	if bitOccurences >= (binaryEntries / 2) {
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
