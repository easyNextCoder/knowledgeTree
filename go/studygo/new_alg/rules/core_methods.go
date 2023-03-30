package rules

import (
	"studygo/new_alg/consts"
)

var CardsTagMap map[rune]int = map[rune]int{
	'\u2660': ColorBlue,
	'\u2666': ColorYellow,
	'\u2663': ColorBlack,
	'\u2665': ColorRed,
}

var CardsNumMap map[string]int = map[string]int{
	"A":      1,
	"2":      2,
	"3":      3,
	"4":      4,
	"5":      5,
	"6":      6,
	"7":      7,
	"8":      8,
	"9":      9,
	"10":     10,
	"J":      11,
	"Q":      12,
	"K":      13,
	"Joker":  14,
	"Joker2": 15,
}

func CardStringToRuleCard(scard string) Card {

	runes := []rune(scard)
	if len(runes) <= 0 {
		return InvalidCard
	}
	if len(runes) > 3 {

		if _, ok := CardsNumMap[scard]; !ok {
			return InvliadCard
		}

		if CardsNumMap[scard] == 14 {
			return OkeyCard
		} else if CardsNumMap[scard] == 15 {
			return OkeyCard2
		}
	} else {
		color, ok := CardsTagMap[runes[0]]
		if !ok {
			return InvliadCard
		}
		num, ok2 := CardsNumMap[string(runes[1:])]
		if !ok2 {
			return InvliadCard
		}
		return NewCard(color, num)
	}
	return InvalidCard
}

//一局游戏结束的时候算分数函数
func CoreGameOverHandCardsScore(cards []Card) int {
	res := 0
	for _, v := range cards {
		res += v.CoreGameOverHandScore()
	}
	return res
}

// 手牌留在手里的分数
func (c Card) CoreGameOverHandScore() int {
	if c == OkeyCard || c == OkeyCard2 {
		return GameOverJokerInHandScore
	}
	if c.Num() > 0 && c.Num() < 14 {
		return gameOverHandCardScoreArr[c.Num()]
	}
	return 0
}

//顺子project的检校
func CoreFlushProjectCheck(proj *Project) bool {
	if proj == nil {
		return false
	}
	if len(proj.Raw) != len(proj.Preset) {
		return false
	}

	if len(proj.Raw) < ProjectMinSize {
		return false
	}

	var startIdx = -1
	for i, c := range proj.Raw {
		if c != OkeyCard && c != OkeyCard2 {
			startIdx = i
			break
		}
	}
	// 检查是否满足条件
	var startPlainCard = proj.Raw[startIdx]
	var color = startPlainCard.Color()
	var startNum = startPlainCard.Num() - startIdx
	var endNum = startNum + len(proj.Raw) - 1

	if startNum < CardNumStart {
		return false
	}

	if endNum > CardNumEnd+1 {
		return false
	}

	//handle ... Q K A condition
	QKA_type := false
	var shouldBeAce Card
	if endNum == CardNumEnd+1 {
		QKA_type = true
		shouldBeAce = proj.Raw[len(proj.Raw)-1]
		if shouldBeAce != NewCard(color, 1) && shouldBeAce != OkeyCard && shouldBeAce != OkeyCard2 {
			return false
		}
	}

	usedPreset := proj.Preset
	if QKA_type {
		if usedPreset[len(usedPreset)-1] != NewCard(color, 1) {
			return false
		}
		usedPreset = usedPreset[:len(usedPreset)-1]
		for i := 0; i < len(usedPreset)-1; i++ {
			if usedPreset[i].Num()+1 != usedPreset[i+1].Num() {
				return false
			}
		}
	} else {
		for i := 0; i < len(usedPreset)-1; i++ {
			if usedPreset[i].Num()+1 != usedPreset[i+1].Num() {
				return false
			}
		}
	}
	return true
}

