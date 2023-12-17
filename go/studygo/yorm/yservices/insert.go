package yservices

import (
	"fmt"
	"math/rand"
	"studygo/yorm/ydaos"
)

func InsertUser(item ...*ydaos.User) error {

	svc := NewSvc()
	defer svc.Close()

	session := svc.GetDaos().GetDb()

	insert, err := session.Insert(item)
	if err != nil {
		return err
	}

	if insert != int64(len(item)) {
		return fmt.Errorf("insert num(%d) != input item num(%d)", insert, len(item))
	}

	return nil
}

func Insert(n int) {

	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	um0 := &ydaos.UserMessage{Uid: 10009}
	um1 := &ydaos.UserMessage{Uid: 10010}
	um2 := &ydaos.UserMessage{Uid: 10011}
	affected, err := session.Insert(um0)

	fmt.Println("insert um0")
	if err != nil {
		fmt.Println("Insert err", err)
	}

	affected1, err1 := session.Insert(um1, um2)
	fmt.Println("insert um1 um2", affected1)
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
	fmt.Println("insert user0 user1 um3")
	if err2 != nil {
		fmt.Println("Insert err2", err2, affected2)
	}

}

func insertMulti() {

	svc := NewSvc()
	defer svc.Close()

	session := svc.GetDaos().GetDb()

	users := []*ydaos.User{
		&ydaos.User{Uid: rand.Int63(), Name: "tom", Age: 19},
		&ydaos.User{Uid: rand.Int63(), Name: "bob", Age: 21},
		&ydaos.User{Uid: rand.Int63(), Name: "hank", Age: 21},
		&ydaos.User{Uid: rand.Int63(), Name: "jim", Age: 19},
		&ydaos.User{Uid: rand.Int63(), Name: "alex", Age: 39},
		&ydaos.User{Uid: rand.Int63(), Name: "smith", Age: 43},
	}

	multi, err := session.InsertMulti(users)
	if err != nil {
		return
	}

	if err != nil {
		fmt.Printf("insertMulti err(%s)", err)
		return
	}

	fmt.Printf("insertMulti num %d", multi)

	for _, v := range users {
		fmt.Printf("user id(%d) uid(%d) name(%s) age(%d)\n", v.Id, v.Uid, v.Name, v.Age)
	}

}

func insert() {

	svc := NewSvc()
	defer svc.Close()

	session := svc.GetDaos().GetDb()

	item := &ydaos.User{Uid: rand.Int63(), Name: "tom", Age: 19}

	insert, err := session.Insert(item)
	if err != nil {
		return
	}
	if err != nil {
		fmt.Printf("insert err(%s)", err)
		return
	}

	fmt.Printf("insert num %d", insert)

	fmt.Printf("user id(%d) uid(%d) name(%s) age(%d)\n", item.Id, item.Uid, item.Name, item.Age)
}
