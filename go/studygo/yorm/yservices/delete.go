package yservices

import (
	"fmt"
	"studygo/yorm/ydaos"
)

func DeleteAll() error {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	deleteSql := "delete  from user_message"

	var res []ydaos.UserMessage
	err := session.SQL(deleteSql).Find(&res) //不加find是不会执行这个sql语句的,但是直接Query的时候会执行sql语句
	lastSql, _ := session.LastSQL()
	if err != nil {
		return err
	}

	fmt.Println("delete all lastsql", lastSql)

	var res2 []ydaos.User
	deleteUserSql := "delete from user"
	err1 := session.SQL(deleteUserSql).Find(&res2)
	lastSql, _ = session.LastSQL()
	if err1 != nil {
		return err1
	}

	fmt.Println("delete all lastsql", lastSql)

	return nil
}

func Delete1() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := new(ydaos.UserMessage)
	item.Id = 4
	get, err := session.Get(item)
	if err != nil {
		fmt.Println("Delete1 err", get)
		return
	}

	if get {
		acd, err := session.Delete(item)
		if err != nil {
			fmt.Println("Delete1 err1", acd)
			return
		}
	}

}

func deleteAll() error {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := &ydaos.User{Id: 136}

	deleteNum, err := session.Delete(item)
	if err != nil {
		return fmt.Errorf("delete err(%s)", err)
	}

	fmt.Printf("deleteAll num(%d)", deleteNum)
	return nil
}
