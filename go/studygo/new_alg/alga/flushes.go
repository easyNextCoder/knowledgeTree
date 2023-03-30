package alga

func (self *AlgX) searchFlush() {

	//defer func() {
	//	if recover() != nil {
	//		fmt.Println(self.hand_cards, self.tmpFlushs, self.cardsXY, self.xy)
	//	}
	//}()

	inputMap := self.states
	self.xy = XY{}
	for i := 0; i < 72; i++ {
		self.cardsXY[i] = 0
		self.freq[i] = 0
	}
	freq := self.freq
	for i := 0; i < self.flushLen; i++ {

		oneRow := self.flushCards[i]

		for j := 0; j < self.flushCardsLen[i]; j++ {
			v := oneRow[j]

			if inputMap[CARD_OFFSET(&v)][0]+inputMap[CARD_OFFSET(&v)][1] > 0 { //[0]代表给顺子
				if int(inputMap[CARD_OFFSET(&v)][0]) == 0 {
					continue
				}
				idx := v.First + (v.Second-1)*18
				val := int(inputMap[CARD_OFFSET(&v)][0])
				self.cardsXY[idx] = val
				self.freq[idx] = val
				continue
			}
		}
	}

	//fmt.Println("this is self.cardsXY:", self.cardsXY)

	self.tmpFlushs.clear()
	self.tmpMaxScoreProjects.clear()
	self.tmpMaxLenProjects.clear()

	self.travelsWork()

	self.freq = freq

}

func (self *AlgX) travelsWork() {

	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1}, {9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}}
	//var cards []Card = []Card{{4, 1}, {7, 1}, {9, 1}, {10, 1}, {11, 1}, {5, 2}, {6, 2}} // {12, 2}, {13, 2}} //, {2, 3}, {5, 3}, {11, 3}, {4, 4}}
	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}}
	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {1, 1}, {2, 1}, {3, 1}}
	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {12, 1}, {13, 1}}
	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {1, 1}, {2, 1}, {3, 1}, {12, 1}, {13, 1}, {12, 1}, {13, 1}}
	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {4, 1}}
	//var cards []Card = []Card{{6, 1}, {7, 1}, {8, 1}, {3, 2}, {4, 2}, {4, 3}, {6, 3}, {11, 3}, {12, 3}, {12, 3}, {3, 4}, {10, 4}, {13, 4}}
	//var cards []Card = []Card{{4, 1}, {7, 1}, {9, 1}, {10, 1}, {11, 1}, {5, 2}, {6, 2}}
	//var cards []Card = []Card{{8, 1}, {9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}} //这个是错的//
	//var cards []Card = []Card{{9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}}
	//var cards []Card = []Card{{7, 1}, {8, 1}, {9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}}
	//var cards []Card = []Card{{12, 1}, {13, 1}, {14, 1}}
	//var cards []Card = []Card{{6, 1}, {7, 1}, {8, 1}, {3, 2}, {4, 2}, {4, 3}, {6, 3}, {11, 3}, {12, 3}, {12, 3}, {3, 4}, {10, 4}, {13, 4}}

	//
	//for i := range cards {
	//	cards[i].First = cards[i].First + (cards[i].Second-1)*18
	//	self.cardsXY[cards[i].First]++
	//}

	self.moveToFirstCard()
	self.travels()
	//fmt.Println("search", self.tmpMaxScoreProjects)
	for i := 0; i < self.tmpMaxScoreProjects.projLen; i++ {
		self.tmpMaxScoreProjects.projs[i].max = self.tmpMaxScoreProjects.projs[i].max % 18
		self.tmpMaxScoreProjects.projs[i].min = self.tmpMaxScoreProjects.projs[i].min % 18
	}

	for i := 0; i < self.tmpMaxLenProjects.projLen; i++ {
		self.tmpMaxLenProjects.projs[i].max = self.tmpMaxLenProjects.projs[i].max % 18
		self.tmpMaxLenProjects.projs[i].min = self.tmpMaxLenProjects.projs[i].min % 18
	}

}

type XY struct {
	first  int
	second int
}

func (self *AlgX) moveToFirstCard() {
	self.moveToNext()
}

func (self *AlgX) moveToNextN(n int) {
	for i := 0; i < n; i++ {
		self.moveToNext()
	}
}

func (self *AlgX) moveToNext() bool {
	self.lastXy = self.xy

	if self.xy.first == len(self.cardsXY) {
		return true //直接还是当前状态进入到下个递归中，这样就走到了isEnd
	}
	if !(self.xy.first < len(self.cardsXY)) {
		panic("xy坐标超界")
	}

	m := self.cardsXY[self.xy.first]

	if m == 0 {
		var n int = self.xy.first
		for ; n < len(self.cardsXY) && self.cardsXY[n] == 0; n++ {
		}
		if n == len(self.cardsXY)-1 {
			panic("xy不能移动到此位置")
		}
		self.xy.first = n
		self.xy.second = 1
	} else {

		if self.xy.second == m {
			var n int = self.xy.first + 1
			for ; n < len(self.cardsXY) && self.cardsXY[n] == 0; n++ {
			}
			if n == len(self.cardsXY)-1 {
				panic("xy不能移动到此位置")
			}
			self.xy.first = n
			self.xy.second = 1
		} else if self.xy.second < m {
			self.xy.second++
		} else {
			panic("moveToNext err")
		}
	}
	if self.xy.first == 72 {
		return false
	}
	return true
}

func (self *AlgX) backToBefore() {
	self.xy = self.lastXy
}

func (self *AlgX) isEnd() bool {
	left := 0

	//fmt.Println("isEnd", self.xy, self.cardsXY)

	if self.xy.first >= len(self.cardsXY) {
		return true
	}

	left += self.cardsXY[self.xy.first] - self.xy.second

	if left >= 0 {
		return false
	}

	//for i := self.xy.first + 1; i < len(self.cardsXY); i++ {
	//	left += self.cardsXY[i]
	//	if left > 0 {
	//		return false
	//	}
	//}

	return true
}

