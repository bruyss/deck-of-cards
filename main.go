package deck

import (
	"fmt"
)

func main() {
	values := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"spades", "diamonds", "clubs", "hearts"}
	cards, _ := NewDeck(values, suits)
	fmt.Println(cards[0].numericValue())
	fmt.Println(cards[0].sameSuit(cards[1]))
}
