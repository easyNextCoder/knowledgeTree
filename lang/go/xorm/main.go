package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm/daos"
)

type UserMessage struct {
	Id             int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Uid            int64     `xorm:"not null INT(11) index"`
	FromUid        int64     `xorm:"not null BIGINT(20)"`    // 发送者uid，系统邮件为0
	SysKey         string    `xorm:"not null VARCHAR(1024)"` // 系统邮件的唯一标识
	TplId          int64     `xorm:"not null BIGINT(20)"`
	Type           int       `xorm:"not null INT(11) DEFAULT 0"`        // 消息类型
	Title          string    `xorm:"not null text"`                     // 消息标题
	ContentData    string    `xorm:"not null text"`                     // 当TplId非0时，为模版参数json，TplId为0时，为消息体本身
	AttachmentData string    `xorm:"not null VARCHAR(1024) DEFAULT ''"` // 附件内容描述
	Args           string    `xorm:"not null VARCHAR(1024) DEFAULT ''"` // 用户客户端的一些附加的参数，比如跳转
	LogTag         int       `xorm:"not null INT(11) DEFAULT 0"`
	Status         int       `xorm:"not null INT(11) DEFAULT 1"` // 消息状态，值参考user_svr/app/common/consts的MessageStatusEnum
	DayReset       time.Time `xorm:"not null DATETIME"`
	Created        time.Time `xorm:"not null DATETIME created index"`
	Updated        time.Time `xorm:"not null DATETIME updated"`
	Deleted        time.Time `xorm:"deleted index"`
}

type Svc struct {
	daos *daos.Daos
}

func NewSvc() *Svc {
	return &Svc{
		daos: daos.NewDaos(),
	}
}

func (svc *Svc) GetDaos() *daos.Daos {
	return svc.daos
}

func (svc *Svc) Close() {
	svc.daos.Close()
}

func p(x ...interface{}) {

	fmt.Println(x)

}

//创建orm引擎
func Test_NoAutoTime() {
	svc := NewSvc()
	defer svc.Close()

	session := svc.GetDaos().GetDb()
	if err := session.Sync2(new(UserMessage)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}

	thisMsg := new(UserMessage)
	thisMsg.Uid = 23050755
	thisMsg.FromUid = 0
	thisMsg.TplId = 0
	thisMsg.Type = 0
	thisMsg.Title = "test"
	thisMsg.ContentData = "test content"
	thisMsg.AttachmentData = "1"
	thisMsg.DayReset = time.Now().Add(-time.Hour*24*8 + time.Minute*30)
	thisMsg.Created = time.Now().Add(-time.Hour*24*8 + time.Minute*30)
	thisMsg.Updated = time.Now().Add(-time.Hour*24*8 + time.Minute*30)
	session.NoAutoTime().Insert(thisMsg)
	//cnt, err := session.Update(thisMsg)
	//if err != nil {
	//	fmt.Println("Test_NoAutoTime update failed")
	//}
	//fmt.Println("updated:", cnt)
}

func Test_Delete() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := new(UserMessage)
	item.Id = 4
	session.Get(item)

	//删除id=5的行
	affected1, err := session.Where("id=?", 5).Delete(new(UserMessage))
	if err != nil {
		fmt.Println("Test_Delete err", err)
	}
	fmt.Println("affected rows is %d", affected1)

	//删除所有的行(⚠️)
	//affected, err := session.Delete(item)
	//if err != nil {
	//	fmt.Println("Test_Delete err", err)
	//}
	//fmt.Println("affected rows is %d", affected)

}

func Test_Get() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := &(UserMessage{Id: 4})
	has, err := session.Where("id=?", 4).Get(item)
	if err != nil {
		fmt.Println("Test_Get err")
	}
	fmt.Println(has, item)

	item = new(UserMessage)
	item.Id = 4
	has, err = session.Get(item)
	if err != nil {
		fmt.Println("Test_Get err")
	}
	fmt.Println(has, item)

}

func Test_Update() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := new(UserMessage)
	session.Where("Id = ?", 4).Get(item)
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

func main() {
	//Test_NoAutoTime()
	//Test_Delete()
	//Test_Get()
	//Test_Update()

	fmt.Println("main done!")
}
