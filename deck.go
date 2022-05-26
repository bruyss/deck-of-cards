package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Generic playing card type
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

// Create a new deck of cards with a given set of values and suits
func New(values []string, suits []string, options ...func(*[]Card)) ([]Card, error) {
	if len(values) == 0 {
		err := errors.New("Length of values slice must be greater than 0!")
		log.Fatal(err)
		return nil, err
	}
	if len(suits) == 0 {
		err := errors.New("Length of suits slice must be greater than 0!")
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
	return deck, nil
}

// func Shuffle(deck *[]Card) {

// }
