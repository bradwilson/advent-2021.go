package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	gamma, epsilon := 0, 0

	for idx := 0; idx < len(data[0]); idx++ {
		zeros, ones := 0, 0

		for _, line := range data {
			if line[idx] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		gamma *= 2
		epsilon *= 2

		if zeros > ones {
			epsilon++
		} else {
			gamma++
		}
	}

	return gamma * epsilon
}

func getPart2() int {
	oxygen := getValue(false)
	co2 := getValue(true)

	return oxygen * co2
}

func getValue(leastCommon bool) int {
	values := make([]string, len(data))
	copy(values, data)

	for idx := 0; idx < len(data[0]); idx++ {
		zeros, ones := 0, 0

		for _, line := range values {
			if line[idx] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if (leastCommon && ones >= zeros) || (!leastCommon && zeros > ones) {
			values = removeAll(values, '1', idx)
		} else {
			values = removeAll(values, '0', idx)
		}

		if len(values) == 1 {
			break
		}
	}

	result := 0
	for _, ch := range values[0] {
		result *= 2
		if ch == '1' {
			result++
		}
	}

	return result
}

func removeAll(values []string, element byte, idx int) []string {
	for idx2 := len(values) - 1; idx2 >= 0; idx2-- {
		if values[idx2][idx] == element {
			values = append(values[:idx2], values[idx2+1:]...)
		}
	}

	return values
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
