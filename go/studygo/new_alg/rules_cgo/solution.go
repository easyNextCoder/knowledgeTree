package rules_cgo

import (
	"fmt"
	"studygo/new_alg/rules"

	//"gamehand/hand_robot/app/common/utils"
	//"github.com/ziipin-server/niuhe"
	"sync"
)

func packCards(plainCards []rules.Card) IntVector {
	var vec = NewIntVector()
	vec.Reserve(int64(len(plainCards) * 2))
	for _, card := range plainCards {
		vec.Add(card.Num())
		vec.Add(card.Color() >> 4)
		//niuhe.LogInfo("Card:{%d, %d}", card.Num(), card.Color() >> 4)
	}
	return vec
}

func unpackCards(cardVec IntVector) [][]rules.Card {
	var group [][]rules.Card
	var current []rules.Card
	for i := 0; i < int(cardVec.Size()); i += 2 {
		var num = cardVec.Get(i)
		var color = cardVec.Get(i + 1)
		if color == 1024 && num == 1024 {
			if len(current) > 0 {
				group = append(group, current)
				current = nil
				continue
			}
		} else {
			if color == 0 {
				current = append(current, rules.OkeyCard)
			} else {
				current = append(current, rules.NewCard(color<<4, num))
			}
		}
		//niuhe.LogInfo("returnCard:{%d,:%d}", num, color)
	}
	return group
}

var Cc func(int, []rules.Card) [][]rules.Card
var CurTestFlag TestFlag

type TestFlag struct {
	DisableCompareCgo bool `json:"disable_compare_cgo"` // todo cc hand ai
}

var Usego = false

//func CardArrScore(arr []rules.Card) int {
//	n := len(arr)
//	c1 := arr[0]
//	c2 := arr[1]
//	if rules.IsOkCard(c1) {
//		c1 = arr[2]
//	} else {
//		if rules.IsOkCard(c2) {
//			c2 = arr[2]
//		}
//	}
//	if c1.Num() == c2.Num() {
//		return n * c1.HandScore()
//	}
//
//	sum := 0
//	for i, v := range arr {
//		if rules.IsOkCard(v) {
//			if i == 0 {
//				next, _ := arr[1].Prev()
//				sum += next.CoreGameOverHandScore()
//			} else {
//				pre := arr[i-1]
//				if pre.Num() == rules.CardNumK {
//					sum += rules.NewCard(pre.Color(), rules.CardNumA).CoreGameOverHandScore()
//				} else {
//					next, _ := pre.Next()
//					sum += next.CoreGameOverHandScore()
//				}
//			}
//		} else {
//			sum += v.CoreGameOverHandScore()
//		}
//	}
//
//	return sum
//}
//
//func GetMaxCardGroupCcCompare(okey int, plain []rules.Card, res [][]rules.Card) {
//	score1 := 0
//	for _, v := range res {
//		score1 += CardArrScore(v)
//	}
//
//	res2 := GetMaxCardGroupCgo(okey, plain)
//	score2 := 0
//	for _, v := range res2 {
//		score2 += CardArrScore(v)
//	}
//
//	diff := score1 - score2
//	if diff == 0 {
//		return
//	}
//	zplog.OutErr("GetMaxCardGroupCcCompare diff(%d) score1(%d) score2(%d) okey(%d) plain(%s) res(%s) res2(%s)", diff, score1, score2, okey, plain, res, res2)
//}

//func GetMaxCardGroupCc(okey int, plain []rules.Card) [][]rules.Card {
//	res := Cc(okey, plain)
//
//	if zpenv.IsDebug() && !CurTestFlag.DisableCompareCgo {
//		GetMaxCardGroupCcCompare(okey, plain, res)
//	}
//
//	return res
//}

