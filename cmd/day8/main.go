package main

import (
	"fmt"
	"github.com/evenlwanvik/adventofcode/internal/math"
	"log"
	"strings"
	"sync"

	"github.com/evenlwanvik/adventofcode/internal/graph"
	"github.com/evenlwanvik/adventofcode/internal/utils"
)

func main() {
	data, err := utils.ReadFile("data/day8_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(data)
	fmt.Println("Part 1 answer: ", result)

	result = part2(data)
	fmt.Println("Part 2 answer: ", result)

}

func Part1FinalValCheck(currentNodeVal string) bool {
	return currentNodeVal == "ZZZ"
}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	instructions := []rune(lines[0])
	//printInstructions(instructions)

	nodes := createNodesMap(lines[2:])

	g := graph.NewGraph()
	g.CreateNodeMap(nodes)
	startNode := g.Vertices["AAA"]

	result := g.NumOfInstructions(
		instructions,
		startNode,
		Part1FinalValCheck,
	)

	return result
}

func Part2FinalValCheck(currentNodeVal string) bool {
	return currentNodeVal[len(currentNodeVal)-1] == 'Z'
}

func part2(data string) int {
	lines := strings.Split(data, "\n")

	instructions := []rune(lines[0])

	nodes := createNodesMap(lines[2:])

	g := graph.NewGraph()
	g.CreateNodeMap(nodes)

	// Create a slice of all start nodes
	var startNodeKeys []string
	for k, _ := range nodes {
		if k[len(k)-1] == 'A' {
			startNodeKeys = append(startNodeKeys, k)
		}
	}

	counters := make([]int, len(startNodeKeys))

	var wg sync.WaitGroup

	for i, startNodeKey := range startNodeKeys {
		wg.Add(1)
		go func(
			instructions []rune,
			graph *graph.Graph,
			startNodeKey string,
			counter *int,
		) {
			defer wg.Done()
			startNode := g.Vertices[startNodeKey]
			*counter = graph.NumOfInstructions(
				instructions,
				startNode,
				Part2FinalValCheck,
			)
		}(instructions, g, startNodeKey, &counters[i])
	}
	wg.Wait()

	// After looking at reddit I found tips to use LCM (Least Common Multiple)
	return math.LCM(counters)
}

func cleanNodeLine(line string) (string, []string) {
	parts := strings.Split(line, " = ")
	v := parts[0]
	es := strings.Replace(parts[1], "(", "", -1)
	es = strings.Replace(es, ")", "", -1)
	e := strings.Split(es, ", ")
	return v, e
}

func createNodesMap(lines []string) map[string][]string {
	nodes := make(map[string][]string)

	for _, line := range lines {
		v, e := cleanNodeLine(line)
		nodes[v] = e
	}

	return nodes
}

func printInstructions(i []rune) {
	fmt.Println("Instructions: ")
	for _, v := range i {
		fmt.Printf("%c", v)
	}
	fmt.Println("\n")
}
