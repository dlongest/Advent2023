package main

import (
	"testing"
)

func TestDefaultCamelHandTyper(t *testing.T) {

	sut := DefaultCamelHandTyper{}

	hands := []string{"32T3K", "T55J5", "KK677", "KTJJT", "QQQJJ", "74777", "TTTTT", "A4982"}
	expected := []CamelHandType{ONE_PAIR, THREE_OF_KIND, TWO_PAIR, TWO_PAIR, FULL_HOUSE, FOUR_OF_KIND, FIVE_OF_KIND, HIGH_CARD}

	for i := range hands {

		h, err := NewCamelCardHand(hands[i], 100, DefaultCamelCardFactory{})

		if err != nil {
			t.Error("Unable to create Camel Hand and should have been able to")
			return
		}

		actual, err := sut.Type(h)

		if err != nil {
			t.Error("Received error trying to type hand. Hand = " + hands[i] + " Error = " + err.Error())
			return
		}

		if expected[i] != actual {
			t.Errorf("Expected type %v :: Got type :: %v", expected[i], actual)
			return
		}
	}
}

type BeatsTestCase struct {
	first, second string
	firstWins     bool
}

func TestDefaultRulesCamelHandBeats(t *testing.T) {

	for _, tt := range []BeatsTestCase{
		{first: "AAA23", second: "AA778", firstWins: true},  // 3 of a kind vs. two pair
		{first: "AA778", second: "AAA23", firstWins: false}, // two pair vs. 3 of a kind
		{first: "2573A", second: "KQ432", firstWins: false}, //high card vs. high card
		{first: "7777K", second: "K7788", firstWins: true},  // 4 of a kind vs. two pair
		{first: "T55J5", second: "QQQJA", firstWins: false}, // 3 of a kind vs. 3 of a kind
		{first: "QQQJA", second: "T55J5", firstWins: true},  // 3 of a kind vs. 3 of a kind
		{first: "QQQJA", second: "T55J5", firstWins: true},  // first has tiebreaker
		{first: "32T3K", second: "KTJJT", firstWins: false}, // one pair < two pair
		{first: "32T3K", second: "KK677", firstWins: false}, // one pair < two pair
		{first: "3JT3K", second: "KK677", firstWins: false}, // one pair < 2 pair
		{first: "3JT3K", second: "33JTK", firstWins: true},  // same cards, but first wins because J > 3 in tiebreaker
		{first: "TTTTT", second: "JTTTT", firstWins: true},  // 5 of a kind > 4 of a kind
		{first: "4488J", second: "4488K", firstWins: false}, // second has two pair tiebreaker
		{first: "8886J", second: "999TT", firstWins: false}, // 3 of a kind < full house
		{first: "9JTJA", second: "9TJAJ", firstWins: true},  // one pair, first wins tiebreaker
		{first: "88823", second: "32888", firstWins: true},  // three of a kind, first wins tiebreaker
		{first: "J8823", second: "32888", firstWins: false}, // one pair < 3 of a kind
		{first: "TTTJ2", second: "AAATT", firstWins: false}, // 3 of a kind < full house
		{first: "JJJJJ", second: "2JJJJ", firstWins: true},  // 5 of a kind > 4 of a kind
		{first: "JJJJ8", second: "QQQQQ", firstWins: false}, // 4 of a kind < 5 of a kind
		{first: "JTJTT", second: "TTTTT", firstWins: false}, // full house < 5 of a kind
		{first: "JKKK2", second: "QQQQ2", firstWins: false}, // JKKK2 is weaker than QQQQ2 because J is weaker than Q.

		{first: "82828", second: "84848", firstWins: false}} { // full house vs. full house {

		fh, err := NewCamelCardHand(tt.first, 100, DefaultCamelCardFactory{})

		if err != nil {
			t.Error(err)
			return
		}

		sh, err := NewCamelCardHand(tt.second, 100, DefaultCamelCardFactory{})

		if err != nil {
			t.Error(err)
			return
		}

		actual, err := fh.Beats(sh, DefaultCamelHandTyper{})

		if err != nil {
			t.Error(err)
			return
		}

		if actual != tt.firstWins {
			t.Error("Expected first hand to win but it didn't: " + fh.String() + " vs. " + sh.String())
			return
		}
	}
}

