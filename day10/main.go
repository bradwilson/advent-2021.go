package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type part1Match struct {
	opening rune
	value   int
}

var data []string
var incompleteLines = make([]Stack, 0)

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

	incompleteLines = make([]Stack, 0)
}

func getPart1() int {
	scores := map[rune]part1Match{
		')': {'(', 3},
		']': {'[', 57},
		'}': {'{', 1197},
		'>': {'<', 25137},
	}
	score := 0

	for _, line := range data {
		var stack Stack
		incomplete := true

		for _, r := range line {
			runeScore := scores[r]

			if runeScore.value == 0 {
				stack.Push(r)
			} else {
				if value, popped := stack.Pop(); popped && value == runeScore.opening {
					continue
				}

				score += runeScore.value
				incomplete = false
				break
			}
		}

		if incomplete {
			incompleteLines = append(incompleteLines, stack)
		}
	}

	return score
}

func getPart2() int {
	scores := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	var lineScores []int

	for _, line := range incompleteLines {
		lineScore := 0

		for {
			r, popped := line.Pop()
			if !popped {
				break
			}
			lineScore = lineScore*5 + scores[r]
		}

		lineScores = append(lineScores, lineScore)
	}

	sort.Slice(lineScores, func(i, j int) bool {
		return lineScores[i] < lineScores[j]
	})

	return lineScores[len(lineScores)/2]
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

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}
