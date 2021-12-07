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

type lineData struct {
	direction string
	distance  int
}

var data []lineData

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction := line[0]
		distance, _ := strconv.Atoi(line[1])
		data = append(data, lineData{direction, distance})
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}
}

func getPart1() int {
	horizontal := 0
	depth := 0

	for _, instruction := range data {
		switch instruction.direction {
		case "down":
			depth += instruction.distance
		case "up":
			depth -= instruction.distance
		default:
			horizontal += instruction.distance
		}
	}

	return horizontal * depth
}

func getPart2() int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, instruction := range data {
		switch instruction.direction {
		case "down":
			aim += instruction.distance
		case "up":
			aim -= instruction.distance
		default:
			horizontal += instruction.distance
			depth += instruction.distance * aim
		}
	}

	return horizontal * depth
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
