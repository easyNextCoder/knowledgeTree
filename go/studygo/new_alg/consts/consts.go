package consts

import (
	"time"
)

const (
	Hand_Crush_Score     = 50
	Hand_Max_Project_Len = 5
)

const (
	Entrust_Open  = true
	Entrust_Close = false
)

//ai version
const (
	AiStableVersion = iota
	AiTestVersion
)

// entrust time
const (
	NormalEntrustInterval = time.Second * 15
)

// 托管触发原因
const (
	EntrustTriggerReason_ManualOperation = iota
	EntrustTriggerReason_TimeOut
)

// pack roomInfo type
const (
	RoomInfoType_Join = iota
	RoomInfoType_Reenter
)

//结算
const (
	SumType_Normal = iota
	SumType_Bankrupt
	SumType_Cap
)

//聊天类型
const (
	CHAT_NONE = iota
	CHAT_WORDS
	CHAT_VOICE
	CHAT_EMOJI

	CHAT_TEMPLATE = 7
	//如果再7的后边定义常量必须设定等于几，否则都等于7
)

//强制流局时间常量
const (
	MaxForceRoundOverTimeDuration        = time.Minute * 60
	AllEntrustForceRoundOverTimeDuration = time.Minute * 15
)

//客户端组合展示顺序类型
const (
	DisplayOrder_Asc = iota
	DisplayOrder_Des
)

//api 相关
const (
	JoinType_Normal = iota
	JoinType_Reenter
)

//控牌类型
const (
	CardCtrlType_NoCtrl = iota
	CardCtrlType_Test
	CardCtrlType_Warm
)

//robot entrust action
const (
	Action_Pick = iota
	Action_Crush
	Action_Add
	Action_Drop
	Action_AddSp
)
