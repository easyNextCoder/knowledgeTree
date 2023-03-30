package alga

import (
	_ "net/http/pprof"
)

func (self *AlgX) useTwoColorsOneJokerAsProject(ca, cb, val int16, index int) {
	//return
	thisBombColorMap := self.bombRepeatMap[index]
	thisBombCards := self.bombCards[index]

	//尝试将这个joker组成的project放到最大的jokerproj中
	cc := int16(0)
	for color := int16(1); color < 5; color++ {
		if color != ca && color != cb {
			cc = color
			break
		}
	}

	carda, cardb := Card{val, ca}, Card{val, cb}

	statesaBak := self.states[CARD_OFFSET(&carda)]
	statesbBak := self.states[CARD_OFFSET(&cardb)]

	if self.states[CARD_OFFSET(&carda)][0] < 0 {
		panic("states[CARD_OFFSET(&ca)][0] err")
	}

	if self.states[CARD_OFFSET(&cardb)][0] < 0 {
		panic("states[CARD_OFFSET(&cb)][0] err")
	}

	jokerProj := NewBombProject(val, ca, cb, cc)
	jokerProj.tagIndex[jokerProj.tagLen] = cc
	jokerProj.tagLen++

	//清楚joker占用的牌，接着跳转到普通的dfs中进行剩下的递归
	colorMapRemoveCard(&thisBombColorMap, carda)
	colorMapRemoveCard(&thisBombColorMap, cardb)

	oldLen := self.bombCardsLen[index]
	thisBombCardsNew, newLen := colorMapToCards(&thisBombColorMap, val)
	self.bombCards[index] = thisBombCardsNew
	self.bombCardsLen[index] = newLen

	self.bombRepeatMap[index] = thisBombColorMap
	if self.jokerLeft <= 0 {
		panic("self.jokerLeft must bigger than 0")
	}
	self.jokerLeft--

	if self.states[CARD_OFFSET(&carda)][1] > 0 {
		self.states[CARD_OFFSET(&carda)][1]--
	} else if self.states[CARD_OFFSET(&carda)][0] > 0 {
		self.states[CARD_OFFSET(&carda)][0]--
	} else {
		panic("")
	}

	if self.states[CARD_OFFSET(&cardb)][1] > 0 {
		self.states[CARD_OFFSET(&cardb)][1]--
	} else if self.states[CARD_OFFSET(&cardb)][0] > 0 {
		self.states[CARD_OFFSET(&cardb)][0]--
	} else {
		panic("")
	}

	self.bombTakeAheadMaxScoreProjects.appendProject(jokerProj)

	self.dfs(index)

	self.bombCards[index] = thisBombCards
	self.bombCardsLen[index] = oldLen
	self.bombRepeatMap[index].colors[carda.Second]++
	self.bombRepeatMap[index].cnt++
	if self.bombRepeatMap[index].colors[carda.Second] == 1 {
		self.bombRepeatMap[index].colorCnt++
	}

	self.bombRepeatMap[index].colors[cardb.Second]++
	self.bombRepeatMap[index].cnt++
	if self.bombRepeatMap[index].colors[cardb.Second] == 1 {
		self.bombRepeatMap[index].colorCnt++
	}

	self.bombTakeAheadMaxScoreProjects.projLen--
	self.bombTakeAheadMaxScoreProjects.updateScoreLen()
	self.jokerLeft++

	self.states[CARD_OFFSET(&carda)] = statesaBak
	self.states[CARD_OFFSET(&cardb)] = statesbBak

}

