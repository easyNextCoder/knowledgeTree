package alga

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"studygo/new_alg/rules"
	"testing"
	"time"
)

//
//import (
//	"encoding/json"
//	"fmt"
//	"gamehand/common/utils"
//	"gamehand/hand_common/rules"
//	"gamehand/hand_robot/rules_cgo"
//	"log"
//	"math/rand"
//	"net/http"
//	_ "net/http/pprof"
//	"testing"
//	"time"
//)
//
//func TestAlgX_TopSearch(t *testing.T) {
//	type fields struct {
//		hand_cards                        Deck
//		jokerNum                          int
//		addJokerToColumn                  int
//		haveProcessedInOneJokerMiddleNode bool
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//		{name: "测试一个joker的补牌情况"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			self := &AlgX{
//				hand_cards:                        tt.fields.hand_cards,
//				jokerLeft:                         tt.fields.jokerNum,
//				addJokerToColumn:                  tt.fields.addJokerToColumn,
//				haveProcessedInOneJokerMiddleNode: tt.fields.haveProcessedInOneJokerMiddleNode,
//			}
//			self.TopSearch()
//		})
//	}
//}
//
//func GetMaxCardGroupPair(cards Deck, okeyNum int) [][]rules.Card {
//	var ccards []rules.Card
//	for _, c := range cards {
//		ccards = append(ccards, rules.NewCard(int(0x01<<(4)*c.Second), int(c.First)))
//	}
//	if okeyNum == 1 {
//		ccards = append(ccards, rules.OkeyCard)
//	}
//
//	if okeyNum == 2 {
//		ccards = append(ccards, rules.OkeyCard)
//		ccards = append(ccards, rules.OkeyCard)
//	}
//	//fmt.Println("ccards", ccards, cards)
//
//	return rules_cgo.GetMaxCardGroup(ccards)
//}
//
//func unmarshalToCards(s string) []rules.Card {
//	var ic []rules.Card
//	err := json.Unmarshal([]byte(s), &ic)
//	if err != nil {
//		return nil
//	}
//	return ic
//}
//
//func unmarshalAllToCards(ss []string) [][]rules.Card {
//	var ret [][]rules.Card
//	for _, s := range ss {
//		one := unmarshalToCards(s)
//		ret = append(ret, one)
//	}
//	return ret
//}
//
func genNoJokerDecks() DeckWithJoker {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(AllCards), func(i, j int) {
		res := AllCards[i]
		AllCards[i] = AllCards[j]
		AllCards[j] = res
	})

	n := RandRange(3, 17)

	cs := make([]int, n)
	copy(cs, AllCards)
	//cs = append(cs, 14, 14)

	jokerNum := 0
	var vertCards []Card
	for _, c := range cs {
		if c == 14 || c == 15 {
			jokerNum++
		} else {
			vertCards = append(vertCards, Card{int16(c % 16), int16(c / 16)})
		}
	}
	//fmt.Println(vertCards)
	return DeckWithJoker{vertCards, jokerNum}
}

func genOneColorDecks() DeckWithJoker {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(OneColorCards), func(i, j int) {
		res := OneColorCards[i]
		OneColorCards[i] = OneColorCards[j]
		OneColorCards[j] = res
	})

	n := RandRange(1, 14)

	cs := make([]int, n)
	copy(cs, OneColorCards)
	cs = append(cs, 14, 14)

	jokerNum := 0
	var vertCards []Card
	for _, c := range cs {
		if c == 14 || c == 15 {
			jokerNum++
		} else {
			vertCards = append(vertCards, Card{int16(c % 16), int16(c / 16)})
		}
	}
	//fmt.Println(vertCards)
	return DeckWithJoker{vertCards, jokerNum}

}

func genOneJokerDecks() DeckWithJoker {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(OneColorCards), func(i, j int) {
		res := OneColorCards[i]
		OneColorCards[i] = OneColorCards[j]
		OneColorCards[j] = res
	})

	n := RandRange(2, 16)

	cs := make([]int, n)
	copy(cs, OneColorCards)
	cs = append(cs, 14)

	jokerNum := 0
	var vertCards []Card
	for _, c := range cs {
		if c == 14 || c == 15 {
			jokerNum++
		} else {
			vertCards = append(vertCards, Card{int16(c % 16), int16(c / 16)})
		}
	}
	//fmt.Println(vertCards)
	return DeckWithJoker{vertCards, jokerNum}

}

