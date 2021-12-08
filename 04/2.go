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
	day := "04"
	challengePart := "2"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/dev/training/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	board := 0
	var bingoDraw []int
	var bingoBoards = make(map[int][5][5]int)

	for scanner.Scan() {
		inputLine := scanner.Text()
		if index == 0 {
			for _, draw := range strings.Split(inputLine, ",") {
				if drawNumber, err := strconv.Atoi(draw); err == nil {
					bingoDraw = append(bingoDraw, drawNumber)
				}
			}
		} else if index > 1 && inputLine != "" {
			var numberLine [5]int
			numberIndex := 0
			for _, input := range strings.Split(inputLine, " ") {
				if number, err := strconv.Atoi(input); err == nil {
					numberLine[numberIndex] = number
					numberIndex++
				}
			}
			boardIndex := (index - 2) % 5
			currentBoard := bingoBoards[board]
			currentBoard[boardIndex] = numberLine
			bingoBoards[board] = currentBoard
		}
		if index > 1 {
			if inputLine == "" {
				board += 1
			}
		}
		index++
	}

	bingoBoard := -1
	var finishedBoards []int

	for _, bingoNumber := range bingoDraw {
		bingoBoard, bingoBoards, finishedBoards = checkWithBingoNumber(bingoNumber, bingoBoards, finishedBoards)
		if bingoBoard > -1 {
			sumUnmarked := 0
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					value := bingoBoards[bingoBoard][x][y]
					if value > 0 {
						sumUnmarked += value
					}
				}
			}
			fmt.Println("Bingo! ", bingoBoard)
			fmt.Println("Final score: ", sumUnmarked*bingoNumber)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func checkWithBingoNumber(bingoNumber int, bingoBoards map[int][5][5]int, finishedBoards []int) (int, map[int][5][5]int, []int) {
	for boardNumber, board := range bingoBoards {
		finished := false
		for _, finishedBoardNumber := range finishedBoards {
			if boardNumber == finishedBoardNumber {
				finished = true
			}
		}
		if finished {
			continue
		}
		bingoMarksX := [5]int{0, 0, 0, 0, 0}
		for x := 0; x < 5; x++ {
			bingoMarksY := 0
			for y := 0; y < 5; y++ {
				if board[x][y] == bingoNumber {
					board[x][y] = -1
				}
				if board[x][y] == -1 {
					bingoMarksY++
					bingoMarksX[y]++
				}
				if bingoMarksY == 5 || bingoMarksX[y] == 5 {
					bingoBoards[boardNumber] = board
					finishedBoards = append(finishedBoards, boardNumber)
					finished = true
					if len(finishedBoards) == len(bingoBoards) {
						return boardNumber, bingoBoards, finishedBoards
					}
					break
				}
			}
			if finished {
				break
			}
		}
		bingoBoards[boardNumber] = board
	}
	return -1, bingoBoards, finishedBoards
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
