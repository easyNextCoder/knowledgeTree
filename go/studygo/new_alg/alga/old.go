package alga

import "sort"

func (self *AlgX) searchBombOld(cards []Card, inputMap [][2]int16) [][]Card {

	//fmt.Println("searchBomb input cards", cards)

	if len(cards) <= 0 {
		panic("searchBomb input cards error")
	}

	colorFreq := make(map[int16]int16)

	for _, c := range cards {
		if inputMap[CARD_OFFSET(&c)][1]+inputMap[CARD_OFFSET(&c)][0] > 0 { // 这里用map得到false会产生巨大错误//todo !!
			//inputMap 是true的时候给bomb用
			if inputMap[CARD_OFFSET(&c)][1] == 0 {
				continue
			}
			colorFreq[c.Second] = inputMap[CARD_OFFSET(&c)][1]

			continue
		}
		colorFreq[c.Second]++

	}
	//fmt.Println("inputMap", inputMap, colorFreq, cards)

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Second == cards[j].Second {
			return cards[i].First < cards[j].First
		} else if cards[i].Second < cards[j].Second {
			return true
		}
		return false
	})

	cardVal := cards[0].First
	var cc [][]Card
	var c []Card

	cnt := int16(0)
	for _, v := range colorFreq {
		cnt += v
	}

	if cnt == 8 {
		e := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
		f := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
		cc = append(cc, e)
		cc = append(cc, f)
		return cc
	}

	switch len(colorFreq) { //那现在牌值相同的牌求得最多的炸弹个数和分数
	case 3:
		switch len(cards) {
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			for color, _ := range colorFreq {
				c = append(c, Card{cardVal, color})
			}
			cc = append(cc, c)
		case 6:
			for color, _ := range colorFreq {
				c = append(c, Card{cardVal, color})
			}
			cc = append(cc, c)
			c = make([]Card, 3)
			copy(c, cc[0])
			cc = append(cc, c)
		default:
			panic("colorFreq 3 len(cards) not right")
		}
	case 4:
		switch len(cards) {
		case 4:
			fallthrough
		case 5:
			for color, _ := range colorFreq {
				c = append(c, Card{cardVal, color})
			}
			cc = append(cc, c)
		case 6: //6张牌可以拆成2个炸弹
			var c2 []Card
			for color, cnt := range colorFreq {
				if cnt == 2 {
					c2 = append(c2, Card{cardVal, color})
					colorFreq[color]--
				}
			}
			for color, _ := range colorFreq {
				if len(c) == 3 {
					c2 = append(c2, Card{cardVal, color})
					break
				}
				c = append(c, Card{cardVal, color})
			}
			cc = append(cc, c, c2)
		case 7:
			fallthrough
		case 8:
			var c2 []Card
			for color, cnt := range colorFreq {
				card := Card{cardVal, color}
				c = append(c, card)
				if cnt > 1 {
					c2 = append(c2, card)
				}
			}
			cc = append(cc, c, c2)
		}
	}

	return cc
}

