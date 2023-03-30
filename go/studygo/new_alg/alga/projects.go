package alga

import "fmt"

//对于有joker牌的处理，这里仍然不区分joker和joker2
type Bomb struct {
	_        int16
	val      int16    //牌值
	cnt      int16    //牌张数量
	colors   [5]int16 //颜色
	colorCnt int16    //这个炸弹中一共有多少种颜色
}

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

type Projects struct {
	projs   [10]Project
	score   int16 //所有project的分数和
	len     int16 //所有project的长度的和
	projLen int
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

	if min(len(self.projs), self.projLen+len(p)) >= 10 {
		panic("")
	}

	for i := self.projLen; i < min(len(self.projs), self.projLen+len(p)); i++ {
		self.projs[i] = p[i-self.projLen]
	}

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
}

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
	//llog(-10, "val %d, color %d\n", val, color)

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
