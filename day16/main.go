package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var data string
var version int
var result int

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data = parseLine(scanner.Text())

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	idx := 0
	version, result = parsePacket(&idx)
}

func getPart1() int {
	return version
}

func getPart2() int {
	return result
}

func parsePacket(idx *int) (versionTotal int, result int) {
	version, _ := strconv.ParseInt(data[*idx:*idx+3], 2, 4)
	versionTotal += int(version)

	packetType, _ := strconv.ParseInt(data[*idx+3:*idx+6], 2, 4)

	*idx += 6

	if packetType == 4 {
		var valueText string

		for {
			firstBit := data[*idx]
			valueText += data[*idx+1 : *idx+5]
			*idx += 5

			if firstBit == '0' {
				break
			}
		}

		resultValue, _ := strconv.ParseInt(valueText, 2, 64)
		result = int(resultValue)
	} else {
		var results []int

		lengthType := data[*idx]
		*idx++

		if lengthType == '0' {
			length, _ := strconv.ParseInt(data[*idx:*idx+15], 2, 64)
			*idx += 15
			targetIndex := *idx + int(length)

			for *idx < targetIndex {
				v, r := parsePacket(idx)
				versionTotal += v
				results = append(results, r)
			}
		} else {
			packetCount, _ := strconv.ParseInt(data[*idx:*idx+11], 2, 64)
			*idx += 11

			for ; packetCount > 0; packetCount-- {
				v, r := parsePacket(idx)
				versionTotal += v
				results = append(results, r)
			}
		}

		switch packetType {
		case 0:
			result = sum(results)
		case 1:
			result = mult(results)
		case 2:
			result = min(results)
		case 3:
			result = max(results)
		case 5:
			if results[0] > results[1] {
				result = 1
			}
		case 6:
			if results[0] < results[1] {
				result = 1
			}
		case 7:
			if results[0] == results[1] {
				result = 1
			}
		}
	}

	return versionTotal, result
}

func sum(results []int) int {
	result := 0

	for _, v := range results {
		result += v
	}

	return result
}

func mult(results []int) int {
	result := 1

	for _, v := range results {
		result *= v
	}

	return result
}

func min(results []int) int {
	result := math.MaxInt

	for _, v := range results {
		if v < result {
			result = v
		}
	}

	return result
}

func max(results []int) int {
	result := math.MinInt

	for _, v := range results {
		if v > result {
			result = v
		}
	}

	return result
}

func parseLine(line string) string {
	var result string

	for idx := 0; idx < len(line); idx += 2 {
		i, _ := strconv.ParseInt(line[idx:idx+2], 16, 9)
		result += fmt.Sprintf("%08b", i)
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
