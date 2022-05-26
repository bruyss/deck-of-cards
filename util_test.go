package main

import "testing"

func TestContainIntegersTrue(t *testing.T) {
	s_ints := []int{3, 4, 20, 50}
	value := 4
	if !Contains(s_ints, value) {
		t.Fatalf("The slice %v contains the value %d", s_ints, value)
	}
}

func TestContainIntegersFalse(t *testing.T) {
	s_ints := []int{3, 4, 20, 50}
	value := 5
	if Contains(s_ints, value) {
		t.Fatalf("The slice %v does not contain the value %d", s_ints, value)
	}
}

func TestContainStringTrue(t *testing.T) {
	s_strs := []string{"allo", "daar", "test", "123"}
	values := []string{"allo", "test"}
	for _, v := range values {
		if !Contains(s_strs, v) {
			t.Fatalf("The slice %v contains the value %s", s_strs, v)
		}
	}
}

func TestContainStringFalse(t *testing.T) {
	s_strs := []string{"allo", "daar", "test", "123"}
	values := []string{"llo", "Jeroen", "yawadded"}
	for _, v := range values {
		if Contains(s_strs, v) {
			t.Fatalf("The slice %v does not contain the value %s", s_strs, v)
		}
	}
}

func TestContainsCards(t *testing.T) {
	values := []string{"A", "5", "K"}
	suits := []string{"diamonds", "clubs", "hearts"}
	deck, _ := New(values, suits)
	check_cards := []Card{
		{"A", "diamonds"},
		{"5", "hearts"},
	}
	false_card := Card{"3", "piekens"}
	for _, c := range check_cards {
		if !Contains(deck, c) {
			t.Fatalf("The deck %v contains card %s", deck, c)
		}
	}
	if Contains(deck, false_card) {
		t.Fatalf("The deck %v does not contain card %s", deck, false_card)
	}
}
