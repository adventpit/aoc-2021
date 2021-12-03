package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	day := "01"
	challengePart := "1"
	file, err := os.Open(fmt.Sprintf("/home/pitm-geshdo/dev/training/aoc/aoc-2021/%s/input.txt", day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}

	fmt.Println("Advent of Code challenge Day " + day + " Part " + challengePart)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
