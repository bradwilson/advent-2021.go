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
var fish [9]int

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

	for _, age := range atoi(strings.Split(data[0], ",")) {
		fish[age]++
	}
}

func getPart1() int {
	for idx := 0; idx < 80; idx++ {
		ageFish()
	}

	return fishSum()
}

func getPart2() int {
	for idx := 80; idx < 256; idx++ {
		ageFish()
	}

	return fishSum()
}

func ageFish() {
	cycle := fish[0]

	for age := 1; age < 9; age++ {
		fish[age-1] = fish[age]
	}

	fish[8] = cycle
	fish[6] += cycle
}

func fishSum() int {
	result := 0

	for idx := 0; idx < 9; idx++ {
		result += fish[idx]
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
