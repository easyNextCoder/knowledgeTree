package alga

func (self *AlgX) findAllFlush() {

	var tag [15][5]int

	for i := 0; i < len(self.hand_cards); i++ {

		tag[self.hand_cards[i].First][self.hand_cards[i].Second]++
	}

	var cnt int
	var i, j int16
	for j = 0; j < 5; j++ {
		for i = 1; i < 15; i++ {
			n := tag[i][j]
			for ; n > 0; n-- {
				self.hand_cards[cnt] = Card{i, j}
				cnt++
			}
		}
	}

	//fmt.Println("flush:", self.hand_cards)

	//sort.Slice(self.hand_cards, func(i, j int) bool {
	//
	//	if self.hand_cards[i].Second == self.hand_cards[j].Second {
	//		return self.hand_cards[i].First < self.hand_cards[j].First
	//	} else if self.hand_cards[i].Second < self.hand_cards[j].Second {
	//		return true
	//	}
	//	return false
	//})
	////fmt.Println("sorted cards before ", self.hand_cards)

	var cards [30]Card
	for i := 0; i < len(self.hand_cards); i++ {
		if self.flushCardsLen[self.flushLen] == 0 {
			cards[self.flushCardsLen[self.flushLen]] = self.hand_cards[i]
			self.flushCardsLen[self.flushLen]++
		} else {
			if /*(self.hand_cards[i].first == cards[len(cards)-1].first+1 || self.hand_cards[i].first == cards[len(cards)-1].first) && */ self.hand_cards[i].Second == cards[self.flushCardsLen[self.flushLen]-1].Second {
				cards[self.flushCardsLen[self.flushLen]] = self.hand_cards[i]
				self.flushCardsLen[self.flushLen]++
			} else {

				if self.flushCardsLen[self.flushLen] > 0 {

					for j := 0; j < self.flushCardsLen[self.flushLen]; j++ {
						c := cards[j]
						if self.recordUsed[CARD_OFFSET(&c)] {

							self.states[CARD_OFFSET(&c)][0]++
						}
					}

					self.flushCards[self.flushLen] = cards
					self.flushLen++
					cards[self.flushCardsLen[self.flushLen]] = self.hand_cards[i]
					self.flushCardsLen[self.flushLen]++
				} else {
					cards[self.flushCardsLen[self.flushLen]] = self.hand_cards[i]
					self.flushCardsLen[self.flushLen]++
				}
			}
		}
	}
	if self.flushCardsLen[self.flushLen] > 0 {
		for j := 0; j < self.flushCardsLen[self.flushLen]; j++ {
			c := cards[j]
			if self.recordUsed[CARD_OFFSET(&c)] {

				self.states[CARD_OFFSET(&c)][0]++
			}
		}
		self.flushCards[self.flushLen] = cards
		self.flushLen++
	}

}

func (self *AlgX) findAllBomb() {

	//sort.Slice(self.hand_cards, func(i, j int) bool {
	//
	//	if self.hand_cards[i].second == self.hand_cards[j].second {
	//		return self.hand_cards[i].first < self.hand_cards[j].first
	//	} else if self.hand_cards[i].second < self.hand_cards[j].second {
	//		return true
	//	}
	//	return false
	//})

	var cards [9]Card

	////fmt.Println("findAllBomb ori cards", self.hand_cards)
	for i := 0; i < len(self.hand_cards); i++ {

		if self.bombCardsLen[self.bombLen] == 0 {
			cards[self.bombCardsLen[self.bombLen]] = self.hand_cards[i]
			self.bombCardsLen[self.bombLen]++
		} else {
			if self.hand_cards[i].First == cards[self.bombCardsLen[self.bombLen]-1].First {
				cards[self.bombCardsLen[self.bombLen]] = self.hand_cards[i]
				self.bombCardsLen[self.bombLen]++
			} else {

				if self.bombCardsLen[self.bombLen] > 0 { //检查至少有三种颜色，这一列牌才有可能和顺子共用，这样才去做选择

					for j := 0; j < self.bombCardsLen[self.bombLen]; j++ {
						c := cards[j]
						self.recordUsed[CARD_OFFSET(&c)] = true
					}

					self.bombCards[self.bombLen] = cards
					self.bombLen++

					cards[self.bombCardsLen[self.bombLen]] = self.hand_cards[i]
					self.bombCardsLen[self.bombLen]++

				} else {
					cards[self.bombCardsLen[self.bombLen]] = self.hand_cards[i]
					self.bombCardsLen[self.bombLen]++
				}
			}
		}
	}
	if self.bombCardsLen[self.bombLen] > 0 {
		for j := 0; j < self.bombCardsLen[self.bombLen]; j++ {
			c := cards[j]
			self.recordUsed[CARD_OFFSET(&c)] = true
		}
		self.bombCards[self.bombLen] = cards
		self.bombLen++
	}

	//if OpenPrint {
	//	cnt := 0
	//	check := make(map[int16]bool)
	//	for _, cards := range self.bombCards {
	//		cnt += len(cards)
	//		for k := range check {
	//			delete(check, k)
	//		}
	//		for _, card := range cards {
	//			check[card.First] = true
	//		}
	//		if len(check) > 1 {
	//			panic("findAllBombsJ not same color")
	//		}
	//
	//	}
	//
	//}
	////fmt.Println("findAllBomb:", bombCards)

}
