package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

type Node struct {
	X, Y, RiskLevel, MinRiskToStart int
	PathBack                        *Node
	Visited                         bool
	Connections                     []*Node
}

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
	nodes := parseNodes()

	return getLowestRisk(nodes)
}

func getPart2() int {
	nodes := parseNodes()
	maxX := len(nodes)
	maxY := len(nodes[0])

	for idx := 1; idx < 5; idx++ {
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				newRiskLevel := nodes[x][y].RiskLevel + idx
				for ; newRiskLevel > 9; newRiskLevel -= 9 {
				}
				nodes[x] = append(nodes[x], newNode(newRiskLevel))
			}
		}
	}

	maxX = len(nodes)
	maxY = len(nodes[0])
	currentX := 0

	for idx := 1; idx < 5; idx++ {
		for x := 0; x < maxX; x++ {
			var row []*Node

			for y := 0; y < maxY; y++ {
				newRiskLevel := nodes[currentX][y].RiskLevel + 1
				for ; newRiskLevel > 9; newRiskLevel -= 9 {
				}
				row = append(row, newNode(newRiskLevel))
			}

			nodes = append(nodes, row)
			currentX++
		}
	}

	return getLowestRisk(nodes)
}

func parseNodes() [][]*Node {
	var result [][]*Node

	for _, line := range data {
		var row []*Node

		for _, c := range line {
			riskLevel := int(c) - int('0')
			row = append(row, newNode(riskLevel))
		}

		result = append(result, row)
	}

	return result
}

func newNode(riskLevel int) *Node {
	result := Node{-1, -1, riskLevel, math.MaxInt, nil, false, nil}
	return &result
}

func printNodes(nodes [][]*Node) {
	for idxRow, row := range nodes {
		if idxRow > 0 && idxRow%10 == 0 {
			fmt.Println()
		}

		for idxNode, node := range row {
			if idxNode > 0 && idxNode%10 == 0 {
				fmt.Print(" ")
			}

			fmt.Print(node.RiskLevel)
		}

		fmt.Println()
	}
}

func getLowestRisk(nodes [][]*Node) int {
	maxX := len(nodes)
	maxY := len(nodes[0])

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			node := nodes[x][y]
			node.X = x
			node.Y = y

			if y > 0 {
				node.Connections = append(node.Connections, nodes[x][y-1])
			}
			if y < maxY-1 {
				node.Connections = append(node.Connections, nodes[x][y+1])
			}
			if x > 0 {
				node.Connections = append(node.Connections, nodes[x-1][y])
			}
			if x < maxX-1 {
				node.Connections = append(node.Connections, nodes[x+1][y])
			}

			sort.Slice(node.Connections, func(i, j int) bool {
				return node.Connections[i].RiskLevel < node.Connections[j].RiskLevel
			})
		}
	}

	start := nodes[0][0]
	end := nodes[maxX-1][maxY-1]

	mapPaths(nodes, start, end)

	shortestPath := []*Node{end}
	buildShortestPath(&shortestPath, end)

	result := -start.RiskLevel
	for _, nodeOnPath := range shortestPath {
		result += nodeOnPath.RiskLevel
	}

	return result
}

func buildShortestPath(path *[]*Node, node *Node) {
	if node.PathBack == nil {
		return
	}

	*path = append(*path, node.PathBack)
	buildShortestPath(path, node.PathBack)
}

func mapPaths(nodes [][]*Node, start *Node, end *Node) {
	start.MinRiskToStart = 0
	queue := []*Node{start}

	for {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].MinRiskToStart < queue[j].MinRiskToStart
		})
		node := queue[0]
		queue = queue[1:]

		for _, childNode := range node.Connections {
			if childNode.Visited {
				continue
			}

			if node.MinRiskToStart+childNode.RiskLevel < childNode.MinRiskToStart {
				childNode.MinRiskToStart = node.MinRiskToStart + childNode.RiskLevel
				childNode.PathBack = node

				if !contains(queue, childNode) {
					queue = append(queue, childNode)
				}
			}
		}

		node.Visited = true

		if node == end || len(queue) == 0 {
			return
		}
	}
}

func contains(queue []*Node, node *Node) bool {
	for _, n := range queue {
		if n == node {
			return true
		}
	}
	return false
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
