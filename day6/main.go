package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(data)
	fmt.Println("Part 1 answer: ", result)

}

func part1(data string) int {
	raceMs, recordDistances := preprocessData(data)

	fmt.Printf("Times %v\n", raceMs)
	fmt.Printf("Distances %v\n", recordDistances)

	MoE := 1
	for i := 0; i < len(raceMs); i++ {

		possibleWins := getPossibleWins(raceMs[i], recordDistances[i])
		println("Game ", i+1, " has ", MoE, " possible wins")
		MoE *= possibleWins
	}

	return MoE
}

func getPossibleWins(raceSeconds int, recordDistance int) int {

	possibleWins := 0

	for i := 0; i <= raceSeconds; i++ {
		speed := i * 1
		dist := speed * (raceSeconds - i)
		if dist > recordDistance {
			possibleWins++
		}
	}

	return possibleWins

}

func preprocessData(data string) ([]int, []int) {
	lines := strings.Split(data, "\n")

	raceMs := strings.Fields(strings.Split(lines[0], ":")[1])
	recordDistances := strings.Fields(strings.Split(lines[1], ":")[1])

	tm := make([]int, len(raceMs))
	dm := make([]int, len(recordDistances))

	for i := 0; i < len(raceMs); i++ {
		tm[i], _ = strconv.Atoi(raceMs[i])
		dm[i], _ = strconv.Atoi(recordDistances[i])
	}

	return tm, dm
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
