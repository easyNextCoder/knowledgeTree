package ydaos

import (
	"fmt"
	"xorm.io/xorm"
)

type Daos struct {
	db *xorm.Session
}

func NewDaos() *Daos {
	return &Daos{}
}

func (daos *Daos) GetDb() *xorm.Session {
	if daos.db == nil {
		db, err := xorm.NewEngine("mysql", "root:1234@tcp(0.0.0.0:3336)/test?charset=utf8")
		if err != nil {
			fmt.Println("NewEngine failed.")
		}
		db.Get()
		daos.db = db.NewSession()
	}
	return daos.db
}

func (daos *Daos) Close() {
	daos.db.Close()
}