func (self *AlgX) travels( /*, inputMap [][2]int16*/ ) {

	//fmt.Println("==>start index", index, len(cards))
	//if index == len(cards) {
	//
	//	fmt.Println(self.tmpFlushs, self.tmpMaxScoreProjects)
	//	if self.tmpFlushs.score > self.tmpMaxScoreProjects.score {
	//		self.tmpMaxScoreProjects.copy(&self.tmpFlushs)
	//		fmt.Println(self.tmpMaxScoreProjects, self.tmpFlushs, self.tmpMaxScoreProjects)
	//	}
	//}

	//for _, v := range cards {
	//
	//	cardNumAsFlush := inputMap[CARD_OFFSET(&v)][0]
	//
	//	if cardNumAsFlush > 0 {
	//		self.freq[v.First] = int(cardNumAsFlush)
	//	}
	//}
	//llog(-1, "search mp %v %d\n", self.freq, self.jokerLeft)

	var outCards Project

	//self.freq := make([]int, 20) //[20]int{}   //make(map[int16]int, 0)
	//for i := index; i < len(cards); i++ {

	//:#已经被用到其他子序列中

	if self.isEnd() {

		twoJokerInOneFlush := false
		for i := 0; i < self.tmpFlushs.projLen; i++ {
			if self.tmpFlushs.projs[i].tagLen > 1 {
				if !self.tmpFlushs.FlushOnlyTrySplit(i) {
					twoJokerInOneFlush = true
					break
				}
			}
		}
		//fmt.Println("copy 成功", "1", self.tmpMaxScoreProjects, "11", self.tmpMaxLenProjects, "2", self.tmpFlushs, "3", self.tmpMaxScoreProjects, "4", self.maxLenProjects, "5", twoJokerInOneFlush)
		if self.tmpFlushs.score > self.tmpMaxScoreProjects.score && !twoJokerInOneFlush {
			self.tmpMaxScoreProjects = self.tmpFlushs
			//fmt.Println("final 成功", self.tmpMaxScoreProjects, self.tmpFlushs, self.tmpMaxScoreProjects, twoJokerInOneFlush)
		}

		if (self.tmpFlushs.len > self.tmpMaxLenProjects.len || (self.tmpFlushs.len >= self.tmpMaxLenProjects.len && self.tmpFlushs.score > self.tmpMaxLenProjects.score)) && !twoJokerInOneFlush {
			self.tmpMaxLenProjects = self.tmpFlushs
		}

		return
	}

	if self.cardsXY[self.xy.first] == 0 || self.xy.second == 0 || self.xy.second > self.cardsXY[self.xy.first] {
		panic("xy 坐标不合法")
	}

	v := int16(self.xy.first)
	colorVal := int16(self.xy.first/18 + 1)

	if v%18 == 1 && self.xy.second == 1 {

		As := self.cardsXY[self.xy.first]
		//这些A都有可能移动到K的后面形成QKA组合
		//移动i个A到最末尾

		Ks := self.cardsXY[self.xy.first+12]
		Qs := self.cardsXY[self.xy.first+11]
		QJokers := min(Qs, self.jokerLeft)

		for i := 1; i <= min(Ks+QJokers, As); i++ {
			//fmt.Println(i)
			xy := self.xy
			self.freq[v] -= i
			self.freq[v+13] += i
			self.cardsXY[self.xy.first+13] += i
			self.moveToNextN(i)
			self.travels()
			self.xy = xy
			self.cardsXY[self.xy.first+13] -= i
			self.freq[v+13] -= i
			self.freq[v] += i
			//fmt.Println("A的for循环中的一次迭代", self.tmpFlushs)
		}
		//fmt.Println("A的不动的迭代", self.tmpFlushs, As, v)
	}

	//fmt.Println("index", v%18, v/18, self.xy, self.flushCards)

	cnt, cnt1 := 0, 0
	if self.getFreq(v) > 0 {
		cnt++
	}
	if self.getFreq(v+1) > 0 {
		cnt++
		cnt1++
	}
	if self.getFreq(v+2) > 0 {
		cnt++
		cnt1++
	}
	if self.getFreq(v+3) > 0 {
		cnt1++
	}

	//fmt.Println("v is:", v, self.xy, self.tmpFlushs, "cnt", cnt, "cnt1", cnt1)
	//_, _, _ := getFreq(self.freq, v), getFreq(self.freq, v+1), getFreq(self.freq, v+2)

	//joker补的原则
	//1.补做争用 2.补做开头和结尾

	if self.freq[v] <= 0 { //这张牌已经用过了则跳过

		self.moveToNext()
		self.travels()
		//self.travels(index+1, self.freq, cards)
		return
	}

	//如果产生争用，那么可能垫一个joker也可能垫两个joker
	//这里是优先把当前的牌补给上一个组合的而不是优先开新的组合
	//fmt.Println("NNNNNNNNNNEEEED", self.xy, self.freq, v%18, self.need[v], self.need, cnt, self.tmpFlushs, self.cardsXY)
	if self.need[v] > 0 { //这里可能和cnt=2存在争用所以只能断开{Deck{{3, 1}, {4, 1}, {5, 1}, {8, 1}, {9, 2}, {1, 3}, {7, 3}, {10, 3}, {11, 3}, {13, 3}, {6, 4}, {12, 4}}, 2},
		//fmt.Println("继续往后添加", v%18, self.tmpFlushs, self.jokerLeft, self.freq, self.needid[v], self.needidIndex[v], self.needid[v+1], self.needidIndex[v+1])
		//if v == 13 {
		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
		//}
		////这里判断是否产生争用
		//if cnt1 >= 2 { //这里有可能产生了争用//所以这里的选择就是如果当前有joker就使用joker来代替，或者不使用joker来代替而是自己用掉
		//
		//} else { //即使使用joker代替了也没有意义
		//
		//}
		//
		////判断要不要补在v的后一个，这样子会重开一个新组合，这样子也就判断了要不要补在一个组合的后面

		ifreqv, needv, needv1 := self.freq[v], self.need[v], self.need[v+1]

		llast := self.needidIndex[v] - 1
		ml := llast
		if self.freq[v] == 2 && llast == 1 { //如果刚好有两个v则刚好平均分，不用在分两种情况了
			ml = 0
		}

		for s := llast; s >= (llast - ml); s-- { //可能有两个组合都可以把这个v添加到自身的后边，所以这里要遍历,所以就是从以前的一种情况，变成了现在的两种情况//{Deck{{9, 1}, {9, 1}, {2, 2}, {1, 3}, {2, 3}, {2, 3}, {3, 3}, {3, 3}, {4, 3}, {7, 3}, {3, 4}, {13, 4}}, 2},

			self.freq[v] -= 1   //#用掉一个v
			self.need[v] -= 1   //#对v的需求减1
			self.need[v+1] += 1 //#对v + 1

			latestProjIndex := self.needid[v][s]

			if llast == 1 && s == 0 && ml == 1 { //主要是for s := llast; s >= (llast - ml); s--{在遍历s=0的时候，self.needidIndex[v]--取消的不是s=0对应的组合，所以要反过来，用完再反过来
				tmp := self.needid[v][0]
				self.needid[v][0] = self.needid[v][1]
				self.needid[v][1] = tmp
			}

			self.needid[v+1][self.needidIndex[v+1]] = latestProjIndex // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
			self.needidIndex[v+1]++
			//fmt.Println("a末尾添加一个遍历之前前", s, llast, self.tmpFlushs, v, self.freq, self.needid[v+1], self.needidIndex[v+1])
			self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v, colorVal)
			self.tmpFlushs.updateScoreLen()
			self.needidIndex[v]--

			//fmt.Println("a末尾添加一个遍历之前无joker", s, llast, self.xy, self.tmpFlushs, v, self.freq)

			xy := self.xy
			self.moveToNext()
			//fmt.Println("为什么不进入", self.xy)
			self.travels()
			self.xy = xy
			//fmt.Println("a末尾添加一个遍历之后", s, llast, self.tmpFlushs, v, self.freq)

			self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
			self.tmpFlushs.updateScoreLen()
			if llast == 1 && s == 0 && ml == 1 {
				tmp := self.needid[v][1]
				self.needid[v][1] = self.needid[v][0]
				self.needid[v][0] = tmp
			}
			self.needidIndex[v]++
			self.needidIndex[v+1]--
			self.needid[v+1][self.needidIndex[v+1]] = 0
			//if v+1 == 8 {
			//fmt.Println("pttttttttt退掉", self.need[8], needv1)
			//}
			self.need[v+1] = needv1

			self.need[v] = needv
			self.freq[v] = ifreqv

			//needidIndex[v]--
			//先简单处理,只要这里的joker跟本组合的joker之间间隔大于3都尝试下替换joker|不替换joker|如果v+1不存在则把joker补在最末尾
		}

		//一个组合中有两个joker的可以先产生之后再分裂看效果
		if self.jokerLeft > 0 { //替换当前的或者插入到最末尾
			//替换当前
			_, needv, needv1 := self.freq[v], self.need[v], self.need[v+1]

			//self.freq[v] -= 1        //#用掉一个v

			llast := self.needidIndex[v] - 1
			ml := llast
			if (self.freq[v] == 2 && llast == 1) || (self.freq[v] == 1 && llast == 1 /*一个v和一个joker所以第一遍和第二遍运算是一样的*/) { //如果刚好有两个v则刚好平均分，不用在分两种情况了
				ml = 0
			}

			for s := llast; s >= (llast - ml); s-- {
				self.need[v] -= 1   //#对v的需求减1
				self.need[v+1] += 1 //#对v + 1

				latestProjIndex := self.needid[v][s]

				if llast == 1 && s == 0 && ml == 1 {
					tmp := self.needid[v][0]
					self.needid[v][0] = self.needid[v][1]
					self.needid[v][1] = tmp
				}

				self.needid[v+1][self.needidIndex[v+1]] = latestProjIndex // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
				self.needidIndex[v+1]++
				//fmt.Println("b末尾添加一个遍历之前前jj", self.tmpFlushs, self.tmpFlushs.projs[latestProjIndex], v, self.tmpFlushs.projs[latestProjIndex].max, v, latestProjIndex, self.freq, self.needid[v], self.needidIndex[v], self.needid[v+1], self.needidIndex[v+1], s, llast, ml)
				self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v, colorVal)
				self.tmpFlushs.projs[latestProjIndex].addJokerIndex(self.tmpFlushs.projs[latestProjIndex].len - 1)
				self.tmpFlushs.updateScoreLen()
				//fmt.Println("b末尾添加一个遍历之前j", self.tmpFlushs, v, self.freq, self.need[v], self.needid[v], self.needidIndex[v], self.needid[v+1], self.needidIndex[v+1])
				self.jokerLeft--
				//self.travels(index, self.freq, cards) //self.travels(index+1, self.freq, cards)=>产生这种错误  {cards[[{7,1} {8,1} joker] [{9,1} {10,1} {11,1} {12,1} {13,1} {14,1}]] score:84 len:9} 2346 56
				self.needidIndex[v]--
				xy := self.xy
				self.moveToNext()
				self.travels() //还来拿这张牌
				self.xy = xy
				self.needidIndex[v]++
				self.jokerLeft++
				//fmt.Println("b末尾添加一个遍历之后", self.tmpFlushs, v, self.freq)

				self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
				self.tmpFlushs.projs[latestProjIndex].removeLastJokerIndex()
				self.tmpFlushs.updateScoreLen()
				if llast == 1 && s == 0 && ml == 1 {
					tmp := self.needid[v][1]
					self.needid[v][1] = self.needid[v][0]
					self.needid[v][0] = tmp
				}
				self.needidIndex[v+1]--
				self.needid[v+1][self.needidIndex[v+1]] = 0
				self.need[v+1] = needv1
				self.need[v] = needv
				//self.freq[v] = ifreqv

				//可以把joker补在末尾
				//fmt.Println("nttttttttt", v%18, (v+1)%18 <= 14, self.freq[v+2])
			}

		}

		if self.jokerLeft > 0 && (v+1)%18 <= 14 { //可以把joker补在末尾
			ifreqv, needv, needv2 := self.freq[v], self.need[v], self.need[v+2]

			llast := self.needidIndex[v] - 1
			ml := llast
			if self.freq[v] == 2 && llast == 1 && ml == 0 { //如果刚好有两个v则刚好平均分，不用在分两种情况了解决这个问题{Deck{{12, 1}, {12, 1}, {13, 1}, {13, 1}, {2, 4}, {10, 1}, {9, 3}, {4, 1}, {8, 4}, {11, 1}, {8, 4}, {7, 3}}, 2},添加的时候panic问题
				ml = 0
			}

			for s := llast; s >= (llast - ml); s-- {
				self.freq[v] -= 1   //#用掉一个v
				self.need[v] -= 1   //#对v的需求减1
				self.need[v+2] += 1 //#对v + 1

				latestProjIndex := self.needid[v][s]

				if llast == 1 && s == 0 && ml == 1 {
					tmp := self.needid[v][0]
					self.needid[v][0] = self.needid[v][1]
					self.needid[v][1] = tmp
				}

				self.needid[v+2][self.needidIndex[v+2]] = latestProjIndex // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
				self.needidIndex[v+2]++
				//fmt.Println("c末尾添加一个遍历之前前", self.tmpFlushs, v, self.freq)
				self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v, colorVal)
				self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v+1, colorVal)
				self.tmpFlushs.projs[latestProjIndex].addJokerIndex(self.tmpFlushs.projs[latestProjIndex].len - 1)
				self.tmpFlushs.updateScoreLen()
				//fmt.Println("c末尾添加一个遍历之前", self.tmpFlushs, v%18, self.freq)
				self.jokerLeft--
				self.needidIndex[v]--
				xy := self.xy
				//fmt.Println("c要往后走了！！", self.xy)
				self.moveToNext()
				//fmt.Println("c要往后走了2！！", self.xy)
				self.travels()
				self.needidIndex[v]++
				//fmt.Println("要往后走了之后！！", self.xy)
				self.xy = xy
				self.jokerLeft++
				//fmt.Println("c末尾添加一个遍历之后", self.tmpFlushs, v%18, self.freq)

				self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
				self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
				self.tmpFlushs.projs[latestProjIndex].removeLastJokerIndex()
				self.tmpFlushs.updateScoreLen()
				if llast == 1 && s == 0 && ml == 1 {
					tmp := self.needid[v][1]
					self.needid[v][1] = self.needid[v][0]
					self.needid[v][0] = tmp
				}
				self.needidIndex[v+2]--
				self.needid[v+2][self.needidIndex[v+2]] = 0
				self.need[v+2] = needv2
				self.need[v] = needv
				self.freq[v] = ifreqv
			}

		}

	}
	if cnt == 3 {
		//fmt.Println("起新组合", v%18, v/18+1)
		//一个都不替换

		//这个是一定要做且不回退的
		f, f1, f2 := self.freq[v], self.freq[v+1], self.freq[v+2]
		n3 := self.need[v+3]

		self.freq[v] -= 1
		self.freq[v+1] -= 1
		self.freq[v+2] -= 1
		self.need[v+3] += 1 //#对v+3的需求加1
		//if v+3 == 8 {
		//	fmt.Println("3pttttttttt退掉", self.need[8], self.tmpFlushs, self.needidIndex[v+3])
		//}
		self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen
		self.needidIndex[v+3]++
		//var p Project
		outCards = NewFlushProject(v, v+2, colorVal)
		self.tmpFlushs.appendProject(outCards)
		xy := self.xy
		self.moveToNext()
		self.travels()
		self.xy = xy
		//fmt.Println("3 joker", self.jokerLeft)

		if self.jokerLeft > 0 {
			//替换掉第一个
			//fmt.Println("替换掉第一个", v+2, self.freq)
			//cnt := 0

			self.freq[v] += 1
			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].addJokerIndex(0)
			self.jokerLeft--
			//fmt.Println("替换掉第一个", v, self.tmpFlushs)
			xy := self.xy
			self.moveToNext()
			self.travels()
			self.xy = xy
			self.jokerLeft++
			self.freq[v] -= 1

			//替换掉第二个
			//fmt.Println("替换掉第二个", v+1, self.tmpFlushs)
			self.freq[v+1] += 1
			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(1)
			self.jokerLeft--
			//fmt.Println("moveToNext之前", self.xy, self.need[self.xy.first], self.tmpFlushs)
			xy = self.xy
			self.moveToNext()
			//fmt.Println("moveToNext", self.xy, self.need[self.xy.first], self.tmpFlushs)
			self.travels()
			self.xy = xy
			self.jokerLeft++
			self.freq[v+1] -= 1
			//fmt.Println("替换掉第二个之后", v, self.freq, self.tmpFlushs)
			//替换掉第三个
			//fmt.Println("替换掉最后一个", v+2, self.freq)
			self.freq[v+2] += 1
			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(2)
			self.jokerLeft--
			//fmt.Println("替换掉最后一个", v, self.tmpFlushs)

			xy = self.xy
			self.moveToNext()
			self.travels()
			self.xy = xy
			self.jokerLeft++
			self.freq[v+2] -= 1

			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
			//尝试放在头部
			if v%18 > 1 { //针对{12,13,14}这种牌型
				//old无法应对以下牌型
				//per us:  39  最大分: 55   {cards[[{11,3} joker {13,3}] [joker {2,4} {3,4} {4,4} {5,4}]] score:55 len:8} 最大长度 {cards[] score:0 len:0} dfsTimes: 0
				//per us: 101  最大分: 58   [[Joker ♦2 ♦3] [♦3 ♦4 ♦5] [♣J Joker ♣K]] 49 89
				//[{10,1} {7,2} {3,3} {8,3} {11,3} {13,3} {2,4} {3,4} {3,4} {4,4} {5,4} {13,4}]
				//outCards.FlushAddHead()
				//self.tmpFlushs.updateScoreLen()
				//outCards.addJokerIndex(0)
				//self.jokerLeft--
				//xy := self.xy
				//self.moveToNext()
				//self.travels(self.freq)
				//self.xy = xy
				//self.jokerLeft++
				//outCards.removeLastJokerIndex()
				//outCards.FlushRemoveHead()
				//
				//self.tmpFlushs.updateScoreLen()
				self.need[v+3] -= 1 //#对v+3的需求加1
				self.needidIndex[v+3]--
				self.needid[v+3][self.needidIndex[v+3]] = 0

				self.need[v+2] += 1
				self.needid[v+2][self.needidIndex[v+2]] = self.tmpFlushs.projLen - 1
				self.needidIndex[v+2]++

				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushRemoveTail()
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushAddHead()
				self.tmpFlushs.updateScoreLen()
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].addJokerIndex(0)

				self.freq[v+2] += 1
				self.jokerLeft--
				xy := self.xy
				//self.moveToNext()
				self.travels()
				self.xy = xy
				self.jokerLeft++
				self.freq[v+2] -= 1

				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushAddTailAuto()
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushRemoveHead()
				self.tmpFlushs.updateScoreLen()
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()

				self.need[v+2] -= 1
				self.needidIndex[v+2]--
				self.needid[v+2][self.needidIndex[v+2]] = 0

				self.need[v+3] += 1
				self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen - 1
				self.needidIndex[v+3]++
			}

			//尝试放在末尾
			if (v+3)%18 <= 14 { //针对{{1, 1}, {2, 1}, {3, 1}, {1, 1}, {2, 1}, {3, 1}}这种类型

				self.need[v+3] -= 1 //#对v+3的需求加1
				self.needidIndex[v+3]--
				self.needid[v+3][self.needidIndex[v+3]] = 0

				self.need[v+4] += 1
				//if v+4 == 8 {
				//	fmt.Println("v4xpttttttttt退掉", self.need[8])
				//}
				self.needid[v+4][self.needidIndex[v+4]] = self.tmpFlushs.projLen - 1
				self.needidIndex[v+4]++
				//fmt.Println("here2")
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushAddTail(v+3, colorVal)
				self.tmpFlushs.updateScoreLen()
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].addJokerIndex(3)
				self.jokerLeft--
				xy := self.xy
				self.moveToNext()
				self.travels()
				self.xy = xy
				self.jokerLeft++
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushRemoveTail()
				self.tmpFlushs.updateScoreLen()
				self.need[v+4] -= 1
				self.needidIndex[v+4]--
				self.needid[v+4][self.needidIndex[v+4]] = 0

				self.need[v+3] += 1
				//if v+3 == 8 {
				//	fmt.Println("v3pttttttttt退掉", self.need[8])
				//}
				self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen - 1
				self.needidIndex[v+3]++
			}

		}

		self.freq[v] = f
		self.freq[v+1] = f1
		self.freq[v+2] = f2
		self.need[v+3] = n3
		//if v+3 == 8 {
		//	fmt.Println("v38pttttttttt退掉", self.need[8], n3)
		//}
		self.needidIndex[v+3]--
		self.needid[v+3][self.needidIndex[v+3]] = 0
		self.tmpFlushs.removeLastProject()
	}
	if cnt == 2 { //分两种情况，一种用joker在这里一种是不用
		//fmt.Println("2个的新组合", self.xy, v%18, v/18+1, self.freq[v], self.freq[v+1], self.freq[v+2], self.jokerLeft)
		xy := self.xy
		self.moveToNext()
		self.travels()
		self.xy = xy
		//fmt.Println("-2个的新组合", self.xy, v%18, v/18+1, self.freq[v], self.freq[v+1], self.freq[v+2], self.jokerLeft)
		if self.jokerLeft > 0 {

			if self.freq[v] == 0 {

			} else if self.freq[v+1] == 0 {
				//joker作为中间一个
				a, b, c := self.freq[v], self.freq[v+1], self.freq[v+2]
				if self.freq[v] > 0 {
					self.freq[v] -= 1
				}

				if self.freq[v+1] > 0 {
					self.freq[v+1] -= 1
				}

				if self.freq[v+2] > 0 {
					self.freq[v+2] -= 1
				}

				self.need[v+3] += 1 //#对v+3的需求加1
				//if v+3 == 8 {
				//	fmt.Println("2pttttttttt退掉", self.need[8])
				//}
				self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen
				self.needidIndex[v+3]++

				outCards = NewFlushProject(v, v+2, colorVal)
				outCards.addJokerIndex(v + 1)
				self.tmpFlushs.updateScoreLen()
				outCards.replaceLastJokerIndex(1)
				self.tmpFlushs.appendProject(outCards)
				//fmt.Println("两个的开始凑集合1", self.xy, self.freq, self.tmpFlushs)

				self.jokerLeft--
				xy := self.xy
				self.moveToNext()
				self.travels()
				self.xy = xy
				//fmt.Println("两个的开始凑集合1 结束后", index, self.freq, self.tmpFlushs)
				self.jokerLeft++
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
				self.tmpFlushs.updateScoreLen()
				self.tmpFlushs.removeLastProject()
				self.needidIndex[v+3]--
				self.needid[v+3][self.needidIndex[v+3]] = 0
				self.need[v+3] -= 1
				self.freq[v], self.freq[v+1], self.freq[v+2] = a, b, c
				//fmt.Println("步出两个开始凑集合", self.tmpFlushs, self.need, self.freq)
			} else if self.freq[v+2] == 0 {
				//joker为最后一个
				//fmt.Println("2v3pttttttttt退掉", self.tmpFlushs, v%18)
				if (v+3)%18 <= 15 {
					a, b, c := self.freq[v], self.freq[v+1], self.freq[v+2]
					if self.freq[v] > 0 {
						self.freq[v] -= 1
					}

					if self.freq[v+1] > 0 {
						self.freq[v+1] -= 1
					}

					if self.freq[v+2] > 0 {
						self.freq[v+2] -= 1
					}

					self.need[v+3] += 1 //#对v+3的需求加1
					//if v+1 == 8 {

					//}
					self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen
					self.needidIndex[v+3]++

					outCards = NewFlushProject(v, v+2, colorVal)
					outCards.addJokerIndex(v + 1)
					self.tmpFlushs.appendProject(outCards)
					self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(2)

					//fmt.Println("两个的开始凑集合2尾部", self.freq, self.tmpFlushs, self.xy, self.cardsXY, self.jokerLeft)

					self.jokerLeft--
					xy := self.xy
					self.moveToNext()
					self.travels()
					self.xy = xy
					self.jokerLeft++
					self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
					self.tmpFlushs.removeLastProject()
					self.needidIndex[v+3]--
					self.needid[v+3][self.needidIndex[v+3]] = 0
					self.need[v+3] -= 1
					self.freq[v], self.freq[v+1], self.freq[v+2] = a, b, c
					//fmt.Println("两个的开始凑集合2尾部之后", self.xy, self.freq, self.tmpFlushs)

				}

				//joker作为最前面一个
				a, b, c := self.freq[v], self.freq[v+1], self.freq[v+2]
				if self.freq[v] > 0 {
					self.freq[v] -= 1
				}

				if self.freq[v+1] > 0 {
					self.freq[v+1] -= 1
				}

				if self.freq[v+2] > 0 {
					self.freq[v+2] -= 1
				}

				self.need[v+2] += 1 //#对v+3的需求加1
				//if v+2 == 8 {
				//	fmt.Println("22v2pttttttttt退掉", self.need[8])
				//}
				self.needid[v+2][self.needidIndex[v+2]] = self.tmpFlushs.projLen
				self.needidIndex[v+2]++

				outCards = NewFlushProject(v-1, v+1, colorVal)
				outCards.addJokerIndex(v)
				self.tmpFlushs.appendProject(outCards)

				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(0)
				//fmt.Println("两个的开始凑集合2头部", self.xy, self.freq, self.tmpFlushs, v)
				self.jokerLeft--
				xy = self.xy
				self.moveToNext()
				self.travels()
				self.xy = xy
				self.jokerLeft++
				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
				self.tmpFlushs.removeLastProject()
				self.needidIndex[v+2]--
				self.needid[v+2][self.needidIndex[v+2]] = 0
				self.need[v+2] -= 1
				self.freq[v], self.freq[v+1], self.freq[v+2] = a, b, c

			} else {
				panic("")
			}

			//self.tmpFlushs.removeLastProject()
			self.tmpFlushs.updateScoreLen()
		}
		//有和无joker的情况下不在这里补joker

	}

	//正常往后迭代//独立的牌
	//fmt.Println("单独一张牌", v%18, v/18)
	if cnt == 1 {
		//fmt.Println("单独一张牌", v%18, v/18)
		self.moveToNext()
		self.travels()
	}

	//	if need[v] > 0 {
	//		//:#先判断v是否能接到其他子序列后面
	//
	//		self.freq[v] -= 1   //#用掉一个v
	//		need[v] -= 1   //#对v的需求减1
	//		need[v+1] += 1 //#对v + 1
	//		llast := needidIndex[v] - 1
	//
	//		needid[v+1][needidIndex[v+1]] = needid[v][llast] // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
	//		needidIndex[v+1]++
	//		cc.projs[needid[v][llast]].FlushAddTail(v, colorVal)
	//
	//		needidIndex[v]--
	//
	//	} else if self.freq[v] > 0 && self.freq[v+1] > 0 && self.freq[v+2] > 0 { //#v作为开头，新建一个长度为3的子序列[v, v + 1, v + 2]
	//
	//		self.freq[v] -= 1
	//		self.freq[v+1] -= 1
	//		self.freq[v+2] -= 1
	//		need[v+3] += 1 //#对v+3的需求加1
	//		needid[v+3][needidIndex[v+3]] = len(cc.projs)
	//		needidIndex[v+3]++
	//		outCards = NewFlushProject(v, v+2, colorVal)
	//
	//		cc.appendProject(outCards)
	//
	//	} else {
	//		//cc = append(cc, outCards)
	//		////fmt.Println("XXXXXXXXXXXXXXXX  here!")
	//	}
	//}
	//cc.updateScoreLen()
	//
	////llog(-100, "\nget0FlushProjects freqSrc(%v) self.freq(%v) Cards(%v) cc is(%v)\n", freqSrc, self.freq, cards, cc) //{1,1} {4,1} {6,2} {9,2} {10,2} {1,3} {1,3} {2,3} {3,3} {10,3} {1,4} {2,4} {12,4}
	//
	//if cc.len > 0 {
	//	for i := 0; i < len(cc.projs); i++ {
	//		if len(cc.projs[i].tagIndex) > 1 { //碰到一个project中有两个joker的需要分裂
	//			if !cc.FlushTrySplit(i) {
	//				llog(-100, "\ngetFlushProjects freqSrc(%v) self.freq(%v) Cards(%v) cc is(%v)\n", freqSrc, self.freq, cards, cc) //{1,1} {4,1} {6,2} {9,2} {10,2} {1,3} {1,3} {2,3} {3,3} {10,3} {1,4} {2,4} {12,4}
	//				return nil
	//			}
	//		}
	//	}
	//	//llog(-100, "\nget2FlushProjects freqSrc(%v) self.freq(%v) Cards(%v) cc is(%v)\n", freqSrc, self.freq, cards, cc) //{1,1} {4,1} {6,2} {9,2} {10,2} {1,3} {1,3} {2,3} {3,3} {10,3} {1,4} {2,4} {12,4}
	//	return &cc
	//}

}