//
//func genDecks() DeckWithJoker {
//
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(AllCards), func(i, j int) {
//		res := AllCards[i]
//		AllCards[i] = AllCards[j]
//		AllCards[j] = res
//	})
//
//	n := utils.RandRange(1, 14)
//
//	cs := make([]int, n)
//	//cs := make([]int, 13)
//	copy(cs, AllCards)
//	cs = append(cs, 14, 14)
//
//	jokerNum := 0
//	var vertCards []Card
//	for _, c := range cs {
//		if c == 14 || c == 15 {
//			jokerNum++
//		} else {
//			vertCards = append(vertCards, Card{int16(c % 16), int16(c / 16)})
//		}
//	}
//	//fmt.Println(vertCards)
//	return DeckWithJoker{vertCards, jokerNum}
//
//}
//
func TestAlgX_handleOneJoker(t *testing.T) {
	type fields struct {
		hand_cards                        Deck
		jokerNum                          int
		addJokerToColumn                  int
		haveProcessedInOneJokerMiddleNode bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "测试一个joker的补牌情况"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			timeNew := int64(0)
			timeOld := int64(0)
			totalTimes := 0

			WORKTIMES := 100000000
			//TimesPerCase := int64(1000)
			TimesPerCase := int64(1)

			go func() {
				http.ListenAndServe("localhost:6060", nil)
			}()

			for x := 0; x < WORKTIMES; x++ {

				//c := genDecks() //15张牌，2个joker
				//c := genOneColorDecks() //3-15张牌，2个joker
				//c := genNoJokerDecks() //1-15张牌，无joker
				c := genOneJokerDecks() //3-16张牌，一个joker
				total := int64(0)
				total2 := int64(0)
				var res3 [][]rules.Card

				//fmt.Println("this", c.deck)
				var score int
				var haveCal bool
				var maxScore, maxLen Projects
				var get bool
				for i := int64(0); i < TimesPerCase; i++ {
					s := time.Now()
					//res, _, _ = self.Entry(c.deck, c.jokerNum)
					var solver AlgX
					solver.Entry(c.deck, c.jokerNum)
					if !get {
						get = true
						maxScore = solver.maxScoreProjects
						maxLen = solver.maxLenProjects
					}
					e := time.Now()
					total += e.Sub(s).Microseconds()
					res3 = GetMaxCardGroupPair(c.deck, c.jokerNum)
					f := time.Now()
					total2 += f.Sub(e).Microseconds()
					if !haveCal {
						haveCal = true
						for _, a := range res3 {
							p := rules.CoreNewProject(a)
							score += p.Score()
						}
					}

				}

				timeNew += total
				timeOld += total2

				totalTimes++
				if totalTimes%50000 == 0 {
					fmt.Println(totalTimes)
					fmt.Println("case: ", x)
					fmt.Println("   cards:", c.deck)
					fmt.Println("   per us: ", total/TimesPerCase, " 最大分:", maxScore.score, " ", maxScore, "最大长度", maxLen)
					fmt.Println("   per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))

				}

				if int(maxScore.score) > score {
					log.Println(totalTimes)
					log.Println("new alg big case: ", x)
					log.Println("   cards", c.deck)
					log.Println("   per us: ", total/TimesPerCase, " 最大分:", maxScore.score, " ", maxScore, "最大长度", maxLen)
					log.Println("   per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))

				} else if int(maxScore.score) < score {
					log.Println(totalTimes)
					log.Println("old alg big case: ", x)
					log.Println("   cards", c.deck)
					log.Println("   per us: ", total/TimesPerCase, " 最大分:", maxScore.score, " ", maxScore, "最大长度", maxLen)
					log.Println("   per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
					fmt.Println("   deck", c.deck)
					panic("")
				} else {

				}
			}

		})
	}
}

func TestAlgX_handleOneJoker_singleCase(t *testing.T) {
	type fields struct {
		hand_cards                        Deck
		jokerNum                          int
		addJokerToColumn                  int
		haveProcessedInOneJokerMiddleNode bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "测试一个joker的补牌情况"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//self := &AlgX{
			//	hand_cards:                        tt.fields.hand_cards,
			//	jokerLeft:                         tt.fields.jokerNum,
			//	addJokerToColumn:                  tt.fields.addJokerToColumn,
			//	haveProcessedInOneJokerMiddleNode: tt.fields.haveProcessedInOneJokerMiddleNode,
			//	//bombCards:                         [][]Card{},
			//}

			//var PrintLocal = true

			timeNew := int64(0)
			timeOld := int64(0)
			totalTimes := 0

			//var res2 *Projects
			//var dfsTimes int
			for i, c := range cases {
				totalTimes++
				if totalTimes%1000 == 0 {
					fmt.Println(totalTimes)
				}
				total := int64(0)
				total2 := int64(0)
				var res3 [][]rules.Card
				n := int64(1)

				//fmt.Println("this", c.deck)
				var maxScore, maxLen Projects
				var get bool
				var score int
				for i := int64(0); i < n; i++ {
					s := time.Now()
					//res, _, _ = self.Entry(c.deck, c.jokerNum)
					var solver AlgX
					solver.Entry(c.deck, c.jokerNum)
					if !get {
						get = true
						maxScore = solver.maxScoreProjects
						maxLen = solver.maxLenProjects
					}
					e := time.Now()
					total += e.Sub(s).Microseconds()

					res3 = GetMaxCardGroupPair(c.deck, c.jokerNum)
					f := time.Now()
					total2 += f.Sub(e).Microseconds()
					score = 0
					for _, a := range res3 {
						p := rules.CoreNewProject(a)
						score += p.Score()
					}
				}

				timeNew += total / n
				timeOld += total2 / n

				if int(maxScore.score) >= score {
					fmt.Println("new alg big case: ", i)
					fmt.Println("	per us: ", total/n, " 最大分:", maxScore.score, " ", maxScore, "最大长度", maxLen)
					fmt.Println("	per us:", total2/n, " 最大分:", score, " ", res3, timeNew/int64(i+1), timeOld/int64(i+1))
					fmt.Println(c.deck)
				} else {
					fmt.Println("old alg big case: ", i)
					fmt.Println("	per us: ", total/n, " 最大分:", maxScore.score, " ", maxScore, "最大长度", maxLen)
					fmt.Println("	per us:", total2/n, " 最大分:", score, " ", res3, timeNew/int64(i+1), timeOld/int64(i+1))
					fmt.Println(c.deck)
					panic(" ")
				}

			}

		})
	}
}

