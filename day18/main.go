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
	var current string

	for idx, line := range data {
		if idx == 0 {
			current = line
		} else {
			current = evaluate(fmt.Sprintf("[%s,%s]", current, line))
		}
	}

	return getMagnitude(current)
}

func getPart2() int {
	largest := 0

	for idx := 0; idx < len(data); idx++ {
		for idx2 := 0; idx2 < len(data); idx2++ {
			if idx != idx2 {
				source := fmt.Sprintf("[%s,%s]", data[idx], data[idx2])
				result := evaluate(source)
				magnitude := getMagnitude(result)

				if magnitude > largest {
					largest = magnitude
				}
			}
		}
	}

	return largest
}

func evaluate(value string) string {
	depth, idx, lastOpen := 0, 0, -1
	var digitLocations []int

	// Look for first explosion

	for idx < len(value) {
		switch value[idx] {
		case '[':
			lastOpen = idx
			depth++
			idx++

		case ']':
			if depth == 5 {
				pieces := strings.Split(value[lastOpen+1:idx], ",")
				toExplodeLeft, _ := strconv.Atoi(pieces[0])
				toExplodeRight, _ := strconv.Atoi(pieces[1])

				prevDigitEndIdx := -1
				for len(digitLocations) > 0 {
					digitIdx := digitLocations[len(digitLocations)-1]
					digitLocations = digitLocations[:len(digitLocations)-1]

					if digitIdx < lastOpen {
						prevDigitEndIdx = digitIdx
						break
					}
				}

				prevDigitStartIdx := prevDigitEndIdx
				for len(digitLocations) > 0 {
					digitIdx := digitLocations[len(digitLocations)-1]
					digitLocations = digitLocations[:len(digitLocations)-1]

					if digitIdx == prevDigitStartIdx-1 {
						prevDigitStartIdx = digitIdx
					} else {
						break
					}
				}

				nextDigitStartIdx := idx
				for ; nextDigitStartIdx < len(value); nextDigitStartIdx++ {
					if value[nextDigitStartIdx] >= '0' && value[nextDigitStartIdx] <= '9' {
						break
					}
				}

				nextDigitEndIdx := nextDigitStartIdx
				for ; nextDigitEndIdx < len(value); nextDigitEndIdx++ {
					if value[nextDigitEndIdx+1] < '0' || value[nextDigitEndIdx+1] > '9' {
						break
					}
				}

				newValue := ""

				if prevDigitEndIdx == -1 {
					newValue += value[:lastOpen]
				} else {
					prevValue, _ := strconv.Atoi(value[prevDigitStartIdx : prevDigitEndIdx+1])
					newValue += fmt.Sprintf("%s%d%s", value[:prevDigitStartIdx], toExplodeLeft+prevValue, value[prevDigitEndIdx+1:lastOpen])
				}

				newValue += "0"

				if nextDigitStartIdx == len(value) {
					newValue += value[idx+1:]
				} else {
					nextValue, _ := strconv.Atoi(value[nextDigitStartIdx : nextDigitEndIdx+1])
					newValue += fmt.Sprintf("%s%d%s", value[idx+1:nextDigitStartIdx], toExplodeRight+nextValue, value[nextDigitEndIdx+1:])
				}

				value = newValue
				idx = 0
				depth = 0
				digitLocations = make([]int, 0)
			} else {
				idx++
				depth--
			}

		case ',':
			idx++

		default:
			digitLocations = append(digitLocations, idx)
			idx++
		}
	}

	// Look for first split

	firstDigitIdx := math.MaxInt

	for _, digitIdx := range digitLocations {
		if digitIdx-1 == firstDigitIdx {
			overSizeValue, _ := strconv.Atoi(value[firstDigitIdx : digitIdx+1])
			newValue := fmt.Sprintf("%s%s%s", value[:firstDigitIdx], split(overSizeValue), value[digitIdx+1:])
			return evaluate(newValue)
		}

		firstDigitIdx = digitIdx
	}

	return value
}

func split(value int) string {
	return fmt.Sprintf("[%d,%d]", value/2, value/2+value%2)
}

func getMagnitude(value string) int {
	var stack []int

	for _, c := range value {
		switch c {
		case '[', ',':
			break

		case ']':
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			result := left*3 + right*2
			stack = append(stack[:len(stack)-2], result)

		default:
			stack = append(stack, int(c)-'0')
		}
	}

	if len(stack) != 1 {
		panic(fmt.Sprintf("Stack count was %d", len(stack)))
	}

	return stack[0]
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
