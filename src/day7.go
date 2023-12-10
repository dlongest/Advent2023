package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day7RunA() {
	fmt.Println("Running Day 7 Part A")

	hands, err := LoadCamelHandsFromFile("../Data/day7.txt", DefaultCamelCardFactory{})

	if err != nil {
		fmt.Println("Error loading file: " + err.Error())
		return
	}

	score := ScoreUnordered(hands, DefaultCamelHandTyper{})

	fmt.Printf("Score = %v\n", score)
}

func Day7RunB() {
	fmt.Println("Running Day 7 Part B")

	hands, err := LoadCamelHandsFromFile("../Data/day7.txt", JokerAwareCamelCardFactory{})

	if err != nil {
		fmt.Println("Error loading file: " + err.Error())
		return
	}

	score := ScoreUnordered(hands, JokerAwareCamelHandTyper{})

	fmt.Printf("Score = %v\n", score)
}

type CamelCard struct {
	value int
	label string
}

func NewDefaultRulesCamelCard(card string) (CamelCard, error) {

	cards := map[string]CamelCard{}

	cards["2"] = CamelCard{value: 2, label: "2"}
	cards["3"] = CamelCard{value: 3, label: "3"}
	cards["4"] = CamelCard{value: 4, label: "4"}
	cards["5"] = CamelCard{value: 5, label: "5"}
	cards["6"] = CamelCard{value: 6, label: "6"}
	cards["7"] = CamelCard{value: 7, label: "7"}
	cards["8"] = CamelCard{value: 8, label: "8"}
	cards["9"] = CamelCard{value: 9, label: "9"}
	cards["T"] = CamelCard{value: 10, label: "T"}
	cards["J"] = CamelCard{value: 11, label: "J"}
	cards["Q"] = CamelCard{value: 12, label: "Q"}
	cards["K"] = CamelCard{value: 13, label: "K"}
	cards["A"] = CamelCard{value: 14, label: "A"}

	c, found := cards[card]

	if !found {
		return CamelCard{}, errors.New("Card not found")
	}

	return c, nil
}

func NewJokerAwareCamelCard(card string) (CamelCard, error) {

	cards := map[string]CamelCard{}

	cards["J"] = CamelCard{value: 1, label: "J"}
	cards["2"] = CamelCard{value: 2, label: "2"}
	cards["3"] = CamelCard{value: 3, label: "3"}
	cards["4"] = CamelCard{value: 4, label: "4"}
	cards["5"] = CamelCard{value: 5, label: "5"}
	cards["6"] = CamelCard{value: 6, label: "6"}
	cards["7"] = CamelCard{value: 7, label: "7"}
	cards["8"] = CamelCard{value: 8, label: "8"}
	cards["9"] = CamelCard{value: 9, label: "9"}
	cards["T"] = CamelCard{value: 10, label: "T"}
	cards["Q"] = CamelCard{value: 12, label: "Q"}
	cards["K"] = CamelCard{value: 13, label: "K"}
	cards["A"] = CamelCard{value: 14, label: "A"}

	c, found := cards[card]

	if !found {
		return CamelCard{}, errors.New("Card not found")
	}

	return c, nil
}

type CamelHand struct {
	cardsInOrder []CamelCard
	bid          int
	originalHand string
}

type CamelCardFactory interface {
	Create(cardLabel string) (CamelCard, error)
}

type DefaultCamelCardFactory struct {
}

type JokerAwareCamelCardFactory struct {
}

func (f DefaultCamelCardFactory) Create(cardLabel string) (CamelCard, error) {
	return NewDefaultRulesCamelCard(cardLabel)
}

func (f JokerAwareCamelCardFactory) Create(cardLabel string) (CamelCard, error) {
	return NewJokerAwareCamelCard(cardLabel)
}

func NewCamelCardHand(cards string, bid int, camelCardFactory CamelCardFactory) (CamelHand, error) {

	cardsInOrder := []CamelCard{}
	for _, c := range cards {

		card, err := camelCardFactory.Create(string(c))

		if err != nil {
			return CamelHand{}, err
		}

		cardsInOrder = append(cardsInOrder, card)
	}

	return CamelHand{cardsInOrder: cardsInOrder, bid: bid, originalHand: cards}, nil
}

