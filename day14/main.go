package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

var data []string
var polymer string
var instructions = map[string]byte{}

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

	polymer = data[0]

	for _, line := range data[2:] {
		instructions[line[0:2]] = line[6]
	}
}

func getPart1() int {
	return countPolymer(10)
}

func getPart2() int {
	return countPolymer(40)
}

func countPolymer(iterationCount int) int {
	pairCounts := map[string]int{}

	for idx := 0; idx < len(polymer)-1; idx++ {
		pair := polymer[idx : idx+2]
		pairCounts[pair] = pairCounts[pair] + 1
	}

	for idx := 0; idx < iterationCount; idx++ {
		newCounts := map[string]int{}

		for pair, currentCount := range pairCounts {
			instruction := instructions[pair]

			pair1 := string([]byte{pair[0], instruction})
			newCounts[pair1] = newCounts[pair1] + currentCount

			pair2 := string([]byte{instruction, pair[1]})
			newCounts[pair2] = newCounts[pair2] + currentCount
		}

		pairCounts = newCounts
	}

	result := map[byte]int{}

	for pair, count := range pairCounts {
		result[pair[1]] = result[pair[1]] + count
	}

	result[polymer[0]] = result[polymer[0]] + 1

	values := make([]int, 0, len(result))
	for _, value := range result {
		values = append(values, value)
	}

	sort.Slice(values, func(x, y int) bool {
		return values[x] < values[y]
	})

	return values[len(values)-1] - values[0]
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