//用于core程序生成新的project
//检查条件1：长度不可小于project的最小长度
//检查条件2：okey牌个数<2
func CoreNewProject(cards []Card) *Project { //xtk todo okey在中间时候生成的好像不正确
	if len(cards) == 2 {
		return nil
	} else {
		straightProj := coreNewStaightFlush(cards)
		bombProj := coreNewBomb(cards)
		if straightProj == nil {
			return bombProj
		} else {
			if bombProj != nil && straightProj.Score() < bombProj.Score() {
				return bombProj
			}
			return straightProj
		}
	}
}

func coreNewBomb(cards []Card) *Project {
	if len(cards) < ProjectMinSize || len(cards) > ProjectBombMaxCount {
		return nil
	}
	if OkeyCardCnt(cards) > 1 {
		return nil
	}
	var colorMap = map[int]int{}
	var numMap = map[int]int{}
	var okCount = 0
	for _, c := range cards {
		if c == OkeyCard || c == OkeyCard2 {
			okCount += 1
		} else {
			colorMap[c.Color()] += 1
			numMap[c.Num()] += 1
		}
	}
	if len(colorMap) != len(cards)-okCount || len(numMap) != 1 {
		return nil
	}
	return &Project{Type: ProjectBomb, Raw: CopyCards(cards)}
}

//可以处理QKA也可以组成顺子这种特殊情况
//当前会检查是否有两个Okey牌
//会检查最小长度是否符合要求，不会检查最长长度（因为补牌操作会让project长度超过限制长度5）
func coreNewStaightFlush(cards []Card) *Project {
	if len(cards) < ProjectMinSize {
		return nil
	}
	if OkeyCardCnt(cards) > 1 {
		return nil
	}
	// 找出第一个非OKEY牌
	var startIdx = -1
	for i, c := range cards {
		if c != OkeyCard && c != OkeyCard2 {
			startIdx = i
			break
		}
	}
	// 检查是否满足条件
	var startPlainCard = cards[startIdx]
	var color = startPlainCard.Color()
	var startNum = startPlainCard.Num() - startIdx
	if startNum < CardNumStart {
		return nil
	}

	var endNum = startNum + len(cards) - 1
	if endNum > CardNumEnd+1 {
		return nil
	}

	//handle ... Q K A condition
	QKA_type := false
	var shouldBeAce Card
	if endNum == CardNumEnd+1 {
		QKA_type = true
		shouldBeAce = cards[len(cards)-1]
		cards = cards[:len(cards)-1]
		if shouldBeAce != NewCard(color, CardNumA) && shouldBeAce != OkeyCard && shouldBeAce != OkeyCard2 {
			return nil
		}
	}

	for i := 0; i < len(cards); i++ {
		presetCard := NewCard(color, startNum+i)
		rawCard := cards[i]
		if rawCard != OkeyCard && rawCard != OkeyCard2 && rawCard != presetCard {
			return nil
		}

	}
	// 生成Preset
	preset := make([]Card, 0, len(cards))
	for i := 0; i < len(cards); i++ {
		preset = append(preset, NewCard(color, startNum+i))
	}
	proj := new(Project)
	proj.Type = ProjectStraightFlush

	if QKA_type {
		cards = append(cards, shouldBeAce)
		preset = append(preset, NewCard(color, 1))
	}

	proj.Raw = CopyCards(cards)
	proj.Preset = preset
	return proj
}

func coreNormalPreSplitProject(cards []Card) ([][]Card, bool) {

	if len(cards) < ProjectMinSize {
		return nil, false
	}

	res := make([][]Card, 0)
	for len(cards) > consts.Hand_Max_Project_Len {
		res = append(res, cards[:3])
		cards = cards[3:]
	}
	if len(cards) < ProjectMinSize || len(cards) > consts.Hand_Max_Project_Len {
		return nil, false
	}
	res = append(res, cards)
	return res, true
}