func (self *AlgX) searchFlushOld(cards []Card, inputMap [][2]int16) [][]Card {

	need := [60]int{}   //make(map[int16]int, 0)
	needid := [20]int{} //make(map[int16]int, 0)
	freq := [20]int{}   //make(map[int16]int, 0)

	for _, v := range cards {

		if inputMap[CARD_OFFSET(&v)][0]+inputMap[CARD_OFFSET(&v)][1] > 0 { //[0]代表给顺子
			if int(inputMap[CARD_OFFSET(&v)][0]) == 0 {
				continue
			}
			freq[v.First] = int(inputMap[CARD_OFFSET(&v)][0])
			continue
		}

		freq[v.First] += 1
	}

	colorVal := cards[0].Second

	var outCards []Card
	var cc [][]Card
	//fmt.Println("outCards is:", cards)
	for _, v := range cards {
		v := v.First
		//fmt.Println("v is:", v)
		if freq[v] == 0 {
			//:#已经被用到其他子序列中
			continue
		}

		if need[v] > 0 {
			//:#先判断v是否能接到其他子序列后面
			freq[v] -= 1   //#用掉一个v
			need[v] -= 1   //#对v的需求减1
			need[v+1] += 1 //#对v + 1
			needid[v+1] = len(cc) - 1
			cc[needid[v]] = append(cc[needid[v]], Card{v, colorVal})
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 { //#v作为开头，新建一个长度为3的子序列[v, v + 1, v + 2]

			freq[v] -= 1
			freq[v+1] -= 1
			freq[v+2] -= 1
			need[v+3] += 1 //#对v+3的需求加1
			needid[v+3] = len(cc)
			outCards = make([]Card, 0)
			outCards = append(outCards, Card{v, colorVal}, Card{v + 1, colorVal}, Card{v + 2, colorVal})
			cc = append(cc, outCards)

		} else {
			//cc = append(cc, outCards)
			//fmt.Println("XXXXXXXXXXXXXXXX  here!")
		}
	}
	//if len(outCards) > 0 {
	//	cc = append(cc, outCards)
	//}
	//fmt.Println("this is flushSearch cc", cc)
	//在每一个project的后面寻找frep[max+1]如果存在就还往后找，记录下这个id，之后用散牌看看拼凑出组合将组合记录下来
	return cc
}

func searchFlushOld(cards []Card, inputMap [][2]int16) [][]Card {

	need := [60]int{}   //make(map[int16]int, 0)
	needid := [20]int{} //make(map[int16]int, 0)
	freq := [20]int{}   //make(map[int16]int, 0)

	for _, v := range cards {

		if inputMap[CARD_OFFSET(&v)][0]+inputMap[CARD_OFFSET(&v)][1] > 0 { //[0]代表给顺子
			if int(inputMap[CARD_OFFSET(&v)][0]) == 0 {
				continue
			}
			freq[v.First] = int(inputMap[CARD_OFFSET(&v)][0])
			continue
		}

		freq[v.First] += 1
	}

	colorVal := cards[0].Second

	var outCards []Card
	var cc [][]Card
	//fmt.Println("outCards is:", cards)
	for _, v := range cards {
		v := v.First
		//fmt.Println("v is:", v)
		if freq[v] == 0 {
			//:#已经被用到其他子序列中
			continue
		}

		if need[v] > 0 {
			//:#先判断v是否能接到其他子序列后面
			freq[v] -= 1   //#用掉一个v
			need[v] -= 1   //#对v的需求减1
			need[v+1] += 1 //#对v + 1
			needid[v+1] = len(cc) - 1
			cc[needid[v]] = append(cc[needid[v]], Card{v, colorVal})
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 { //#v作为开头，新建一个长度为3的子序列[v, v + 1, v + 2]

			freq[v] -= 1
			freq[v+1] -= 1
			freq[v+2] -= 1
			need[v+3] += 1 //#对v+3的需求加1
			needid[v+3] = len(cc)
			outCards = make([]Card, 0)
			outCards = append(outCards, Card{v, colorVal}, Card{v + 1, colorVal}, Card{v + 2, colorVal})
			cc = append(cc, outCards)

		} else {
			//cc = append(cc, outCards)
			//fmt.Println("XXXXXXXXXXXXXXXX  here!")
		}
	}
	//if len(outCards) > 0 {
	//	cc = append(cc, outCards)
	//}
	//fmt.Println("this is flushSearch cc", cc)
	//在每一个project的后面寻找frep[max+1]如果存在就还往后找，记录下这个id，之后用散牌看看拼凑出组合将组合记录下来
	return cc
}

func getFlushProjectsOld(freqSrc []int, cards []Card, colorVal int16) [][]Card {

	freq := [16]int16{}
	for k, v := range freqSrc {
		freq[k] = int16(v)
	}
	//copy(freq, freqSrc)

	var cc [][]Card
	var outCards []Card
	need := [16]int{}     //make(map[int16]int, 0)
	needid := [16][]int{} //make(map[int16]int, 0)

	//freq := make([]int, 20) //[20]int{}   //make(map[int16]int, 0)
	for _, v := range cards {
		v := v.First
		////fmt.Println("v is:", v)
		if freq[v] == 0 {
			//:#已经被用到其他子序列中
			continue
		}
		//if v == 14 {
		//	llog(-1, "\nXXXXXXXXXXXXX cards %v freq %v v %v, need %v needid %v cc %v\n", cards, freq, v, need, needid, cc)
		//}

		if need[v] > 0 {
			//:#先判断v是否能接到其他子序列后面
			freq[v] -= 1   //#用掉一个v
			need[v] -= 1   //#对v的需求减1
			need[v+1] += 1 //#对v + 1
			llast := len(needid[v]) - 1
			needid[v+1] = append(needid[v+1], needid[v][llast]) //todo 是否存在两个
			cc[needid[v][llast]] = append(cc[needid[v][llast]], Card{v, colorVal})
			needid[v] = needid[v][:llast]
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 { //#v作为开头，新建一个长度为3的子序列[v, v + 1, v + 2]

			freq[v] -= 1
			freq[v+1] -= 1
			freq[v+2] -= 1
			need[v+3] += 1 //#对v+3的需求加1
			needid[v+3] = append(needid[v+3], len(cc))
			outCards = make([]Card, 0, 15)
			outCards = append(outCards, Card{v, colorVal}, Card{v + 1, colorVal}, Card{v + 2, colorVal})
			cc = append(cc, outCards)

		} else {
			//cc = append(cc, outCards)
			////fmt.Println("XXXXXXXXXXXXXXXX  here!")
		}
	}
	//llog(-1, "\ngetFlushProjects freqSrc(%v) freq(%v) Cards(%v) cc is(%v)\n", freqSrc, freq, cards, cc)
	if len(cc) > 0 {
		return cc
	}

	return nil
}

//func dfsOld(index int){
//	if index == len(xxx){
//		self.dfsTimes++
//		//var cc [][]Card
//		//for _, cards := range flushCards {
//		//	cc = append(cc, self.searchFlushOld(cards, states)...)
//		//}
//		////fmt.Println("This is CC before", cc)
//		////fmt.Println("this is bombCards", bombCards)
//		//for _, cards := range bombCards {
//		//	cc = append(cc, self.searchBombOld(cards, states)...)
//		//}
//		////fmt.Println("This is CC:", cc)
//		//score := 0
//		//for _, cards := range cc {
//		//	for _, card := range cards {
//		//		score += scoreArr[card.first]
//		//	}
//		//}
//		//
//		//if score > maxScoreProjects.score {
//		//	maxScoreProjects.score = score
//		//	maxScoreProjects.proj = cc
//		//}
//		//
//		//return
//
//		var idx int
//		var noJoker [20]*Projects
//		var oneJoker [20]*Projects
//		var twoJoker [20]*Projects
//
//		for _, cards := range self.flushCards {
//			//if true {
//			//llog(-100, "HHHandle flush cards %v\n", cards)
//			//	for _, card := range cards {
//			//		llog(-1, "states %v", states[CARD_OFFSET(&card)])
//			//	}
//			//}
//
//			a, b, c := self.searchFlushNew(cards, self.states)
//			//if a != nil && len(a) > 0 {
//			noJoker[idx] = a
//			//}
//
//			//if b != nil && len(b) > 0 {
//			oneJoker[idx] = b
//
//			//}
//			//if c != nil && len(c) > 0 {
//			twoJoker[idx] = c
//
//			idx++
//
//			if !(idx < 20) {
//				panic("too long")
//			}
//			//}
//			//fmt.Printf("start dfs searchFlush no(%v)\n one(%v)\n two(%v)\n\n", a, b, c)
//		}
//
//		if self.jokerLeft > 0 && len(noJoker) != len(oneJoker) /*|| len(noJoker) != len(twoJoker) */ {
//			//log("nojoker %v onejoker %v twojoker %v", noJoker, oneJoker, twoJoker)
//			panic("joker len not equal")
//		}
//
//		//llog(-1, "8777777 %v jokerLeft %v  nojoker %v", tmpMaxScoreProjects, self.jokerLeft, noJoker)
//		tmpMaxScoreProjectsLen := len(self.tmpMaxScoreProjects.projs)
//		tmpMaxLenProjectsLen := len(self.tmpMaxLenProjects.projs)
//		var cc *Projects
//		cc = self.searchBomb(self.bombCards, self.states, 0) //todo 这里废掉了在bomb中加入joker的情况20230313-22:49
//		//fmt.Printf("final dfs searchFlush no(%v)\n one(%v)\n two(%v)\n cc(%v)\n\n", noJoker, oneJoker, twoJoker, cc)
//		self.tmpMaxScoreProjects.merge(cc)
//		self.tmpMaxLenProjects.merge(cc)
//		//llog(-100, "after search bomb cc %s tmpMaxLenProjects %s tmpMaxLenProjects %s nojoker %s\n", cc, self.tmpMaxLenProjects, self.tmpMaxLenProjects, noJoker)
//		switch self.jokerLeft {
//		case 0:
//			//llog(-1, "self.jokerLeft = 0\n")
//			var CC Projects
//			for _, p := range noJoker {
//				CC.merge(p)
//			}
//			self.tmpMaxScoreProjects.merge(&CC)
//			self.tmpMaxLenProjects.merge(&CC)
//		case 1:
//			//llog(-1, "self.jokerLeft = 1\n")
//			var noJokerMaxScore, tmpOneJokerMaxScore int16
//			var noJokerMaxLen, tmpOneJokerMaxLen int16
//			var noJokerRowScore []int16
//			var noJokerRowLen []int16
//			var oneJokerCards Projects
//			var oneJokerCardsLen Projects
//
//			for _, row := range noJoker {
//				var s, l int16
//				if row != nil {
//					s = row.score
//					l = row.len
//				}
//				noJokerRowScore = append(noJokerRowScore, s)
//				noJokerMaxScore += s
//				noJokerRowLen = append(noJokerRowLen, l)
//				noJokerMaxLen += l
//			}
//			tmpOneJokerMaxScore, tmpOneJokerMaxLen = noJokerMaxScore, noJokerMaxLen
//			theMaxIndex, theMaxLenIndex := 0, 0
//			for i, p := range oneJoker {
//				var pscore, plen int16
//				if p != nil {
//					pscore = p.score
//					plen = p.len
//				}
//
//				s := noJokerMaxScore - noJokerRowScore[i] + pscore
//
//				if s > tmpOneJokerMaxScore {
//					tmpOneJokerMaxScore = s
//					theMaxIndex = i
//				}
//
//				l := noJokerMaxLen - noJokerRowLen[i] + plen
//				if l > tmpOneJokerMaxLen {
//					tmpOneJokerMaxLen = l
//					theMaxLenIndex = i
//				}
//			}
//			//fmt.Println("noJoker", noJoker)
//			for i, p := range noJoker {
//				if i != theMaxIndex {
//					oneJokerCards.merge(p)
//				} else {
//					oneJokerCards.merge(oneJoker[i])
//				}
//
//				if i != theMaxLenIndex {
//					oneJokerCardsLen.merge(p)
//				} else {
//					oneJokerCardsLen.merge(oneJoker[i])
//				}
//
//			}
//
//			self.tmpMaxScoreProjects.merge(&oneJokerCards)  //  .proj = append(self.tmpMaxScoreProjects.proj, oneJokerCards...)
//			self.tmpMaxLenProjects.merge(&oneJokerCardsLen) //.proj = append(self.tmpMaxLenProjects.proj, oneJokerCardsLen...)
//			//llog(-1, "\nnojoker(%v) oneJoker(%v) twoJoker(%v) oneJokerCards(%v) bombTakeAheadMaxScoreProjects(%v) tmpMaxScoreProjects(%v)\n", noJoker, oneJoker, twoJoker, oneJokerCards, bombTakeAheadMaxScoreProjects, tmpMaxScoreProjects)
//
//		case 2:
//			//llog(-1, "self.jokerLeft = 2\n")
//			var noJokerMaxScore, noJokerMaxLen, tmpTwoJokerMaxScore, tmpTwoJokerMaxLen int16
//			var noJokerRowScore, noJokerRowLen []int16
//			var twoJokerCards1, twoJokerCards2, twoJokerMaxScoreCards Projects
//			var twoJokerLenCards1, twoJokerLenCards2, twoJokerMaxLenCards Projects
//
//			for _, row := range noJoker {
//				var s, l int16
//				if row != nil {
//					s = row.score
//					l = row.len
//				}
//
//				noJokerRowScore = append(noJokerRowScore, s)
//				noJokerMaxScore += s
//
//				noJokerRowLen = append(noJokerRowLen, l)
//				noJokerMaxLen += l
//			}
//
//			tmpTwoJokerMaxScore = noJokerMaxScore
//			theMaxIndex := 0
//
//			tmpTwoJokerMaxLen = noJokerMaxLen
//			theMaxLenIndex := 0
//			//第一种情况找到顺子中可以连续补两个的
//			for i, p := range twoJoker {
//
//				var pscore, plen int16
//				if p != nil {
//					pscore = p.score
//					plen = p.len
//				}
//
//				s := noJokerMaxScore - noJokerRowScore[i] + pscore
//				if s > tmpTwoJokerMaxScore {
//					tmpTwoJokerMaxScore = s
//					theMaxIndex = i
//				}
//
//				l := noJokerMaxLen - noJokerRowLen[i] + plen
//				if l > tmpTwoJokerMaxLen {
//					tmpTwoJokerMaxLen = l
//					theMaxLenIndex = i
//				}
//			}
//			for i, p := range noJoker {
//				if i != theMaxIndex {
//					twoJokerCards1.merge(p) // = append(twoJokerCards1, cardss...)
//				} else {
//					twoJokerCards1.merge(twoJoker[i]) //= append(twoJokerCards1, twoJoker[i]...)
//				}
//
//				if i != theMaxLenIndex {
//					twoJokerLenCards1.merge(p) //= append(twoJokerLenCards1, cardss...)
//				} else {
//					twoJokerLenCards1.merge(twoJoker[i]) //= append(twoJokerLenCards1, twoJoker[i]...)
//				}
//			}
//
//			//找到顺子中两个含有joker的最大的
//			var maxi, maxj, maxli, maxlj int
//			for i := 0; i < len(oneJoker); i++ {
//				for j := i + 1; j < len(oneJoker); j++ {
//					var score, l int16
//					for k := 0; k < len(noJoker); k++ {
//						if k == i || k == j {
//							if oneJoker[k] != nil {
//								l += oneJoker[k].len
//								score += oneJoker[k].score
//							}
//						} else {
//							if noJoker[k] != nil {
//								l += noJoker[k].len
//								score += noJoker[k].score
//							}
//						}
//					}
//					if score > tmpTwoJokerMaxScore {
//						tmpTwoJokerMaxScore = score
//						maxi = i
//						maxj = j
//					}
//
//					if l > tmpTwoJokerMaxLen {
//						tmpTwoJokerMaxLen = l
//						maxli = i
//						maxlj = j
//					}
//
//				}
//			}
//
//			//llog(-1, "\nxoneJoker %v\n", oneJoker)
//
//			for i, p := range noJoker {
//				if i != maxi && i != maxj {
//					twoJokerCards2.merge(p)
//				} else {
//					twoJokerCards2.merge(oneJoker[i])
//				}
//
//				if i != maxli && i != maxlj {
//					twoJokerLenCards2.merge(p)
//				} else {
//					twoJokerLenCards2.merge(oneJoker[i])
//				}
//			}
//
//			score1 := twoJokerCards1.score
//			score2 := twoJokerCards2.score
//
//			l1, l2 := twoJokerLenCards1.len, twoJokerLenCards2.len
//
//			if score1 > score2 {
//				twoJokerMaxScoreCards = twoJokerCards1
//			} else {
//				twoJokerMaxScoreCards = twoJokerCards2
//			}
//
//			if l1 > l2 {
//				twoJokerMaxLenCards = twoJokerLenCards1
//			} else {
//				twoJokerMaxLenCards = twoJokerLenCards2
//			}
//
//			self.tmpMaxScoreProjects.merge(&twoJokerMaxScoreCards) // .proj = append(self.tmpMaxScoreProjects.proj, twoJokerMaxScoreCards...)
//			self.tmpMaxLenProjects.merge(&twoJokerMaxLenCards)     //.proj = append(self.tmpMaxLenProjects.proj, twoJokerMaxLenCards...)
//		}
//
//		if self.bombTakeAheadMaxScoreProjects.len > 0 {
//			self.tmpMaxScoreProjects.merge(&self.bombTakeAheadMaxScoreProjects) // .proj = append(self.tmpMaxScoreProjects.proj, self.bombTakeAheadMaxScoreProjects.proj...)
//			self.tmpMaxLenProjects.merge(&self.bombTakeAheadMaxScoreProjects)   //.proj = append(self.tmpMaxLenProjects.proj, self.bombTakeAheadMaxScoreProjects.proj...)
//		}
//
//		self.updateMaxScoreProjects(&self.tmpMaxScoreProjects)
//		self.updateMaxLenProjects(&self.tmpMaxLenProjects)
//
//		llog(-1, "dfs final bombTakeAheadMaxScoreProjects(%v) tmpMaxS%v  maxScoreProjects%v searchBomb %v\n", self.bombTakeAheadMaxScoreProjects, self.tmpMaxScoreProjects, self.maxScoreProjects, cc)
//		self.tmpMaxScoreProjects.projs = self.tmpMaxScoreProjects.projs[:tmpMaxScoreProjectsLen]
//
//		self.tmpMaxLenProjects.projs = self.tmpMaxLenProjects.projs[:tmpMaxLenProjectsLen]
//		self.tmpMaxScoreProjects.updateScoreLen()
//		self.tmpMaxLenProjects.updateScoreLen()
//	}
//
//}

//func (self *AlgX) searchBomb(bombCards [][]Card, inputMap [][2]int16, jokerNum int) *Projects {
//
//	switch jokerNum {
//	case 0:
//		var cc []*Project
//		for _, cards := range bombCards {
//			//fmt.Println("searchBombs", cards)
//			cc = append(cc, self.searchBombWork(cards, self.states, 0)...)
//
//		}
//		var res Projects
//
//		res.appendProject(cc...)
//		//log("searchBomb() 0 joker cc %v", cc)
//		return &res
//	case 1:
//		panic("jokerNum can't be 1")
//		//var cc [][]Card
//		//columnScoreNoJoker := make([]int, len(bombCards))
//		//var totalScoreNoJoker int
//		//
//		//noJokerProjects := make([][][]Card, len(bombCards))
//		//
//		//for i, cards := range bombCards {
//		//	tmp := self.searchBombWork(cards, self.states, 0)
//		//	columnScoreNoJoker[i] = scoreProjects(tmp)
//		//	totalScoreNoJoker += columnScoreNoJoker[i]
//		//	noJokerProjects[i] = tmp
//		//}
//		//
//		//g := 0
//		//maxScoreProjectsWithJokerIndex := -1
//		//var maxScoreProjectsWithJoker [][]Card
//		//for i, cards := range bombCards {
//		//	withJoker := self.searchBombWork(cards, self.states, jokerNum)
//		//
//		//	score := scoreProjects(withJoker)
//		//	if score-columnScoreNoJoker[i] > g {
//		//		g = score - columnScoreNoJoker[i]
//		//		maxScoreProjectsWithJokerIndex = i
//		//		maxScoreProjectsWithJoker = withJoker
//		//	}
//		//}
//		//
//		//for i := range bombCards {
//		//	thisColumn := noJokerProjects[i]
//		//	if maxScoreProjectsWithJokerIndex == i {
//		//		thisColumn = maxScoreProjectsWithJoker
//		//	}
//		//	if len(thisColumn) > 0 {
//		//		cc = append(cc, thisColumn...)
//		//	}
//		//}
//		//log("searchBomb() 1 joker cc %v", cc)
//		//return cc
//	}
//	return nil
//}
