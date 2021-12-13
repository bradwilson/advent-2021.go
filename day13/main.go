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

type instruction struct {
	direction string
	line      int
}
type point struct {
	x int
	y int
}

var instructions []instruction = make([]instruction, 0, 12)
var grid [][]bool
var maxX int
var maxY int

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := make([]point, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "fold along ") {
			pieces := strings.Split(line[11:], "=")
			line, _ := strconv.Atoi(pieces[1])
			instructions = append(instructions, instruction{pieces[0], line})
		} else if line != "" {
			pieces := strings.Split(line, ",")
			x, _ := strconv.Atoi(pieces[0])
			y, _ := strconv.Atoi(pieces[1])
			points = append(points, point{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	grid = make([][]bool, maxX+1)
	for idx := range grid {
		grid[idx] = make([]bool, maxY+1)
	}

	for _, point := range points {
		grid[point.x][point.y] = true
	}
}

func getPart1() int {
	fold(instructions[0])

	result := 0

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[x][y] {
				result++
			}
		}
	}

	return result
}

func getPart2() {
	for idx := 1; idx < len(instructions); idx++ {
		fold(instructions[idx])
	}

	fmt.Println()

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func fold(i instruction) {
	if i.direction == "x" {
		foldAlongX(i.line)
	} else {
		foldAlongY(i.line)
	}
}

func foldAlongX(line int) {
	offset := (line * 2) - maxX

	for x := line + 1; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if grid[x][y] {
				grid[maxX-x+offset][y] = true
			}
		}
	}

	maxX = line - 1
}

func foldAlongY(line int) {
	offset := (line * 2) - maxY

	for y := line + 1; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[x][y] {
				grid[x][maxY-y+offset] = true
			}
		}
	}

	maxY = line - 1
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
	getPart2()
	fmt.Printf("[%s] Part 2\n", time.Since(start))
}