func LoadCamelHandsFromFile(filePath string, camelCardFactory CamelCardFactory) ([]CamelHand, error) {

	lines := ReadLines(filePath)

	hands := []CamelHand{}

	for _, line := range lines {

		parts := strings.Split(line, " ")

		bid, err := strconv.Atoi(parts[1])

		if err != nil {
			return nil, err
		}

		hand, err := NewCamelCardHand(parts[0], bid, camelCardFactory)

		if err != nil {
			return nil, err
		}

		hands = append(hands, hand)
	}

	return hands, nil
}

func (h CamelHand) String() string {

	return h.originalHand
}

type CamelHandType int

const (
	HIGH_CARD     = iota // A2367
	ONE_PAIR             // AA246
	TWO_PAIR             // AA773
	THREE_OF_KIND        // AAA79
	FULL_HOUSE           // AAA33
	FOUR_OF_KIND         // AAAA7
	FIVE_OF_KIND         // AAAAA
)

type CamelHandTyper interface {
	Type(CamelHand) (CamelHandType, error)
}

type DefaultCamelHandTyper struct {
}

func (d DefaultCamelHandTyper) Type(hand CamelHand) (CamelHandType, error) {

	m := map[CamelCard]int{}

	for _, c := range hand.cardsInOrder {
		m[c]++
	}

	if len(m) < 1 || len(m) > 5 {
		return HIGH_CARD, errors.New("Hand seems to have an irregular number of cards: " + hand.String())
	}

	// If dict only has 1 entry, all cards must be the same type ergo 5 of a kind
	if len(m) == 1 {
		return FIVE_OF_KIND, nil
	}

	// If dict has 5 values, all cards must be different ergo High Card
	if len(m) == 5 {
		return HIGH_CARD, nil
	}

	// This case handles full house (e.g. AAA77) or four of a kind (AAAA4)
	if len(m) == 2 {

		keys := Keys(m)

		if len(keys) > 2 {
			panic("No way we should have gotten here")
		}

		if m[keys[0]] == 3 || m[keys[1]] == 3 {
			return FULL_HOUSE, nil
		}

		if m[keys[0]] == 4 || m[keys[1]] == 4 {
			return FOUR_OF_KIND, nil
		}
	}

	// If the map contains 3 keys, then we can have either 2 pair or 3 of a kind. Ex: XXYYZ or XXXYZ.
	if len(m) == 3 {

		keys := Keys(m)

		if m[keys[0]] == 3 || m[keys[1]] == 3 || m[keys[2]] == 3 {
			return THREE_OF_KIND, nil
		}

		return TWO_PAIR, nil
	}

	// If the map contains 4 keys, we can only have a one pair hand Ex: AA456.
	if len(m) == 4 {
		return ONE_PAIR, nil
	}

	// At this point we've run through every possible combination of lengths so this line should never be hit.
	// It's here to satisfy the compiler and would be an error.
	return HIGH_CARD, errors.New("No idea what this hand is: " + hand.String())
}

type JokerAwareCamelHandTyper struct{}

