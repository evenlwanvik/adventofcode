package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(data)
	fmt.Println("\nPart 1 answer: ", result)

	result = part2(data)
	fmt.Println("\nPart 2 answer: ", result)

}

func part1(data string) int {
	lines := strings.Split(data, "\n")
	nLines := len(lines)

	//fmt.Printf("\nCard ranks: %s", string(game.CardRanks))

	hands := make([]Hand, nLines)
	scores := make([]Scores, nLines)

	for i, hand := range hands {
		hand.CreateHand(lines[i])
		hand.getScore()

		scores[i] = Scores{
			Score: hand.Score,
			Bid:   hand.Bid,
		}

	}

	// Sort scores
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score < scores[j].Score
	})

	total := 0
	for i, score := range scores {
		rank := i + 1
		result := rank * score.Bid
		total += result
		fmt.Printf("\nRank %d - %d - bid %d =  %d", i+1, score.Score, score.Bid, result)
	}

	return total
}

func part2(data string) int {
	return 2
}

type Scores struct {
	Score int
	Bid   int
}

type Hand struct {
	Rank       int
	Score      int
	Cards      []int
	Bid        int
	CardCounts map[int]int
}

func (h *Hand) CreateHand(line string) {
	cardRanks := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'1': 1,
	}

	lineSplit := strings.Split(line, " ")

	for _, card := range []rune(lineSplit[0]) {
		h.Cards = append(h.Cards, cardRanks[card])
	}

	counts := make(map[int]int)
	for _, num := range h.Cards {
		counts[num]++
	}
	h.CardCounts = counts

	h.Bid, _ = strconv.Atoi(lineSplit[1])
}

func (h *Hand) getScore() {
	threeALike := false
	twoALike := false

	// 5 alike
	if len(h.CardCounts) == 1 {
		h.Score = 7 * 10e13 //7000000
		h.getSecondScore()
		return
	}

	for _, cardCount := range h.CardCounts {
		if cardCount == 4 {
			// 4 alike
			h.Score = 6 * 10e13
			h.getSecondScore()
			return
		}
		if cardCount == 3 {
			threeALike = true
		}
		if cardCount == 2 && twoALike {
			// Two pairs
			h.Score = 3 * 10e13
			h.getSecondScore()
			return
		}
		if cardCount == 2 {
			twoALike = true
		}
		if threeALike && twoALike {
			// Full house
			h.Score = 5 * 10e13
			h.getSecondScore()
			return
		}
	}
	if threeALike {
		// Three alike
		h.Score = 4 * 10e13
		h.getSecondScore()
		return
	}
	if twoALike {
		// One pair
		h.Score = 2 * 10e13
		h.getSecondScore()
		return
	}

	// High card
	h.Score = 1 * 10e13
	h.getSecondScore()
}

func (h *Hand) getSecondScore() {
	for i := 4; i >= 0; i-- {
		coeff := int(math.Pow(10, float64(10-i*2)))

		h.Score += coeff * h.Cards[i]
	}
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
