package rules

import (
	"fmt"
)

type ProjType int

const (
	ProjectStraightFlush ProjType = 1 // 同花顺
	ProjectBomb          ProjType = 2 // 炸弹💣

)

const ProjectMinSize = 3
const ProjectBombMaxCount = 4

type Project struct {
	// 共用部分
	Type ProjType `json:"type"` // Project类型
	Raw  []Card   `json:"raw"`  // 当前牌的原始牌
	// 顺子牌所需要的额外数据
	Preset []Card `json:"preset"` // 当前牌所代表的牌 没有Okey牌的时候为空
}

func (p *Project) String() string {
	return fmt.Sprintf("%s", p.Raw)
}

func (p *Project) Cards() []Card {
	return p.Raw
}

func (p *Project) Lens() int {
	return len(p.Raw)
}

func (p *Project) OkeyCount() int {
	total := 0
	if p == nil {
		return 0
	}
	for _, v := range p.Raw {
		if v == OkeyCard {
			total += 1
		}
	}
	return total
}

func (p *Project) ContainCard(card Card) bool {
	if p == nil {
		return false
	}
	if !card.Valid() {
		return false
	}
	for _, v := range p.Raw {
		if v == card {
			return true
		}
	}
	return false
}

func (p *Project) GetBehalfCards() []Card {
	if len(p.Preset) > 0 {
		return p.Preset
	}
	return p.Raw
}

func (p *Project) Clone() *Project {
	newP := new(Project)
	newP.Type = p.Type
	newP.Raw = CopyCards(p.Raw)
	newP.Preset = CopyCards(p.Preset)
	return newP
}