//func GetMaxCardGroupCcCompare2(okey int, plain []rules.Card, res2 [][]rules.Card) {
//	score1 := 0
//	res := Cc(okey, plain)
//	for _, v := range res {
//		score1 += CardArrScore(v)
//	}
//
//	score2 := 0
//	for _, v := range res2 {
//		score2 += CardArrScore(v)
//	}
//
//	diff := score1 - score2
//	if diff == 0 {
//		return
//	}
//	zplog.OutErr("GetMaxCardGroupCc2Compare diff(%d) score1(%d) score2(%d) okey(%d) plain(%s) res(%s) res2(%s)", diff, score1, score2, okey, plain, res, res2)
//}
//
//func GetMaxCardGroupCc2(okey int, plain []rules.Card) [][]rules.Card {
//	//res := Cc(okey, plain)
//	res2 := GetMaxCardGroupCgo(okey, plain)
//
//	if zpenv.IsDebug() && !CurTestFlag.DisableCompareCgo {
//		GetMaxCardGroupCcCompare2(okey, plain, res2)
//	}
//
//	return res2
//}
//var PrintLocal bool = false
//var PrintLimit = datetime.TimeLimit(0)
//var PrintLimitB = datetime.TimeLimit(0)
//var newStat utils.TimeCostStatis = utils.TimeCostStatis{Tag: "go", PrintInterval: 1000}
//var oldStat utils.TimeCostStatis = utils.TimeCostStatis{Tag: "cgo", PrintInterval: 1000}
//var emptyInputCnt int64

var container_us0 map[int64]string = map[int64]string{}
var container_us10 map[int64]string = map[int64]string{}
var container_us100 map[int64]string = map[int64]string{}
var container_us500 map[int64]string = map[int64]string{}
var container_ms10 map[int64]string = map[int64]string{}
var container_ms50 map[int64]string = map[int64]string{}
var container_ms100 map[int64]string = map[int64]string{}
var container_ms200 map[int64]string = map[int64]string{}
var container_ms500 map[int64]string = map[int64]string{}
var container_s1 map[int64]string = map[int64]string{}

var totalContainer []map[int64]string = []map[int64]string{
	container_us0,
	container_us10,
	container_us100,
	container_us500,
	container_ms10,
	container_ms50,
	container_ms100,
	container_ms200,
	container_ms500,
	container_s1,
}

var lock sync.RWMutex = sync.RWMutex{}
var containerPrintIndex int = 0

func hashCard(c rules.Card) int64 {
	switch c.Color() {
	case rules.ColorNone:
		return int64(c.Num() * 4)
	case rules.ColorBlue:
		return int64(c.Num() * 1)
	case rules.ColorRed:
		return int64(c.Num() * 2)
	case rules.ColorBlack:
		return int64(c.Num() * 3)
	case rules.ColorYellow:
		return int64(c.Num() * 4)
	}
	return 0
}

//将传入的一手牌哈希成唯一一个值
func hashCards(cards []rules.Card) int64 {
	var ret int64
	for _, card := range cards {
		ret |= 1 << hashCard(card)
	}
	return ret
}

var total int64 = 0
var total2 int64 = 0
var cnt int64 = 0

