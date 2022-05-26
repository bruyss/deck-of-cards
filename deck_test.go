package deck

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
	deck, err := NewDeck(values, suits)
	if err != nil {
		t.Fatal(err)
	}
	checkCards := []Card{
		{"A", "diamonds"},
		{"5", "hearts"},
		{"10", "clubs"},
		{"Q", "spades"},
	}
	falseCard := Card{"3", "piekens"}
	for _, c := range checkCards {
		if !Contains(deck, c) {
			t.Fatalf("The deck %v contains card %s", deck, c)
		}
	}
	if Contains(deck, falseCard) {
		t.Fatalf("The deck %v does not contain card %s", deck, falseCard)
	}
}

func TestCardOrder(t *testing.T) {
	cards := []Card{
		{"A", "hearts"},
		{"A", "spades"},
		{"3", "hearts"},
	}
	if cards[0].Less(cards[1]) {
		t.Fatalf("%s is bigger than %s", cards[0], cards[1])
	}
	if cards[2].Less(cards[1]) {
		t.Fatalf("%s is smaller than %s", cards[1], cards[2])
	}
	if cards[0].Less(cards[2]) {
		t.Fatalf("%s is bigger than %s", cards[0], cards[2])
	}
}

// func TestSort(t *testing.T) {
// 	values := []string{"K", "3", "4"}
// 	suits := []string{"hearts", "diamonds"}
// 	cardValues := func(c1, c2 *Card) bool {
// 		return c1.Less(c2)
// 	}
// }
