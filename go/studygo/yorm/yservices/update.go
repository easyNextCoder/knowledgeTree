package yservices

import (
	"fmt"
	"strconv"
	"studygo/yorm/ydaos"
	"time"
)

func updateRows() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	var list []ydaos.User
	for i := 0; i < 50; i++ {
		list = append(list, ydaos.User{
			Id:   0,
			Uid:  int64(100 + i),
			Name: "xuyongkang",
		})
	}

	insert, err := session.Insert(&list)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("updateRows insert", insert)

	list = make([]ydaos.User, 0)
	err = session.Where("name=?", "xuyongkang").Find(&list)
	if err != nil {
		fmt.Println("find err", err)
		return
	}
	fmt.Println(list)

	for i := 0; i < 50; i++ {
		list[i].Name = "xuyongqing" + strconv.Itoa(i)
	}

	finalUpdate := 0
	for _, v := range list {
		update, err := session.Where("uid=?", v.Uid).Cols("name").Update(ydaos.User{Name: "xuyongqing"})
		if err != nil {
			fmt.Println("update err", update, err)
			return
		}
		finalUpdate++
	}

	fmt.Println("updateRows update", finalUpdate)

}

func Update1() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := new(ydaos.UserMessage)
	session.Where("Id = ?", 10010).Get(item)
	item.FromUid = 2
	item.SysKey = "2"
	item.TplId = 2
	item.Type = 2
	item.Created = time.Now().Add(-time.Hour*24*2 + 30*time.Minute)
	item.Updated = time.Now().Add(-time.Hour*24*2 + 30*time.Minute)
	item.DayReset = time.Now().Add(-time.Hour*24*2 + 30*time.Minute)
	//更新全部行的type和tpl_id字段（⚠️）
	affected, err := session.Cols("type, tpl_id").Update(item)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("affected rows is: ", affected)

	item.FromUid = 3
	item.SysKey = "3"
	item.TplId = 3
	item.Type = 3
	item.Created = time.Now().Add(-time.Hour*24*3 + 30*time.Minute)
	item.Updated = time.Now().Add(-time.Hour*24*3 + 30*time.Minute)
	item.DayReset = time.Now().Add(-time.Hour*24*3 + 30*time.Minute)
	//更新全部行的全部字段（⚠️）
	affected1, err := session.Update(item)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("affected rows is: ", affected1)
}

func UpdateNoAutoTime() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	thisMsg := new(ydaos.UserMessage)
	thisMsg.Uid = 10010
	thisMsg.FromUid = 0
	thisMsg.TplId = 0
	thisMsg.Type = 0
	thisMsg.Title = "test"
	thisMsg.ContentData = "test content"
	thisMsg.AttachmentData = "1"
	thisMsg.DayReset = time.Now().Add(-time.Hour*24*8 + time.Minute*30)
	thisMsg.Created = time.Now().Add(-time.Hour*24*8 + time.Minute*30)
	thisMsg.Updated = time.Now().Add(-time.Hour*24*8 + time.Minute*30)

	cnt, err := session.NoAutoTime().Update(thisMsg)
	if err != nil {
		fmt.Println("Test_NoAutoTime update failed")
	}
	fmt.Println("updated:", cnt)
}
