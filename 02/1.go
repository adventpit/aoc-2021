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
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := 0
	y := 0

	for scanner.Scan() {
		command := scanner.Text()
		x, y = move(command, x, y)
	}

	fmt.Println(fmt.Sprintf("Horizontal position multiplied by depth: %d", x*y))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func move(command string, x int, y int) (int, int) {
	split := strings.Split(command, " ")
	commandName := split[0]
	commandStrength := split[1]

	if strength, err := strconv.Atoi(commandStrength); err == nil {
		switch commandName {
		case "forward":
			x += strength
			break
		case "down":
			y += strength
			break
		case "up":
			y -= strength
			break
		}
	}

	return x, y
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
