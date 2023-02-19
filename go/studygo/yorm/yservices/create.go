package yservices

import (
	"fmt"
	"studygo/yorm/ydaos"
)

func Insert(n int) {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	um0 := &ydaos.UserMessage{Uid: 10009}
	um1 := &ydaos.UserMessage{Uid: 10010}
	um2 := &ydaos.UserMessage{Uid: 10011}
	affected, err := session.Insert(um0)
	if err != nil {
		fmt.Println("Insert err", err)
	}
	affected1, err1 := session.Insert(um1, um2)
	if err1 != nil {
		fmt.Println("Insert err1", err1)
	}
	if affected+affected1 != 3 {
		fmt.Println("Insert affect row err", affected, affected1)
	}

	user0 := &ydaos.User{Uid: 10009, Name: "user0"}
	user1 := &ydaos.User{Uid: 10010, Name: "user1"}
	um3 := &ydaos.UserMessage{Uid: 10012}
	affected2, err2 := session.Insert(user0, user1, um3)
	if err2 != nil {
		fmt.Println("Insert err2", err2, affected2)
	}

}
