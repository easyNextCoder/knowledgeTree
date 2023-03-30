package alga

func (self *AlgX) searchBomb() Projects {

	var res Projects
	for i := 0; i < self.bombLen; i++ {
		//cards := self.bombCards[i]
		ps, n := self.searchBombWork(i, 0)
		for j := 0; j < n; j++ {
			res.appendProject(ps[j])
		}
	}

	return res

}

func (self *AlgX) searchBombWork(cardsIndex int, jokerNum int) ([3]Project, int) {

	if self.bombCardsLen[cardsIndex] <= 0 {
		return [3]Project{}, 0
		panic("searchBomb input cards error")
	}

	if jokerNum == 1 {
		panic("jokerNum can't be 1")
		//return self.searchBombWorkWithJoker(cards, inputMap)
	} else if jokerNum == 2 {
		panic("searchBombWork error jokerLeft is 2")
	} else if jokerNum == 0 {
		//正常执行以下的代码
	} else {
		panic("searchBombWork error jokerLeft is 0")
	}

	colorFreq := make([]int16, 5)
	colorCnt := 0
	cnt := int16(0)

	for i := 0; i < self.bombCardsLen[cardsIndex]; i++ {
		c := self.bombCards[cardsIndex][i]
		cardNumAsBomb := self.states[CARD_OFFSET(&c)][1]
		if cardNumAsBomb > 0 {
			colorFreq[c.Second] = cardNumAsBomb
		}

	}

	for _, v := range colorFreq {
		if v > 0 {
			colorCnt++
			cnt += v
		}
	}

	//sort.Slice(cards, func(i, j int) bool {
	//	if cards[i].second == cards[j].second {
	//		return cards[i].first < cards[j].first
	//	} else if cards[i].second < cards[j].second {
	//		return true
	//	}
	//	return false
	//})

	cardVal := self.bombCards[cardsIndex][0].First
	//llog(-1, "bombColorFreq %v %d\n", colorFreq, cardVal)
	//fmt.Println("search Bomb mp", colorFreq, cardVal)

	if cnt == 8 {
		var a, b Project
		NewBombProjectAllInit(&a, cardVal)
		NewBombProjectAllInit(&b, cardVal)

		//e := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
		//f := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
		//cc = append(cc, e)
		//cc = append(cc, f)
		return [3]Project{a, b}, 2
	}
	//llog(-100, "searchBomb colorCnt %d cnt %d colorFreq %v cardVal %v cards %v\n", colorCnt, cnt, colorFreq, cardVal, cards)
	switch colorCnt { //那现在牌值相同的牌求得最多的炸弹个数和分数
	case 3:
		switch cnt { //todo 这里之前用的是len(cards)这样是错误的，因为cards在这种states改变的情况下并不会变化，只有提前处理joker或者搞顺子的A的时候会传指针进去改变原cards
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			dontHaveColor := 0
			for color, ct := range colorFreq {
				if color == 0 {
					continue
				}
				if ct == 0 {
					dontHaveColor = color
					break
				}
			}
			return [3]Project{NewBombProjectExceptOneColor(cardVal, int16(dontHaveColor))}, 1
		case 6:
			dontHaveColor := 0
			for color, ct := range colorFreq {
				if color == 0 {
					continue
				}
				if ct == 0 {
					dontHaveColor = color
					break
				}
			}

			return [3]Project{NewBombProjectExceptOneColor(cardVal, int16(dontHaveColor)), NewBombProjectExceptOneColor(cardVal, int16(dontHaveColor))}, 2

		default:
			panic("colorFreq 3 len(cards) not right")
		}
	case 4:
		switch cnt {
		case 4:
			fallthrough
		case 5:
			return [3]Project{NewBombProjectAll(cardVal)}, 1

		case 6: //6张牌可以拆成2个炸弹

			var np int
			var resx [3]Project
			for i := int16(1); i < 5; i++ {
				color := i
				cnt := colorFreq[i]
				if cnt == 1 {
					resx[np] = NewBombProjectExceptOneColor(cardVal, color)
					np++
				}
			}
			return resx, np
		case 7:
			var np int
			var resx [3]Project
			for i := int16(1); i < 5; i++ {
				color := i
				cnt := colorFreq[i]
				if cnt == 1 {
					resx[np] = NewBombProjectExceptOneColor(cardVal, color)
					np++
				}
			}
			resx[np] = NewBombProjectAll(cardVal)
			np++
			return resx, np

		case 8:
			return [3]Project{NewBombProjectAll(cardVal), NewBombProjectAll(cardVal)}, 2

		}
	}

	return [3]Project{}, 0
}
