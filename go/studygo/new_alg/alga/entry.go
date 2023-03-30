package alga

func (self *AlgX) Entry(deck Deck, jokerNum int) {

	self.hand_cards = deck
	self.jokerLeft = jokerNum

	var tag [15][5]int

	for i := 0; i < len(self.hand_cards); i++ {

		tag[self.hand_cards[i].First][self.hand_cards[i].Second]++
	}

	var cnt int
	var i, j int16
	for i = 1; i < 15; i++ {
		for j = 0; j < 5; j++ {
			n := tag[i][j]
			for ; n > 0; n-- {
				self.hand_cards[cnt] = Card{i, j}
				cnt++
			}
		}
	}

	//fmt.Println("bomb:", self.hand_cards)
	//sort.Slice(self.hand_cards, func(i, j int) bool {
	//	if self.hand_cards[i].First == self.hand_cards[j].First {
	//		return self.hand_cards[i].Second < self.hand_cards[j].Second
	//	} else if self.hand_cards[i].First < self.hand_cards[j].First {
	//		return true
	//	}
	//	return false
	//})

	//self.states = make([][2]int16, 60) //[0]代表顺子，全部给顺子
	//self.bombCards = [15][]Card{}
	//self.flushCards = [5][]Card{}
	//self.recordUsed = [60]bool{}
	//self.bombRepeatMap = [15]Bomb{}
	self.bombTakeAheadMaxScoreProjects = Projects{}

	self.tmpMaxScoreProjects = Projects{}
	self.maxScoreProjects = Projects{}
	self.tmpMaxLenProjects = Projects{}
	self.maxLenProjects = Projects{}

	self.tmpFlushs = Projects{}

	//self.cardsXY = make([]int, 72)
	//self.freq = make([]int, 72)

	self.findAllBomb()

	for i := 0; i < self.bombLen; i++ {
		bombcs := self.bombCards[i]
		bp := Bomb{}

		for j := 0; j < self.bombCardsLen[i]; j++ {
			c := bombcs[j]
			bp.colors[c.Second]++
			bp.cnt++
			bp.val = c.First
		}

		for _, v := range bp.colors {
			if v > 0 {
				bp.colorCnt++
			}
		}
		self.bombRepeatMap[i] = bp
	}

	self.findAllFlush()

	//fmt.Printf("flushCards %v\n\n bombCards %v\n\n", self.flushCards, self.bombCards)
	//self.jokerLeft = 0
	self.jokerNum = self.jokerLeft

	//if self.jokerNum == 2 { //这里要进行所有牌抽出单张+2张joker凑成炸弹的预先计算
	//	self.handTwoJokerFirstNode()
	//}

	self.dfs(0)

	//self.cardsXY = nil
	//self.freq = nil
	//self.states = nil
	//self.bombTakeAheadMaxScoreProjects.projs = nil
	//return &self.maxScoreProjects, &self.maxLenProjects, self.dfsTimes
	//llog(LogLevelN100, "handleOneJoker final maxScoreProject %v ", maxScoreProjects)

}