func (self *AlgX) useThreeColorsOneJokerAsProject(ca, cb, cc, val int16, index int) {
	//return
	thisBombColorMap := self.bombRepeatMap[index]
	thisBombCards := self.bombCards[index]

	var cd int16

	for color := int16(1); color < 5; color++ {
		if color != ca && color != cb && color != cc {
			cd = color
			break
		}
	}

	carda, cardb, cardc := Card{val, ca}, Card{val, cb}, Card{val, cc}

	statesaBak := self.states[CARD_OFFSET(&carda)]
	statesbBak := self.states[CARD_OFFSET(&cardb)]
	statescBak := self.states[CARD_OFFSET(&cardc)]

	if self.states[CARD_OFFSET(&carda)][0] < 0 {
		panic("states[CARD_OFFSET(&ca)][0] err")
	}

	if self.states[CARD_OFFSET(&cardb)][0] < 0 {
		panic("states[CARD_OFFSET(&cb)][0] err")
	}
	if self.states[CARD_OFFSET(&cardc)][0] < 0 {
		panic("states[CARD_OFFSET(&cb)][0] err")
	}

	//尝试将这个joker组成的project放到最大的jokerproj中

	jokerProj := NewBombProjectAll(val)
	jokerProj.tagIndex[jokerProj.tagLen] = cd
	jokerProj.tagLen++
	//fmt.Println("carda, cardb", carda, cardb, cardc)

	//清楚joker占用的牌，接着跳转到普通的dfs中进行剩下的递归
	colorMapRemoveCard(&thisBombColorMap, carda)
	colorMapRemoveCard(&thisBombColorMap, cardb)
	colorMapRemoveCard(&thisBombColorMap, cardc)
	oldLen := self.bombCardsLen[index]
	thisBombCardsNew, newLen := colorMapToCards(&thisBombColorMap, val)
	self.bombCards[index] = thisBombCardsNew
	self.bombRepeatMap[index] = thisBombColorMap
	self.bombCardsLen[index] = newLen
	if self.jokerLeft <= 0 {
		panic("self.jokerLeft must bigger than 0")
	}
	self.jokerLeft--

	if self.states[CARD_OFFSET(&carda)][1] > 0 {
		self.states[CARD_OFFSET(&carda)][1]--
	} else if self.states[CARD_OFFSET(&carda)][0] > 0 {
		self.states[CARD_OFFSET(&carda)][0]--
	} else {
		panic("")
	}

	if self.states[CARD_OFFSET(&cardb)][1] > 0 {
		self.states[CARD_OFFSET(&cardb)][1]--
	} else if self.states[CARD_OFFSET(&cardb)][0] > 0 {
		self.states[CARD_OFFSET(&cardb)][0]--
	} else {
		panic("")
	}

	if self.states[CARD_OFFSET(&cardc)][1] > 0 {
		self.states[CARD_OFFSET(&cardc)][1]--
	} else if self.states[CARD_OFFSET(&cardc)][0] > 0 {
		self.states[CARD_OFFSET(&cardc)][0]--
	} else {
		panic("")
	}

	self.bombTakeAheadMaxScoreProjects.appendProject(jokerProj)

	self.dfs(index)

	self.bombCards[index] = thisBombCards
	self.bombCardsLen[index] = oldLen
	self.bombRepeatMap[index].colors[carda.Second]++
	self.bombRepeatMap[index].cnt++
	if self.bombRepeatMap[index].colors[carda.Second] == 1 {
		self.bombRepeatMap[index].colorCnt++
	}

	self.bombRepeatMap[index].colors[cardb.Second]++
	self.bombRepeatMap[index].cnt++
	if self.bombRepeatMap[index].colors[cardb.Second] == 1 {
		self.bombRepeatMap[index].colorCnt++
	}

	self.bombRepeatMap[index].colors[cardc.Second]++
	self.bombRepeatMap[index].cnt++
	if self.bombRepeatMap[index].colors[cardc.Second] == 1 {
		self.bombRepeatMap[index].colorCnt++
	}

	self.bombTakeAheadMaxScoreProjects.projLen--

	//thisBombCardsNew = nil
	self.bombTakeAheadMaxScoreProjects.updateScoreLen()
	self.jokerLeft++

	self.states[CARD_OFFSET(&carda)] = statesaBak
	self.states[CARD_OFFSET(&cardb)] = statesbBak
	self.states[CARD_OFFSET(&cardc)] = statescBak
}

