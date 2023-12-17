package main

import (
	"fmt"
	"time"
)

type UserProp struct {
	Uid     int64     `xorm:"not null INT(11) unique(unique) index index(pair)"`
	PropId  int       `xorm:"not null INT(11) unique(unique) index index(pair)"`
	Count   int       `xorm:"not null INT(11)"`
	Created time.Time `xorm:"not null DATETIME created"`
	Updated time.Time `xorm:"not null DATETIME updated"`
}

type vipProp UserProp

func main() {
	v := &UserProp{
		Uid:     1,
		PropId:  2,
		Count:   3,
		Created: time.Now(),
		Updated: time.Now(),
	}

	p := vipProp(*v)

	fmt.Println("p is:", p)
}
