package rules

import "fmt"

type CrushType int

const (
	CrushTypeNone             CrushType = 0
	CrushTypeProject          CrushType = 1 // 普通组合破冰
	CrushTypeProjectHand      CrushType = 2 //Hand类型的破冰（drop a meld）
	CrushTypeProjectJokerHand CrushType = 3 //JokerHand破冰
)

var CrushTypeId2String map[CrushType]string = map[CrushType]string{
	CrushTypeNone:             "none",
	CrushTypeProject:          "normal",
	CrushTypeProjectHand:      "hand",
	CrushTypeProjectJokerHand: "joker hand",
}

func (c CrushType) ToInt() int {
	return int(c)
}

type CoreCrushIceResult struct {
	Type     CrushType
	Value    int
	Projects []*Project
}

func (self CoreCrushIceResult) String() string {
	return fmt.Sprintf("crushType(%d) crushValue(%d) crushProjects(%s)", self.Type, self.Value, self.Projects)
}

type CrushIceResult struct {
	Type     int
	Value    int
	Projects []*Project
}

type WinType int

const (
	WinTypeNormal              WinType = 0
	WinTypeHand                WinType = 1
	WinTypeJokerHand           WinType = 2
	WinTypeForceRoundOver      WinType = 3
	WinTypeAllEntrustRoundOver WinType = 4
)
