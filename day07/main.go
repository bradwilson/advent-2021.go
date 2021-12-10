package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var data []string
var crabs []int
var low int
var high int

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

	crabs = atoi(strings.Split(data[0], ","))
	for i, crab := range crabs {
		if i == 0 || crab < low {
			low = crab
		}
		if i == 0 || crab > high {
			high = crab
		}
	}
}

func getPart1() int {
	return findBestFuel(part1Calc)
}

func getPart2() int {
	return findBestFuel(part2Calc)
}

func part1Calc(crab int, current int) int {
	return abs(crab - current)
}

func part2Calc(crab int, current int) int {
	distance := abs(crab - current)
	return distance * (distance + 1) / 2
}

func findBestFuel(fuelCost func(int, int) int) int {
	bestFuel := math.MaxInt

	for current := low; current <= high; current++ {
		target := 0

		for _, crab := range crabs {
			target += fuelCost(crab, current)
		}

		if bestFuel > target {
			bestFuel = target
		}
	}

	return bestFuel
}

func abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func atoi(strValues []string) []int {
	result := make([]int, len(strValues))

	for idx, strValue := range strValues {
		intValue, _ := strconv.Atoi(strValue)
		result[idx] = intValue
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
