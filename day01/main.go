package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var data []int

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, value)
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}
}

func getPart1() int {
	var increments int = 0
	lastValue := data[0]

	for idx := 1; idx < len(data); idx++ {
		value := data[idx]
		if value > lastValue {
			increments++
		}
		lastValue = value
	}

	return increments
}

func computeWindow(startIndex int) int {
	return data[startIndex] + data[startIndex+1] + data[startIndex+2]
}

func getPart2() int {
	lastValue := computeWindow(0)
	increments := 0

	for idx := 3; idx < len(data); idx++ {
		value := computeWindow(idx - 2)
		if value > lastValue {
			increments++
		}
		lastValue = value
	}

	return increments
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
