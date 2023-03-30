package alga

type AlgX struct {
	hand_cards                        Deck
	jokerLeft                         int
	jokerNum                          int
	addJokerToColumn                  int
	haveProcessedInOneJokerMiddleNode bool

	tmpFlushs Projects
	xy        XY      //牌的坐标
	lastXy    XY      //牌的上次的坐标
	cardsXY   [72]int //所有牌的分布
	freq      [72]int

	need        [80]int    //make(map[int16]int, 0)
	needid      [80][3]int //make(map[int16]int, 0)
	needidIndex [80]int

	states                        [60][2]int16 //[0]代表给顺子[1]代表给炸弹
	bombCards                     [15][9]Card
	bombCardsLen                  [15]int
	bombLen                       int
	flushCards                    [5][30]Card
	flushCardsLen                 [5]int
	flushLen                      int
	recordUsed                    [60]bool
	bombRepeatMap                 [15]Bomb
	bombTakeAheadMaxScoreProjects Projects
	tmpMaxScoreProjects           Projects
	maxScoreProjects              Projects
	tmpMaxLenProjects             Projects
	maxLenProjects                Projects
}

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

	self.bombTakeAheadMaxScoreProjects = Projects{}

	self.tmpMaxScoreProjects = Projects{}
	self.maxScoreProjects = Projects{}
	self.tmpMaxLenProjects = Projects{}
	self.maxLenProjects = Projects{}

	self.tmpFlushs = Projects{}

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

	self.jokerNum = self.jokerLeft

	self.dfs(0)

}
