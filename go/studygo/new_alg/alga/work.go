package alga

import (
	//"gamehand/common/utils"
	//"gamehand/hand_common/rules"
	//"gamehand/hand_robot/rules_cgo"

	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"studygo/new_alg/rules"
	"studygo/new_alg/rules_cgo"
	"time"
)

//func genNoJokerDecks() DeckWithJoker {
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(AllCards), func(i, j int) {
//		res := AllCards[i]
//		AllCards[i] = AllCards[j]
//		AllCards[j] = res
//	})
//
//	n := utils.RandRange(1, 16)
//
//	cs := make([]int, n)
//	copy(cs, AllCards)
//	//cs = append(cs, 14, 14)
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
//}
//
//func genOneColorDecks() DeckWithJoker {
//
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(OneColorCards), func(i, j int) {
//		res := OneColorCards[i]
//		OneColorCards[i] = OneColorCards[j]
//		OneColorCards[j] = res
//	})
//
//	n := utils.RandRange(1, 14)
//
//	cs := make([]int, n)
//	copy(cs, OneColorCards)
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
func genDecks() DeckWithJoker {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(AllCards), func(i, j int) {
		res := AllCards[i]
		AllCards[i] = AllCards[j]
		AllCards[j] = res
	})

	n := RandRange(1, 14)

	cs := make([]int, n)
	//cs := make([]int, 13)
	copy(cs, AllCards)
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

func GetMaxCardGroupPair(cards Deck, okeyNum int) [][]rules.Card {
	var ccards []rules.Card
	for _, c := range cards {
		ccards = append(ccards, rules.NewCard(int(0x01<<(4)*c.Second), int(c.First)))
	}
	if okeyNum == 1 {
		ccards = append(ccards, rules.OkeyCard)
	}

	if okeyNum == 2 {
		ccards = append(ccards, rules.OkeyCard)
		ccards = append(ccards, rules.OkeyCard)
	}
	//fmt.Println("ccards", ccards, cards)

	return rules_cgo.GetMaxCardGroup(ccards)
}

//
//func (self *AlgX) WorkOld() {
//
//	//PrintLocal := false
//
//	timeNew := int64(0)
//	timeOld := int64(0)
//	totalTimes := 0
//
//	WORKTIMES := 100000000
//	//TimesPerCase := int64(1000)
//	TimesPerCase := int64(10)
//
//	go func() {
//		http.ListenAndServe("localhost:6061", nil)
//	}()
//
//	for x := 0; x < WORKTIMES; x++ {
//
//		c := genDecks() //15张牌，2个joker
//		//c := genOneColorDecks() //3-15张牌，2个joker
//		//c := genNoJokerDecks() //1-15张牌，无joker
//
//		total := int64(0)
//		total2 := int64(0)
//		var res3 [][]rules.Card
//
//		var res *Projects
//		var res2 *Projects
//		var dfsTimes int
//		//fmt.Println("this", c.deck)
//		var score int
//		//var haveCal bool
//		for i := int64(0); i < TimesPerCase; i++ {
//			s := time.Now()
//
//			self.Entry(c.deck, c.jokerNum)
//			e := time.Now()
//			total += e.Sub(s).Microseconds()
//			res3 = GetMaxCardGroupPair(c.deck, c.jokerNum)
//			//f := time.Now()
//			//total2 += f.Sub(e).Microseconds()
//			//if !haveCal {
//			//	haveCal = true
//			//	for _, a := range res3 {
//			//		p := rules.CoreNewProject(a)
//			//		score += p.Score()
//			//	}
//			//}
//
//		}
//
//		timeNew += total
//		timeOld += total2
//
//		totalTimes++
//		if totalTimes%10000 == 0 {
//			fmt.Println(totalTimes)
//			fmt.Println("	per us: ", total/TimesPerCase, " 最大分:", res, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
//			fmt.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1)/TimesPerCase, timeOld/int64(x+1)/TimesPerCase)
//			fmt.Println("   deck", c.deck)
//		}
//
//		//if PrintLocal {
//		//	fmt.Println("new alg big case: ", x)
//		//	fmt.Println("	per us: ", total/TimesPerCase, " 最大分:", res.score, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
//		//	fmt.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
//		//	fmt.Println(c.deck)
//		//}
//		//
//		//if int(res.score) > score {
//		//	log.Println("new alg big case: ", x)
//		//	log.Println("	per us: ", total/TimesPerCase, " 最大分:", res.score, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
//		//	log.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
//		//	log.Println(c.deck)
//		//} else if int(res.score) < score {
//		//	log.Println("old alg big case: ", x, time.Now())
//		//	log.Println("	per us: ", total/TimesPerCase, " 最大分:", res.score, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
//		//	log.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
//		//	log.Println(c.deck)
//		//	fmt.Println(c.deck)
//		//	//panic("")
//		//} else {
//		//
//		//}
//	}
//
//}
func (self *AlgX) Work() {

	//PrintLocal := false

	timeNew := int64(0)
	timeOld := int64(0)
	totalTimes := 0

	WORKTIMES := 100000000
	//WORKTIMES := 1
	//TimesPerCase := int64(1000)
	TimesPerCase := int64(10)

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	var cnt int

	for x := 0; x < WORKTIMES; x++ {

		c := genDecks() //15张牌，2个joker
		//c := genOneColorDecks() //3-15张牌，2个joker
		//c := genNoJokerDecks() //1-15张牌，无joker

		total := int64(0)
		total2 := int64(0)

		//fmt.Println("this", c.deck)

		//var haveCal bool

		var maxScore, maxLen Projects
		var set bool
		for i := int64(0); i < TimesPerCase; i++ {
			s := time.Now()
			cnt++
			var solver AlgX
			solver.Entry(c.deck, c.jokerNum)
			if !set {
				set = true
				maxScore = solver.maxScoreProjects
				maxLen = solver.maxLenProjects
			}

			e := time.Now()
			total += e.Sub(s).Microseconds()
			//res3 = GetMaxCardGroupPair(c.deck, c.jokerNum)
			//f := time.Now()
			//total2 += f.Sub(e).Microseconds()
			//if !haveCal {
			//	haveCal = true
			//	for _, a := range res3 {
			//		p := rules.CoreNewProject(a)
			//		score += p.Score()
			//	}
			//}

		}

		//time.Sleep(5 * time.Microsecond)
		timeNew += total
		timeOld += total2

		totalTimes++
		if totalTimes%10000 == 0 {
			fmt.Println(totalTimes)
			fmt.Println("   deck", c.deck)
			fmt.Println("	per us: ", total/TimesPerCase, " 最大分:", maxScore, "最大长度", maxLen)
			fmt.Println(timeNew/int64(x+1)/TimesPerCase, timeOld/int64(x+1)/TimesPerCase)

		}

		//if PrintLocal {
		//	fmt.Println("new alg big case: ", x)
		//	fmt.Println("	per us: ", total/TimesPerCase, " 最大分:", res.score, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
		//	fmt.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
		//	fmt.Println(c.deck)
		//}
		//
		//if int(res.score) > score {
		//	log.Println("new alg big case: ", x)
		//	log.Println("	per us: ", total/TimesPerCase, " 最大分:", res.score, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
		//	log.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
		//	log.Println(c.deck)
		//} else if int(res.score) < score {
		//	log.Println("old alg big case: ", x, time.Now())
		//	log.Println("	per us: ", total/TimesPerCase, " 最大分:", res.score, " ", res, "最大长度", res2, "dfsTimes:", dfsTimes)
		//	log.Println("	per us:", total2/TimesPerCase, " 最大分:", score, " ", res3, timeNew/int64(x+1), timeOld/int64(x+1))
		//	log.Println(c.deck)
		//	fmt.Println(c.deck)
		//	//panic("")
		//} else {
		//
		//}
	}

}
