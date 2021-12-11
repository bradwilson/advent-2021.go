package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var data [][]int
var maxX int
var maxY int
var firstFullFlash = -1

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, parseLine(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	maxX = len(data)
	maxY = len(data[0])
}

func getPart1() int {
	result := 0

	for idx := 1; idx <= 100; idx++ {
		result += step(idx)
	}

	return result
}

func getPart2() int {
	idx := 101

	for firstFullFlash == -1 {
		step(idx)
		idx++
	}

	return firstFullFlash
}

func step(step int) int {
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			processOctopus(x, y)
		}
	}

	result := 0

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if data[x][y] == 10 {
				result++
				data[x][y] = 0
			}
		}
	}

	if firstFullFlash == -1 && result == maxX*maxY {
		firstFullFlash = step
	}

	return result
}

func processOctopus(x, y int) {
	if x < 0 || x >= maxX || y < 0 || y >= maxY {
		return
	}

	if data[x][y] < 9 {
		data[x][y]++
	} else if data[x][y] == 9 {
		data[x][y] = 10
		processOctopus(x-1, y-1)
		processOctopus(x-1, y)
		processOctopus(x-1, y+1)
		processOctopus(x, y-1)
		processOctopus(x, y+1)
		processOctopus(x+1, y-1)
		processOctopus(x+1, y)
		processOctopus(x+1, y+1)
	}
}

func parseLine(line string) []int {
	result := make([]int, len(line))

	for idx, c := range line {
		result[idx] = int(c) - int('0')
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