//在生成project之前先对cards进行拆分，主要是拆分含有2个okey牌或者长度过长的project
func CorePreSplitProject(cards []Card) ([][]Card, bool) {

	okeyCnt := OkeyCardCnt(cards)

	res := make([][]Card, 0)

	if okeyCnt == 2 {
		if len(cards) <= consts.Hand_Max_Project_Len {
			return nil, false
		}
		firstIdx := getFirstOkeyCardIndex(cards)
		lastIdx := getLastOkeyCardIndex(cards)
		if firstIdx == -1 || lastIdx == -1 || firstIdx == lastIdx {
			return nil, false
		}
		splitEnd := firstIdx + 1
		for i := firstIdx; i < lastIdx; i++ { //左闭右开
			if i-0+1 >= ProjectMinSize && len(cards)-1-i >= ProjectMinSize {
				splitEnd = i + 1
				break
			}
		}

		cards1 := cards[:splitEnd]
		cards2 := cards[splitEnd:]
		ret1, ok1 := coreNormalPreSplitProject(cards1)
		if !ok1 {
			return nil, false
		}

		ret2, ok2 := coreNormalPreSplitProject(cards2)
		if !ok2 {
			return nil, false
		}

		res = append(res, ret1...)
		res = append(res, ret2...)
		return res, true
	}

	if okeyCnt == 0 || okeyCnt == 1 {
		ret, ok := coreNormalPreSplitProject(cards)
		if !ok {
			return nil, false
		}
		res = append(res, ret...)
		return res, true
	}
	return res, false
}

func getFirstOkeyCardIndex(cards []Card) int {
	idx := -1
	for i, v := range cards {
		if v == OkeyCard || v == OkeyCard2 {
			idx = i
			break
		}
	}
	return idx
}

func getLastOkeyCardIndex(cards []Card) int {
	idx := -1
	for i := len(cards) - 1; i >= 0; i-- {
		v := cards[i]
		if v == OkeyCard || v == OkeyCard2 {
			idx = i
			break
		}
	}
	return idx
}

func OkeyCardCnt(cards []Card) int {
	cnt := 0
	for _, card := range cards {
		if card == OkeyCard || card == OkeyCard2 {
			cnt++
		}
	}
	return cnt
}

//核心框架中的破冰程序
func CoreCrushIceByProjects(projList []*Project, crushLimit int) (*CoreCrushIceResult, bool) {
	if len(projList) == 0 {
		return nil, false
	}
	total := getProjectScore(projList)
	if crushLimit > 0 && total <= crushLimit {
		return nil, false
	}
	return &CoreCrushIceResult{Type: CrushTypeProject, Value: total, Projects: projList}, true
}

//核心框架的换牌程序（与机器人用过的解耦）
func (p *Project) CoreReplaceCard(cards []Card) ([]Card, bool) {
	if p.Type == ProjectStraightFlush {
		return p.coreReplaceSrtaightCard(cards)
	} else if p.Type == ProjectBomb {
		return p.coreReplaceBombCard(cards)
	} else {
		return nil, false
	}
}

// 用普通牌换OKEY牌，又重复的牌肯定
func (p *Project) coreReplaceSrtaightCard(cards []Card) ([]Card, bool) {
	if len(cards) == 0 {
		return nil, false
	}
	cardsMap := ToCardMap(cards)
	replacedCards := make([]Card, 0)
	// 有相同牌的肯定是有问题的
	if len(cardsMap) != len(cards) {
		return nil, false
	}
	var replacedMap = map[int]Card{}
	for i := 0; i < len(p.Raw); i++ {
		// 只能替换OKEY牌
		rawCard := p.Raw[i]
		if rawCard != OkeyCard && rawCard != OkeyCard2 {
			continue
		}
		replacedCards = append(replacedCards, p.Raw[i])
		presetCard := p.Preset[i]
		if count := cardsMap[presetCard]; count > 0 {
			replacedMap[i] = presetCard
			delete(cardsMap, presetCard)
		}
	}
	if len(cardsMap) != 0 {
		return nil, false
	}
	for i, presetCard := range replacedMap {
		p.Raw[i] = presetCard
	}
	return replacedCards, true
}

