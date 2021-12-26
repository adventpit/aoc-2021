package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var caveConnections = make(map[string][]string)

func main() {
	day := "12"
	challengePart := "1"
	defer Duration(Track(fmt.Sprintf("Advent of Code challenge Day %s Part %s", day, challengePart)))
	file, err := os.Open(fmt.Sprintf("%s/data/dev/aoc/aoc-2021/%s/input.txt", os.Getenv("HOME"), day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputLine := scanner.Text()
		caveConn := strings.Split(inputLine, "-")

		connectionsOne := caveConnections[caveConn[0]]
		if !contains(connectionsOne, caveConn[1]) {
			caveConnections[caveConn[0]] = append(connectionsOne, caveConn[1])
		}
		connectionsTwo := caveConnections[caveConn[1]]
		if !contains(connectionsTwo, caveConn[0]) {
			caveConnections[caveConn[1]] = append(connectionsTwo, caveConn[0])
		}
	}

	var pathOptions [][]string
	pathOptions = findPath([]string{"start"}, pathOptions)

	fmt.Println(len(pathOptions))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func findPath(path []string, pathOptions [][]string) [][]string {
	startPoint := path[len(path)-1]
	for _, connectedCave := range caveConnections[startPoint] {
		if connectedCave == "end" {
			pathOptions = append(pathOptions, localPathCopy(path, connectedCave))
			continue
		} else if MyString(connectedCave).IsUpperCase() {
			pathOptions = findPath(localPathCopy(path, connectedCave), pathOptions)
		} else {
			if !contains(localPathCopy(path, ""), connectedCave) {
				pathOptions = findPath(localPathCopy(path, connectedCave), pathOptions)
			}
		}
	}
	return pathOptions
}

func localPathCopy(path []string, connectedCave string) []string {
	localPath := make([]string, len(path))
	copy(localPath, path)
	if connectedCave != "" {
		localPath = append(localPath, connectedCave)
	}
	return localPath
}

type MyString string

func (s MyString) IsUpperCase() bool {
	return s == MyString(strings.ToUpper(string(s)))
}

func contains(elems []string, v string) bool {
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
