package alga

import (
	"fmt"
	_ "net/http/pprof"
)

const MAX_N int16 = 14

const MAX_M int16 = 4

const MAX_LEN int16 = MAX_N * MAX_M

func CARD_OFFSET(card *Card) int16 { return (card.First-1)*MAX_M + (card.Second - 1) }

var scoreArr [15]int = [15]int{0, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
var scoreArrInt16 [15]int16 = [15]int16{0, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

//对于有joker牌的处理，这里仍然不区分joker和joker2

type Bomb struct {
	_        int16
	val      int16    //牌值
	cnt      int16    //牌张数量
	colors   [5]int16 //颜色
	colorCnt int16    //这个炸弹中一共有多少种颜色
}

const (
	ProjectFlush = iota
	ProjectBomb
)

type Project struct { //为了计算分值和长度
	min      int16
	max      int16
	len      int16
	ptype    int16 //ptype 0表示flush 1代表bomb
	score    int16
	pcolor   [5]bool
	tagIndex [3]int16 //flush时候标记joker的位置，bomb标记joker的color，从下标0开始计数
	tagLen   int16
}

func (self *Project) FlushUpdateScoreLen() {
	self.score = 0
	for i := self.min; i <= self.max; i++ {
		self.score += scoreArrInt16[i]
	}
	self.len = self.max - self.min + 1
}

func (self Project) String() string {
	if self.ptype == ProjectFlush {
		var color int16
		for i, v := range self.pcolor {
			if v {
				color = int16(i)
				break
			}
		}
		var cards []Card
		var cnt int16
		var isJoker bool
		for i := self.min; i <= self.max; i++ {
			isJoker = false
			var j int16
			for j = 0; j < self.tagLen; j++ {
				if cnt == self.tagIndex[j] {
					isJoker = true
				}
			}
			cnt++
			if isJoker {
				cards = append(cards, Card{i, 0})
			} else {
				cards = append(cards, Card{i, color})
			}
		}
		return fmt.Sprintf("%v", cards)
	} else {
		var cards []Card
		for i, v := range self.pcolor {
			if v {
				i := int16(i)
				if self.tagLen > 2 {
					panic("bomb have two many joker")
				}
				if self.tagLen > 0 && i == self.tagIndex[0] {
					cards = append(cards, Card{self.min, 0})
					continue
				}
				cards = append(cards, Card{self.min, int16(i)})
			}
		}
		return fmt.Sprintf("%v", cards)
	}
}

func (self Projects) String() string {
	var s string
	for i := 0; i < self.projLen; i++ {
		s += fmt.Sprintf("%s", self.projs[i])
	}
	return fmt.Sprintf("{cards%s score:%d len:%d}", s, self.score, self.len)
}

func (self *Project) clone() *Project {
	var c Project
	c = *self

	return &c
}

type Projects struct {
	projs   [10]Project
	score   int16 //所有project的分数和
	len     int16 //所有project的长度的和
	projLen int
}

//func (self *Projects) clear() {
//	self.projs = self.projs[:0]
//	self.score = 0
//	self.len = 0
//}

func (self *Projects) FlushOnlyTrySplit(index int) bool {
	if self.projs[index].tagLen != 2 {
		panic("trySplit err")
	}

	var frontLen, endLen int16

	for mid := self.projs[index].tagIndex[0]; mid <= self.projs[index].tagIndex[1]; mid++ {

		frontLen, endLen = mid+1, self.projs[index].len-mid-1

		if frontLen >= 3 && endLen >= 3 {
			//self.projs[index].max = self.projs[index].min + frontLen - 1
			//self.projs[index].FlushUpdateScoreLen()
			//
			//nextProjMin := self.projs[index].max + 1
			//self.projs = append(self.projs, NewFlushProject(nextProjMin, nextProjMin+endLen-1, self.projs[index].FlushGetColor()))
			//self.projs[len(self.projs)-1].tagIndex = append(self.projs[len(self.projs)-1].tagIndex, self.projs[index].tagIndex[1]-frontLen)
			//self.projs[index].tagIndex = self.projs[index].tagIndex[:1]
			return true
		}
	}
	return false
}

func (self *Projects) FlushTrySplit(index int) bool {

	if self.projs[index].tagLen != 2 {
		panic("trySplit err")
	}

	var frontLen, endLen int16

	for mid := self.projs[index].tagIndex[0]; mid <= self.projs[index].tagIndex[1]; mid++ {

		frontLen, endLen = mid+1, self.projs[index].len-mid-1

		if frontLen >= 3 && endLen >= 3 {
			self.projs[index].max = self.projs[index].min + frontLen - 1
			self.projs[index].FlushUpdateScoreLen()
			self.updateScoreLen()
			nextProjMin := self.projs[index].max + 1
			var proj Project
			NewFlushProjectInit(nextProjMin, nextProjMin+endLen-1, self.projs[index].FlushGetColor(), &proj)
			self.appendProject(proj)
			//self.projs = append(self.projs, )
			self.projs[self.projLen-1].tagIndex[self.projs[self.projLen-1].tagLen] = self.projs[index].tagIndex[1] - frontLen
			self.projs[self.projLen-1].tagLen++
			self.projs[index].tagLen--
			return true
		}
	}
	return false
}

func (self *Projects) appendProject(p ...Project) {
	//fmt.Println("appendProject", *self, self.projLen, len(self.projs))

	if min(len(self.projs), self.projLen+len(p)) >= 10 {
		panic("")
	}

	for i := self.projLen; i < min(len(self.projs), self.projLen+len(p)); i++ {
		self.projs[i] = p[i-self.projLen]
	}

	//y := min(len(self.projs), self.projLen+len(p))
	//for i := y; i < self.projLen+len(p); i++ {
	//	self.projs = append(self.projs, p[i-y])
	//}

	for _, v := range p {
		self.score += v.score
		self.len += v.len
	}
	self.projLen = self.projLen + len(p)
	//self.projLen = 0
}
func (self *Projects) removeLastProject() {

	last := self.projs[self.projLen-1]
	self.score -= last.score
	self.len -= last.len
	self.projLen--
	//self.projs = self.projs[:len(self.projs)-1]
}

//func (self *Projects) copyold(p *Projects) {
//	self.len = p.len
//	self.score = p.score
//
//	if len(self.projs) <= p.projLen {
//		v := p.projLen - len(self.projs)
//		for i := 0; i < v; i++ {
//			self.projs = append(self.projs, Project{})
//		}
//	}
//	self.projLen = p.projLen
//	for i := 0; i < self.projLen; i++ {
//		self.projs[i] = p.projs[i]
//	}
//
//	//
//	//if len(self.projs) < len(p.projs) {
//	//	x := len(self.projs)
//	//	for i := 0; i < (len(p.projs) - x); i++ {
//	//		self.projs = append(self.projs, nil)
//	//	}
//	//} else {
//	//	self.projs = self.projs[:len(p.projs)]
//	//}
//	//if len(p.projs) != len(self.projs) {
//	//	fmt.Println(self.projs, p.projs)
//	//	panic("err")
//	//}
//	//
//	////self.projs = make([]*Project, len(p.projs))
//	//for i := 0; i < len(p.projs); i++ {
//	//	if self.projs[i] != nil {
//	//		x := self.projs[i].tagIndex
//	//		*self.projs[i] = *p.projs[i]
//	//		self.projs[i].tagIndex = x
//	//
//	//		lenx := len(x)
//	//
//	//		if lenx == len(p.projs[i].tagIndex) {
//	//
//	//		} else if lenx > len(p.projs[i].tagIndex) {
//	//			self.projs[i].tagIndex = self.projs[i].tagIndex[:len(p.projs[i].tagIndex)]
//	//		} else {
//	//			for j := 0; j < (len(p.projs[i].tagIndex) - lenx); j++ {
//	//				self.projs[i].tagIndex = append(self.projs[i].tagIndex, 0)
//	//			}
//	//		}
//	//		for j := 0; j < len(self.projs[i].tagIndex); j++ {
//	//			self.projs[i].tagIndex[j] = p.projs[i].tagIndex[j]
//	//		}
//	//
//	//	} else {
//	//		self.projs[i] = p.projs[i].clone()
//	//	}
//	//
//	//}
//}
//
//func (self *Projects) copy(p *Projects) {
//	self.len = p.len
//	self.score = p.score
//
//	if len(self.projs) <= p.projLen {
//		v := p.projLen - len(self.projs)
//		for i := 0; i < v; i++ {
//			self.projs = append(self.projs, Project{})
//		}
//	}
//	self.projLen = p.projLen
//	for i := 0; i < self.projLen; i++ {
//		self.projs[i] = p.projs[i]
//	}
//
//	//
//	//if len(self.projs) < len(p.projs) {
//	//	x := len(self.projs)
//	//	for i := 0; i < (len(p.projs) - x); i++ {
//	//		self.projs = append(self.projs, nil)
//	//	}
//	//} else {
//	//	self.projs = self.projs[:len(p.projs)]
//	//}
//	//if len(p.projs) != len(self.projs) {
//	//	fmt.Println(self.projs, p.projs)
//	//	panic("err")
//	//}
//	//
//	////self.projs = make([]*Project, len(p.projs))
//	//for i := 0; i < len(p.projs); i++ {
//	//	if self.projs[i] != nil {
//	//		x := self.projs[i].tagIndex
//	//		*self.projs[i] = *p.projs[i]
//	//		self.projs[i].tagIndex = x
//	//
//	//		lenx := len(x)
//	//
//	//		if lenx == len(p.projs[i].tagIndex) {
//	//
//	//		} else if lenx > len(p.projs[i].tagIndex) {
//	//			self.projs[i].tagIndex = self.projs[i].tagIndex[:len(p.projs[i].tagIndex)]
//	//		} else {
//	//			for j := 0; j < (len(p.projs[i].tagIndex) - lenx); j++ {
//	//				self.projs[i].tagIndex = append(self.projs[i].tagIndex, 0)
//	//			}
//	//		}
//	//		for j := 0; j < len(self.projs[i].tagIndex); j++ {
//	//			self.projs[i].tagIndex[j] = p.projs[i].tagIndex[j]
//	//		}
//	//
//	//	} else {
//	//		self.projs[i] = p.projs[i].clone()
//	//	}
//	//
//	//}
//}

func (self *Projects) merge(theOther *Projects) {
	if theOther == nil {
		return
	}

	p := theOther.projs

	if self.projLen+theOther.projLen >= 10 {
		fmt.Println(*self, *theOther)
		panic("")
	}

	for i := self.projLen; i < min(len(self.projs), self.projLen+theOther.projLen); i++ {
		self.projs[i] = p[i-self.projLen]
	}

	//y := min(len(self.projs), self.projLen+len(p))
	//for i := y; i < self.projLen+len(p); i++ {
	//	self.projs = append(self.projs, p[i-y])
	//}

	self.score += theOther.score
	self.len += theOther.len
	self.projLen += theOther.projLen
}

func (self *Projects) clear() {
	//self.projs = nil
	self.score = 0
	self.len = 0
	self.projLen = 0
}

func (self *Projects) updateUse(p *Projects) {
	if p == nil {
		return
	}

	self.clear()
	self.projs = p.projs
	self.score = p.score
	self.len = p.len

	p.clear()
}

func NewBombProjectInit(res *Project, val int16, colors ...int16) *Project {
	res.ptype = ProjectBomb
	if val <= 0 {
		panic("val <= 0")
	}
	res.min = val
	res.max = val
	if len(colors) <= 0 {
		panic("colors len <= 0")
	}
	res.len = int16(len(colors))
	res.score = res.len * scoreArrInt16[val]

	for _, v := range colors {
		res.pcolor[v] = true
	}
	return res
}

func NewBombProject(val int16, colors ...int16) Project {
	res := Project{
		min:    0,
		max:    0,
		len:    0,
		ptype:  ProjectBomb,
		score:  0,
		pcolor: [5]bool{},
	}
	if val <= 0 {
		panic("val <= 0")
	}
	res.min = val
	res.max = val
	if len(colors) <= 0 {
		panic("colors len <= 0")
	}
	res.len = int16(len(colors))
	res.score = res.len * scoreArrInt16[val]

	for _, v := range colors {
		res.pcolor[v] = true
	}
	return res
}

func NewBombProjectAllInit(res *Project, val int16) *Project {
	if val <= 0 {
		panic("val <= 0")
	}

	res.min = val
	res.max = val
	res.len = 4
	res.ptype = ProjectBomb
	res.score = 4 * scoreArrInt16[val]
	for i := 1; i <= 4; i++ {
		res.pcolor[i] = true
	}

	return res
}

func NewBombProjectAll(val int16) Project {
	//panic("")
	if val <= 0 {
		panic("val <= 0")
	}

	res := Project{
		min:    val,
		max:    val,
		len:    4,
		ptype:  ProjectBomb,
		score:  4 * scoreArrInt16[val],
		pcolor: [5]bool{false, true, true, true, true},
	}

	return res
}

func NewBombProjectExceptOneColor(val int16, color int16) Project {
	llog(-10, "val %d, color %d\n", val, color)

	if val <= 0 {
		panic("val <= 0")
	}

	res := Project{
		min:    val,
		max:    val,
		len:    3,
		ptype:  ProjectBomb,
		score:  3 * scoreArrInt16[val],
		pcolor: [5]bool{false, true, true, true, true},
	}

	if color > 0 && color < 5 {
		res.pcolor[color] = false
	}
	return res
}

func NewBombProjectMustContain(val, color, cnt int16) *Project {
	llog(-10, "val %d, color %d, cnt %d\n", val, color, cnt)

	if val <= 0 {
		panic("val <= 0")
	}

	res := &Project{
		min:    val,
		max:    val,
		len:    cnt,
		ptype:  ProjectBomb,
		score:  cnt * scoreArrInt16[val],
		pcolor: [5]bool{},
	}

	if color > 0 && color < 5 {
		res.pcolor[color] = true
	}

	cnt--
	for i := 1; i < 5; i++ {
		if !res.pcolor[i] {
			res.pcolor[i] = true
			cnt--
		}
		if cnt <= 0 {
			break
		}
	}
	return res
}

func NewFlushProjectInit(min, max, color int16, res *Project) *Project {
	len := (max - min + 1)

	if len < 3 {
		panic("not flush")
	}
	res.min = min
	res.max = max
	res.len = len
	res.ptype = ProjectFlush
	res.score = 0

	if color > 0 && color < 5 {
		res.pcolor[color] = true
	}

	for i := min; i <= max; i++ {

		res.score += scoreArrInt16[i%18]
	}
	return res
}

func NewFlushProject(min, max, color int16) Project {
	//fmt.Println("new one")
	len := (max - min + 1)

	if len < 3 {
		panic("not flush")
	}
	res := Project{
		min:    min,
		max:    max,
		len:    len,
		ptype:  ProjectFlush,
		score:  0,
		pcolor: [5]bool{},
	}

	if color > 0 && color < 5 {
		res.pcolor[color] = true
	}

	for i := min; i <= max; i++ {

		res.score += scoreArrInt16[i%18]
	}
	return res
}

func (self *Projects) updateScoreLen() {
	self.len = 0
	self.score = 0
	for i := 0; i < self.projLen; i++ {
		self.len += self.projs[i].len
		self.score += self.projs[i].score
	}
}

func (self *Project) addJokerIndex(index int16) {
	self.tagIndex[self.tagLen] = index
	self.tagLen++
}
func (self *Project) removeLastJokerIndex() {
	if self.tagLen == 0 {
		panic("")
	}
	self.tagLen--
	//self.tagIndex = self.tagIndex[:len(self.tagIndex)-1]
}

func (self *Project) replaceLastJokerIndex(index int16) {
	self.tagIndex[self.tagLen-1] = index
}

func (self *Project) FlushAddTailAuto() {
	//if self.max+1 != val {
	//	panic("flushaddtail err")
	//}
	//
	//if !self.pcolor[color] {
	//	panic("not color")
	//}

	self.max++
	self.len++
	self.score += scoreArrInt16[self.max%18]
}

func (self *Project) FlushAddTail(val, color int16) {
	if self.max+1 != val {
		fmt.Println("panic reason", self.max, val, *self)
		panic("flushaddtail err")
	}

	if !self.pcolor[color] {
		panic("not color")
	}

	self.max++
	self.len++
	self.score += scoreArrInt16[val%18]
}

func (self *Project) FlushRemoveTail() {

	if self.len < 3 {
		panic("len can't < 3")
	}

	self.score -= scoreArrInt16[self.max%18]
	self.max--
	self.len--

}

func (self *Project) FlushAddHead() {
	self.score += scoreArrInt16[(self.min-1)%18]
	self.min--
	self.len++
}

func (self *Project) FlushRemoveHead() {
	if self.len <= 3 {
		panic("len can't <= 3")
	}

	self.score -= scoreArrInt16[(self.min)%18]
	self.min++
	self.len--

}

func (self *Project) FlushGetColor() int16 {
	for i, v := range self.pcolor {
		if v {
			return int16(i)
		}
	}
	panic("no color")
}

//type ProjectWithOneJoker struct {
//	proj           []Card
//	score          int
//	jokerRepresent Card
//}
//type ProjectWithTwoJoker struct {
//	proj                            []Card
//	score                           int
//	jokerRepresent, joker2Represent Card
//}

//var maxOneProjectWithOneJoker ProjectWithOneJoker
//var maxTwoProjectWithOneJoker = struct {
//	projs []ProjectWithOneJoker
//}{}
//var maxOneProjectWithTwoJoker ProjectWithTwoJoker

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

//
//func (self *MaxScoreProjects) appendM(p *MaxScoreProjects) {
//	self.proj = append(self.proj, p.proj...)
//	self.score += p.score
//}
//
//func (self *MaxScoreProjects) append(p ...*Project) {
//	self.proj = append(self.proj, p...)
//	for _, v := range p {
//		self.score += v.score
//	}
//}
//func (self *MaxScoreProjects) clear() {
//	self.proj = self.proj[:0]
//	self.score = 0
//}
//func (self *MaxScoreProjects) updateUse(p *MaxScoreProjects) {
//	self.proj = p.proj
//	self.score = p.score
//
//	p.proj = nil
//	p.score = 0
//}
//
//func (self *MaxLenProjects) appendM(p *MaxLenProjects) {
//	self.proj = append(self.proj, p.proj...)
//	self.len += p.len
//}
//
//func (self *MaxLenProjects) append(p ...*Project) {
//	self.proj = append(self.proj, p...)
//	for _, v := range p {
//		self.len += v.len
//	}
//}
//func (self *MaxLenProjects) clear() {
//	self.proj = self.proj[:0]
//	self.len = 0
//}
//func (self *MaxLenProjects) updateUse(p *MaxLenProjects) {
//	self.proj = p.proj
//	self.len = p.len
//
//	p.proj = nil
//	p.len = 0
//}

func (self *AlgX) TopSearch() {

	self.dfs(0)

	//end := time.Now()
	//fmt.Printf("Result: maxScoreProjects(%+v)\n", maxScoreProjects)
	//costTime := end.Sub(start).Microseconds()
	////fmt.Println("result is:", costTime)

}

//func (self *AlgX) AddToMaxJokerProj(cards []Card, jokerPresent Card) {
//	jokerProjScore := score(cards)
//
//	maxOneProjectWithOneJoker.proj = copyCards(cards)
//	maxOneProjectWithOneJoker.score = jokerProjScore
//	maxOneProjectWithOneJoker.jokerRepresent = Card{jokerPresent.first, jokerPresent.second}
//
//}

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

//func (self *AlgX) useTwoColorsOneJokerAsProject(ca, cb, val int16, index int) {
//
//	thisBombColorMap := self.bombRepeatMap[index]
//	thisBombCards := self.bombCards[index]
//
//	//todo 这块没有专业测试
//	//[case]2023/03/26 05:13:13 joker_test.go:257: 	per us:  52  最大分: 62   {cards[[joker {2,3} {3,3}] [{12,3} joker {14,3}] [{5,1} {5,2} {5,4}]] score:62 len:9} 最大长度 {cards[[{2,3} {3,3} joker] [{12,3} joker {14,3}] [{5,1} {5,2} {5,4}]] score:55 len:9} dfsTimes: 0
//	//	[case]2023/03/26 05:13:13 joker_test.go:258: 	per us: 122  最大分: 64   [[Joker ♠A ♣A] [Joker ♣2 ♣3] [♠5 ♥5 ♦5]] 524 1220
//	//	[case]2023/03/26 05:13:13 joker_test.go:259: [{1,1} {5,1} {7,1} {5,2} {13,2} {1,3} {2,3} {3,3} {12,3} {5,4} {11,4}]
//	//var jokerReplaceColors int16
//	//for k := 0; k < len(colors); k++ {
//	//	if colors[k] != colors[i] && colors[k] != colors[j] {
//	//		if newColors[colors[k]] > 0 || len(colors) == 2 { //这个颜色对组成顺子有意义所以发生代替, len(colors)==2代表对于组成炸弹有益处
//	//			//fmt.Println("XXXXXXXXXXXXXXXXXX")
//	//			jokerReplaceColors = colors[k]
//	//			break
//	//		}
//	//	}
//	//
//	//}
//
//	//尝试将这个joker组成的project放到最大的jokerproj中
//	cc := int16(0)
//	for color := int16(1); color < 5; color++ {
//		if color != ca && color != cb {
//			cc = color
//			break
//		}
//	}
//
//	//cc := int16(0)
//	//if jokerReplaceColors != 0 {
//	//	//fmt.Println("YYYYYYYYYYYY")
//	//	cc = jokerReplaceColors
//	//} else {
//	//	//跑到这里证明这次的joker替换之后对顺子并没有益处，那就要看看对炸弹是否有好处
//	//
//	//	continue
//	//}
//	//ca, cb := colors[i], colors[j]
//
//	carda, cardb := Card{val, ca}, Card{val, cb}
//
//	//fmt.Println("states cards", carda, cardb, states[CARD_OFFSET(&carda)][0], states[CARD_OFFSET(&cardb)][0])
//	statesaBak := self.states[CARD_OFFSET(&carda)]
//	statesbBak := self.states[CARD_OFFSET(&cardb)]
//
//	if self.states[CARD_OFFSET(&carda)][0] < 0 {
//		panic("states[CARD_OFFSET(&ca)][0] err")
//	}
//
//	if self.states[CARD_OFFSET(&cardb)][0] < 0 {
//		panic("states[CARD_OFFSET(&cb)][0] err")
//	}
//
//	jokerProj := NewBombProject(val, ca, cb, cc)
//	jokerProj.tagIndex = append(jokerProj.tagIndex, cc)
//	//fmt.Println("carda, cardb", carda, cardb, jokerProj)
//
//	//清楚joker占用的牌，接着跳转到普通的dfs中进行剩下的递归
//	colorMapRemoveCard(thisBombColorMap, carda)
//	colorMapRemoveCard(thisBombColorMap, cardb)
//	thisBombCardsNew := colorMapToCards(thisBombColorMap, val)
//	self.bombCards[index] = thisBombCardsNew
//	self.bombRepeatMap[index] = thisBombColorMap
//	if self.jokerLeft <= 0 {
//		panic("self.jokerLeft must bigger than 0")
//	}
//	self.jokerLeft--
//
//	if self.states[CARD_OFFSET(&carda)][1] > 0 {
//		self.states[CARD_OFFSET(&carda)][1]--
//	} else if self.states[CARD_OFFSET(&carda)][0] > 0 {
//		self.states[CARD_OFFSET(&carda)][0]--
//	} else {
//		panic("")
//	}
//
//	if self.states[CARD_OFFSET(&cardb)][1] > 0 {
//		self.states[CARD_OFFSET(&cardb)][1]--
//	} else if self.states[CARD_OFFSET(&cardb)][0] > 0 {
//		self.states[CARD_OFFSET(&cardb)][0]--
//	} else {
//		panic("")
//	}
//
//	self.bombTakeAheadMaxScoreProjects.appendProject(jokerProj)
//
//	//llog(-100, "\nOneJokerMiddleNode index %d jokerProj %+v %v thisBombCards(%v) thisBombColorMap(%v)\n", index, self.bombTakeAheadMaxScoreProjects, []int16{self.states[CARD_OFFSET(&carda)][0], self.states[CARD_OFFSET(&carda)][1], self.states[CARD_OFFSET(&cardb)][0], self.states[CARD_OFFSET(&cardb)][1]}, thisBombCardsNew, thisBombColorMap)
//
//	self.dfs(index)
//
//	self.bombCards[index] = thisBombCards
//	self.bombRepeatMap[index].colors[carda.Second]++
//	self.bombRepeatMap[index].cnt++
//	if self.bombRepeatMap[index].colors[carda.Second] == 1 {
//		self.bombRepeatMap[index].colorCnt++
//	}
//
//	self.bombRepeatMap[index].colors[cardb.Second]++
//	self.bombRepeatMap[index].cnt++
//	if self.bombRepeatMap[index].colors[cardb.Second] == 1 {
//		self.bombRepeatMap[index].colorCnt++
//	}
//
//	thisBombCardsNew = nil
//
//	//self.bombTakeAheadMaxScoreProjects.projs = self.bombTakeAheadMaxScoreProjects.projs[:len(self.bombTakeAheadMaxScoreProjects.projs)-1]
//	jokerProj = nil
//	self.bombTakeAheadMaxScoreProjects.projLen--
//	self.bombTakeAheadMaxScoreProjects.updateScoreLen()
//	self.jokerLeft++
//
//	self.states[CARD_OFFSET(&carda)] = statesaBak
//	self.states[CARD_OFFSET(&cardb)] = statesbBak
//
//	//jokerProj.tagIndex = jokerProj.tagIndex[:len(jokerProj.tagIndex)-1]//todo tocheck
//}
//
//func (self *AlgX) useThreeColorsOneJokerAsProject(ca, cb, cc, val int16, index int) {
//
//	thisBombColorMap := self.bombRepeatMap[index]
//	thisBombCards := self.bombCards[index]
//
//	//ca, cb, cc := colors[0], colors[1], colors[2]
//	var cd int16
//
//	for color := int16(1); color < 5; color++ {
//		if color != ca && color != cb && color != cc {
//			cd = color
//			break
//		}
//	}
//
//	carda, cardb, cardc := Card{val, ca}, Card{val, cb}, Card{val, cc}
//
//	//fmt.Println("states cards", carda, cardb, states[CARD_OFFSET(&carda)][0], states[CARD_OFFSET(&cardb)][0])
//	statesaBak := self.states[CARD_OFFSET(&carda)]
//	statesbBak := self.states[CARD_OFFSET(&cardb)]
//	statescBak := self.states[CARD_OFFSET(&cardc)]
//
//	if self.states[CARD_OFFSET(&carda)][0] < 0 {
//		panic("states[CARD_OFFSET(&ca)][0] err")
//	}
//
//	if self.states[CARD_OFFSET(&cardb)][0] < 0 {
//		panic("states[CARD_OFFSET(&cb)][0] err")
//	}
//	if self.states[CARD_OFFSET(&cardc)][0] < 0 {
//		panic("states[CARD_OFFSET(&cb)][0] err")
//	}
//
//	//尝试将这个joker组成的project放到最大的jokerproj中
//
//	jokerProj := NewBombProjectAll(val)
//	jokerProj.tagIndex = append(jokerProj.tagIndex, cd)
//	//fmt.Println("carda, cardb", carda, cardb, cardc)
//
//	//清楚joker占用的牌，接着跳转到普通的dfs中进行剩下的递归
//	colorMapRemoveCard(thisBombColorMap, carda)
//	colorMapRemoveCard(thisBombColorMap, cardb)
//	colorMapRemoveCard(thisBombColorMap, cardc)
//	thisBombCardsNew := colorMapToCards(thisBombColorMap, val)
//	self.bombCards[index] = thisBombCardsNew
//	self.bombRepeatMap[index] = thisBombColorMap
//
//	if self.jokerLeft <= 0 {
//		panic("self.jokerLeft must bigger than 0")
//	}
//	self.jokerLeft--
//
//	if self.states[CARD_OFFSET(&carda)][1] > 0 {
//		self.states[CARD_OFFSET(&carda)][1]--
//	} else if self.states[CARD_OFFSET(&carda)][0] > 0 {
//		self.states[CARD_OFFSET(&carda)][0]--
//	} else {
//		panic("")
//	}
//
//	if self.states[CARD_OFFSET(&cardb)][1] > 0 {
//		self.states[CARD_OFFSET(&cardb)][1]--
//	} else if self.states[CARD_OFFSET(&cardb)][0] > 0 {
//		self.states[CARD_OFFSET(&cardb)][0]--
//	} else {
//		panic("")
//	}
//
//	if self.states[CARD_OFFSET(&cardc)][1] > 0 {
//		self.states[CARD_OFFSET(&cardc)][1]--
//	} else if self.states[CARD_OFFSET(&cardc)][0] > 0 {
//		self.states[CARD_OFFSET(&cardc)][0]--
//	} else {
//		panic("")
//	}
//
//	self.bombTakeAheadMaxScoreProjects.appendProject(jokerProj)
//
//	//llog(-1, "\nOneJokerMiddleNode index %d jokerProj %v %v thisBombCards(%v) thisBombColorMap(%v) cardc %v\n", index, jokerProj, []int16{self.states[CARD_OFFSET(&carda)][0], self.states[CARD_OFFSET(&carda)][1], self.states[CARD_OFFSET(&cardb)][0], self.states[CARD_OFFSET(&cardb)][1]}, thisBombCardsNew, thisBombColorMap, cardc)
//
//	self.dfs(index)
//
//	self.bombCards[index] = thisBombCards
//	self.bombRepeatMap[index].colors[carda.Second]++
//	self.bombRepeatMap[index].cnt++
//	if self.bombRepeatMap[index].colors[carda.Second] == 1 {
//		self.bombRepeatMap[index].colorCnt++
//	}
//
//	self.bombRepeatMap[index].colors[cardb.Second]++
//	self.bombRepeatMap[index].cnt++
//	if self.bombRepeatMap[index].colors[cardb.Second] == 1 {
//		self.bombRepeatMap[index].colorCnt++
//	}
//
//	self.bombRepeatMap[index].colors[cardc.Second]++
//	self.bombRepeatMap[index].cnt++
//	if self.bombRepeatMap[index].colors[cardc.Second] == 1 {
//		self.bombRepeatMap[index].colorCnt++
//	}
//
//	//self.bombTakeAheadMaxScoreProjects.projs = self.bombTakeAheadMaxScoreProjects.projs[:len(self.bombTakeAheadMaxScoreProjects.projs)-1]
//	self.bombTakeAheadMaxScoreProjects.projLen--
//	jokerProj = nil
//	thisBombCardsNew = nil
//	self.bombTakeAheadMaxScoreProjects.updateScoreLen()
//	self.jokerLeft++
//
//	self.states[CARD_OFFSET(&carda)] = statesaBak
//	self.states[CARD_OFFSET(&cardb)] = statesbBak
//	self.states[CARD_OFFSET(&cardc)] = statescBak
//}

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

func (self *AlgX) OneJokerMiddleNode(index int) {
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

//func flushJokerZero(freq []int, cards []Card, colorVal int16, jokerNum int, finalCC *Projects, maxScore1 *int16) {
//
//	cc := getFlushProjects(freq, cards, colorVal) //先更新没有joker的
//
//	var score int16
//
//	if cc != nil {
//		score = cc.score
//	}
//
//	if score > *maxScore1 {
//		*maxScore1 = score
//		finalCC.updateUse(cc)
//	}
//
//}

//func flushJokerOne(freq []int, cards []Card, colorVal int16, jokerNum int, finalCC *Projects, maxScore1 *int16) {
//
//	jokerCanBe, _ := getJokerCanBeCardsOneJoker(freq, cards, colorVal, 1)
//	//llog(-100, "1 joker cardsFreq(%v) jokerCanBe(%v)\n", freq, jokerCanBe)
//	for _, card := range jokerCanBe {
//
//		freq[card.First]++
//		cards := freqArrToCards(freq, card.Second)
//
//		cardCnt := freq[card.First]
//
//		cc := getFlushProjectsOneJokers(freq, cards, colorVal, []Pair{{card.First, int16(cardCnt)}})
//
//		score := cc.score
//		if score > *maxScore1 {
//			*maxScore1 = score
//			finalCC.updateUse(cc)
//		}
//
//		freq[card.First]--
//		if freq[card.First] < 0 {
//			panic("one joker can be")
//		}
//	}
//}

//func (self *AlgX) flushJokertwo(freq []int, cards []Card, colorVal int16, jokerNum int, finalCC *Projects, maxScore1 *int16, finalCC2 *Projects, maxScore2 *int16) {
//
//	jokerCanBe, twoJokerBe := getJokerCanBeCardsOneJoker(freq, cards, colorVal, 2)
//	//llog(-100, "1 2 joker oneJokerBe(%v)\n", jokerCanBe)
//	for _, card := range jokerCanBe {
//
//		freq[card.First]++
//		cards := freqArrToCards(freq, card.Second)
//
//		cardCnt := freq[card.First]
//		for i := 1; i <= cardCnt; i++ {
//			cc := getFlushProjectsOneJokers(freq, cards, colorVal, []Pair{{card.First, int16(i)}})
//
//			if cc == nil {
//				llog(-100, "cards  %v card %v allCards %v\n", cards, card, self.hand_cards)
//				panic("")
//			}
//
//			score := cc.score
//			if score > *maxScore1 {
//				*maxScore1 = score
//				finalCC.updateUse(cc)
//			}
//		}
//
//		freq[card.First]--
//
//	}
//	//两个joker的情况
//	//1. 所有的找到的两个可以放joker的位置
//	var jokerIndexPair []Pair
//
//	for i := 0; i < len(jokerCanBe); i++ {
//		for j := i + 1; j < len(jokerCanBe); j++ {
//			card1, card2 := jokerCanBe[i], jokerCanBe[j]
//
//			freq[card1.First]++
//			freq[card2.First]++
//
//			card1Cnt := freq[card1.First]
//			card2Cnt := freq[card2.First]
//
//			cards := freqArrToCards(freq, card1.Second)
//
//			if card1Cnt == card2Cnt && card2Cnt == 1 {
//				jokerIndexPair = []Pair{{card1.First, 1}, {card2.First, 1}}
//			} else if card1Cnt == card2Cnt && card2Cnt == 2 {
//
//				gap := card1.First - card2.First
//				if gap*gap == 1 { //连着的两个joker分布在两个组合中是最佳的
//					jokerIndexPair = []Pair{{card1.First, 1}, {card2.First, 2}}
//				} else {
//					jokerIndexPair = []Pair{{card1.First, 1}, {card2.First, 1}}
//				}
//			} else {
//				jokerIndexPair = []Pair{{card1.First, int16(card1Cnt)}, {card2.First, int16(card2Cnt)}}
//			}
//
//			cc := getFlushProjectsTwoJokers(freq, cards, colorVal, jokerIndexPair)
//			var score int16
//			if cc != nil {
//				score = cc.score
//			} //cc.score
//			if score > *maxScore2 {
//				*maxScore2 = score
//				finalCC2.updateUse(cc)
//			}
//
//			//llog(-100, "2 11joker card1 %v card2 %v cc(%v) score %d maxScore2 %d finalCC2 %v\n", card1, card2, cc, score, *maxScore2, finalCC2)
//
//			freq[card1.First]--
//			freq[card2.First]--
//			if freq[card1.First] < 0 {
//				panic("freq2")
//			}
//			if freq[card2.First] < 0 {
//				panic("freq2")
//			}
//		}
//	}
//
//	for _, cardss := range twoJokerBe { //todo 不允许两个joker在一个组合中
//		card1, card2 := cardss[0], cardss[1]
//
//		//llog(-100, "2 2joker freq %v %v\n", freq, cardss)
//		freq[card1.First]++
//		freq[card2.First]++
//
//		card1Cnt := freq[card1.First]
//		card2Cnt := freq[card2.First]
//
//		cards := freqArrToCards(freq, card1.Second)
//
//		if card1Cnt == card2Cnt && card2Cnt == 1 {
//			jokerIndexPair = []Pair{{card1.First, 1}, {card2.First, 1}}
//		} else if card1Cnt == card2Cnt && card2Cnt == 2 {
//
//			gap := card1.First - card2.First
//			if gap*gap == 1 {
//				jokerIndexPair = []Pair{{card1.First, 1}, {card2.First, 2}}
//			} else {
//				jokerIndexPair = []Pair{{card1.First, 1}, {card2.First, 1}}
//			}
//
//		} else {
//			jokerIndexPair = []Pair{{card1.First, int16(card1Cnt)}, {card2.First, int16(card2Cnt)}}
//		}
//		//llog(-100, "2 b11joker card1 %v card2 %v cc(%v) cards(%v)\n", card1, card2, cc, cards)
//		cc := getFlushProjectsTwoJokers(freq, cards, colorVal, jokerIndexPair)
//		//llog(-100, "2 11joker card1 %v card2 %v cc(%v) cards(%v)\n", card1, card2, cc, cards)
//		var score int16
//		if cc != nil {
//			score = cc.score
//		} //cc.score
//		if score > *maxScore2 {
//			*maxScore2 = score
//			finalCC2.updateUse(cc)
//		}
//
//		freq[card1.First]--
//		freq[card2.First]--
//		if freq[card1.First] < 0 {
//			panic("freq2")
//		}
//		if freq[card2.First] < 0 {
//			panic("freq2")
//		}
//	}
//
//}

func recoverFreq(oriv, oriv1, oriv2 int, freq []int, v int) {
	freq[v] = oriv
	freq[v+1] = oriv1
	freq[v+2] = oriv2
}

func (self *AlgX) updateMaxLenProjects(tmpMaxLenProjects *Projects) {
	l := tmpMaxLenProjects.len
	if l > self.maxLenProjects.len || (l == self.maxLenProjects.len && tmpMaxLenProjects.score > self.maxLenProjects.score) {
		self.maxLenProjects.updateUse(tmpMaxLenProjects)
	}
}

func (self *AlgX) updateMaxScoreProjects(tmpMaxScoreProjects *Projects) {

	//llog(-100, "self.tmpMaxScoreProjects %v \ntmpMaxScoreProjects %v\n\n", self.maxScoreProjects, tmpMaxScoreProjects)

	score := tmpMaxScoreProjects.score // scoreProjects(tmpMaxScoreProjects.proj)
	if score > self.maxScoreProjects.score {
		self.maxScoreProjects.updateUse(tmpMaxScoreProjects)
	}

	//llog(-100, "after self.tmpMaxScoreProjects %v tmpMaxScoreProjects %v\n\n", self.maxScoreProjects, tmpMaxScoreProjects)
}

func getOneUnSameColor(cards []Card) int16 {
	for i := int16(1); i < 5; i++ {
		isInColors := false
		for j := 0; j < len(cards); j++ {
			if i == cards[j].Second {
				isInColors = true
			}
		}
		if !isInColors {
			return i
		}
	}
	panic("getOneUnSameColor err")
	return 0
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

//func splitSamePointCardsToBombsWithOneJoker(pcc *[][]Card, colorFreq map[int16]int16, cardVal int16) {
//	var c1, c2 []Card
//	var cp *[]Card = &c1
//	for color, n := range colorFreq {
//		if n >= 2 {
//			c1 = append(c1, Card{cardVal, color})
//			c2 = append(c2, Card{cardVal, color})
//		} else if n == 1 {
//			if len(c1) > len(c2) {
//				cp = &c2
//			}
//			*cp = append(*cp, Card{cardVal, color})
//		} else {
//			panic("n == 0")
//		}
//	}
//
//	if len(c1) > len(c2) {
//		cp = &c2
//	} else {
//		cp = &c1
//	}
//
//	*cp = append(*cp, Card{First: cardVal, Second: getOneUnSameColor(*cp)})
//
//	if len(c1) > 2 {
//		*pcc = append(*pcc, c1)
//	}
//
//	if len(c2) > 2 {
//		*pcc = append(*pcc, c2)
//	}
//
//}

//func (self *AlgX) BombInsertThisColorBetter(cards []Card, inputMap [][2]int16) [][]Card {
//
//	colorFreq := make(map[int16]int16)
//
//	for _, c := range cards {
//		cardNumAsBomb := inputMap[CARD_OFFSET(&c)][1]
//		if cardNumAsBomb > 0 {
//			colorFreq[c.Second] = cardNumAsBomb
//		}
//	}
//
//	sort.Slice(cards, func(i, j int) bool {
//		if cards[i].Second == cards[j].Second {
//			return cards[i].First < cards[j].First
//		} else if cards[i].Second < cards[j].Second {
//			return true
//		}
//		return false
//	})
//
//	cardVal := cards[0].First
//	//fmt.Println("searchBombWorkWithJoker mp", colorFreq, cardVal)
//	var cc [][]Card
//	var c []Card
//
//	switch len(colorFreq) { //那现在牌值相同的牌求得最多的炸弹个数和分数
//	case 2: //可能有2,3,4,5张牌，但是加入joker之后也只能形成一个炸弹
//
//		for color, _ := range colorFreq {
//			c = append(c, Card{cardVal, color})
//		}
//
//		c = append(c, Card{cardVal, getOneUnSameColor(c)})
//		cc = append(cc, c)
//	case 3:
//		switch len(cards) {
//		case 3:
//			fallthrough
//		case 4:
//			for color, _ := range colorFreq {
//				c = append(c, Card{cardVal, color})
//			}
//			cc = append(cc, c)
//		case 5: //5张牌加一个joker可能能形成两个炸弹,至少形成一个4张牌的炸弹
//			for color, _ := range colorFreq {
//				c = append(c, Card{cardVal, color})
//				colorFreq[color]--
//				if colorFreq[color] == 0 {
//					delete(colorFreq, color)
//				}
//			}
//			cc = append(cc, c)
//
//			if len(colorFreq) > 1 {
//				c = make([]Card, 3)
//				copy(c, cc[0])
//				cc = append(cc, c)
//			} else {
//				cc[0] = append(cc[0], Card{cardVal, getOneUnSameColor(cc[0])}) //组成4张牌的炸
//			}
//
//		case 6: //加入一个joker之后可能形成4,3格局，至少是3，3格局
//			fallthrough
//		case 7: //加入一个joker之后一定可以形成4，3格局
//			splitSamePointCardsToBombsWithOneJoker(&cc, colorFreq, cardVal)
//		default:
//			panic("colorFreq 3 len(cards) not right")
//		}
//	case 4:
//		switch len(cards) {
//		case 4:
//			for color, _ := range colorFreq {
//				c = append(c, Card{cardVal, color})
//			}
//			cc = append(cc, c)
//		case 5: //加入一个joker之后可以形成两个3炸
//			fallthrough
//		case 6: //6张牌加入一个joker之后可以形成4,3或者3，3组合
//			fallthrough
//		case 7: //加入一个joker之后至少形成43组合，最多可以形成44组合
//			splitSamePointCardsToBombsWithOneJoker(&cc, colorFreq, cardVal)
//		case 8: //最多可以形成333组合，至少形成44组合
//
//			colorHasTwoCardNum := 0
//			for _, n := range colorFreq {
//				if n == 2 {
//					colorHasTwoCardNum++
//				}
//			}
//			if colorHasTwoCardNum == 4 {
//				c1 := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}}
//				c2 := []Card{{cardVal, 4}, {cardVal, 1}, {cardVal, 2}}
//				c3 := []Card{{cardVal, 3}, {cardVal, 4}, {cardVal, 1}}
//				cc = append(cc, c1, c2, c3)
//			} else {
//				c1 := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
//				c2 := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
//				cc = append(cc, c1, c2)
//			}
//		}
//	}
//
//	return cc
//}

//func (self *AlgX) searchBombWorkWithJoker(cards []Card, inputMap [][2]int16) [][]Card {
//
//	colorFreq := make(map[int16]int16)
//
//	for _, c := range cards {
//		cardNumAsBomb := inputMap[CARD_OFFSET(&c)][1]
//		if cardNumAsBomb > 0 {
//			colorFreq[c.Second] = cardNumAsBomb
//		}
//	}
//
//	sort.Slice(cards, func(i, j int) bool {
//		if cards[i].Second == cards[j].Second {
//			return cards[i].First < cards[j].First
//		} else if cards[i].Second < cards[j].Second {
//			return true
//		}
//		return false
//	})
//
//	cardVal := cards[0].First
//	//fmt.Println("searchBombWorkWithJoker mp", colorFreq, cardVal)
//	var cc [][]Card
//	var c []Card
//
//	switch len(colorFreq) { //那现在牌值相同的牌求得最多的炸弹个数和分数
//	case 2: //可能有2,3,4,5张牌，但是加入joker之后也只能形成一个炸弹
//
//		for color, _ := range colorFreq {
//			c = append(c, Card{cardVal, color})
//		}
//
//		c = append(c, Card{cardVal, getOneUnSameColor(c)})
//		cc = append(cc, c)
//	case 3:
//		switch len(cards) {
//		case 3:
//			fallthrough
//		case 4:
//			for color, _ := range colorFreq {
//				c = append(c, Card{cardVal, color})
//			}
//			cc = append(cc, c)
//		case 5: //5张牌加一个joker可能能形成两个炸弹,至少形成一个4张牌的炸弹
//			for color, _ := range colorFreq {
//				c = append(c, Card{cardVal, color})
//				colorFreq[color]--
//				if colorFreq[color] == 0 {
//					delete(colorFreq, color)
//				}
//			}
//			cc = append(cc, c)
//
//			if len(colorFreq) > 1 {
//				c = make([]Card, 3)
//				copy(c, cc[0])
//				cc = append(cc, c)
//			} else {
//				cc[0] = append(cc[0], Card{cardVal, getOneUnSameColor(cc[0])}) //组成4张牌的炸
//			}
//
//		case 6: //加入一个joker之后可能形成4,3格局，至少是3，3格局
//			fallthrough
//		case 7: //加入一个joker之后一定可以形成4，3格局
//			splitSamePointCardsToBombsWithOneJoker(&cc, colorFreq, cardVal)
//		default:
//			panic("colorFreq 3 len(cards) not right")
//		}
//	case 4:
//		switch len(cards) {
//		case 4:
//			for color, _ := range colorFreq {
//				c = append(c, Card{cardVal, color})
//			}
//			cc = append(cc, c)
//		case 5: //加入一个joker之后可以形成两个3炸
//			fallthrough
//		case 6: //6张牌加入一个joker之后可以形成4,3或者3，3组合
//			fallthrough
//		case 7: //加入一个joker之后至少形成43组合，最多可以形成44组合
//			splitSamePointCardsToBombsWithOneJoker(&cc, colorFreq, cardVal)
//		case 8: //最多可以形成333组合，至少形成44组合
//
//			colorHasTwoCardNum := 0
//			for _, n := range colorFreq {
//				if n == 2 {
//					colorHasTwoCardNum++
//				}
//			}
//			if colorHasTwoCardNum == 4 {
//				c1 := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}}
//				c2 := []Card{{cardVal, 4}, {cardVal, 1}, {cardVal, 2}}
//				c3 := []Card{{cardVal, 3}, {cardVal, 4}, {cardVal, 1}}
//				cc = append(cc, c1, c2, c3)
//			} else {
//				c1 := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
//				c2 := []Card{{cardVal, 1}, {cardVal, 2}, {cardVal, 3}, {cardVal, 4}}
//				cc = append(cc, c1, c2)
//			}
//		}
//	}
//
//	return cc
//}