func GetMaxCardGroup(cards []rules.Card) [][]rules.Card {
	cnt++
	var plain = []rules.Card{}
	var okey int
	for _, c := range cards {
		if c == rules.OkeyCard {
			okey += 1
		} else {
			plain = append(plain, c)
		}

	}

	if len(cards) == 0 {
		//emptyInputCnt++
	}

	//if Usego && Cc != nil {
	//	//return GetMaxCardGroupCc2(okey, plain)
	//	return GetMaxCardGroupCc(okey, plain)
	//}

	return GetMaxCardGroupCgo(okey, plain)
	//if zpenv.IsProd() {
	//	return GetMaxCardGroupCgo(okey, plain)
	//} else {
	//	var res, algRes [][]rules.Card
	//	var newCost, oldCost time.Duration
	//	//sw := sync.WaitGroup{}
	//	//
	//	//sw.Add(1)
	//	//go func() {
	//	//	defer sw.Done()
	//	start1 := time.Now()
	//	//alg.TopSearth(plain, int16(okey))
	//	end1 := time.Now()
	//	gap1 := end1.Sub(start1).Microseconds()
	//	total += gap1
	//	newCost = end1.Sub(start1)
	//	//newStat.SetRecordFunc(func(s *utils.TimeCostStatis, index int) {
	//	//	if index < len(totalContainer) {
	//	//		container := totalContainer[index]
	//	//		if container == nil {
	//	//			return
	//	//		}
	//	//		if len(container) < 10 {
	//	//			hashVal := hashCards(cards)
	//	//			if _, ok := container[hashVal]; !ok {
	//	//				lock.Lock()
	//	//				container[hashVal] = zpjson.JsonString(cards)
	//	//				lock.Unlock()
	//	//			}
	//	//		}
	//	//	}
	//	//
	//	//})
	//	newStat.Record(newCost)
	//	//}()
	//
	//	//sw.Add(1)
	//	//go func() {
	//	//	defer sw.Done()
	//	start := time.Now()
	//	res = GetMaxCardGroupCgo(okey, plain)
	//	end := time.Now()
	//	gap2 := end.Sub(start).Microseconds()
	//	total2 += gap2
	//	//fmt.Println("final Res ", total, total2, total/cnt, total2/cnt, gap1, gap2)
	//	//oldCost = end.Sub(start)
	//	//oldStat.SetRecordFunc(func(s *utils.TimeCostStatis, index int) {
	//	//	if index < len(totalContainer) {
	//	//		container := totalContainer[index]
	//	//		if container == nil {
	//	//			return
	//	//		}
	//	//		if len(container) < 10 {
	//	//			hashVal := hashCards(cards)
	//	//			if _, ok := container[hashVal]; !ok {
	//	//				lock.Lock()
	//	//				container[hashVal] = zpjson.JsonString(cards)
	//	//				lock.Unlock()
	//	//			}
	//	//		}
	//	//	}
	//	//})
	//	//oldStat.Record(oldCost)
	//	//}()
	//	//
	//	//sw.Wait()
	//
	//	rules.SortCardsProject(res)
	//	rules.SortCardsProject(algRes)
	//
	//	equal := rules.CardsEquals(rules.CardsProject(res).ToCards(), rules.CardsProject(algRes).ToCards())
	//	//fmt.Println("PrintLocal is", PrintLocal, algRes)
	//
	//	if PrintLocal {
	//		fmt.Printf("algRes(%t) new(%v)(%d)us  old(%v)(%d)us \n", equal, algRes, gap1, res, gap2)
	//		oldStat.PrintTimeCostStatis(utils.Stdout_Print)
	//		newStat.PrintTimeCostStatis(utils.Stdout_Print)
	//		//for i, container := range totalContainer {
	//		//	for key, val := range container {
	//		//		lock.Lock()
	//		//		fmt.Printf("algBig %d %s new(%d) old(%d)\n", i, val, newCost.Microseconds(), oldCost.Microseconds())
	//		//		delete(container, key)
	//		//		lock.Unlock()
	//		//	}
	//		//}
	//
	//	} //[Joker ♠2 ♠3 ♠4 ♠5 ♠6 ♠7] [♠6 ♣6 ♦6] [♠7 ♠8 Joker ♠10 ♠J] [♥4 ♥5 ♥6]
	//
	//	if !equal {
	//		//	fmt.Println("ERR!")
	//		//	fmt.Println("ERR!")
	//		//	fmt.Println("ERR!")
	//		//	fmt.Println("ERR!")
	//		//	fmt.Println("ERR!")
	//		//	fmt.Println("ERR!")
	//		//	niuhe.LogDebug("algResInput %d  %s", okey, zpjson.JsonString(plain))
	//		//	niuhe.LogDebug("algResInput %d  %s", okey, zpjson.JsonString(plain))
	//		//	fmt.Printf("algRes(%t) new(%v)(%d)us  old(%v)(%d)us", equal, algRes, newCost, res, oldCost)
	//	}
	//
	//	if PrintLimitB.Can(100) {
	//		oldStat.PrintTimeCostStatis(utils.Stdout_Log)
	//		newStat.PrintTimeCostStatis(utils.Stdout_Log)
	//		for i, container := range totalContainer {
	//			clen := len(container)
	//			for key, val := range container {
	//				lock.Lock()
	//				niuhe.LogDebug("algBig %d %s new(%d) old(%d)\n", i, val, newCost.Microseconds(), oldCost.Microseconds())
	//				lock.Unlock()
	//				if clen >= 10 {
	//					delete(container, key)
	//				}
	//			}
	//
	//		}
	//	}
	//
	//	return res
	//}
}

