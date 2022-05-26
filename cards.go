//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

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

func (c Card) numericRank() int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].numericRank() < cards[j].numericRank()
	}
}

func More(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].numericRank() > cards[j].numericRank()
	}
}
func (c1 Card) Equal(c2 Card) bool {
	return int(c1.Suit) == int(c2.Suit) && int(c1.Rank) == int(c2.Rank)
}

func Equal(d1, d2 []Card) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i := range d1 {
		if !d1[i].Equal(d2[i]) {
			return false
		}
	}
	return true
}

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

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	shuffled := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		shuffled[i] = cards[j]
	}
	return shuffled
}

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

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var combinedDeck []Card
		for i := 0; i < n; i++ {
			combinedDeck = append(combinedDeck, cards...)
		}
		return combinedDeck
	}
}
