package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/evenlwanvik/adventofcode/internal/data"
)

const (
	FIVE_OF_A_KIND  = 7 * 10e13
	FOUR_OF_A_KIND  = 6 * 10e13
	FULL_HOUSE      = 5 * 10e13
	THREE_OF_A_KIND = 4 * 10e13
	TWO_PAIR        = 3 * 10e13
	ONE_PAIR        = 2 * 10e13
	HIGH_CARD       = 1 * 10e13
)

func main() {
	data, err := data.ReadFile("data/day7.txt")
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
		hand.CreateHand(lines[i], false)
		hand.getScore()
		hand.getSecondScore()
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
		//fmt.Printf("\nRank %d - %d - bid %d =  %d", i+1, score.Score, score.Bid, result)
	}
	return total
}

func part2(data string) int {
	lines := strings.Split(data, "\n")
	nLines := len(lines)

	hands := make([]Hand, nLines)
	scores := make([]Scores, nLines)

	for i, hand := range hands {
		hand.CreateHand(lines[i], true)
		hand.getScore()
		hand.Score = hand.getWildcardScore()
		hand.getSecondScore()
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
	}

	return total
}

type Scores struct {
	Score float64
	Bid   int
}

type Hand struct {
	Rank         int
	Score        float64
	Cards        []rune
	Bid          int
	CardCounts   map[rune]int
	CardStrength map[rune]float64
}

func (h *Hand) CreateHand(line string, jokerIsPresent bool) {
	h.CardStrength = map[rune]float64{
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
	}

	if jokerIsPresent {
		h.CardStrength['J'] = 1 // Don't need to set the rest, their rank wil hold
	}

	lineSplit := strings.Split(line, " ")

	for _, card := range []rune(lineSplit[0]) {
		h.Cards = append(h.Cards, card)
	}

	counts := make(map[rune]int)
	for _, ru := range h.Cards {
		counts[ru]++
	}
	h.CardCounts = counts

	h.Bid, _ = strconv.Atoi(lineSplit[1])
}

func (h *Hand) getScore() {
	threeALike := false
	twoALike := false

	if len(h.CardCounts) == 1 {
		h.Score = FIVE_OF_A_KIND
		return
	}

	for _, cardCount := range h.CardCounts {

		if cardCount == 4 {
			h.Score = FOUR_OF_A_KIND
			return
		}
		if cardCount == 3 {
			threeALike = true
		}
		if cardCount == 2 && twoALike {
			h.Score = TWO_PAIR
			return
		}
		if cardCount == 2 {
			twoALike = true
		}
		if threeALike && twoALike {
			h.Score = FULL_HOUSE
			return
		}
	}
	if threeALike {
		h.Score = THREE_OF_A_KIND
		return
	}
	if twoALike {
		h.Score = ONE_PAIR
		return
	}

	// High card
	h.Score = HIGH_CARD
}

func (h *Hand) getSecondScore() {
	for i := 4; i >= 0; i-- {
		coeff := math.Pow(10, float64(10-i*2))

		h.Score += coeff * h.CardStrength[h.Cards[i]]
	}
}

func (h *Hand) getWildcardScore() float64 {

	cardMap := map[rune]int{}

	// Create map of numbers of cards (runes)
	for _, card := range h.Cards {
		cardMap[card] += 1
	}

	if cardMap['J'] >= 4 {
		return FIVE_OF_A_KIND
	}

	if cardMap['J'] == 3 {
		if len(cardMap) == 2 {
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	}

	if cardMap['J'] == 2 {
		if h.Score == TWO_PAIR {
			return FOUR_OF_A_KIND
		}
		if h.Score == ONE_PAIR {

			return THREE_OF_A_KIND
		}
		if h.Score == FULL_HOUSE {
			return FIVE_OF_A_KIND
		}
	}

	if cardMap['J'] == 1 {
		if h.Score == THREE_OF_A_KIND {
			return FOUR_OF_A_KIND
		}
		if h.Score == HIGH_CARD {
			return ONE_PAIR
		}
		if h.Score == TWO_PAIR {
			return FULL_HOUSE
		}
		if h.Score == ONE_PAIR {
			return THREE_OF_A_KIND
		}
		if h.Score == FOUR_OF_A_KIND {
			return FIVE_OF_A_KIND
		}

	}
	return h.Score
}

func findNewScore(prevScore int, newScore int) int {
	diff := newScore - prevScore
	return prevScore + diff
}