//func (self *AlgX) searchFlushOld() {
//
//	//defer func() {
//	//	if recover() != nil {
//	//		fmt.Println(self.hand_cards, self.tmpFlushs, self.cardsXY, self.xy)
//	//	}
//	//}()
//
//	inputMap := self.states
//	self.xy = XY{}
//	for i := 0; i < 72; i++ {
//		self.cardsXY[i] = 0
//		self.freq[i] = 0
//	}
//	freq := self.freq
//	for _, oneRow := range self.flushCards {
//		for _, v := range oneRow {
//
//			if inputMap[CARD_OFFSET(&v)][0]+inputMap[CARD_OFFSET(&v)][1] > 0 { //[0]代表给顺子
//				if int(inputMap[CARD_OFFSET(&v)][0]) == 0 {
//					continue
//				}
//				idx := v.First + (v.Second-1)*18
//				val := int(inputMap[CARD_OFFSET(&v)][0])
//				self.cardsXY[idx] = val
//				freq[idx] = val
//				continue
//			}
//		}
//	}
//
//	//fmt.Println("this is self.cardsXY:", self.cardsXY)
//
//	self.tmpFlushs.clear()
//	self.tmpMaxScoreProjects.clear()
//	self.tmpMaxLenProjects.clear()
//
//	self.travelsWork(freq)
//
//}