func GetMaxCardGroupCgo(okey int, plain []rules.Card) [][]rules.Card {
	var cardVec = packCards(plain)
	defer DeleteIntVector(cardVec)
	var result = GetSolution(cardVec, okey)
	defer DeleteResult(result)
	return unpackCards(result.GetSolution())
}

//
//func GetBestCountGroup(cards []rules.Card) [][]rules.Card {
//	var plain = []rules.Card{}
//	var okey int
//	for _, c := range cards {
//		if c == rules.OkeyCard {
//			okey += 1
//		} else {
//			plain = append(plain, c)
//		}
//
//	}
//
//	if len(cards) == 0 {
//		emptyInputCnt++
//	}
//
//	//if Usego && Cc != nil {
//	//	//return GetMaxCardGroupCc2(okey, plain)
//	//	return GetMaxCardGroupCc(okey, plain)
//	//}
//
//	if zpenv.IsProd() {
//		return GetBestCountGroupCgo(okey, plain)
//	} else {
//		var res, algRes [][]rules.Card
//		var newCost, oldCost time.Duration
//		//sw := sync.WaitGroup{}
//		//
//		//sw.Add(1)
//		//go func() {
//		//	defer sw.Done()
//		start := time.Now()
//		res = GetBestCountGroupCgo(okey, plain)
//		end := time.Now()
//		oldCost = end.Sub(start)
//		oldStat.Record(oldCost)
//		//}()
//
//		//sw.Add(1)
//		//go func() {
//		//	defer sw.Done()
//		start1 := time.Now()
//		algRes = alg.GetBestCountGroup(okey, plain)
//		end1 := time.Now()
//		newCost = end1.Sub(start1)
//		newStat.Record(newCost)
//		//}()
//		//
//		//sw.Wait()
//
//		rules.SortCardsProject(res)
//		rules.SortCardsProject(algRes)
//
//		equal := rules.CardsEquals(rules.CardsProject(res).ToCards(), rules.CardsProject(algRes).ToCards())
//
//		if PrintLocal {
//			//fmt.Printf("algResBC(%t) new(%v)(%d)us  old(%v)(%d)us\n", equal, algRes, newCost, res, oldCost)
//			newStat.PrintTimeCostStatis(utils.Stdout_Print)
//			oldStat.PrintTimeCostStatis(utils.Stdout_Print)
//		}
//
//		if !equal {
//			niuhe.LogDebug("algResInputBC %d  %s", okey, zpjson.JsonString(plain))
//			niuhe.LogDebug("algResBC(%t) new(%v)(%d)us  old(%v)(%d)us", equal, algRes, newCost, res, oldCost)
//		}
//
//		if PrintLimitB.Can(100) {
//			newStat.PrintTimeCostStatis(utils.Stdout_Log)
//			oldStat.PrintTimeCostStatis(utils.Stdout_Log)
//		}
//
//		return res
//	}
//}

func GetBestCountGroupCgo(okey int, plain []rules.Card) [][]rules.Card {
	var cardVec = packCards(plain)
	defer DeleteIntVector(cardVec)
	var result = GetSolution(cardVec, okey)
	defer DeleteResult(result)
	return unpackCards(result.GetBestCountSolution())
}

func main() {
	fmt.Println("we work")
}
