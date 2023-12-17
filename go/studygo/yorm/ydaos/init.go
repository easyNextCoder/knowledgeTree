package ydaos

import (
	"fmt"
	"reflect"
	"time"
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

type iSession interface {
	Insert(beans ...interface{}) (int64, error)
}

type Decorator struct {
	before func(...interface{}) (interface{}, error)
	after  func(...interface{}) (interface{}, error)
}

type XSession struct {
	*xorm.Session
	iSession
	decorators []Decorator
}

func (s *XSession) Insert(beans ...interface{}) (int64, error) {
	start := time.Now()

	d := reflect.ValueOf(*s)

	re := d.Field(0)

	realSession, ok := re.Interface().(*xorm.Session)
	//
	if !ok {
		fmt.Println("insert err", realSession)
		return 0, nil
	}

	a, b := realSession.Insert(beans...)

	end := time.Now()
	fmt.Println("cost time is:", end.Sub(start).Milliseconds())

	//return a, b

	fmt.Println("value type", re.Type(), realSession, a, b)

	return 0, nil
}
