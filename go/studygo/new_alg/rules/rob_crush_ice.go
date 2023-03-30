package rules

import (
	"fmt"
)

func (cr *CrushIceResult) FreeCards(has []Card) int {
	return len(has) - GetProjectCardCount(cr.Projects)
}

type CrushConfig struct {
	MinPair    int
	MinProject int
}

// CRUSHLIMIT为空则仅组合出结果，不限制破冰的值
func CrushIceByProjects(projList []*Project, crushLimit *CrushConfig) (*CrushIceResult, bool) {
	if len(projList) == 0 {
		return nil, false
	}
	total := getProjectScore(projList)
	if crushLimit != nil && total < crushLimit.MinProject {
		return nil, false
	}
	return &CrushIceResult{Type: int(CrushTypeProject), Value: total, Projects: projList}, true
	//}
}

func CrushProjectsNeedScore(projList []*Project, minProject int) int {
	total := getProjectScore(projList)
	if total > minProject {
		return 0
	}
	return minProject - total
}

func GetProjectCardCount(projList []*Project) int {
	total := 0
	for _, v := range projList {
		if v == nil {
			fmt.Printf("wtf:%d", len(projList))
		}
		if v != nil {
			total += len(v.Cards())
		}
	}
	return total
}

func GetAllProjectCards(projectList []*Project) []Card {
	if len(projectList) == 0 {
		return nil
	}
	result := make([]Card, 0)
	for _, v := range projectList {
		if v == nil {
			fmt.Printf("wtf:%d", len(projectList))
		}
		if v != nil && len(v.Cards()) > 0 {
			result = append(result, v.Cards()...)
		}
	}
	return result
}