func (t JokerAwareCamelHandTyper) Type(hand CamelHand) (CamelHandType, error) {

	joker, err := NewJokerAwareCamelCard("J")

	if err != nil {
		return HIGH_CARD, nil
	}

	m := map[CamelCard]int{}

	for _, c := range hand.cardsInOrder {
		m[c]++
	}

	if len(m) < 1 || len(m) > 5 {
		return HIGH_CARD, errors.New("Hand seems to have an irregular numer of cards: " + hand.String())
	}

	howManyJokers := m[joker]

	// If we don't have any Jokers, hand rules are the default rules.
	if howManyJokers == 0 {
		return DefaultCamelHandTyper{}.Type(hand)
	}

	// 4 or 5 Jokers == 5 of a kind
	if howManyJokers == 4 || howManyJokers == 5 {
		return FIVE_OF_KIND, nil
	}

	// Given a hand XJJJY, then 3 Jokers becomes 4 of a kind
	// Given a hand XXJJJ, then 3 Jokers becomes 5 of a kind.
	// What matters is how many entries are in the map: do we have 3 unique cards (including Joker) or just 2?
	// If we have 3 unique cards i.e. XJJJY, then we've got a 4 of a kind. Otherwise, we're turning a full house into a 5 of a kind
	// Examples:
	// * XXJJJ = 5 of a kind
	// * XYJJJ = 4 of a kind
	if howManyJokers == 3 && len(m) == 3 {
		return FOUR_OF_KIND, nil // This line may never have been tested?
	} else if howManyJokers == 3 && len(m) == 2 {
		return FIVE_OF_KIND, nil // This line may never have been tested?
	}

	// Given XYZJJ, we get 3 of a kind (3 different cards = high card plus 2 jokers)
	// Given XXZJJ, we get 4 of a kind (we have 1 pair plus a card plus 2 Jokers)
	// Given XXXJJ, we get 5 of a kind (we have 3 of a kind plus 2 Jokers)
	if howManyJokers == 2 {
		switch len(m) {
		case 2:
			return FIVE_OF_KIND, nil
		case 3:
			// JJXXY
			return FOUR_OF_KIND, nil
		case 4:
			// JJXYZ
			return THREE_OF_KIND, nil
		}
	}

	// howManyJokers = 1
	// Given XYZWJ, we get 1 pair from high card (len = 5)
	// Given XXYZJ, we get 3 of a kind from 1 pair (we could get 2 pair but 3 of a kind > 2 pair) (len = 4)
	// Given XXYYJ, we get full house from 2 pair (len = 3)
	// Given XXXYJ, we get 4 of a kind from 3 of a kind (len = 3)
	// Given XXXXJ, we get 5 of a kind (len = 2)
	switch len(m) {
	case 5:
		return ONE_PAIR, nil
	case 4:
		return THREE_OF_KIND, nil
	case 3:
		// 1 joker and 3 distinct cards
		// Options:
		//	* XXYYJ = Full House
		// 	* XXXYJ = Four of a Kind
		//

		d := DefaultCamelHandTyper{} // This line might be the problem

		dt, err := d.Type(hand)

		if err != nil {
			return HIGH_CARD, err
		}

		if dt == TWO_PAIR {
			return FULL_HOUSE, nil
		}

		if dt == THREE_OF_KIND {
			return FOUR_OF_KIND, nil
		}

		//return HIGH_CARD, errors.New("Trying to type a 1 Joker hand and not able to do so: " + hand.String())
	case 2: // 1 Joker and 2 distinct cards == JXXXX == 5 of a kind
		return FIVE_OF_KIND, nil
	}

	return HIGH_CARD, errors.New("Typer was not able to type hand " + hand.String())
}

func (h CamelHand) Beats(other CamelHand, typer CamelHandTyper) (bool, error) {

	hType, err := typer.Type(h)

	if err != nil {
		return false, err
	}

	otherType, err := typer.Type(other)

	if err != nil {
		return false, err
	}

	// If the hands are of different types, high type wins
	if hType != otherType {
		return hType > otherType, nil
	}

	// When hands are of the same type, the winner is based on strength of each card, in order
	for i := range h.cardsInOrder {
		if h.cardsInOrder[i].value < other.cardsInOrder[i].value {
			return false, nil
		}

		if h.cardsInOrder[i].value > other.cardsInOrder[i].value {
			return true, nil
		}
	}

	// At this point the hands are the same, and we have no current mechanism for breaking ties so panic
	panic("Unable to determine a winner between these two hands: " + h.String() + " and " + other.String())
}

func Order(hands []CamelHand, typer CamelHandTyper) []CamelHand {

	sort.Slice(hands, func(i, j int) bool {
		firstWins, err := hands[i].Beats(hands[j], typer)

		if err != nil {
			panic("Cannot determine hand winner: " + hands[i].String() + " vs. " + hands[j].String())
		}

		return !firstWins
	})

	return hands
}

func ScoreOrdered(hands []CamelHand) int {

	totalScore := 0

	for i := range hands {

		rank := i + 1

		score := rank * hands[i].bid

		totalScore += score
	}

	return totalScore
}

func ScoreUnordered(hands []CamelHand, typer CamelHandTyper) int {

	h := make([]CamelHand, len(hands))

	copy(h, hands)

	h = Order(h, typer)

	return ScoreOrdered(h)
}