func TestJokerAwareRulesCamelHandBeats(t *testing.T) {

	for _, tt := range []BeatsTestCase{
		{first: "AAA23", second: "AA778", firstWins: true},  // 3 of a kind vs. two pair
		{first: "AA778", second: "AAA23", firstWins: false}, // two pair vs. 3 of a kind
		{first: "2573A", second: "KQ432", firstWins: false}, //high card vs. high card
		{first: "7777K", second: "K7788", firstWins: true},  // 4 of a kind vs. two pair
		{first: "T55J5", second: "QQQJA", firstWins: false}, // 4 of a kind > 4 of a kind, 2nd one wins tie breaker
		{first: "QQQJA", second: "T55J5", firstWins: true},  // 3 of a kind < 4 of a kind
		{first: "32T3K", second: "KTJJT", firstWins: false}, // one pair < 4 of a kind
		{first: "32T3K", second: "KK677", firstWins: false}, // one pair < two pair
		{first: "3JT3K", second: "KK677", firstWins: true},  // three of a kind > 2 pair
		{first: "3JT3K", second: "33JTK", firstWins: false}, // same cards, but second wins because J < 3 in tiebreaker
		{first: "TTTTT", second: "JTTTT", firstWins: true},  // type and values, but T > J so first wins
		{first: "4488J", second: "4488K", firstWins: true},  // full house > 2 pair
		{first: "8886J", second: "999TT", firstWins: true},  // 4 of a kind > full house
		{first: "9JTJA", second: "9TJAJ", firstWins: false},
		{first: "88823", second: "32888", firstWins: true},
		{first: "J8823", second: "32888", firstWins: false},
		{first: "TTTJ2", second: "AAATT", firstWins: true}, // 4 of a kind > full house
		{first: "JJJJJ", second: "2JJJJ", firstWins: false},
		{first: "JJJJ8", second: "QQQQQ", firstWins: false},
		{first: "JTJTT", second: "TTTTT", firstWins: false},
		{first: "JKKK2", second: "QQQQ2", firstWins: false},   // JKKK2 is weaker than QQQQ2 because J is weaker than Q.
		{first: "82828", second: "84848", firstWins: false}} { // full house vs. full house {

		//32T3K", "KTJJT", "KK677", "T55J5", "QQQJA"
		fh, err := NewCamelCardHand(tt.first, 100, JokerAwareCamelCardFactory{})

		if err != nil {
			t.Error(err)
			return
		}

		sh, err := NewCamelCardHand(tt.second, 100, JokerAwareCamelCardFactory{})

		if err != nil {
			t.Error(err)
			return
		}

		actual, err := fh.Beats(sh, JokerAwareCamelHandTyper{})

		if err != nil {
			t.Error(err)
			return
		}

		if actual != tt.firstWins {
			t.Error("Expected first hand to win but it didn't: " + fh.String() + " vs. " + sh.String())
			return
		}
	}
}

func TestDefaultScoresScoreUnordered(t *testing.T) {
	expected := 6440

	input, err := LoadCamelHandsFromFile("../Data/day7-example.txt", DefaultCamelCardFactory{})

	if err != nil {
		t.Error(err)
		return
	}

	actual := ScoreUnordered(input, DefaultCamelHandTyper{})

	if actual != expected {
		t.Errorf("Expected score %v :: Actual score %v\n", expected, actual)
	}
}

// 32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
// KK677 is now the only two pair, making it the second-weakest hand.
// T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.

func TestJokerAwareCamelHandTyper(t *testing.T) {

	sut := JokerAwareCamelHandTyper{}

	hands := []string{"TTJTT", "T8243", "4JJ8J", "55JJJ", "32T3K", "T55J5", "KK677", "KTJJT", "QQQJA", "74777", "TTTTT",
		"A4982", "2J456", "J2452", "JJJJJ", "444AA", "K7KKK", "2JK22", "J2345", "T9876", "AATJT"}
	expected := []CamelHandType{FIVE_OF_KIND, HIGH_CARD, FOUR_OF_KIND, FIVE_OF_KIND, ONE_PAIR, FOUR_OF_KIND, TWO_PAIR, FOUR_OF_KIND, FOUR_OF_KIND, FOUR_OF_KIND,
		FIVE_OF_KIND, HIGH_CARD, ONE_PAIR, THREE_OF_KIND, FIVE_OF_KIND, FULL_HOUSE, FOUR_OF_KIND, FOUR_OF_KIND, ONE_PAIR, HIGH_CARD, FULL_HOUSE}

	for i := range hands {

		h, err := NewCamelCardHand(hands[i], 100, JokerAwareCamelCardFactory{})

		if err != nil {
			t.Error("Unable to create Camel Hand and should have been able to")
			return
		}

		actual, err := sut.Type(h)

		if err != nil {
			t.Error("Received error trying to type hand. Hand = " + hands[i] + " Error = " + err.Error())
			return
		}

		if expected[i] != actual {
			t.Errorf("Expected type %v :: Got type :: %v. Hand = %v", expected[i], actual, h.String())
			return
		}
	}
}

func TestJokerAwareCamelHandOrder(t *testing.T) {

	expected := []string{"32T3K", "KK677", "T55J5", "QQQJA", "KTJJT"}
	hands, err := LoadCamelHandsFromFile("../Data/day7-example.txt", JokerAwareCamelCardFactory{})
	//expected := 5905

	if err != nil {
		t.Error(err)
		return
	}

	ordered := Order(hands, JokerAwareCamelHandTyper{})

	for i := range expected {

		if expected[i] != hands[i].String() {
			t.Errorf("Expected %v :: Got %v", expected, ordered)
			return
		}
	}
}

func TestJokerAwareCamelHandScore(t *testing.T) {
	expected := 5905
	hands, err := LoadCamelHandsFromFile("../Data/day7-example.txt", JokerAwareCamelCardFactory{})

	if err != nil {
		t.Error(err)
		return
	}

	score := ScoreUnordered(hands, JokerAwareCamelHandTyper{})

	if score != expected {
		t.Errorf("Expected Score = %v :: Got Score  = %v", expected, score)
	}
}