func (p *Project) coreReplaceBombCard(cards []Card) ([]Card, bool) {
	if !p.coreCanReplaceBoomb(cards) {
		return nil, false
	}
	var replaceIdx = 0
	replacedCards := make([]Card, 0)
	for i, c := range p.Raw {
		if c == OkeyCard || c == OkeyCard2 {
			replacedCards = append(replacedCards, c)
			p.Raw[i] = cards[replaceIdx]
			replaceIdx += 1
			if replaceIdx >= len(cards) {
				break
			}
		}
	}
	return replacedCards, true
}

func (p *Project) coreCanReplaceBoomb(cards []Card) bool {
	if len(cards) < 1 || len(cards) > 2 || len(p.Raw) != ProjectBombMaxCount {
		return false
	}
	var mainNum int
	var canAddColor = map[int]bool{
		ColorBlue:   true,
		ColorYellow: true,
		ColorRed:    true,
		ColorBlack:  true,
	}
	var okCount int
	for _, c := range p.Raw {
		if c != OkeyCard && c != OkeyCard2 {
			mainNum = c.Num()
			delete(canAddColor, c.Color())
		} else {
			okCount += 1
		}
	}
	// 能替换的牌不会超过OK牌的数目
	if len(cards) > okCount {
		return false
	}
	// 检查是否有重复的牌
	for _, toReplaceCard := range cards {
		if toReplaceCard.Num() != mainNum {
			return false
		}
		// 已经有则表示存在
		if !canAddColor[toReplaceCard.Color()] {
			return false
		}
	}
	return true
}

//核心框架的补牌程序
//如果补完牌不足以分裂的时候会检查补牌后新生成的proj中是否含有两个okey
//如果足以分裂则不检查先往proj中补牌，补牌之后交给后边的splitProject函数去判断合法性若合法则分割，否则此proj则recover到补牌之前的状态
func (p *Project) CoreAddCard(prefix, suffix []Card) bool {
	if p.Type == ProjectStraightFlush {
		return p.coreAddStraightFlushCard(prefix, suffix)
	} else if p.Type == ProjectBomb {
		return p.coreAddBombCard(prefix, suffix)
	} else {
		return false
	}
}

func (p *Project) coreAddBombCard(prefix, suffix []Card) bool {
	if len(prefix)+len(suffix) != 1 {
		return false
	}
	if len(prefix) > 0 {
		if !p.coreCanAddBomb(prefix) {
			return false
		}
		p.Raw = append(CopyCards(prefix), p.Raw...)
	} else {
		if !p.coreCanAddBomb(suffix) {
			return false
		}
		p.Raw = append(p.Raw, suffix...)
	}
	return true
}

func (p *Project) coreCanAddBomb(toAdd []Card) bool {
	if len(toAdd) != 1 || len(p.Raw) != ProjectBombMaxCount-1 {
		return false
	}
	addCard := toAdd[0]
	if addCard == OkeyCard || addCard == OkeyCard2 {
		return true
	}
	addNum := addCard.Num()
	addColor := addCard.Color()
	for _, v := range p.Raw {
		if v == OkeyCard || v == OkeyCard2 {
			continue
		}
		if v.Num() != addNum {
			return false
		}
		if v.Color() == addColor {
			return false
		}
	}
	return true
}

func (p *Project) coreAddStraightFlushCard(prefix, suffix []Card) bool {

	if !CoreFlushProjectCheck(p) {
		return false
	}
	if len(prefix)+len(suffix) == 0 {
		return true
	}
	if len(prefix) > 0 && !p.coreCanAddPrefix(prefix) {
		return false
	}
	if len(suffix) > 0 && !p.coreCanAddSuffix(suffix) {
		return false
	}
	var presetPrefix = p.corePrefixPreset(len(prefix))
	var presetSuffix = p.coreSuffixPreset(len(suffix))
	p.Preset = coreMergeCardList(presetPrefix, p.Preset, presetSuffix)
	p.Raw = coreMergeCardList(prefix, p.Raw, suffix)
	return true
}

