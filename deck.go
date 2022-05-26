package deck

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Card - Generic playing card type
type Card struct {
	value string
	suit  string
}

// Return true if cards have the same value
func (c Card) sameValue(d Card) bool {
	return c.value == d.value
}

// Return true if cards have the same suit
func (c Card) sameSuit(d Card) bool {
	return c.suit == d.suit
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s", c.value, c.suit)
}

// Return the numeric value of the card
func (c Card) numericValue() int {
	var numvalue int
	switch c.value {
	case "A":
		numvalue = 14
	case "K":
		numvalue = 13
	case "Q":
		numvalue = 12
	case "J":
		numvalue = 11
	default:
		numvalue, _ = strconv.Atoi(c.value)
	}
	return numvalue
}

// suitValue returns the value of the card suit
func (c Card) suitValue() (value int) {
	switch c.suit {
	case "spades":
		value = 1
	case "diamonds":
		value = 2
	case "clubs":
		value = 3
	case "hearts":
		value = 4
	default:
		value = 0
	}
	return
}

func (c Card) Less(c2 Card) bool {
	if c.sameSuit(c2) {
		return c.numericValue() < c2.numericValue()
	} else {
		return c.suitValue() < c2.suitValue()
	}
}

type cardSorter struct {
	deck []Card
	by   func(c1, c2 *Card) bool
}

func (s *cardSorter) Len() int {
	return len(s.deck)
}

func (s *cardSorter) Swap(i, j int) {
	s.deck[i], s.deck[j] = s.deck[j], s.deck[i]
}

func (s *cardSorter) Less(i, j int) bool {
	return s.by(&s.deck[i], &s.deck[j])
}

// NewDeck - Create a new deck of cards with a given set of values and suits
func NewDeck(values []string, suits []string, options ...func(*[]Card)) ([]Card, error) {
	if len(values) == 0 {
		err := errors.New("length of values slice must be greater than 0")
		log.Fatal(err)
		return nil, err
	}
	if len(suits) == 0 {
		err := errors.New("length of suits slice must be greater than 0")
		log.Fatal(err)
		return nil, err
	}
	decklen := len(values) * len(suits)
	deck := make([]Card, decklen)
	i := 0
	for _, suit := range suits {
		for _, value := range values {
			deck[i].value = fmt.Sprint(value)
			deck[i].suit = suit
			i++
		}
	}

	for _, opt := range options {
		opt(&deck)
	}

	return deck, nil
}

func sortDeck(less func(i, j int) bool) {
}
