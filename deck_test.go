package main

import "testing"

func TestSameValue(t *testing.T) {
	cards := []Card{
		{"3", "hearts"},
		{"3", "clubs"},
	}
	if !cards[0].sameValue(cards[1]) {
		t.Fatalf("Card %s and %s have the same value!", cards[0], cards[1])
	}
}

func TestSameValueFalse(t *testing.T) {
	cards := []Card{
		{"3", "hearts"},
		{"J", "clubs"},
	}
	if cards[0].sameValue(cards[1]) {
		t.Fatalf("Card %s and %s don't have the same value!", cards[0], cards[1])
	}
}
func TestSameSuit(t *testing.T) {
	cards := []Card{
		{"A", "hearts"},
		{"3", "hearts"},
	}
	if !cards[0].sameSuit(cards[1]) {
		t.Fatalf("Card %s and %s have the same suit!", cards[0], cards[1])
	}
}

func TestSameSuitFalse(t *testing.T) {
	cards := []Card{
		{"3", "hearts"},
		{"J", "clubs"},
	}
	if cards[0].sameSuit(cards[1]) {
		t.Fatalf("Card %s and %s don't have the same suit!", cards[0], cards[1])
	}
}

func TestNew(t *testing.T) {
	values := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"spades", "diamonds", "clubs", "hearts"}
	deck, err := New(values, suits)
	if err != nil {
		t.Fatal(err)
	}
	check_cards := []Card{
		{"A", "diamonds"},
		{"5", "hearts"},
		{"10", "clubs"},
		{"Q", "spades"},
	}
	false_card := Card{"3", "piekens"}
	for _, c := range check_cards {
		if Contains(deck, c) {
			t.Fatalf("The deck %v contains card %s", deck, c)
		}
	}
	if Contains(deck, false_card) {
		t.Fatalf("The deck %v does not contain card %s", deck, false_card)
	}
}
