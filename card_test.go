package deck

import (
	"fmt"
	"testing"
)

type cardTest struct {
	rank     Rank
	suit     Suit
	expected string
}

var cardTests = []cardTest{
	{Ace, Heart, "Ace of Hearts"},
	{Two, Spade, "Two of Spades"},
	{Four, Diamond, "Four of Diamonds"},
	{Nine, Club, "Nine of Clubs"},
	{Ten, Joker, "Joker"},
}

func TestCard(t *testing.T) {
	for _, test := range cardTests {
		output := fmt.Sprint(Card{test.suit, test.rank})
		if output != test.expected {
			t.Errorf("Output of %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 52 {
		t.Errorf("Wrong number of cards in a new deck")
	}
}

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Four, Suit: Diamond})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output
	// Ace of Hearts
	// Two of Spades
	// Four of Diamonds
	// Nine of Clubs
	// Joker
}
