package yservices

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"studygo/yorm/ydaos"
	"time"
)

type Svc struct {
	daos *ydaos.Daos
}

func NewSvc() *Svc {
	return &Svc{
		daos: ydaos.NewDaos(),
	}
}

func (svc *Svc) GetDaos() *ydaos.Daos {
	return svc.daos
}

func (svc *Svc) Close() {
	svc.daos.Close()
}

func init() {
	fmt.Println("db init...")
	Sync()
}

func Sync() {
	svc := NewSvc()
	defer svc.Close()

	session := svc.GetDaos().GetDb()
	if err := session.Sync2(new(ydaos.UserMessage), new(ydaos.User)); err != nil {
		fmt.Println("数据表同步失败:", err)
	}
}

//创建orm引擎

func Delete() {
	Delete1()
}

func Update() {
	//UpdateNoAutoTime()
	updateRows()
}

func Create() {
	Insert(10)
}

func Read() (err error) {

	Get()
	Get1()
	GetSpecColumns()
	err = Join()
	if err != nil {
		fmt.Println("join err", err)
	}
	return
}

func block() {

	go func() {
		dao := NewSvc()
		defer dao.Close()

		session := dao.GetDaos().GetDb()

		session.Begin()

		var users []*ydaos.User

		session.Select("*").Where("age>?", 20).ForUpdate().Find(&users)

		//session.Commit()

		fmt.Println("select users are:", users)
	}()

	//go func() {
	//	dao := NewSvc()
	//	defer dao.Close()
	//
	//	session := dao.GetDaos().GetDb()
	//
	//	session.Begin()
	//
	//	session.Where("id=?", )
	//
	//}()

	<-time.After(time.Second * 20)

}

func UseYorm() {
	//DeleteAll()
	//Create()
	//Read()
	//Update()
	//Delete()

	//insertMulti()
	block()
}