func (self *AlgX) shouldGiveFlush(colors []int16, val int16) []int16 {

	newColors := make([]int16, 5)
	newColorsCnt := 0
	for i := 0; i < len(colors); i++ { //首先看四个中哪个给顺子之后能凑成顺子组合，joker就应该是去替代那个可以凑成顺子的那张牌
		card := Card{val, colors[i]}
		var nexts, nnexts, _ int16
		cardfBak := card.First

		card.First++
		if card.First <= 14 {
			if card.First == 14 {
				nexts = self.states[CARD_OFFSET(&Card{1, colors[i]})][0] + self.states[CARD_OFFSET(&Card{1, colors[i]})][1]
			} else {
				nexts = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			}

		}

		card.First++
		if card.First <= 14 {
			if card.First == 14 {
				nnexts = self.states[CARD_OFFSET(&Card{1, colors[i]})][0] + self.states[CARD_OFFSET(&Card{1, colors[i]})][1]
			} else {
				nnexts = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			}

			//nnextsJokers = minInt16(nnexts, int16(self.jokerLeft-1))
		}

		card.First = cardfBak

		var before, bbefore, _ int16

		card.First--
		if card.First >= 1 {
			before = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
		}

		card.First--
		if card.First >= 1 {
			bbefore = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			//bbeforeJokers = minInt16(bbefore, int16(self.jokerLeft-1))
			//fmt.Println("this is val", val, card.First, bbefore, bbeforeJokers)
		}

		if before > 0 || bbefore > 0 || nexts > 0 || nnexts > 0 {
			newColors[colors[i]]++
			newColorsCnt++
		}

		card.First = cardfBak

	}

	if val == 1 {
		val = 14
		for i := 0; i < len(colors); i++ { //首先看四个中哪个给顺子之后能凑成顺子组合，joker就应该是去替代那个可以凑成顺子的那张牌
			card := Card{val, colors[i]}
			cardfBak := card.First

			var before, bbefore, bbeforeJokers int16

			card.First--
			if card.First >= 1 {
				before = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			}

			card.First--
			if card.First >= 1 {
				bbefore = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
				bbeforeJokers = minInt16(bbefore, int16(self.jokerLeft-1))

			}

			if before > 0 || bbeforeJokers > 0 {
				newColors[colors[i]]++
				newColorsCnt++
			}

			card.First = cardfBak

		}
		val = 1
	}
	return newColors
}

