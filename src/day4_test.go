package main

import (
	"reflect"
	"testing"
)

func TestCardParser(t *testing.T) {
	p := NewCardParser()

	card, err := p.Parse("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	if err != nil {
		t.Error(err)
	}

	if card.id != 1 {
		t.Error("Incorrect Card ID")
	}

	if !reflect.DeepEqual(card.winningNumbers, []int{41, 48, 83, 86, 17}) {
		t.Error("Incorrect winning numbers on card")
	}

	if !reflect.DeepEqual(card.cardNumbers, []int{83, 86, 6, 31, 17, 9, 48, 53}) {
		t.Error("Incorrect card numbers on card")
	}
}

type CardTestCase struct {
	card          Card
	expectedScore int
}

func TestCardScore(t *testing.T) {
	cases := []CardTestCase{
		{card: Card{id: 1,
			winningNumbers: []int{41, 48, 83, 86, 17}, cardNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}, expectedScore: 8},
		{card: Card{id: 2,
			winningNumbers: []int{13, 32, 20, 16, 61}, cardNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19}}, expectedScore: 2},
		{card: Card{id: 3,
			winningNumbers: []int{1, 21, 53, 59, 44}, cardNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1}}, expectedScore: 2},
		{card: Card{id: 4,
			winningNumbers: []int{41, 92, 73, 84, 69}, cardNumbers: []int{59, 84, 76, 51, 58, 5, 54, 83}}, expectedScore: 1},
		{card: Card{id: 5,
			winningNumbers: []int{87, 83, 26, 28, 32}, cardNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36}}, expectedScore: 0},
		{card: Card{id: 6,
			winningNumbers: []int{31, 18, 13, 56, 72}, cardNumbers: []int{74, 77, 10, 23, 35, 67, 36, 11}}, expectedScore: 0}}

	for _, testCase := range cases {
		if testCase.card.Score() != testCase.expectedScore {
			t.Error("Incorrect score on Card ", testCase.card.id)
		}
	}
}

type ScoreTests struct {
	card          string
	expectedScore int
}

func TestScore(t *testing.T) {
	cards := []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 31:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", "Card 216: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}
	expectedScore := 13

	actualScore, err := Score(cards)

	if err != nil {
		t.Error(err)
	}

	if expectedScore != actualScore {
		t.Errorf("Expected score %v does not match actual score %v", expectedScore, actualScore)
	}
}

func TestWon(t *testing.T) {

	card := Card{id: 1,
		winningNumbers: []int{41, 48, 83, 86, 17}, cardNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}

	cardsWon := card.Win()

	if len(cardsWon) != 4 {
		t.Error("Incorrect number of cards won")
	}

	if !reflect.DeepEqual(cardsWon, []int{2, 3, 4, 5}) {
		t.Error("Didn't win the correct cards", cardsWon)
	}
}

func TestCardPileProcess(t *testing.T) {
	cards := []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}
	expected := 30

	sut, err := NewCardPile(cards)

	if err != nil {
		t.Error(err)
		return
	}

	cardsWon, err := sut.Process()

	if err != nil {
		t.Error(err)
		return
	}

	if cardsWon != expected {
		t.Error("Didn't win the right amout of cards")
	}
}
