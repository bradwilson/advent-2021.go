package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

var data [][]int
var maxX, maxY int

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
	total := 0

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			cur := tryGetData(x, y)
			left := tryGetData(x-1, y)
			right := tryGetData(x+1, y)
			top := tryGetData(x, y-1)
			bottom := tryGetData(x, y+1)

			if cur < left && cur < right && cur < top && cur < bottom {
				total += cur + 1
			}
		}
	}

	return total
}

func getPart2() int {
	basinSizes := make([]int, 0, 25)

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			basinSize := calculateBasin(x, y)
			if basinSize != 0 {
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func tryGetData(x, y int) int {
	if x < 0 || x >= maxX || y < 0 || y >= maxY {
		return 10
	}
	return data[x][y]
}

func calculateBasin(x, y int) int {
	if x < 0 || x >= maxX || y < 0 || y >= maxY || data[x][y] > 8 {
		return 0
	}

	data[x][y] = 10
	return 1 + calculateBasin(x-1, y) + calculateBasin(x+1, y) + calculateBasin(x, y-1) + calculateBasin(x, y+1)
}

func parseLine(text string) []int {
	result := make([]int, len(text))

	for idx, c := range text {
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
