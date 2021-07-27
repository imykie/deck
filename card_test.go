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
