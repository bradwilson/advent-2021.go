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
}

func getPart1() int {
	return calculateBoard(false)
}

func getPart2() int {
	return calculateBoard(true)
}

type point struct {
	x, y int
}

func calculateBoard(diagonal bool) int {
	board := make(map[point]int)

	for _, row := range data {
		pieces := strings.Split(row, " ")
		left := atoi(strings.Split(pieces[0], ","))
		right := atoi(strings.Split(pieces[2], ","))
		xInc := direction(left[0], right[0])
		yInc := direction(left[1], right[1])

		if xInc != 0 && yInc != 0 && !diagonal {
			continue
		}

		x := left[0]
		y := left[1]
		limit := max(abs(left[0]-right[0]), abs(left[1]-right[1]))

		for idx := 0; idx <= limit; idx++ {
			coord := point{x, y}
			board[coord] = board[coord] + 1
			x += xInc
			y += yInc
		}
	}

	result := 0

	for _, count := range board {
		if count > 1 {
			result++
		}
	}

	return result
}

func atoi(strValues []string) []int {
	result := make([]int, len(strValues))

	for idx, strValue := range strValues {
		intValue, _ := strconv.Atoi(strValue)
		result[idx] = intValue
	}

	return result
}

func direction(left int, right int) int {
	if left > right {
		return -1
	}
	if left < right {
		return 1
	}
	return 0
}

func abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func max(value1 int, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
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
