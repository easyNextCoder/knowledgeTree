package alga

import "fmt"

const (
	ProjectFlush = iota
	ProjectBomb
)

const MAX_N int16 = 14

const MAX_M int16 = 4

var scoreArr [15]int = [15]int{0, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

var scoreArrInt16 [15]int16 = [15]int16{0, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

type Pair struct {
	First  int16
	Second int16
}
type Card Pair
type Cards []Card
type Deck Cards

func (self Card) String() string {
	if self.Second == 0 {
		return "joker"
	} else {
		return fmt.Sprintf("{%d,%d}", self.First, self.Second)
	}
}

func CARD_OFFSET(card *Card) int16 { return (card.First-1)*MAX_M + (card.Second - 1) }