func (self *AlgX) MakeUpJokerBombFirst(index int) {
	//return
	thisBombColorMap := self.bombRepeatMap[index]
	thisBombCards := self.bombCards[index]

	if thisBombColorMap.colorCnt < 2 { //跳过处理下一个
		//fmt.Println("dfsOneJoker thisBombColorMap < 2")
		panic("handleOneJoker should not handle this")
		return
	}

	//case 2: //这里需要处理而在dfs中就不能再处理了，应该直接全部给顺子

	//为Joker牌新增的

	val := thisBombCards[0].First
	//colors := make([]int16, 0, 4)
	var colors [5]int16
	var differentColorsCnt int
	var colorsCnt [5]int16
	var totalColorCnt int16

	for color, cnt := range thisBombColorMap.colors {
		if cnt > 0 {
			colors[differentColorsCnt] = int16(color)
			differentColorsCnt++
			colorsCnt[color] = cnt
			totalColorCnt += cnt
		}
	}

	//	//[case]2023/03/26 05:13:13 joker_test.go:257: 	per us:  52  最大分: 62   {cards[[joker {2,3} {3,3}] [{12,3} joker {14,3}] [{5,1} {5,2} {5,4}]] score:62 len:9} 最大长度 {cards[[{2,3} {3,3} joker] [{12,3} joker {14,3}] [{5,1} {5,2} {5,4}]] score:55 len:9} dfsTimes: 0
	//	//	[case]2023/03/26 05:13:13 joker_test.go:258: 	per us: 122  最大分: 64   [[Joker ♠A ♣A] [Joker ♣2 ♣3] [♠5 ♥5 ♦5]] 524 1220
	//	//	[case]2023/03/26 05:13:13 joker_test.go:259: [{1,1} {5,1} {7,1} {5,2} {13,2} {1,3} {2,3} {3,3} {12,3} {5,4} {11,4}]
	var newColors [5]int16
	newColorsCnt := 0
	for i := 0; i < differentColorsCnt; i++ { //首先看四个中哪个给顺子之后能凑成顺子组合，joker就应该是去替代那个可以凑成顺子的那张牌
		card := Card{val, colors[i]}
		var nexts, nnexts, nnextsJokers int16
		cardfBak := card.First

		card.First++
		if card.First <= 14 {
			if card.First == 14 {
				nexts = self.states[CARD_OFFSET(&Card{1, colors[i]})][0] + self.states[CARD_OFFSET(&Card{1, colors[i]})][1]
			} else {
				nexts = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			}

		}

		card.First++
		if card.First <= 14 {
			if card.First == 14 {
				nnexts = self.states[CARD_OFFSET(&Card{1, colors[i]})][0] + self.states[CARD_OFFSET(&Card{1, colors[i]})][1]
			} else {
				nnexts = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			}

			nnextsJokers = minInt16(nnexts, int16(self.jokerLeft-1))
		}

		card.First = cardfBak

		var before, bbefore, bbeforeJokers int16

		card.First--
		if card.First >= 1 {
			before = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
		}

		card.First--
		if card.First >= 1 {
			bbefore = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			bbeforeJokers = minInt16(bbefore, int16(self.jokerLeft-1))

		}

		if before > 0 || bbeforeJokers > 0 || nexts > 0 || nnextsJokers > 0 {
			newColors[colors[i]]++
			newColorsCnt++
		}

		card.First = cardfBak

	}

	if val == 1 {
		val = 14
		for i := 0; i < differentColorsCnt; i++ { //首先看四个中哪个给顺子之后能凑成顺子组合，joker就应该是去替代那个可以凑成顺子的那张牌
			card := Card{val, colors[i]}
			cardfBak := card.First

			var before, bbefore, bbeforeJokers int16

			card.First--
			if card.First >= 1 {
				before = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
			}

			card.First--
			if card.First >= 1 {
				bbefore = self.states[CARD_OFFSET(&card)][0] + self.states[CARD_OFFSET(&card)][1]
				bbeforeJokers = minInt16(bbefore, int16(self.jokerLeft-1))

			}

			if before > 0 || bbeforeJokers > 0 {
				newColors[colors[i]]++
				newColorsCnt++
			}

			card.First = cardfBak

		}
		val = 1
	}

	//fmt.Println(newColors, val)

	switch differentColorsCnt {
	case 2: //两张的牌可以加一个joker凑成炸弹
		self.useTwoColorsOneJokerAsProject(colors[0], colors[1], val, index)
	case 3:

		if newColorsCnt > 0 { //有顺子用到某个颜色的牌，则替换一下试试看
			for i := 1; i < len(newColors); i++ {

				if newColors[i] <= 0 {
					continue
				}

				var c2 [5]int16
				var c2Cnt int
				for j := 0; j < differentColorsCnt; j++ {
					if colors[j] != int16(i) {
						c2[c2Cnt] = colors[j]
						c2Cnt++
					}
				}

				self.useTwoColorsOneJokerAsProject(c2[0], c2[1], val, index)
			}
		}

		//炸弹牌不被顺子争用的情况

		self.useThreeColorsOneJokerAsProject(colors[0], colors[1], colors[2], val, index) //先拿joker凑成4个的炸弹

		var one, two [5]int16
		var oneCnt, twoCnt int
		for i := 0; i < differentColorsCnt; i++ {
			if colorsCnt[colors[i]] == 1 {
				one[oneCnt] = colors[i]
				oneCnt++

			} else {
				two[twoCnt] = colors[i]
				twoCnt++
			}
		}

		switch totalColorCnt {
		//case3://单独的三除了用一个joker凑成4张之外其他没有意义
		case 4:
			self.useTwoColorsOneJokerAsProject(two[0], one[0], val, index)

		case 5:
			self.useTwoColorsOneJokerAsProject(two[0], two[1], val, index)

		case 6: //凑整一个4张牌的炸弹

		}

	case 4: //
		switch totalColorCnt {
		case 4: //加两个joker既可以凑成2个三炸弹
			fallthrough
		case 5: //加一个joker就可以凑成2个3炸弹
			//即使顺子没用，也可以加joker组成炸弹
			var one, two [5]int16
			var oneCnt, twoCnt int
			for i := 0; i < differentColorsCnt; i++ {
				if colorsCnt[colors[i]] == 1 {
					one[oneCnt] = colors[i]
					oneCnt++

				} else {
					two[twoCnt] = colors[i]
					twoCnt++

				}
			}

			a, b := one[0], one[1]
			if twoCnt > 0 { //为了兼容4-5的情况
				a = two[0]
			}

			self.useTwoColorsOneJokerAsProject(a, b, val, index)

			//看看是否值得将4炸的组合中省下一个
			if newColorsCnt > 0 {
				for i := 1; i < len(newColors); i++ { //拿joker替换其中一个,然后组成4炸弹即可
					//fmt.Println("vvvvv", val, newColors)
					if newColors[i] <= 0 {
						continue
					}

					var c2 [5]int16
					var c2Cnt int

					for j := 0; j < differentColorsCnt; j++ {
						if colors[j] != int16(i) {
							c2[c2Cnt] = colors[j]
							c2Cnt++
						}
					}

					self.useThreeColorsOneJokerAsProject(c2[0], c2[1], c2[2], val, index)
				}
			}

		case 8: //可以凑成333炸弹
			fallthrough
		case 6: //尽量去组成43炸弹
			//如果顺子用不到就让joker凑炸弹
			if newColorsCnt == 0 {
				var c1, c2 [5]int16
				var c1Cnt, c2Cnt int
				for i := 1; i < 5; i++ {

					if colorsCnt[i] == 2 {
						c2[c2Cnt] = int16(i)
						c2Cnt++
					} else {

						c1[c1Cnt] = int16(i)
						c1Cnt++
					}
				}
				var c3 int16 //2 2 2 2//2 2 1 1
				if c2Cnt > 2 {
					c3 = c2[2]
				} else {
					c3 = c1[0]
				}
				self.useTwoColorsOneJokerAsProject(c2[0], c2[1], val, index)       // [6 6 6 6][joker 6 6]
				self.useThreeColorsOneJokerAsProject(c2[0], c2[1], c3, val, index) //[joker 6 6 6][joker 6 6 6]

			} else { //如果顺子能用到
				//fmt.Println("newColorCnt", newColorsCnt)
				for i := 0; i < differentColorsCnt; i++ {
					ca := colors[i]
					cs := getSubColor(ca)
					c1, c2, c3 := cs[0], cs[1], cs[2]
					self.useThreeColorsOneJokerAsProject(c1, c2, c3, val, index)
				}
				for i := 0; i < differentColorsCnt; i++ {
					for j := i + 1; j < differentColorsCnt; j++ {
						c1, c2 := colors[i], colors[j]
						self.useTwoColorsOneJokerAsProject(c1, c2, val, index)
					}
				}
			}
		case 7: //可以凑成44炸弹
			//if newColorsCnt == 0 { //如果顺子用不到就让joker凑炸弹
			//	fmt.Println(newColorsCnt)
			//	var c2 []int16
			//	for i := 1; i < 5; i++ {
			//
			//		if colorsCnt[i] == 2 {
			//			c2 = append(c2, int16(i))
			//		}
			//	}
			//	self.useThreeColorsOneJokerAsProject(c2[0], c2[1], c2[2], val, index)
			//} else { //如果顺子能用到
			for i := 0; i < differentColorsCnt; i++ {
				ca := colors[i]
				cs := getSubColor(ca)
				c1, c2, c3 := cs[0], cs[1], cs[2]
				self.useThreeColorsOneJokerAsProject(c1, c2, c3, val, index)
			}
			for i := 0; i < differentColorsCnt; i++ {
				for j := i + 1; j < differentColorsCnt; j++ {
					c1, c2 := colors[i], colors[j]
					//fmt.Println(c1, c2)
					self.useTwoColorsOneJokerAsProject(c1, c2, val, index)

				}
			}
			//}

		}

	}

}

func minInt16(a ...int16) int16 {
	m := int16(30000)
	//fmt.Println(a)
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func min(a ...int) int {
	m := 999999
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func splitSamePointCardsToBombs(colorFreq [5]int16) (x, y int) { //x<=y
	a, b := 0, 0
	for _, n := range colorFreq {
		if n >= 2 {
			a++
			b++
		} else if n == 1 {
			if a > b {
				b++
			} else {
				a++
			}
		} else {
			//panic("n == 0")
		}
	}
	if a > b {
		return b, a
	} else {
		return a, b
	}

}

func getSubColor(c int16) [3]int16 {
	var i int16
	var res [3]int16
	var resCnt int
	for i = 1; i <= 4; i++ {
		if c == i {
			continue
		}
		res[resCnt] = i
		resCnt++
	}
	return res
}