//
//func TestAlgX_travelsWork(t *testing.T) {
//	type fields struct {
//		hand_cards                        Deck
//		jokerLeft                         int
//		jokerNum                          int
//		addJokerToColumn                  int
//		haveProcessedInOneJokerMiddleNode bool
//		tmpFlushs                         Projects
//		states                            [][2]int16
//		bombCards                         [][]Card
//		flushCards                        [][]Card
//		recordUsed                        [60]bool
//		bombRepeatMap                     []*Bomb
//		bombTakeAheadMaxScoreProjects     Projects
//		tmpMaxScoreProjects               Projects
//		maxScoreProjects                  Projects
//		tmpMaxLenProjects                 Projects
//		maxLenProjects                    Projects
//		dfsTimes                          int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//		{name: "测试"},
//	}
//	//Joker ♠2 ♠3 ♠4 ♠5 ♠6 ♠7 ♠8 ♠9 ♠10 ♠J ♠Q] [Joker ♠K ♠A]
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			self := &AlgX{
//				hand_cards:                        tt.fields.hand_cards,
//				jokerLeft:                         tt.fields.jokerLeft,
//				jokerNum:                          tt.fields.jokerNum,
//				addJokerToColumn:                  tt.fields.addJokerToColumn,
//				haveProcessedInOneJokerMiddleNode: tt.fields.haveProcessedInOneJokerMiddleNode,
//				tmpFlushs:                         tt.fields.tmpFlushs,
//				states:                            tt.fields.states,
//				bombCards:                         tt.fields.bombCards,
//				flushCards:                        tt.fields.flushCards,
//				recordUsed:                        tt.fields.recordUsed,
//				bombRepeatMap:                     tt.fields.bombRepeatMap,
//				bombTakeAheadMaxScoreProjects:     tt.fields.bombTakeAheadMaxScoreProjects,
//				tmpMaxScoreProjects:               tt.fields.tmpMaxScoreProjects,
//				maxScoreProjects:                  tt.fields.maxScoreProjects,
//				tmpMaxLenProjects:                 tt.fields.tmpMaxLenProjects,
//				maxLenProjects:                    tt.fields.maxLenProjects,
//				dfsTimes:                          tt.fields.dfsTimes,
//			}
//			for i := 0; i < 1000; i++ {
//				s := time.Now()
//				//self.travelsWork()
//				e := time.Now()
//				fmt.Println(self.tmpMaxScoreProjects, e.Sub(s).Microseconds())
//			}
//
//		})
//	}
//}
