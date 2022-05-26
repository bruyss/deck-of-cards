//go:generate stringer -type=Suit,Rank

// Package deck provides a representation of a deck of cards
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents a cards' suit
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents a cards' value
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card is a type representing a playing card
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func (c Card) numericValue() int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Less returns true if card at index i has a lower numerical value than the card at index j
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].numericValue() < cards[j].numericValue()
	}
}

// More returns true if card at index i has a higher numerical value than the card at index j
func More(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].numericValue() > cards[j].numericValue()
	}
}

// EqualCard returns true if c and c2 have the same numerical value
func (c Card) EqualCard(c2 Card) bool {
	return c.numericValue() == c2.numericValue()
}

// Equal returns true if both slices contain the same cards in the same order
func Equal(d1, d2 []Card) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i := range d1 {
		if !d1[i].EqualCard(d2[i]) {
			return false
		}
	}
	return true
}

// NewDeck creates a new deck of cards
// Available options: DefaultSort, Sort, Shuffle, Jokers, Filter, Deck
func NewDeck(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort sorts the cards like a new deck would be
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort sorts the deck with a custom sort function
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Shuffle randomizes the order of the deck
func Shuffle(cards []Card) []Card {
	shuffled := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		shuffled[i] = cards[j]
	}
	return shuffled
}

// Jokers adds n jokers to the deck
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

// Filter removes the cards satisfying the filter function conditions from the deck
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var filtered []Card
		for _, c := range cards {
			if !f(c) {
				filtered = append(filtered, c)
			}
		}
		return filtered
	}
}

// Deck copies the deck n times
// Be aware of the order in which this option is called when making a new deck
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var combinedDeck []Card
		for i := 0; i < n; i++ {
			combinedDeck = append(combinedDeck, cards...)
		}
		return combinedDeck
	}
}
