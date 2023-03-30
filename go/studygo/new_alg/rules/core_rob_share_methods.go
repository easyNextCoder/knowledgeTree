package rules

//核心程序和机器人可共用的方法

func GetProjectScore(projList []*Project) int {
	return getProjectScore(projList)
}

func getProjectScore(projList []*Project) int {
	total := 0
	for _, v := range projList {
		if v.Type == ProjectBomb || v.Type == ProjectStraightFlush {
			total += v.Score()
		}
	}
	return total
}

func (p *Project) Score() int {
	if p.Type == ProjectStraightFlush {
		return p.straightFlushScore()
	} else {
		return p.bombScore()
	}
}

func (p *Project) straightFlushScore() int {
	total := 0
	for _, v := range p.Preset {
		if v.Num() > 0 && v.Num() < 14 {
			total += handScoreArr[v.Num()]
		} else {
			total += 0
			
		}
	}
	return total
}

//Bomb

// 炸弹组中不可能有三个OKEY牌，找出第一个非OKEY就行了
func (p *Project) bombScore() int {
	for _, v := range p.Raw {
		if v != OkeyCard && v != OkeyCard2 {
			return handScoreArr[v.Num()] * len(p.Raw)
		}
	}
	return 0
}
