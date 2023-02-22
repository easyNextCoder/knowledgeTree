package yservices

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"studygo/yorm/ydaos"
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

func UseYorm() {
	DeleteAll()
	Create()
	Read()
	Update()
	//Delete()
}