//func (self *AlgX) travelsWorkOld(freq []int) {
//
//	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1}, {9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}}
//	//var cards []Card = []Card{{4, 1}, {7, 1}, {9, 1}, {10, 1}, {11, 1}, {5, 2}, {6, 2}} // {12, 2}, {13, 2}} //, {2, 3}, {5, 3}, {11, 3}, {4, 4}}
//	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}}
//	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {1, 1}, {2, 1}, {3, 1}}
//	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {12, 1}, {13, 1}}
//	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {1, 1}, {2, 1}, {3, 1}, {12, 1}, {13, 1}, {12, 1}, {13, 1}}
//	//var cards []Card = []Card{{1, 1}, {2, 1}, {3, 1}, {4, 1}}
//	//var cards []Card = []Card{{6, 1}, {7, 1}, {8, 1}, {3, 2}, {4, 2}, {4, 3}, {6, 3}, {11, 3}, {12, 3}, {12, 3}, {3, 4}, {10, 4}, {13, 4}}
//	//var cards []Card = []Card{{4, 1}, {7, 1}, {9, 1}, {10, 1}, {11, 1}, {5, 2}, {6, 2}}
//	//var cards []Card = []Card{{8, 1}, {9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}} //这个是错的//
//	//var cards []Card = []Card{{9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}}
//	//var cards []Card = []Card{{7, 1}, {8, 1}, {9, 1}, {10, 1}, {11, 1}, {12, 1}, {13, 1}, {14, 1}}
//	//var cards []Card = []Card{{12, 1}, {13, 1}, {14, 1}}
//	//var cards []Card = []Card{{6, 1}, {7, 1}, {8, 1}, {3, 2}, {4, 2}, {4, 3}, {6, 3}, {11, 3}, {12, 3}, {12, 3}, {3, 4}, {10, 4}, {13, 4}}
//
//	//
//	//for i := range cards {
//	//	cards[i].First = cards[i].First + (cards[i].Second-1)*18
//	//	self.cardsXY[cards[i].First]++
//	//}
//
//	self.moveToFirstCard()
//	self.travels(freq)
//	//fmt.Println("search", self.tmpMaxScoreProjects)
//	for i := 0; i < self.tmpMaxScoreProjects.projLen; i++ {
//		self.tmpMaxScoreProjects.projs[i].max = self.tmpMaxScoreProjects.projs[i].max % 18
//		self.tmpMaxScoreProjects.projs[i].min = self.tmpMaxScoreProjects.projs[i].min % 18
//	}
//
//	for i := 0; i < self.tmpMaxLenProjects.projLen; i++ {
//		self.tmpMaxLenProjects.projs[i].max = self.tmpMaxLenProjects.projs[i].max % 18
//		self.tmpMaxLenProjects.projs[i].min = self.tmpMaxLenProjects.projs[i].min % 18
//	}
//
//}

