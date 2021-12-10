package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var data []digitsAndValues

type digitsAndValues struct {
	digits [10]map[rune]bool
	values [4]map[rune]bool
}

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
}

func getPart1() int {
	result := 0

	for _, line := range data {
		for _, value := range line.values {
			if len(value) != 5 && len(value) != 6 {
				result++
			}
		}
	}

	return result
}

func getPart2() int {
	result := 0

	for _, line := range data {
		result += getDecodedValue(line)
	}

	return result
}

func getDecodedValue(line digitsAndValues) int {
	var digitMap [10]map[rune]bool

	// Single length matches

	digitMap[1] = ofLength(line.digits, 2)[0] // cf
	digitMap[4] = ofLength(line.digits, 4)[0] // bcdf
	digitMap[7] = ofLength(line.digits, 3)[0] // acf
	digitMap[8] = ofLength(line.digits, 7)[0] // abcdefg

	// Five digit length matches

	fiveDigitValues := ofLength(line.digits, 5)

	digitMap[2], fiveDigitValues = matchSegments(fiveDigitValues, digitMap[4], 2) // acdef (shares cd with 4)
	digitMap[3], fiveDigitValues = matchSegments(fiveDigitValues, digitMap[1], 2) // acdfg (shares cf with 1)
	digitMap[5] = fiveDigitValues[0]

	// Six digit length matches

	sixDigitValues := ofLength(line.digits, 6)

	digitMap[6], sixDigitValues = matchSegments(sixDigitValues, digitMap[1], 1) // abdefg (shares f with 1)
	digitMap[9], sixDigitValues = matchSegments(sixDigitValues, digitMap[4], 4) // abcdfg (shares bcdf with 4)
	digitMap[0] = sixDigitValues[0]

	// Match the digits

	result := 0

	for _, resultDigit := range line.values {
		for idx, digit := range digitMap {
			if len(digit) == len(resultDigit) && getCommonCount(digit, resultDigit) == len(digit) {
				result = result*10 + idx
				break
			}
		}
	}

	return result
}

func getCommonCount(val1, val2 map[rune]bool) int {
	result := 0

	for r := range val1 {
		if val2[r] {
			result++
		}
	}

	return result
}

func matchSegments(digits []map[rune]bool, toMatch map[rune]bool, intersectionCount int) (map[rune]bool, []map[rune]bool) {
	for idx, digit := range digits {
		if getCommonCount(digit, toMatch) == intersectionCount {
			match := digits[idx]
			return match, append(digits[:idx], digits[idx+1:]...)
		}
	}

	panic("Couldn't find it")
}

func ofLength(digits [10]map[rune]bool, length int) []map[rune]bool {
	result := make([]map[rune]bool, 0, 3)

	for _, digit := range digits {
		if len(digit) == length {
			result = append(result, digit)
		}
	}

	return result
}

func parseLine(text string) digitsAndValues {
	pieces := strings.Split(text, "|")

	var digits [10]map[rune]bool
	for idx, digit := range strings.Split(strings.TrimSpace(pieces[0]), " ") {
		digitMap := make(map[rune]bool)
		for _, r := range digit {
			digitMap[r] = true
		}
		digits[idx] = digitMap
	}

	var values [4]map[rune]bool
	for idx, value := range strings.Split(strings.TrimSpace(pieces[1]), " ") {
		valueMap := make(map[rune]bool)
		for _, v := range value {
			valueMap[v] = true
		}
		values[idx] = valueMap
	}

	return digitsAndValues{digits, values}
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
