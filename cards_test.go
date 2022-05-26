package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Three, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Spade})
	fmt.Println(Card{Rank: Eight, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Three of Diamonds
	// King of Spades
	// Eight of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 13*4 {
		t.Fatal("wrong number of cards in deck")
	}
}

func TestEqual(t *testing.T) {
	c1 := Card{Rank: Ace, Suit: Heart}
	c2 := Card{Rank: Ace, Suit: Heart}
	c3 := Card{Rank: Two, Suit: Heart}
	c4 := Card{Rank: Ace, Suit: Club}
	if !c1.EqualCard(c2) {
		t.Fatalf("%s and %s are equal cards", c1, c2)
	}
	if c1.EqualCard(c3) {
		t.Fatalf("%s and %s have a different rank", c1, c3)
	}
	if c1.EqualCard(c4) {
		t.Fatalf("%s and %s have a different suit", c1, c4)
	}
	if c3.EqualCard(c4) {
		t.Fatalf("%s and %s are different cards", c3, c4)
	}
}

func TestDefaultSort(t *testing.T) {
	cards := []Card{
		{Rank: Jack, Suit: Club},
		{Rank: Five, Suit: Heart},
		{Rank: Two, Suit: Club},
		{Rank: Two, Suit: Heart},
		{Rank: Nine, Suit: Diamond},
		{Rank: Three, Suit: Spade},
	}
	sortedCards := []Card{
		{Rank: Three, Suit: Spade},
		{Rank: Nine, Suit: Diamond},
		{Rank: Two, Suit: Club},
		{Rank: Jack, Suit: Club},
		{Rank: Two, Suit: Heart},
		{Rank: Five, Suit: Heart},
	}
	cards = DefaultSort(cards)
	if !Equal(cards, sortedCards) {
		t.Fatalf("%s not sorted correctly", cards)
	}
}

func TestSort(t *testing.T) {
	cards := []Card{
		{Rank: Jack, Suit: Club},
		{Rank: Five, Suit: Heart},
		{Rank: Two, Suit: Club},
		{Rank: Two, Suit: Heart},
		{Rank: Nine, Suit: Diamond},
		{Rank: Three, Suit: Spade},
	}
	sortedCards := []Card{
		{Rank: Five, Suit: Heart},
		{Rank: Two, Suit: Heart},
		{Rank: Jack, Suit: Club},
		{Rank: Two, Suit: Club},
		{Rank: Nine, Suit: Diamond},
		{Rank: Three, Suit: Spade},
	}
	cards = Sort(More)(cards)
	if !Equal(cards, sortedCards) {
		t.Fatalf("%s not sorted correctly", cards)
	}
}

func TestJokers(t *testing.T) {
	cards := NewDeck(Jokers(4))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 4 {
		t.Fatal("expecting 4 jokers, received:", count)
	}
}

func TestFilter(t *testing.T) {
	noTwosThrees := func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}
	cards := NewDeck(Filter(noTwosThrees))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Fatal("2s and 3s should be filtered out of the deck")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := NewDeck(Deck(3))
	if len(cards) != 13*4*3 {
		t.Fatalf("expected %d cards, received %d cards", 13*4*3, len(cards))
	}
}
