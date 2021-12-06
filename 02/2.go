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
	day := "02"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := 0
	y := 0
	aim := 0

	for scanner.Scan() {
		command := scanner.Text()
		x, y, aim = move(command, x, y, aim)
	}

	fmt.Println(fmt.Sprintf("Horizontal position multiplied by depth: %d", x*y))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func move(command string, x int, y int, aim int) (int, int, int) {
	split := strings.Split(command, " ")
	commandName := split[0]
	commandStrength := split[1]

	if strength, err := strconv.Atoi(commandStrength); err == nil {
		switch commandName {
		case "forward":
			x += strength
			y += aim * strength
			break
		case "down":
			aim += strength
			break
		case "up":
			aim -= strength
			break
		}
	}

	return x, y, aim
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
