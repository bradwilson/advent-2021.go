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

var data []string
var calledNumbers []int

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	for _, calledNumberText := range strings.Split(data[0], ",") {
		calledNumber, _ := strconv.Atoi(calledNumberText)
		calledNumbers = append(calledNumbers, calledNumber)
	}
}

func getPart1() int {
	return findWinner(true)
}

func getPart2() int {
	return findWinner(false)
}

func findWinner(winnerIsFirst bool) int {
	var boards []board
	currentIdx := 2

	for currentIdx < len(data) {
		boards = append(boards, makeBoard(data, currentIdx))
		currentIdx += 6
	}

	for _, calledNumber := range calledNumbers {
		for idx := len(boards) - 1; idx >= 0; idx-- {
			boards[idx].mark(calledNumber)
			if boards[idx].isWinner() {
				if winnerIsFirst || len(boards) == 1 {
					return boards[idx].sumUnmarked() * calledNumber
				}

				boards = append(boards[:idx], boards[idx+1:]...)
			}
		}
	}

	panic("should not get here")
}

type board struct {
	numbers [5][5]int
}

func makeBoard(data []string, firstLineIdx int) board {
	var numbers [5][5]int

	for idx := 0; idx < 5; idx++ {
		lineNumbers := strings.Split(strings.TrimSpace(data[firstLineIdx+idx]), " ")
		for idx2 := 0; idx2 < 5; idx2++ {
			lineNumber, _ := strconv.Atoi(lineNumbers[idx2])
			numbers[idx][idx2] = lineNumber
		}
	}

	return board{numbers}
}

func (b *board) mark(number int) {
	for idx := 0; idx < 5; idx++ {
		for idx2 := 0; idx2 < 5; idx2++ {
			if b.numbers[idx][idx2] == number {
				b.numbers[idx][idx2] = -1
			}
		}
	}
}

func (b *board) isWinner() bool {
	for idx := 0; idx < 5; idx++ {
		if b.numbers[idx][0]+b.numbers[idx][1]+b.numbers[idx][2]+b.numbers[idx][3]+b.numbers[idx][4] == -5 {
			return true
		}
		if b.numbers[0][idx]+b.numbers[1][idx]+b.numbers[2][idx]+b.numbers[3][idx]+b.numbers[4][idx] == -5 {
			return true
		}
	}

	return false
}

func (b *board) sumUnmarked() int {
	result := 0

	for idx := 0; idx < 5; idx++ {
		for idx2 := 0; idx2 < 5; idx2++ {
			if b.numbers[idx][idx2] != -1 {
				result += b.numbers[idx][idx2]
			}
		}
	}

	return result
}

func main() {
	var start time.Time

	start = time.Now()
	preCompute()
	fmt.Printf("[%s] Pre-compute\n", time.Since(start))

	start = time.Now()
	part1 := getPart1()
	fmt.Printf("[%s] Part 1: %d\n", time.Since(start), part1)

	start = time.Now()
	part2 := getPart2()
	fmt.Printf("[%s] Part 2: %d\n", time.Since(start), part2)
}