//func (self *AlgX) travelsold(freq []int /*, inputMap [][2]int16*/) {
//
//	//fmt.Println("==>start index", index, len(cards))
//	//if index == len(cards) {
//	//
//	//	fmt.Println(self.tmpFlushs, self.tmpMaxScoreProjects)
//	//	if self.tmpFlushs.score > self.tmpMaxScoreProjects.score {
//	//		self.tmpMaxScoreProjects.copy(&self.tmpFlushs)
//	//		fmt.Println(self.tmpMaxScoreProjects, self.tmpFlushs, self.tmpMaxScoreProjects)
//	//	}
//	//}
//
//	//for _, v := range cards {
//	//
//	//	cardNumAsFlush := inputMap[CARD_OFFSET(&v)][0]
//	//
//	//	if cardNumAsFlush > 0 {
//	//		freq[v.First] = int(cardNumAsFlush)
//	//	}
//	//}
//	//llog(-1, "search mp %v %d\n", freq, self.jokerLeft)
//
//	var outCards Project
//
//	//freq := make([]int, 20) //[20]int{}   //make(map[int16]int, 0)
//	//for i := index; i < len(cards); i++ {
//
//	//:#已经被用到其他子序列中
//
//	if self.isEnd() {
//
//		twoJokerInOneFlush := false
//		for i := 0; i < self.tmpFlushs.projLen; i++ {
//			if len(self.tmpFlushs.projs[i].tagIndex) > 1 {
//				if !self.tmpFlushs.FlushOnlyTrySplit(i) {
//					twoJokerInOneFlush = true
//					break
//				}
//			}
//		}
//		//fmt.Println("copy 成功", self.tmpMaxScoreProjects, self.tmpFlushs, self.tmpMaxScoreProjects, twoJokerInOneFlush)
//		if self.tmpFlushs.score > self.tmpMaxScoreProjects.score && !twoJokerInOneFlush {
//			self.tmpMaxScoreProjects.copy(&self.tmpFlushs)
//			//fmt.Println("final 成功", self.tmpMaxScoreProjects, self.tmpFlushs, self.tmpMaxScoreProjects, twoJokerInOneFlush)
//		}
//
//		if self.tmpFlushs.len > self.tmpMaxLenProjects.len && !twoJokerInOneFlush {
//			self.tmpMaxLenProjects.copy(&self.tmpFlushs)
//		}
//
//		return
//	}
//
//	if self.cardsXY[self.xy.first] == 0 || self.xy.second == 0 || self.xy.second > self.cardsXY[self.xy.first] {
//		panic("xy 坐标不合法")
//	}
//
//	v := int16(self.xy.first)
//	colorVal := int16(self.xy.first/18 + 1)
//
//	if v%18 == 1 && self.xy.second == 1 {
//
//		As := self.cardsXY[self.xy.first]
//		//这些A都有可能移动到K的后面形成QKA组合
//		//移动i个A到最末尾
//
//		Ks := self.cardsXY[self.xy.first+12]
//		Qs := self.cardsXY[self.xy.first+11]
//		QJokers := min(Qs, self.jokerLeft)
//
//		for i := 1; i <= min(Ks+QJokers, As); i++ {
//			//fmt.Println(i)
//			xy := self.xy
//			freq[v] -= i
//			freq[v+13] += i
//			self.cardsXY[self.xy.first+13] += i
//			self.moveToNextN(i)
//			self.travels(freq)
//			self.xy = xy
//			self.cardsXY[self.xy.first+13] -= i
//			freq[v+13] -= i
//			freq[v] += i
//			//fmt.Println("A的for循环中的一次迭代", self.tmpFlushs)
//		}
//		//fmt.Println("A的不动的迭代", self.tmpFlushs, As, v)
//	}
//
//	//fmt.Println("index", v%18, v/18, self.xy, self.flushCards)
//
//	cnt, cnt1 := 0, 0
//	if getFreq(freq, v) > 0 {
//		cnt++
//	}
//	if getFreq(freq, v+1) > 0 {
//		cnt++
//		cnt1++
//	}
//	if getFreq(freq, v+2) > 0 {
//		cnt++
//		cnt1++
//	}
//	if getFreq(freq, v+3) > 0 {
//		cnt1++
//	}
//
//	//fmt.Println("v is:", v, self.xy, self.tmpFlushs, "cnt", cnt, "cnt1", cnt1)
//	//_, _, _ := getFreq(freq, v), getFreq(freq, v+1), getFreq(freq, v+2)
//
//	//joker补的原则
//	//1.补做争用 2.补做开头和结尾
//
//	if freq[v] <= 0 { //这张牌已经用过了则跳过
//
//		self.moveToNext()
//		self.travels(freq)
//		//self.travels(index+1, freq, cards)
//		return
//	}
//
//	//如果产生争用，那么可能垫一个joker也可能垫两个joker
//	//这里是优先把当前的牌补给上一个组合的而不是优先开新的组合
//	//fmt.Println("NNNNNNNNNNEEEED", self.xy, freq, v%18, self.need[v], self.need, cnt, self.tmpFlushs, self.cardsXY)
//	if self.need[v] > 0 { //这里可能和cnt=2存在争用所以只能断开{Deck{{3, 1}, {4, 1}, {5, 1}, {8, 1}, {9, 2}, {1, 3}, {7, 3}, {10, 3}, {11, 3}, {13, 3}, {6, 4}, {12, 4}}, 2},
//		//fmt.Println("继续往后添加", v%18, self.tmpFlushs, self.jokerLeft, freq, self.needid[v], self.needidIndex[v], self.needid[v+1], self.needidIndex[v+1])
//		//if v == 13 {
//		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
//		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
//		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
//		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
//		//	fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKK")
//		//}
//		////这里判断是否产生争用
//		//if cnt1 >= 2 { //这里有可能产生了争用//所以这里的选择就是如果当前有joker就使用joker来代替，或者不使用joker来代替而是自己用掉
//		//
//		//} else { //即使使用joker代替了也没有意义
//		//
//		//}
//		//
//		////判断要不要补在v的后一个，这样子会重开一个新组合，这样子也就判断了要不要补在一个组合的后面
//
//		ifreqv, needv, needv1 := freq[v], self.need[v], self.need[v+1]
//
//		llast := self.needidIndex[v] - 1
//		ml := llast
//		if freq[v] == 2 && llast == 1 { //如果刚好有两个v则刚好平均分，不用在分两种情况了
//			ml = 0
//		}
//
//		for s := llast; s >= (llast - ml); s-- { //可能有两个组合都可以把这个v添加到自身的后边，所以这里要遍历,所以就是从以前的一种情况，变成了现在的两种情况//{Deck{{9, 1}, {9, 1}, {2, 2}, {1, 3}, {2, 3}, {2, 3}, {3, 3}, {3, 3}, {4, 3}, {7, 3}, {3, 4}, {13, 4}}, 2},
//
//			freq[v] -= 1        //#用掉一个v
//			self.need[v] -= 1   //#对v的需求减1
//			self.need[v+1] += 1 //#对v + 1
//
//			latestProjIndex := self.needid[v][s]
//
//			if llast == 1 && s == 0 && ml == 1 { //主要是for s := llast; s >= (llast - ml); s--{在遍历s=0的时候，self.needidIndex[v]--取消的不是s=0对应的组合，所以要反过来，用完再反过来
//				tmp := self.needid[v][0]
//				self.needid[v][0] = self.needid[v][1]
//				self.needid[v][1] = tmp
//			}
//
//			self.needid[v+1][self.needidIndex[v+1]] = latestProjIndex // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
//			self.needidIndex[v+1]++
//			//fmt.Println("a末尾添加一个遍历之前前", s, llast, self.tmpFlushs, v, freq, self.needid[v+1], self.needidIndex[v+1])
//			self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v, colorVal)
//			self.tmpFlushs.updateScoreLen()
//			self.needidIndex[v]--
//
//			//fmt.Println("a末尾添加一个遍历之前无joker", s, llast, self.xy, self.tmpFlushs, v, freq)
//
//			xy := self.xy
//			self.moveToNext()
//			//fmt.Println("为什么不进入", self.xy)
//			self.travels(freq)
//			self.xy = xy
//			//fmt.Println("a末尾添加一个遍历之后", s, llast, self.tmpFlushs, v, freq)
//
//			self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
//			self.tmpFlushs.updateScoreLen()
//			if llast == 1 && s == 0 && ml == 1 {
//				tmp := self.needid[v][1]
//				self.needid[v][1] = self.needid[v][0]
//				self.needid[v][0] = tmp
//			}
//			self.needidIndex[v]++
//			self.needidIndex[v+1]--
//			self.needid[v+1][self.needidIndex[v+1]] = 0
//			//if v+1 == 8 {
//			//fmt.Println("pttttttttt退掉", self.need[8], needv1)
//			//}
//			self.need[v+1] = needv1
//
//			self.need[v] = needv
//			freq[v] = ifreqv
//
//			//needidIndex[v]--
//			//先简单处理,只要这里的joker跟本组合的joker之间间隔大于3都尝试下替换joker|不替换joker|如果v+1不存在则把joker补在最末尾
//		}
//
//		//一个组合中有两个joker的可以先产生之后再分裂看效果
//		if self.jokerLeft > 0 { //替换当前的或者插入到最末尾
//			//替换当前
//			_, needv, needv1 := freq[v], self.need[v], self.need[v+1]
//
//			//freq[v] -= 1        //#用掉一个v
//
//			llast := self.needidIndex[v] - 1
//			ml := llast
//			if (freq[v] == 2 && llast == 1) || (freq[v] == 1 && llast == 1 /*一个v和一个joker所以第一遍和第二遍运算是一样的*/) { //如果刚好有两个v则刚好平均分，不用在分两种情况了
//				ml = 0
//			}
//
//			for s := llast; s >= (llast - ml); s-- {
//				self.need[v] -= 1   //#对v的需求减1
//				self.need[v+1] += 1 //#对v + 1
//
//				latestProjIndex := self.needid[v][s]
//
//				if llast == 1 && s == 0 && ml == 1 {
//					tmp := self.needid[v][0]
//					self.needid[v][0] = self.needid[v][1]
//					self.needid[v][1] = tmp
//				}
//
//				self.needid[v+1][self.needidIndex[v+1]] = latestProjIndex // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
//				self.needidIndex[v+1]++
//				//fmt.Println("b末尾添加一个遍历之前前jj", self.tmpFlushs, self.tmpFlushs.projs[latestProjIndex], v, self.tmpFlushs.projs[latestProjIndex].max, v, latestProjIndex, freq, self.needid[v], self.needidIndex[v], self.needid[v+1], self.needidIndex[v+1], s, llast, ml)
//				self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v, colorVal)
//				self.tmpFlushs.projs[latestProjIndex].addJokerIndex(self.tmpFlushs.projs[latestProjIndex].len - 1)
//				self.tmpFlushs.updateScoreLen()
//				//fmt.Println("b末尾添加一个遍历之前j", self.tmpFlushs, v, freq, self.need[v], self.needid[v], self.needidIndex[v], self.needid[v+1], self.needidIndex[v+1])
//				self.jokerLeft--
//				//self.travels(index, freq, cards) //self.travels(index+1, freq, cards)=>产生这种错误  {cards[[{7,1} {8,1} joker] [{9,1} {10,1} {11,1} {12,1} {13,1} {14,1}]] score:84 len:9} 2346 56
//				self.needidIndex[v]--
//				xy := self.xy
//				self.moveToNext()
//				self.travels(freq) //还来拿这张牌
//				self.xy = xy
//				self.needidIndex[v]++
//				self.jokerLeft++
//				//fmt.Println("b末尾添加一个遍历之后", self.tmpFlushs, v, freq)
//
//				self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
//				self.tmpFlushs.projs[latestProjIndex].removeLastJokerIndex()
//				self.tmpFlushs.updateScoreLen()
//				if llast == 1 && s == 0 && ml == 1 {
//					tmp := self.needid[v][1]
//					self.needid[v][1] = self.needid[v][0]
//					self.needid[v][0] = tmp
//				}
//				self.needidIndex[v+1]--
//				self.needid[v+1][self.needidIndex[v+1]] = 0
//				self.need[v+1] = needv1
//				self.need[v] = needv
//				//freq[v] = ifreqv
//
//				//可以把joker补在末尾
//				//fmt.Println("nttttttttt", v%18, (v+1)%18 <= 14, freq[v+2])
//			}
//
//		}
//
//		if self.jokerLeft > 0 && (v+1)%18 <= 14 { //可以把joker补在末尾
//			ifreqv, needv, needv2 := freq[v], self.need[v], self.need[v+2]
//
//			llast := self.needidIndex[v] - 1
//			ml := llast
//			if freq[v] == 2 && llast == 1 && ml == 0 { //如果刚好有两个v则刚好平均分，不用在分两种情况了解决这个问题{Deck{{12, 1}, {12, 1}, {13, 1}, {13, 1}, {2, 4}, {10, 1}, {9, 3}, {4, 1}, {8, 4}, {11, 1}, {8, 4}, {7, 3}}, 2},添加的时候panic问题
//				ml = 0
//			}
//
//			for s := llast; s >= (llast - ml); s-- {
//				freq[v] -= 1        //#用掉一个v
//				self.need[v] -= 1   //#对v的需求减1
//				self.need[v+2] += 1 //#对v + 1
//
//				latestProjIndex := self.needid[v][s]
//
//				if llast == 1 && s == 0 && ml == 1 {
//					tmp := self.needid[v][0]
//					self.needid[v][0] = self.needid[v][1]
//					self.needid[v][1] = tmp
//				}
//
//				self.needid[v+2][self.needidIndex[v+2]] = latestProjIndex // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
//				self.needidIndex[v+2]++
//				//fmt.Println("c末尾添加一个遍历之前前", self.tmpFlushs, v, freq)
//				self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v, colorVal)
//				self.tmpFlushs.projs[latestProjIndex].FlushAddTail(v+1, colorVal)
//				self.tmpFlushs.projs[latestProjIndex].addJokerIndex(self.tmpFlushs.projs[latestProjIndex].len - 1)
//				self.tmpFlushs.updateScoreLen()
//				//fmt.Println("c末尾添加一个遍历之前", self.tmpFlushs, v%18, freq)
//				self.jokerLeft--
//				self.needidIndex[v]--
//				xy := self.xy
//				//fmt.Println("c要往后走了！！", self.xy)
//				self.moveToNext()
//				//fmt.Println("c要往后走了2！！", self.xy)
//				self.travels(freq)
//				self.needidIndex[v]++
//				//fmt.Println("要往后走了之后！！", self.xy)
//				self.xy = xy
//				self.jokerLeft++
//				//fmt.Println("c末尾添加一个遍历之后", self.tmpFlushs, v%18, freq)
//
//				self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
//				self.tmpFlushs.projs[latestProjIndex].FlushRemoveTail()
//				self.tmpFlushs.projs[latestProjIndex].removeLastJokerIndex()
//				self.tmpFlushs.updateScoreLen()
//				if llast == 1 && s == 0 && ml == 1 {
//					tmp := self.needid[v][1]
//					self.needid[v][1] = self.needid[v][0]
//					self.needid[v][0] = tmp
//				}
//				self.needidIndex[v+2]--
//				self.needid[v+2][self.needidIndex[v+2]] = 0
//				self.need[v+2] = needv2
//				self.need[v] = needv
//				freq[v] = ifreqv
//			}
//
//		}
//
//	}
//	if cnt == 3 {
//		//fmt.Println("起新组合", v%18, v/18+1)
//		//一个都不替换
//
//		//这个是一定要做且不回退的
//		f, f1, f2 := freq[v], freq[v+1], freq[v+2]
//		n3 := self.need[v+3]
//
//		freq[v] -= 1
//		freq[v+1] -= 1
//		freq[v+2] -= 1
//		self.need[v+3] += 1 //#对v+3的需求加1
//		//if v+3 == 8 {
//		//	fmt.Println("3pttttttttt退掉", self.need[8], self.tmpFlushs, self.needidIndex[v+3])
//		//}
//		self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen
//		self.needidIndex[v+3]++
//		//var p Project
//		outCards = NewFlushProject(v, v+2, colorVal)
//		self.tmpFlushs.appendProject(outCards)
//		xy := self.xy
//		self.moveToNext()
//		self.travels(freq)
//		self.xy = xy
//		//fmt.Println("3 joker", self.jokerLeft)
//
//		if self.jokerLeft > 0 {
//			//替换掉第一个
//			//fmt.Println("替换掉第一个", v+2, freq)
//			//cnt := 0
//
//			freq[v] += 1
//			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].addJokerIndex(0)
//			self.jokerLeft--
//			//fmt.Println("替换掉第一个", v, self.tmpFlushs)
//			xy := self.xy
//			self.moveToNext()
//			self.travels(freq)
//			self.xy = xy
//			self.jokerLeft++
//			freq[v] -= 1
//
//			//替换掉第二个
//			//fmt.Println("替换掉第二个", v+1, self.tmpFlushs)
//			freq[v+1] += 1
//			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(1)
//			self.jokerLeft--
//			//fmt.Println("moveToNext之前", self.xy, self.need[self.xy.first], self.tmpFlushs)
//			xy = self.xy
//			self.moveToNext()
//			//fmt.Println("moveToNext", self.xy, self.need[self.xy.first], self.tmpFlushs)
//			self.travels(freq)
//			self.xy = xy
//			self.jokerLeft++
//			freq[v+1] -= 1
//			//fmt.Println("替换掉第二个之后", v, freq, self.tmpFlushs)
//			//替换掉第三个
//			//fmt.Println("替换掉最后一个", v+2, freq)
//			freq[v+2] += 1
//			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(2)
//			self.jokerLeft--
//			//fmt.Println("替换掉最后一个", v, self.tmpFlushs)
//
//			xy = self.xy
//			self.moveToNext()
//			self.travels(freq)
//			self.xy = xy
//			self.jokerLeft++
//			freq[v+2] -= 1
//
//			self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
//			//尝试放在头部
//			if v%18 > 1 { //针对{12,13,14}这种牌型
//				//old无法应对以下牌型
//				//per us:  39  最大分: 55   {cards[[{11,3} joker {13,3}] [joker {2,4} {3,4} {4,4} {5,4}]] score:55 len:8} 最大长度 {cards[] score:0 len:0} dfsTimes: 0
//				//per us: 101  最大分: 58   [[Joker ♦2 ♦3] [♦3 ♦4 ♦5] [♣J Joker ♣K]] 49 89
//				//[{10,1} {7,2} {3,3} {8,3} {11,3} {13,3} {2,4} {3,4} {3,4} {4,4} {5,4} {13,4}]
//				//outCards.FlushAddHead()
//				//self.tmpFlushs.updateScoreLen()
//				//outCards.addJokerIndex(0)
//				//self.jokerLeft--
//				//xy := self.xy
//				//self.moveToNext()
//				//self.travels(freq)
//				//self.xy = xy
//				//self.jokerLeft++
//				//outCards.removeLastJokerIndex()
//				//outCards.FlushRemoveHead()
//				//
//				//self.tmpFlushs.updateScoreLen()
//				self.need[v+3] -= 1 //#对v+3的需求加1
//				self.needidIndex[v+3]--
//				self.needid[v+3][self.needidIndex[v+3]] = 0
//
//				self.need[v+2] += 1
//				self.needid[v+2][self.needidIndex[v+2]] = self.tmpFlushs.projLen - 1
//				self.needidIndex[v+2]++
//
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushRemoveTail()
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushAddHead()
//				self.tmpFlushs.updateScoreLen()
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].addJokerIndex(0)
//
//				freq[v+2] += 1
//				self.jokerLeft--
//				xy := self.xy
//				//self.moveToNext()
//				self.travels(freq)
//				self.xy = xy
//				self.jokerLeft++
//				freq[v+2] -= 1
//
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushAddTailAuto()
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushRemoveHead()
//				self.tmpFlushs.updateScoreLen()
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
//
//				self.need[v+2] -= 1
//				self.needidIndex[v+2]--
//				self.needid[v+2][self.needidIndex[v+2]] = 0
//
//				self.need[v+3] += 1
//				self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen - 1
//				self.needidIndex[v+3]++
//			}
//
//			//尝试放在末尾
//			if (v+3)%18 <= 14 { //针对{{1, 1}, {2, 1}, {3, 1}, {1, 1}, {2, 1}, {3, 1}}这种类型
//
//				self.need[v+3] -= 1 //#对v+3的需求加1
//				self.needidIndex[v+3]--
//				self.needid[v+3][self.needidIndex[v+3]] = 0
//
//				self.need[v+4] += 1
//				//if v+4 == 8 {
//				//	fmt.Println("v4xpttttttttt退掉", self.need[8])
//				//}
//				self.needid[v+4][self.needidIndex[v+4]] = self.tmpFlushs.projLen - 1
//				self.needidIndex[v+4]++
//				//fmt.Println("here2")
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushAddTail(v+3, colorVal)
//				self.tmpFlushs.updateScoreLen()
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].addJokerIndex(3)
//				self.jokerLeft--
//				xy := self.xy
//				self.moveToNext()
//				self.travels(freq)
//				self.xy = xy
//				self.jokerLeft++
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].FlushRemoveTail()
//				self.tmpFlushs.updateScoreLen()
//				self.need[v+4] -= 1
//				self.needidIndex[v+4]--
//				self.needid[v+4][self.needidIndex[v+4]] = 0
//
//				self.need[v+3] += 1
//				//if v+3 == 8 {
//				//	fmt.Println("v3pttttttttt退掉", self.need[8])
//				//}
//				self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen - 1
//				self.needidIndex[v+3]++
//			}
//
//		}
//
//		freq[v] = f
//		freq[v+1] = f1
//		freq[v+2] = f2
//		self.need[v+3] = n3
//		//if v+3 == 8 {
//		//	fmt.Println("v38pttttttttt退掉", self.need[8], n3)
//		//}
//		self.needidIndex[v+3]--
//		self.needid[v+3][self.needidIndex[v+3]] = 0
//		self.tmpFlushs.removeLastProject()
//	}
//	if cnt == 2 { //分两种情况，一种用joker在这里一种是不用
//		//fmt.Println("2个的新组合", self.xy, v%18, v/18+1, freq[v], freq[v+1], freq[v+2], self.jokerLeft)
//		xy := self.xy
//		self.moveToNext()
//		self.travels(freq)
//		self.xy = xy
//		//fmt.Println("-2个的新组合", self.xy, v%18, v/18+1, freq[v], freq[v+1], freq[v+2], self.jokerLeft)
//		if self.jokerLeft > 0 {
//
//			if freq[v] == 0 {
//
//			} else if freq[v+1] == 0 {
//				//joker作为中间一个
//				a, b, c := freq[v], freq[v+1], freq[v+2]
//				if freq[v] > 0 {
//					freq[v] -= 1
//				}
//
//				if freq[v+1] > 0 {
//					freq[v+1] -= 1
//				}
//
//				if freq[v+2] > 0 {
//					freq[v+2] -= 1
//				}
//
//				self.need[v+3] += 1 //#对v+3的需求加1
//				//if v+3 == 8 {
//				//	fmt.Println("2pttttttttt退掉", self.need[8])
//				//}
//				self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen
//				self.needidIndex[v+3]++
//
//				outCards = NewFlushProject(v, v+2, colorVal)
//				outCards.addJokerIndex(v + 1)
//				self.tmpFlushs.updateScoreLen()
//				outCards.replaceLastJokerIndex(1)
//				self.tmpFlushs.appendProject(outCards)
//				//fmt.Println("两个的开始凑集合1", self.xy, freq, self.tmpFlushs)
//
//				self.jokerLeft--
//				xy := self.xy
//				self.moveToNext()
//				self.travels(freq)
//				self.xy = xy
//				//fmt.Println("两个的开始凑集合1 结束后", index, freq, self.tmpFlushs)
//				self.jokerLeft++
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
//				self.tmpFlushs.updateScoreLen()
//				self.tmpFlushs.removeLastProject()
//				self.needidIndex[v+3]--
//				self.needid[v+3][self.needidIndex[v+3]] = 0
//				self.need[v+3] -= 1
//				freq[v], freq[v+1], freq[v+2] = a, b, c
//				//fmt.Println("步出两个开始凑集合", self.tmpFlushs, self.need, freq)
//			} else if freq[v+2] == 0 {
//				//joker为最后一个
//				//fmt.Println("2v3pttttttttt退掉", self.tmpFlushs, v%18)
//				if (v+3)%18 <= 15 {
//					a, b, c := freq[v], freq[v+1], freq[v+2]
//					if freq[v] > 0 {
//						freq[v] -= 1
//					}
//
//					if freq[v+1] > 0 {
//						freq[v+1] -= 1
//					}
//
//					if freq[v+2] > 0 {
//						freq[v+2] -= 1
//					}
//
//					self.need[v+3] += 1 //#对v+3的需求加1
//					//if v+1 == 8 {
//
//					//}
//					self.needid[v+3][self.needidIndex[v+3]] = self.tmpFlushs.projLen
//					self.needidIndex[v+3]++
//
//					outCards = NewFlushProject(v, v+2, colorVal)
//					outCards.addJokerIndex(v + 1)
//					self.tmpFlushs.appendProject(outCards)
//					self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(2)
//
//					//fmt.Println("两个的开始凑集合2尾部", freq, self.tmpFlushs, self.xy, self.cardsXY, self.jokerLeft)
//
//					self.jokerLeft--
//					xy := self.xy
//					self.moveToNext()
//					self.travels(freq)
//					self.xy = xy
//					self.jokerLeft++
//					self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
//					self.tmpFlushs.removeLastProject()
//					self.needidIndex[v+3]--
//					self.needid[v+3][self.needidIndex[v+3]] = 0
//					self.need[v+3] -= 1
//					freq[v], freq[v+1], freq[v+2] = a, b, c
//					//fmt.Println("两个的开始凑集合2尾部之后", self.xy, freq, self.tmpFlushs)
//
//				}
//
//				//joker作为最前面一个
//				a, b, c := freq[v], freq[v+1], freq[v+2]
//				if freq[v] > 0 {
//					freq[v] -= 1
//				}
//
//				if freq[v+1] > 0 {
//					freq[v+1] -= 1
//				}
//
//				if freq[v+2] > 0 {
//					freq[v+2] -= 1
//				}
//
//				self.need[v+2] += 1 //#对v+3的需求加1
//				//if v+2 == 8 {
//				//	fmt.Println("22v2pttttttttt退掉", self.need[8])
//				//}
//				self.needid[v+2][self.needidIndex[v+2]] = self.tmpFlushs.projLen
//				self.needidIndex[v+2]++
//
//				outCards = NewFlushProject(v-1, v+1, colorVal)
//				outCards.addJokerIndex(v)
//				self.tmpFlushs.appendProject(outCards)
//
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].replaceLastJokerIndex(0)
//				//fmt.Println("两个的开始凑集合2头部", self.xy, freq, self.tmpFlushs, v)
//				self.jokerLeft--
//				xy = self.xy
//				self.moveToNext()
//				self.travels(freq)
//				self.xy = xy
//				self.jokerLeft++
//				self.tmpFlushs.projs[self.tmpFlushs.projLen-1].removeLastJokerIndex()
//				self.tmpFlushs.removeLastProject()
//				self.needidIndex[v+2]--
//				self.needid[v+2][self.needidIndex[v+2]] = 0
//				self.need[v+2] -= 1
//				freq[v], freq[v+1], freq[v+2] = a, b, c
//
//			} else {
//				panic("")
//			}
//
//			//self.tmpFlushs.removeLastProject()
//			self.tmpFlushs.updateScoreLen()
//		}
//		//有和无joker的情况下不在这里补joker
//
//	}
//
//	//正常往后迭代//独立的牌
//	//fmt.Println("单独一张牌", v%18, v/18)
//	if cnt == 1 {
//		//fmt.Println("单独一张牌", v%18, v/18)
//		self.moveToNext()
//		self.travels(freq)
//	}
//
//	//	if need[v] > 0 {
//	//		//:#先判断v是否能接到其他子序列后面
//	//
//	//		freq[v] -= 1   //#用掉一个v
//	//		need[v] -= 1   //#对v的需求减1
//	//		need[v+1] += 1 //#对v + 1
//	//		llast := needidIndex[v] - 1
//	//
//	//		needid[v+1][needidIndex[v+1]] = needid[v][llast] // append(needid[v+1], needid[v][llast]) //todo 是否存在两个
//	//		needidIndex[v+1]++
//	//		cc.projs[needid[v][llast]].FlushAddTail(v, colorVal)
//	//
//	//		needidIndex[v]--
//	//
//	//	} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 { //#v作为开头，新建一个长度为3的子序列[v, v + 1, v + 2]
//	//
//	//		freq[v] -= 1
//	//		freq[v+1] -= 1
//	//		freq[v+2] -= 1
//	//		need[v+3] += 1 //#对v+3的需求加1
//	//		needid[v+3][needidIndex[v+3]] = len(cc.projs)
//	//		needidIndex[v+3]++
//	//		outCards = NewFlushProject(v, v+2, colorVal)
//	//
//	//		cc.appendProject(outCards)
//	//
//	//	} else {
//	//		//cc = append(cc, outCards)
//	//		////fmt.Println("XXXXXXXXXXXXXXXX  here!")
//	//	}
//	//}
//	//cc.updateScoreLen()
//	//
//	////llog(-100, "\nget0FlushProjects freqSrc(%v) freq(%v) Cards(%v) cc is(%v)\n", freqSrc, freq, cards, cc) //{1,1} {4,1} {6,2} {9,2} {10,2} {1,3} {1,3} {2,3} {3,3} {10,3} {1,4} {2,4} {12,4}
//	//
//	//if cc.len > 0 {
//	//	for i := 0; i < len(cc.projs); i++ {
//	//		if len(cc.projs[i].tagIndex) > 1 { //碰到一个project中有两个joker的需要分裂
//	//			if !cc.FlushTrySplit(i) {
//	//				llog(-100, "\ngetFlushProjects freqSrc(%v) freq(%v) Cards(%v) cc is(%v)\n", freqSrc, freq, cards, cc) //{1,1} {4,1} {6,2} {9,2} {10,2} {1,3} {1,3} {2,3} {3,3} {10,3} {1,4} {2,4} {12,4}
//	//				return nil
//	//			}
//	//		}
//	//	}
//	//	//llog(-100, "\nget2FlushProjects freqSrc(%v) freq(%v) Cards(%v) cc is(%v)\n", freqSrc, freq, cards, cc) //{1,1} {4,1} {6,2} {9,2} {10,2} {1,3} {1,3} {2,3} {3,3} {10,3} {1,4} {2,4} {12,4}
//	//	return &cc
//	//}
//
//}
