package alga

import "fmt"

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
