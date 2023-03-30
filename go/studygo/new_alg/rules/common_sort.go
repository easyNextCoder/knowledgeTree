package rules

type SortByNumColor []Card

func (s SortByNumColor) Len() int { return len(s) }

func (s SortByNumColor) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s SortByNumColor) Less(i, j int) bool {
	numI := s[i].Num()
	numJ := s[j].Num()
	if numI != numJ {
		return numI < numJ
	} else {
		return s[i].Color() < s[j].Color()
	}
}

type SortByColorNum []Card

func (s SortByColorNum) Len() int { return len(s) }

func (s SortByColorNum) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s SortByColorNum) Less(i, j int) bool {
	c1 := s[i].Color()
	c2 := s[j].Color()
	if c1 != c2 {
		return c1 < c2
	}
	return s[i].Num() < s[j].Num()
}

type SortByNumDesc []Card

func (s SortByNumDesc) Len() int { return len(s) }

func (s SortByNumDesc) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s SortByNumDesc) Less(i, j int) bool { return s[i].Num() > s[j].Num() }

type SortByNumAsc []Card

func (s SortByNumAsc) Len() int { return len(s) }

func (s SortByNumAsc) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s SortByNumAsc) Less(i, j int) bool { return s[i].Num() < s[j].Num() }

type pairListAsc [][]Card

func (p pairListAsc) Len() int { return len(p) }

func (p pairListAsc) Less(i, j int) bool {
	var num1, num2 int
	for _, c := range p[i] {
		if c != OkeyCard {
			num1 = c.Num()
			break
		}
	}
	for _, c := range p[j] {
		if c != OkeyCard {
			num2 = c.Num()
			break
		}
	}
	return num1 < num2
}

func (p pairListAsc) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
