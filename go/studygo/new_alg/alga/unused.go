package alga

//func (self *AlgX) fixedStateSearch() [][]Card {
//
//	var cc [][]Card
//	switch self.jokerLeft {
//	case 0:
//
//		cc = append(cc, self.searchFlush(flushCards, states, 0)...)
//		cc = append(cc, self.searchBomb(bombCards, states, 0)...)
//
//	case 1:
//
//		////fmt.Println("一个joker")
//		//var finalCC [][]Card
//		//score := 0
//		//
//		//cd := self.searchFlush(flushCards, states, 1)
//		//cd = append(cd, self.searchBomb(bombCards, states, 0)...)
//		//scored := scoreProjects(cd)
//		//fmt.Println("fixedStateSearch cd", cd)
//		////fmt.Println("cd", cd)
//		//if scored > score {
//		//	finalCC = cd
//		//	score = scored
//		//}
//		//
//		//ce := self.searchFlush(flushCards, states, 0)
//		//fmt.Println("fixedStateSearch ce", ce)
//		//ce = append(ce, self.searchBomb(bombCards, states, 1)...)
//		//fmt.Println("fixedStateSearch ce1", ce)
//		//scoree := scoreProjects(ce)
//		//if scoree > score {
//		//	finalCC = ce
//		//	score = scoree
//		//}
//		//cc = finalCC
//
//	}
//
//	return cc
//}

//作为一个中间层插入到dfs路径过程中
//func (self *AlgX) OneJokerMiddleNode2(index int) {
//
//	thisBombColorMap := bombRepeatMap[index]
//	thisBombCards := bombCards[index]
//
//	if len(thisBombColorMap) < 2 { //跳过处理下一个
//		//fmt.Println("dfsOneJoker thisBombColorMap < 2")
//		panic("handleOneJoker should not handle this")
//		return
//	}
//
//	//case 2: //这里需要处理而在dfs中就不能再处理了，应该直接全部给顺子
//
//	//为Joker牌新增的
//
//	val := thisBombCards[0].first
//	colors := []int16{}
//	for color, _ := range thisBombColorMap {
//		colors = append(colors, color)
//	}
//
//	for i := 0; i < len(colors); i++ { //如果是三种颜色就是有三个分支，如果是四种颜色就有6个分支
//		for j := i + 1; j < len(colors); j++ { //任意取两张和joker凑成炸弹
//			ca, cb := colors[i], colors[j]
//			carda, cardb := Card{val, ca}, Card{val, cb}
//
//			//每迭代一层的话要首先要进入到这里
//			states[CARD_OFFSET(&carda)][0]--
//			states[CARD_OFFSET(&cardb)][0]--
//
//			//fmt.Println("states cards", carda, cardb, states[CARD_OFFSET(&carda)][0], states[CARD_OFFSET(&cardb)][0])
//
//			if states[CARD_OFFSET(&carda)][0] < 0 {
//				panic("states[CARD_OFFSET(&ca)][0] err")
//			}
//
//			if states[CARD_OFFSET(&cardb)][0] < 0 {
//				panic("states[CARD_OFFSET(&cb)][0] err")
//			}
//
//			//尝试将这个joker组成的project放到最大的jokerproj中
//			cc := int16(0)
//			for color := int16(1); color < 5; color++ {
//				if color != ca && color != cb {
//					cc = color
//					break
//				}
//			}
//
//			cardc := Card{val, cc}
//			jokerProj := []Card{carda, cardb, cardc}
//			//fmt.Println("carda, cardb", carda, cardb, cardc)
//			var bak ProjectWithOneJoker
//			bak.jokerRepresent = maxOneProjectWithOneJoker.jokerRepresent
//			bak.proj = maxOneProjectWithOneJoker.proj
//			bak.score = maxOneProjectWithOneJoker.score
//
//			self.AddToMaxJokerProj(jokerProj, cardc)
//
//			//清楚joker占用的牌，接着跳转到普通的dfs中进行剩下的递归
//			colorMapRemoveCard(thisBombColorMap, carda)
//			colorMapRemoveCard(thisBombColorMap, cardb)
//			thisBombCardsNew := colorMapToCards(thisBombColorMap, val)
//			bombCards[index] = thisBombCardsNew
//			bombRepeatMap[index] = thisBombColorMap
//			self.haveProcessedInOneJokerMiddleNode = true
//			if self.jokerLeft <= 0 {
//				panic("self.jokerLeft must bigger than 0")
//			}
//			self.jokerLeft--
//			self.dfs(index)
//			bombCards[index] = thisBombCards
//			bombRepeatMap[index][carda.second]++
//			bombRepeatMap[index][cardb.second]++
//
//			states[CARD_OFFSET(&carda)][0]++
//			states[CARD_OFFSET(&cardb)][0]++
//			maxOneProjectWithOneJoker.jokerRepresent = bak.jokerRepresent
//			maxOneProjectWithOneJoker.proj = bak.proj
//			maxOneProjectWithOneJoker.score = bak.score
//			self.jokerLeft++
//		}
//	}
//}
