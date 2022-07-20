package daos

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

type Daos struct {
	db *xorm.Session
}

func NewDaos() *Daos {
	return &Daos{}
}

func (daos *Daos) GetDb() *xorm.Session {
	if daos.db == nil {
		db, err := xorm.NewEngine("mysql", "root:mM13137276827_@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			fmt.Println("NewEngine failed.")
		}
		daos.db = db.NewSession()
	}
	return daos.db
}

func (daos *Daos) Close() {
	daos.db.Close()
}
