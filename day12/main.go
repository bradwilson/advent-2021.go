package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var data [][]string
var nodes map[string][]string

func preCompute() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), "-"))
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	nodes = make(map[string][]string)

	for _, line := range data {
		list := nodes[line[0]]
		list = append(list, line[1])
		nodes[line[0]] = list

		list = nodes[line[1]]
		list = append(list, line[0])
		nodes[line[1]] = list
	}
}

func getPart1() int {
	return len(findAllPaths(false))
}

func getPart2() int {
	return len(findAllPaths(true))
}

func findAllPaths(canHaveDoubleSmallCaveVisit bool) [][]string {
	return findAllSubPaths(canHaveDoubleSmallCaveVisit, []string(nil), "start")
}

func findAllSubPaths(canHaveDoubleSmallCaveVisit bool, path []string, node string) [][]string {
	if strings.ToLower(node) == node {
		for _, existingNode := range path {
			if existingNode == node {
				if canHaveDoubleSmallCaveVisit && node != "start" && node != "end" {
					canHaveDoubleSmallCaveVisit = false
					break
				}
				return nil
			}
		}
	}

	subPath := append(path, node)

	var result [][]string

	if node == "end" {
		return append(result, subPath)
	}

	for _, neighbor := range nodes[node] {
		result = append(result, findAllSubPaths(canHaveDoubleSmallCaveVisit, subPath, neighbor)...)
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
