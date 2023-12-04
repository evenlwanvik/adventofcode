package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

	result = part2(data)
	fmt.Println("Part 2 answer: ", result)
}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	totalScore := 0

	for _, card := range lines {
		wn := strings.Split(card, "|")[1]
		cn := strings.Split(card, "|")[0]
		cardNumbers := strings.Split(cn, ":")[1]
		winningNumbers := getWinningNumbers(wn)
		cardScoreBit := 1
		for _, n := range strings.Fields(cardNumbers) {
			cardNumber, err := strconv.Atoi(n)
			if err != nil {
				continue
			}
			if slices.Contains(winningNumbers, cardNumber) {
				cardScoreBit <<= 1 // shift to left by 1 to double the value
			}
		}
		cardScoreBit >>= 1 // adjust for one too many flips
		totalScore += cardScoreBit
	}

	return totalScore
}

func part2(data string) int {

	lines := strings.Split(data, "\n")
	copiesOfCards := make([]int, len(lines))
	for idx, _ := range lines {
		copiesOfCards[idx] = 1
	}

	for i, card := range lines {
		wn := strings.Split(card, "|")[1]
		cn := strings.Split(card, "|")[0]
		cardNumbers := strings.Split(cn, ":")[1]
		winningNumbers := getWinningNumbers(wn)
		wonCopies := 0
		for _, n := range strings.Fields(cardNumbers) {
			cardNumber, err := strconv.Atoi(n)
			if err != nil {
				continue
			}
			if slices.Contains(winningNumbers, cardNumber) {
				wonCopies++ // shift to left by 1 to double the value
			}

		}

		countTil := i + wonCopies
		// If EOL stop at last line
		if countTil >= len(lines) {
			countTil = len(lines) - 1
		}

		// Add copies to next cards
		for j := i + 1; j <= countTil; j++ {
			copiesOfCards[j] += copiesOfCards[i]
		}
	}

	total := 0
	for _, c := range copiesOfCards {
		total += c
	}

	return total
}

func getWinningNumbers(line string) []int {
	var winningNumbers []int
	winningNumSlice := strings.Split(line, "|")[0]
	for _, n := range strings.Fields(winningNumSlice) {
		winningNumber, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		winningNumbers = append(winningNumbers, winningNumber)
	}
	return winningNumbers
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
