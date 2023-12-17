package xinterface

import "fmt"

type (
	Reaction struct {
		SelfUid int64       `json:"self_uid"`
		Val     interface{} `json:"val"`
		ToUid   int64       `json:"to_uid"`
	}
	Message struct {
		Type int    `json:"type"`
		Msg  string `json:"msg"`
	}

	Prop struct {
		PropId  int `json:"prop_id"`
		PropNum int `json:"prop_num"`
	}
)

var (
	AiReactions [6][]Reaction = [6][]Reaction{
		{{Val: Message{Type: 3, Msg: "Fox-Happy"}}, {Val: &Message{Type: 3, Msg: "Fox-Love"}}, {Val: Message{Type: 3, Msg: "Fox-Nice"}}, {Val: Message{Type: 7, Msg: "hand_room_quick_chat_lang.langHandRoomQuickChat6"}}},
		{{Val: Message{Type: 3, Msg: "Fox-Cry"}}, {Val: Message{Type: 3, Msg: "Fox-Angry"}}, {Val: Message{Type: 3, Msg: "Fox-Amazed"}}, {Val: Message{Type: 3, Msg: "Fox-Sad"}}, {Val: Message{Type: 7, Msg: "hand_room_quick_chat_lang.langHandRoomQuickChat7"}}},
		{{Val: Prop{PropId: 30004, PropNum: 1}}, {Val: Message{Type: 7, Msg: "hand_room_quick_chat_lang.langHandRoomQuickChat8"}}},
		{{Val: Message{Type: 3, Msg: "Fox-Thank"}}, {Val: Message{Type: 3, Msg: "Fox-Ok"}}, {Val: Message{Type: 3, Msg: "Fox-Nice"}}, {Val: Message{Type: 3, Msg: "Fox-Love"}}, {Val: Message{Type: 7, Msg: "hand_room_quick_chat_lang.langHandRoomQuickChat6"}}},
		{{Val: Prop{PropId: 30002, PropNum: 1}}, {Val: Prop{PropId: 30003, PropNum: 1}}},
		{{Val: Prop{PropId: 30004, PropNum: 1}}, {Val: Prop{PropId: 30005, PropNum: 1}}, {Val: Prop{PropId: 30006, PropNum: 1}}},
	}
)

func assignmentInterfaceWork() {
	rea := AiReactions[0][1]
	rea.ToUid = 100
	rea.SelfUid = 999

	fmt.Printf("%+v %p %p %+v\n", rea, &rea.Val, &rea.Val, AiReactions[0][1])
	fmt.Printf("%p %p\n", &rea, &AiReactions[0][1])

	newInterface := AiReactions[0][1].Val
	v := newInterface.(*Message)
	v.Msg = "100"
	fmt.Printf("%+v %+v %+v\n", v, newInterface, AiReactions[0][1].Val)

}