func (p *Project) coreCanAddPrefix(prefix []Card) bool {
	var color = p.Preset[0].Color()
	var endNum = p.Preset[0].Num()
	canAdd := endNum - CardNumStart
	if canAdd < len(prefix) {
		return false
	}
	for i := 1; i <= len(prefix); i++ { //xyk todo 目前一次性加多张可能会失效
		paddingCard := prefix[len(prefix)-i]
		presetCard := NewCard(color, endNum-i)
		if paddingCard != presetCard && paddingCard != OkeyCard && paddingCard != OkeyCard2 {
			return false
		}
	}

	//不可以向已经有Okey的组合中补Okey，目前仍保留一次补多张的能力，如果补的超过了6张交给后边的分裂函数处理
	preNewRaw := make([]Card, 0)
	preNewRaw = append(preNewRaw, p.Raw...)
	preNewRaw = append(preNewRaw, prefix...)
	if OkeyCardCnt(preNewRaw) > 1 {
		return false
	}
	return true
}

func (p *Project) coreCanAddSuffix(prefix []Card) bool {
	var endCard = p.Preset[len(p.Preset)-1]
	var endColor = endCard.Color()
	var startNum = endCard.Num()

	//避免补牌出现...QKA23这种类型的,普通顺子不可能存在在A的后边补牌
	if startNum == CardNumA {
		return false
	}

	canAdd := CardNumEnd + 1 - startNum //hand Q K A type
	if canAdd < len(prefix) {
		return false
	}

	if len(prefix) == canAdd { //存在QKA类型
		for i := 1; i <= len(prefix)-1; i++ {
			presetCard := NewCard(endColor, startNum+i)
			paddingCard := prefix[i-1]
			if paddingCard != presetCard && paddingCard != OkeyCard && paddingCard != OkeyCard2 {
				return false
			}
		}
		lastCard := prefix[len(prefix)-1]
		if lastCard != NewCard(endColor, CardNumA) && lastCard != OkeyCard2 && lastCard != OkeyCard {
			return false
		}
	} else {
		for i := 1; i <= len(prefix); i++ {
			presetCard := NewCard(endColor, startNum+i)
			paddingCard := prefix[i-1]
			if paddingCard != presetCard && paddingCard != OkeyCard && paddingCard != OkeyCard2 {
				return false
			}
		}
	}
	//不可以向已经有Okey的组合中补Okey，目前仍保留一次补多张的能力，如果补的超过了6张交给后边的分裂函数处理
	preNewRaw := make([]Card, 0)
	preNewRaw = append(preNewRaw, p.Raw...)
	preNewRaw = append(preNewRaw, prefix...)
	if OkeyCardCnt(preNewRaw) > 1 {
		return false
	}

	return true
}

func (p *Project) corePrefixPreset(n int) []Card {
	var startColor = p.Preset[0].Color()
	var startNum = p.Preset[0].Num()
	var maxAdd = startNum - CardNumStart
	if maxAdd > n {
		maxAdd = n
	}
	var preset = make([]Card, 0, maxAdd)
	for i := maxAdd; i >= 1; i-- {
		preset = append(preset, NewCard(startColor, startNum-i))
	}
	return preset
}

func (p *Project) coreSuffixPreset(n int) []Card {
	var endColor = p.Preset[len(p.Preset)-1].Color()
	var endNum = p.Preset[len(p.Preset)-1].Num()
	var maxAdd = CardNumEnd + 1 - endNum
	var preset = make([]Card, 0, maxAdd)

	if maxAdd == n {
		for i := 1; i < n; i++ {
			preset = append(preset, NewCard(endColor, endNum+i))
		}
		preset = append(preset, NewCard(endColor, 1)) //append A
	} else if maxAdd > n {
		for i := 1; i <= n; i++ {
			preset = append(preset, NewCard(endColor, endNum+i))
		}
	}

	return preset
}

func coreMergeCardList(cards ...[]Card) []Card {
	total := 0
	for _, v := range cards {
		total += len(v)
	}
	all := make([]Card, 0, total)
	for _, v := range cards {
		all = append(all, v...)
	}
	return all
}
