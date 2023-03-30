package alga

func (self *AlgX) flushWillUse() {

}

func (self *AlgX) dfs(index int) {

	if index == self.bombLen { //todo check 进入到这里所有的可能产生冲突的牌状态已经定了，后边可根据state的指引分别求顺子和炸弹最大即可

		bps := self.searchBomb()
		self.searchFlush()

		twoJokerInOneFlush := false

		for i := 0; i < self.tmpMaxScoreProjects.projLen; i++ {
			if self.tmpMaxScoreProjects.projs[i].tagLen > 1 {
				if !self.tmpMaxScoreProjects.FlushTrySplit(i) {
					twoJokerInOneFlush = true
					break
				}
			}
		}

		newScore := bps.score + self.tmpMaxScoreProjects.score + self.bombTakeAheadMaxScoreProjects.score
		newProjectsNum := bps.projLen + self.tmpMaxScoreProjects.projLen + self.bombTakeAheadMaxScoreProjects.projLen

		if (newScore > self.maxScoreProjects.score || (newScore == self.maxScoreProjects.score && newProjectsNum < self.maxScoreProjects.projLen)) && !twoJokerInOneFlush {

			self.maxScoreProjects = self.tmpMaxScoreProjects

			if bps.score > 0 {
				self.maxScoreProjects.merge(&bps)
			}

			if self.bombTakeAheadMaxScoreProjects.score > 0 {
				self.maxScoreProjects.merge(&self.bombTakeAheadMaxScoreProjects)
			}
		}

		twoJokerInOneFlush = false

		for i := 0; i < self.tmpMaxLenProjects.projLen; i++ {
			if self.tmpMaxLenProjects.projs[i].tagLen > 1 {
				if !self.tmpMaxLenProjects.FlushTrySplit(i) {
					twoJokerInOneFlush = true
					break
				}
			}
		}

		newLen := bps.len + self.tmpMaxLenProjects.len + self.bombTakeAheadMaxScoreProjects.len
		newProjectsNum = bps.projLen + self.tmpMaxLenProjects.projLen + self.bombTakeAheadMaxScoreProjects.projLen

		if (newLen > self.maxLenProjects.len || (newLen == self.maxLenProjects.len && newProjectsNum < self.maxLenProjects.projLen)) && !twoJokerInOneFlush {

			self.maxLenProjects = self.tmpMaxLenProjects

			if bps.len > 0 {
				self.maxLenProjects.merge(&bps)
			}

			if self.bombTakeAheadMaxScoreProjects.score > 0 {
				self.maxLenProjects.merge(&self.bombTakeAheadMaxScoreProjects)
			}
		}

		if self.maxScoreProjects.len == self.maxLenProjects.len && self.maxScoreProjects.score > self.maxLenProjects.score {
			self.maxLenProjects = self.maxScoreProjects
		}

		return
	}

	thisBombColorMap := self.bombRepeatMap[index]
	thisBombCards := self.bombCards[index]

	//log("dfs() 处理bomb jokerLeft %d thisBombCards %v ", self.jokerLeft, thisBombCards)

	if self.jokerLeft == 0 {
		if thisBombColorMap.colorCnt < 3 { //没有joker且同数字牌数量无法组成炸弹，直接跳向下一个炸弹的选择
			//log("index %d bombCardsColorNum %d %+v\n", index, len(thisBombColorMap), thisBombColorMap)
			self.dfs(index + 1) //这里没办法给炸弹了，那就全部给顺子吧
			return
		}
	} else if self.jokerLeft == 1 || self.jokerLeft == 2 {
		if thisBombColorMap.colorCnt < 2 { //有joker牌，但是同点数花色牌只有两张，也无法组成炸弹
			//log("index %d bombCardsColorNum %d %+v\n", index, len(thisBombColorMap), thisBombColorMap)
			self.dfs(index + 1) //这里没办法给炸弹了，那就全部给顺子吧
			return
		}
	} else {
		panic("self.jokerLeft not legal")
	}
	//llog(-100, "dfs index %d thisBombColorMap %+v %+v\n", index, thisBombColorMap, thisBombCards)
	switch thisBombColorMap.colorCnt {

	case 2: //这种情况发生在有joker的情况下 两种情况1.拿两张给炸弹剩下的给顺子 2.把所有的牌都给顺子
		//都给顺子
		cardVal := self.bombCards[index][0].First

		for color, cnt := range thisBombColorMap.colors {

			if cnt == 0 {
				continue
			}

			if cnt < 0 {
				panic("bomb neg")
			}

			code := CARD_OFFSET(&Card{First: cardVal, Second: int16(color)})
			self.states[code][0] = cnt
			self.states[code][1] = 0
		}
		self.dfs(index + 1)

		//llog(-100, "hhhh  %+v\n", thisBombColorMap)

		if self.jokerLeft > 0 {

			for i := 0; i < self.bombCardsLen[index]; i++ {
				c := thisBombCards[i]

				cnt := thisBombColorMap.colors[c.Second]
				self.states[CARD_OFFSET(&c)][0] = cnt //给顺子
				self.states[CARD_OFFSET(&c)][1] = 0
			}
			self.MakeUpJokerBombFirst(index)
		}

	case 3:
		switch self.bombCardsLen[index] {
		case 3: //在无joker的情况下只有两种情况，在有joker的情况下又多了3种情况
			////fmt.Println("thisBombColorMap thisBombCards:", thisBombColorMap, thisBombCards)

			for i := 0; i < self.bombCardsLen[index]; i++ {
				c := thisBombCards[i]
				cnt := thisBombColorMap.colors[c.Second]
				self.states[CARD_OFFSET(&c)][0] = cnt //给顺子
				self.states[CARD_OFFSET(&c)][1] = 0
			}
			self.dfs(index + 1)
			//fmt.Println("都给炸弹!")
			for i := 0; i < self.bombCardsLen[index]; i++ {
				c := thisBombCards[i]
				cnt := thisBombColorMap.colors[c.Second]
				self.states[CARD_OFFSET(&c)][0] = 0
				self.states[CARD_OFFSET(&c)][1] = cnt //把这三张牌给炸弹
			}
			self.dfs(index + 1)

			if self.jokerLeft > 0 { //任选两个组成炸弹，共三种情况
				for i := 0; i < self.bombCardsLen[index]; i++ {
					c := thisBombCards[i]
					cnt := thisBombColorMap.colors[c.Second]
					self.states[CARD_OFFSET(&c)][0] = cnt //给顺子
					self.states[CARD_OFFSET(&c)][1] = 0
				}

				if self.jokerLeft > 0 {
					self.MakeUpJokerBombFirst(index) //把本层预处理之后再回到本层
				}

			}

		case 4: //共2种情况
			//把其中一个颜色有两张的拿一张给顺子
			cardVal := self.bombCards[index][0].First
			for color, cnt := range thisBombColorMap.colors {

				if cnt == 0 {
					continue
				}

				if cnt < 0 {
					panic("bomb neg")
				}

				code := CARD_OFFSET(&Card{First: cardVal, Second: int16(color)})
				if cnt == 2 {
					self.states[code][0] = 1
					self.states[code][1] = 1
				} else {
					self.states[code][0] = 0
					self.states[code][1] = 1
				}
			}
			self.dfs(index + 1)
			//把所有的牌都给顺子
			for color, cnt := range thisBombColorMap.colors {

				if cnt == 0 {
					continue
				}

				if cnt < 0 {
					panic("bomb neg")
				}

				code := CARD_OFFSET(&Card{First: cardVal, Second: int16(color)})
				self.states[code][0] = cnt
				self.states[code][1] = 0
			}
			self.dfs(index + 1)

			if self.jokerLeft > 0 { //任选两个组成炸弹，共三种情况
				for color, cnt := range thisBombColorMap.colors {

					if cnt == 0 {
						continue
					}

					if cnt < 0 {
						panic("bomb neg")
					}

					code := CARD_OFFSET(&Card{First: cardVal, Second: int16(color)})
					self.states[code][0] = cnt
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}

		case 5: //共2种情况

			cardVal := self.bombCards[index][0].First

			//第一种情况：将有两张牌花色的牌各给一张给顺子，炸弹仍然可用
			var card Card
			for c1, cnt := range thisBombColorMap.colors {
				if cnt == 0 {
					continue
				}

				if cnt < 0 {
					panic("bomb neg")
				}

				card = Card{cardVal, int16(c1)}
				code := CARD_OFFSET(&card)
				self.states[code][0] = cnt - 1
				self.states[code][1] = 1 //一种颜色给炸弹一个
			}
			self.dfs(index + 1)
			//第二种情况，无法组成炸弹所有的牌全部交给顺子去用
			for c1, cnt := range thisBombColorMap.colors {
				if cnt == 0 {
					continue
				}

				if cnt < 0 {
					panic("bomb neg")
				}
				card = Card{cardVal, int16(c1)}
				code := CARD_OFFSET(&card)
				self.states[code][0] = cnt //所有的都给顺子
				self.states[code][1] = 0
			}
			self.dfs(index + 1)

			if self.jokerLeft > 0 { //任选两个组成炸弹，共三种情况
				for c1, cnt := range thisBombColorMap.colors {

					if cnt == 0 {
						continue
					}

					if cnt < 0 {
						panic("bomb neg")
					}

					card = Card{cardVal, int16(c1)}
					code := CARD_OFFSET(&card)
					self.states[code][0] = cnt //所有的都给顺子
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}

		case 6:
			fallthrough
		case 7: //无joker的情况下最多三种情况
			//只能组成一组炸弹无法组成更多的炸弹
			cardVal := self.bombCards[index][0].First
			colors2 := [3]int16{}
			cnt := 0
			for k, v := range thisBombColorMap.colors {
				if v == int16(2) {
					colors2[cnt] = int16(k)
					cnt++
				}
			}
			//把所有的炸弹都保留一张都不给顺子,刚好可以组成两个炸弹
			if cnt == 3 {
				var card Card
				for c1, cnt1 := range thisBombColorMap.colors {

					if cnt1 == 0 {
						continue
					}

					if cnt1 < 0 {
						panic("bomb neg")
					}

					card = Card{cardVal, int16(c1)}
					code := CARD_OFFSET(&card)
					self.states[code][0] = cnt1 - 2 //处理3种颜色7张牌的情况
					self.states[code][1] = 2        //都给炸弹可以组成两个炸弹
					//llog(-1, "use here %v, %v\n", index, self.states[code])
				}
				self.dfs(index + 1)

			}

			//保留一个炸弹所有剩余的牌交给顺子去使用
			var card Card
			for c1, cnt1 := range thisBombColorMap.colors {

				if cnt1 == 0 {
					continue
				}

				if cnt1 < 0 {
					panic("bomb neg")
				}

				card = Card{cardVal, int16(c1)}
				code := CARD_OFFSET(&card)
				self.states[code][0] = cnt1 - 1
				self.states[code][1] = 1
			}
			self.dfs(index + 1)
			//炸弹不再保留包所有的牌交给顺子使用
			for c1, cnt1 := range thisBombColorMap.colors {

				if cnt1 == 0 {
					continue
				}

				if cnt1 < 0 {
					panic("bomb neg")
				}

				card = Card{cardVal, int16(c1)}
				code := CARD_OFFSET(&card)
				self.states[code][0] = cnt1
				self.states[code][1] = 0
			}
			self.dfs(index + 1)

			//处理有joker，利用joker加两张牌组成炸弹的

			if self.jokerLeft > 0 { //任选两个组成炸弹，共三种情况
				for c1, cnt1 := range thisBombColorMap.colors {

					if cnt1 == 0 {
						continue
					}

					if cnt1 < 0 {
						panic("bomb neg")
					}

					card = Card{cardVal, int16(c1)}
					code := CARD_OFFSET(&card)
					self.states[code][0] = cnt1
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}

		}

	case 4:
		switch self.bombCardsLen[index] {
		case 4: //共6中情况

			val := self.bombCards[index][0].First
			var color, c2 int16
			//全部都给顺子
			for color = 1; color < 5; color++ {
				code := CARD_OFFSET(&Card{First: val, Second: color})
				self.states[code][0] = 1
				self.states[code][1] = 0
			}
			self.dfs(index + 1)
			//把其中一个拆出来给顺子，剩下的给炸弹//4种情况

			//var colors []int16 = []int16{1, 2, 3, 4}

			//newColors := self.shouldGiveFlush(colors, val)
			//fmt.Println(newColors, val)
			//fmt.Printf("%s", *NewBombProjectAll(5))
			for color = 1; color < 5; color++ {

				//if newColors[int16(color)] <= 0 {
				//	continue
				//}

				for c2 = 1; c2 < 5; c2++ {
					code := CARD_OFFSET(&Card{First: val, Second: c2})

					if c2 == color {
						self.states[code][0] = 1
						self.states[code][1] = 0
					} else {
						self.states[code][0] = 0
						self.states[code][1] = 1
					}
				}
				self.dfs(index + 1)
			}
			//全部都给炸弹
			for color = 1; color < 5; color++ {
				code := CARD_OFFSET(&Card{First: val, Second: color})
				self.states[code][0] = 0
				self.states[code][1] = 1
			}
			////fmt.Println("全部都给炸弹 give all bomb")
			self.dfs(index + 1)

			if self.jokerLeft > 0 { //任选两个组成炸弹去跟joker组成炸弹共6种情况
				for color = 1; color < 5; color++ {
					code := CARD_OFFSET(&Card{First: val, Second: color})
					self.states[code][0] = 1
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}

		case 5: //共6种
			val := self.bombCards[index][0].First
			var code, color, c2 int16
			//组成一个4炸，剩下的给顺子
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				if thisBombColorMap.colors[color] == 2 {
					self.states[code][0] = 1
					self.states[code][1] = 1
				} else {
					self.states[code][0] = 0
					self.states[code][1] = 1
				}
			}
			//llog(-100, "66666666666666\n")

			self.dfs(index + 1)
			//组成一个3炸4种
			//var colors []int16 = []int16{1, 2, 3, 4}
			//
			//newColors := self.shouldGiveFlush(colors, val)

			for color = 1; color < 5; color++ {
				//if newColors[int16(color)] <= 0 { //给当前颜色的顺子也没用
				//	continue
				//}
				for c2 = 1; c2 < 5; c2++ {
					cnt := thisBombColorMap.colors[c2]
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					if c2 == color { //这个都给顺子
						self.states[code][0] = cnt
						self.states[code][1] = 0

					} else {
						self.states[code][0] = cnt - 1
						self.states[code][1] = 1
					}
				}
				//llog(-1, "8888888888")
				self.dfs(index + 1)
			}

			//全部都给顺子1种
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				cnt := thisBombColorMap.colors[color]
				self.states[code][0] = cnt
				self.states[code][1] = 0
			}
			//llog(-1, "99999999999999")
			self.dfs(index + 1)
			if self.jokerLeft > 0 { //任选两个组成炸弹去跟joker组成炸弹共6种情况
				for color = 1; color < 5; color++ {
					code = CARD_OFFSET(&Card{First: val, Second: color})
					cnt := thisBombColorMap.colors[color]
					self.states[code][0] = cnt
					self.states[code][1] = 0
					//llog(-1, "thisBombColorMap states %v val %d\n", states[code][1], states[code][1])
				}
				//llog(-1, "thisBombColorMap before jokerLeft %d MakeUpJokerBombFirst %v val %d\n", self.jokerLeft, thisBombColorMap, val)
				//llog(-1, "1010101010101")
				self.MakeUpJokerBombFirst(index)
			}
		case 6: //最多共7种

			val := self.bombCards[index][0].First
			var color, code, c2 int16

			//组成一个4炸1种
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				cnt := thisBombColorMap.colors[color]
				self.states[code][0] = cnt - 1
				self.states[code][1] = 1
			}

			self.dfs(index + 1)
			//组成一个3炸4种
			//var colors []int16 = []int16{1, 2, 3, 4}
			//newColors := self.shouldGiveFlush(colors, val)

			for color = 1; color < 5; color++ {
				//if newColors[int16(color)] <= 0 { //给当前颜色的顺子也没用
				//	continue
				//}
				for c2 = 1; c2 < 5; c2++ {
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					cnt := thisBombColorMap.colors[c2]
					if c2 == color { //这个都给顺子
						self.states[code][0] = cnt
						self.states[code][1] = 0
					} else {
						self.states[code][0] = cnt - 1
						self.states[code][1] = 1
					}
				}
				self.dfs(index + 1)
			}
			//全部都给顺子1
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				cnt := thisBombColorMap.colors[color]
				self.states[code][0] = cnt
				self.states[code][1] = 0
			}
			self.dfs(index + 1)

			a, b := splitSamePointCardsToBombs(thisBombColorMap.colors)
			if a == 3 && b == 3 { //可以组成两个三炸
				for color = 1; color < 5; color++ {
					code = CARD_OFFSET(&Card{First: val, Second: color})
					cnt := thisBombColorMap.colors[color]
					self.states[code][0] = 0
					self.states[code][1] = cnt //全部都给炸弹组成2个三炸
				}
				self.dfs(index + 1)
			}

			if self.jokerLeft > 0 { //任选两个组成炸弹去跟joker组成炸弹共6种情况
				for color = 1; color < 5; color++ {
					code = CARD_OFFSET(&Card{First: val, Second: color})
					cnt := thisBombColorMap.colors[color]
					self.states[code][0] = cnt
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}
		case 7: //共11种

			val := self.bombCards[index][0].First
			var color, c2, code int16
			//把所有的牌给顺子1种
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				cnt := thisBombColorMap.colors[color]
				self.states[code][0] = cnt
				self.states[code][1] = 0
			}
			self.dfs(index + 1)

			//组成4 3炸弹1种
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				if thisBombColorMap.colors[color] == 2 {
					self.states[code][0] = 0
					self.states[code][1] = 2
				} else {
					self.states[code][0] = 0
					self.states[code][1] = 1
				}
			}
			self.dfs(index + 1)

			//组成1个4的炸弹1种
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				if thisBombColorMap.colors[color] == 2 {
					self.states[code][0] = 1
					self.states[code][1] = 1
				} else {
					self.states[code][0] = 0
					self.states[code][1] = 1
				}
			}
			//llog(-100, "here!!  %v\n", thisBombColorMap.colors)
			self.dfs(index + 1)

			//组成33炸弹4种
			for color = 1; color < 5; color++ {
				for c2 = 1; c2 < 5; c2++ {
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					if c2 == color {
						self.states[code][0] = 1 //把其中一张牌给顺子用
						self.states[code][1] = thisBombColorMap.colors[c2] - 1
					} else {
						self.states[code][0] = 0
						self.states[code][1] = thisBombColorMap.colors[c2] //把所有的牌都给炸弹
					}
				}
				self.dfs(index + 1)
			}

			//组成一个3炸弹的牌剩下的都给顺子4种
			for color = 1; color < 5; color++ {
				for c2 = 1; c2 < 5; c2++ {
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					cnt := thisBombColorMap.colors[c2]
					if c2 == color {
						self.states[code][0] = cnt //把其中一种颜色的牌都给顺子用
						self.states[code][1] = 0
					} else { //把剩下的三张牌都拿出一张给炸弹，剩下的都给顺子
						self.states[code][0] = 1
						self.states[code][1] = thisBombColorMap.colors[c2] - 1
					}
				}
				self.dfs(index + 1)
			}

			//处理有joker的情况
			if self.jokerLeft > 0 { //任选两个组成炸弹去跟joker组成炸弹共6种情况
				for color = 1; color < 5; color++ {
					code = CARD_OFFSET(&Card{First: val, Second: color})
					cnt := thisBombColorMap.colors[color]
					self.states[code][0] = cnt
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}
		case 8: //共20种
			val := self.bombCards[index][0].First

			var code, color, c2 int16
			//组成两个4炸1
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				self.states[code][0] = 0
				self.states[code][1] = 2
			}
			self.dfs(index + 1)
			//组成一个4炸一个3炸4
			for color = 1; color < 5; color++ {
				for c2 = 1; c2 < 5; c2++ {
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					if c2 == color { //选中的颜色拿出一个给顺子剩下的都交给炸弹
						self.states[code][0] = 1
						self.states[code][1] = 1
					} else {
						self.states[code][0] = 0
						self.states[code][1] = 2
					}
				}
				self.dfs(index + 1)
			}
			//组成两个3炸10种（因为有可能两张牌都给顺子）
			//	1.把同一种颜色的两张牌都给顺子一共四种情况
			for color = 1; color < 5; color++ {
				for c2 = 1; c2 < 5; c2++ {
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					if c2 == color { //选中的颜色拿出一个给顺子剩下的都交给炸弹
						self.states[code][0] = 2
						self.states[code][1] = 0
					} else {
						self.states[code][0] = 0
						self.states[code][1] = 2
					}
				}
				self.dfs(index + 1)
			}
			//	2.把不同颜色的两张牌给顺子一共6种
			var c3 int16
			for color = 1; color < 5; color++ {
				for c2 = color + 1; c2 < 5; c2++ {
					for c3 = 1; c3 < 5; c3++ {
						code = CARD_OFFSET(&Card{First: val, Second: c3})
						if c3 == color || c3 == c2 { //拿出两种颜色的各一张给顺子剩下的都给炸弹
							self.states[code][0] = 1
							self.states[code][1] = 1
						} else {
							self.states[code][0] = 0
							self.states[code][1] = 2
						}
					}
					self.dfs(index + 1)
				}
			}
			//组成一个3炸中剩下的都给顺子4种
			for color = 1; color < 5; color++ {
				for c2 = 1; c2 < 5; c2++ {
					code = CARD_OFFSET(&Card{First: val, Second: c2})
					if c2 == color { //拿其中一个颜色的两张牌都给顺子，剩下的3种颜色的两边各一张
						self.states[code][0] = 2
						self.states[code][1] = 0
					} else {
						self.states[code][0] = 1
						self.states[code][1] = 1
					}
				}
				self.dfs(index + 1)
			}
			//全部都给顺子1种
			for color = 1; color < 5; color++ {
				code = CARD_OFFSET(&Card{First: val, Second: color})
				cnt := thisBombColorMap.colors[color]
				self.states[code][0] = cnt
				self.states[code][1] = 0
			}
			self.dfs(index + 1)
			//处理有joker的情况
			if self.jokerLeft > 0 { //任选两个组成炸弹去跟joker组成炸弹共6种情况
				for color = 1; color < 5; color++ {
					code = CARD_OFFSET(&Card{First: val, Second: color})
					cnt := thisBombColorMap.colors[color]
					self.states[code][0] = cnt
					self.states[code][1] = 0
				}
				self.MakeUpJokerBombFirst(index)
			}
		}
	}

}

