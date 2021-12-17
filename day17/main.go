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

type Range struct {
	minX, maxX, minY, maxY int
}

var hitCount = 0
var highestY = math.MinInt

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	pieces := strings.Split(scanner.Text()[13:], ", ")

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	rangeX := strings.Split(pieces[0][2:], "..")
	minX, _ := strconv.Atoi(rangeX[0])
	maxX, _ := strconv.Atoi(rangeX[1])
	rangeY := strings.Split(pieces[1][2:], "..")
	minY, _ := strconv.Atoi(rangeY[0])
	maxY, _ := strconv.Atoi(rangeY[1])
	data := Range{minX, maxX, minY, maxY}

	for startVelocityX := 0; startVelocityX <= data.maxX; startVelocityX++ {
		for startVelocityY := data.minY; startVelocityY <= -data.minY; startVelocityY++ {
			curX, curY := 0, 0
			velocityX, velocityY := startVelocityX, startVelocityY
			maxY := math.MinInt

			for curX < data.maxX && curY > data.minY {
				if curY > maxY {
					maxY = curY
				}

				curX += velocityX
				curY += velocityY

				if curX >= data.minX && curX <= data.maxX && curY >= data.minY && curY <= data.maxY {
					hitCount++

					if maxY > highestY {
						highestY = maxY
					}

					break
				}

				if velocityX > 0 {
					velocityX--
				}
				velocityY--
			}
		}
	}
}

func getPart1() int {
	return highestY
}

func getPart2() int {
	return hitCount
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
