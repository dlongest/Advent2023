package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day4RunA() {
	fmt.Println("Running Day 4 Part A")

	cards := ReadLines("../Data/day4.txt")

	score, err := Score(cards)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Score = ", score)
}

func Day4RunB() {
	fmt.Println("Running Day 4 Part B")

	cards := ReadLines("../Data/day4.txt")

	pile, err := NewCardPile(cards)

	if err != nil {
		fmt.Println(err)
		return
	}

	cardsWon, err := pile.Process()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cards won = ", cardsWon)
}

type CardParser struct {
	number *regexp.Regexp
}

func NewCardParser() *CardParser {
	return &CardParser{number: regexp.MustCompile("[0-9]+")}
}

type Card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func Score(cards []string) (int, error) {
	p := NewCardParser()
	score := 0

	for _, card := range cards {
		c, err := p.Parse(card)

		if err != nil {
			return -1, err
		}

		score += c.Score()
	}

	return score, nil
}

func (c Card) HowManyMatches() int {
	winningNumberPresentCount := 0

	for _, winningNumber := range c.winningNumbers {
		if slices.Contains(c.cardNumbers, winningNumber) {
			winningNumberPresentCount++
		}
	}

	return winningNumberPresentCount
}

func (c Card) Win() []int {

	matchCount := c.HowManyMatches()

	if matchCount == 0 {
		return []int{}
	}

	cardsWonIds := []int{}

	for id := c.id + 1; id <= c.id+matchCount; id++ {
		cardsWonIds = append(cardsWonIds, id)
	}

	return cardsWonIds
}

func (p *CardParser) Parse(s string) (Card, error) {

	firstSplit := strings.Split(s, ": ")

	secondSplit := strings.Split(firstSplit[1], " | ")

	cardID, err := strconv.Atoi(p.number.FindString(firstSplit[0]))

	if err != nil {
		return Card{}, err
	}

	winningNumbers, err := ToIntSlice(p.number.FindAllString(secondSplit[0], -1))

	if err != nil {
		return Card{}, err
	}

	cardNumbers, err := ToIntSlice(p.number.FindAllString(secondSplit[1], -1))

	if err != nil {
		return Card{}, err
	}

	return Card{id: cardID, winningNumbers: winningNumbers, cardNumbers: cardNumbers}, nil
}

func (c Card) Score() int {

	matchingNumbers := c.HowManyMatches()

	return int(math.Pow(float64(2), float64(matchingNumbers-1)))
}

func ToIntSlice(numbers []string) ([]int, error) {
	n := []int{}

	for _, number := range numbers {

		num, err := strconv.Atoi(number)

		if err != nil {
			return nil, err
		}

		n = append(n, num)
	}

	return n, nil
}

type CardPile struct {
	pile map[int]Card
}

func (pile *CardPile) Card(id int) Card {
	return pile.pile[id]
}

func (pile *CardPile) Add(card Card) {
	pile.pile[card.id] = card
}

func NewCardPile(cards []string) (*CardPile, error) {
	pile := &CardPile{pile: map[int]Card{}}

	p := NewCardParser()

	for _, card := range cards {
		c, err := p.Parse(card)

		if err != nil {
			return &CardPile{}, err
		}

		pile.pile[c.id] = c
	}

	return pile, nil
}

type CardID int

func (pile *CardPile) Process() (int, error) {

	pileSize := len(pile.pile)

	copiesWon := map[CardID]int{} // Key: Card ID of the card that was won. Value: How many copies of it have been won.

	totalCardsWon := 0

	// Our card IDs start at 1 and go up monotonically.
	for id := 1; id <= pileSize; id++ {

		currentCard := pile.pile[id]

		// How many copies of the current card do we have? Returns 0 if we have none
		currentCopies := copiesWon[CardID(id)]

		// This is how many total versions of the current card we have: 1 for the original + however many copies
		currentVersions := currentCopies + 1

		// What cards IDs are won by the current card?
		idsWon := currentCard.Win()

		// For each card won by the current card, add it to the copiesWon map.
		// However, we're not just adding 1 copy: we're adding 1 copy for every version of the current card we've got.
		for _, wonId := range idsWon {
			// We don't need to check for existence of a key because it will be 0 if not present. So just adding
			// currentVersions to whatever's there will give us the right quantity
			copiesWon[CardID(wonId)] += currentVersions
		}

		// This is the last step of processing for the current card. We know how many versions of it we have.
		// We simply add that to our running total of all cards
		totalCardsWon += currentVersions
	}

	return totalCardsWon, nil
}
