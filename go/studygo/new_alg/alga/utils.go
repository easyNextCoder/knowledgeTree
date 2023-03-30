package alga

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

/// gRander need lock() and defer unlock
var randMutex = sync.Mutex{}
var gRander = rand.New(rand.NewSource(time.Now().UnixNano()))

/**
rand [min,max)
*/
func RandRange(i, j int) int {
	min := i
	max := j
	if min > max {
		min, max = max, min
	}
	if (max - min) <= 0 {
		panic("invalid argument to randrange max cant equal min")
	}
	randMutex.Lock()
	rand := gRander.Intn(max-min) + min
	randMutex.Unlock()
	return rand
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a ...int) int {
	min := 99999
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func mapCopy(freqSrc map[int16]int) map[int16]int {
	var dst map[int16]int = map[int16]int{}
	for k, v := range freqSrc {
		dst[k] = v
	}
	return dst
}

func freqMapToCards(freq map[int16]int, color int16) []Card {
	var cards []Card
	for k, v := range freq {
		if v > 0 {
			n := v
			for n > 0 {
				cards = append(cards, Card{int16(k), int16(color)})
				n--
			}
		}
	}
	return cards
}

func freqArrToCards(freq []int, color int16) []Card {
	var cards []Card
	for k, v := range freq { //这个函数是在range的顺序一定遵循0,1,2,3,4,的前提下才会产生正确的结果
		if v > 0 {
			n := v
			for n > 0 {
				cards = append(cards, Card{int16(k), int16(color)})
				n--
			}
		}
	}
	return cards
}

func freqToCards(freq []int, color int16) []Card {
	var cards []Card
	for i := 0; i < len(freq); i++ {
		if freq[i] > 0 {
			n := freq[i]
			for n > 0 {
				cards = append(cards, Card{int16(i), int16(color)})
				n--
			}
		}
	}
	return cards
}

func canFilterOneJokerDesc(a1, b1, c1 *int, v int) (int, bool) {
	if *a1 < 0 || *b1 < 0 || *c1 < 0 {
		return 0, false
	}
	a, b, c := *a1, *b1, *c1
	min := getMin(a, b, c)
	if min > 0 {
		a -= min
		b -= min
		c -= min
	}

	noZeroCnt := 0
	resV := 0

	if a != 0 {
		noZeroCnt++
	} else {
		resV = v
	}

	if b != 0 {
		noZeroCnt++
	} else {
		resV = v - 1
	}

	if c != 0 {
		noZeroCnt++
	} else {
		resV = v - 2
	}

	if noZeroCnt == 2 {
		*a1--
		*b1--
		*c1--
		return resV, true
	}

	return 0, false
}

func canFilterOneJoker(a1, b1, c1 *int, v int) (int, bool) {
	if *a1 < 0 || *b1 < 0 || *c1 < 0 {
		return 0, false
	}
	a, b, c := *a1, *b1, *c1
	min := getMin(a, b, c)
	if min > 0 {
		a -= min
		b -= min
		c -= min
	}

	noZeroCnt := 0
	resV := 0

	if a != 0 {
		noZeroCnt++
	} else {
		resV = v
	}

	if b != 0 {
		noZeroCnt++
	} else {
		resV = v + 1
	}

	if c != 0 {
		noZeroCnt++
	} else {
		resV = v + 2
	}

	if noZeroCnt == 2 {
		*a1--
		*b1--
		*c1--
		return resV, true
	}

	return 0, false
}

func (self *AlgX) getFreq(index int16) int {
	if index < 0 || int(index) >= len(self.freq) {
		return 0
	} else {
		return self.freq[index]
	}
}

func ProjectsScore(p []*Project) int16 {
	var score int16
	for _, v := range p {
		score += v.score
	}
	return score
}

func scoreProjects(cardss [][]Card) int {
	s := 0
	for i := 0; i < len(cardss); i++ {
		s += score(cardss[i])
	}
	return s
}

func lenProjects(cardss [][]Card) int {
	s := 0
	for i := 0; i < len(cardss); i++ {
		if len(cardss[i]) < 3 {
			panic("xxx")
		}
		s += len(cardss[i])
	}
	return s
}

func score(cards []Card) int {
	s := 0
	for i := 0; i < len(cards); i++ {
		s += scoreArr[cards[i].First]
	}
	return s
}

func copyCards(cards []Card) []Card {
	ret := make([]Card, len(cards))
	copy(ret, cards)
	return ret
}

func colorMapToCards(b *Bomb, val int16) ([9]Card, int) {
	var res [9]Card
	var x int
	for color, cnt := range b.colors {
		i := cnt
		for i > 0 && color > 0 {
			i--
			res[x] = Card{val, int16(color)}
			x++
			//res = append(res, Card{val, int16(color)})
		}
	}
	return res, x
}

func colorMapRemoveCard(b *Bomb, card Card) {

	if b.colors[card.Second] <= 0 {
		fmt.Println("card.second", card.Second)
		panic("bomb color <= 0")
	}
	b.colors[card.Second]--
	if b.colors[card.Second] == 0 {
		b.colorCnt--
	}
	b.cnt--
}

func colorMapRemoveColor(b *Bomb, c int16) {

	if b.colors[c] <= 0 {
		fmt.Println("card.second", b.colors, c)
		panic("bomb color <= 0")
	}
	b.colors[c]--
	if b.colors[c] == 0 {
		b.colorCnt--
	}
	b.cnt--
}

func sortByColor(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {

		if cards[i].Second == cards[j].Second {
			return cards[i].First < cards[j].First
		} else if cards[i].Second < cards[j].Second {
			return true
		}
		return false
	})
}

func sortByPoint(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {

		if cards[i].First == cards[j].First {
			return cards[i].Second < cards[j].Second
		} else if cards[i].First < cards[j].First {
			return true
		}
		return false
	})
}
