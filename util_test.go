package deck

import "testing"

func TestContainIntegersTrue(t *testing.T) {
	sInts := []int{3, 4, 20, 50}
	value := 4
	if !Contains(sInts, value) {
		t.Fatalf("The slice %v contains the value %d", sInts, value)
	}
}

func TestContainIntegersFalse(t *testing.T) {
	sInts := []int{3, 4, 20, 50}
	value := 5
	if Contains(sInts, value) {
		t.Fatalf("The slice %v does not contain the value %d", sInts, value)
	}
}

func TestContainStringTrue(t *testing.T) {
	sStrs := []string{"allo", "daar", "test", "123"}
	values := []string{"allo", "test"}
	for _, v := range values {
		if !Contains(sStrs, v) {
			t.Fatalf("The slice %v contains the value %s", sStrs, v)
		}
	}
}

func TestContainStringFalse(t *testing.T) {
	sStrs := []string{"allo", "daar", "test", "123"}
	values := []string{"llo", "Jeroen", "yawadded"}
	for _, v := range values {
		if Contains(sStrs, v) {
			t.Fatalf("The slice %v does not contain the value %s", sStrs, v)
		}
	}
}

func TestContainsCards(t *testing.T) {
	values := []string{"A", "5", "K"}
	suits := []string{"diamonds", "clubs", "hearts"}
	deck, _ := NewDeck(values, suits)
	checkCards := []Card{
		{"A", "diamonds"},
		{"5", "hearts"},
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
