package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/evenlwanvik/adventofcode/internal/graph"
	"github.com/evenlwanvik/adventofcode/internal/utils"
)

func main() {
	data, err := utils.ReadFile("data/day8_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(data)
	fmt.Println("\nPart 1 answer: ", result)

	//result = part2(data)
	//fmt.Println("\nPart 2 answer: ", result)

}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	instructions := []rune(lines[0])
	//printInstructions(instructions)

	nodes := createNodesMap(lines[2:])

	g := graph.NewGraph()
	g.CreateNodeMap(nodes)
	//g.PrintVertices()

	return g.NumOfInstructions(instructions)
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
