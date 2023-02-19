package yservices

import (
	"fmt"
	"reflect"
	"strconv"
	"studygo/yorm/ydaos"
	"time"
)

func Get() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	item := &(ydaos.UserMessage{Id: 4})
	has, err := session.Where("id=?", 4).Get(item)
	if err != nil {
		fmt.Println("Get err", err)
	}
	fmt.Println(has, item)

	item = new(ydaos.UserMessage)
	item.Id = 4
	has, err = session.Get(item)
	if err != nil {
		fmt.Println("Get err1", err)
	}
	fmt.Println(has, item)
}

func Get1() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	var id, uid int64
	affected, err := session.Where("uid=?", 10010).Cols("id", "uid").Get(&id, &uid)
	if err != nil {
		fmt.Println("Get1 err", affected)
	}
}

type PartColumn struct {
	Uid      int
	Updated  string
	DayReset string
}

func GetSpecColumns() {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	querySql := "select uid,updated,day_reset from user_message"
	query, err := session.Query(querySql)
	if err != nil {
		return
	}

	var Res []PartColumn
	ParsePartColumn(&Res, query) //实现了Find的功能
	fmt.Println("GetSpecColumns ", Res)

	var Res1 []struct {
		Uid      int
		Updated  time.Time
		DayReset time.Time
	}
	querySql1 := "select uid,updated,day_reset from user_message"
	err1 := session.SQL(querySql1).Find(&Res1)
	if err1 != nil {
		return
	}

	fmt.Println("GetSpecColumns1 ", Res1)

}

func Join() (err error) {
	svc := NewSvc()
	defer svc.Close()
	session := svc.GetDaos().GetDb()

	querySql := "select distinct a.name, b.updated from user as a join user_message as b on a.uid = b.uid"
	query, err := session.Query(querySql)
	if err != nil {
		return
	}

	fmt.Println(query)

	var Res []struct {
		Name    string
		Updated string
	}

	ParsePartColumn(&Res, query)

	fmt.Println("Join ", Res)
	return
}

func ParsePartColumn(slicePointer interface{}, querys []map[string][]byte) {

	slice := reflect.ValueOf(slicePointer).Elem()
	elemType := slice.Type().Elem()
	nf := elemType.NumField()
	fieldMap := make(map[string]reflect.Type)
	nameMap := make(map[string]string)
	for i := 0; i < nf; i++ {
		fi := elemType.Field(i)
		fieldMap[fi.Name] = fi.Type

		var replace []byte
		for j := 0; j < len(fi.Name); j++ {
			c := fi.Name[j]
			if j == 0 {
				if c < 'a' && c > '9' {
					c = c + 32
				}
				replace = append(replace, c)
			} else {
				if fi.Name[j] >= 'A' && fi.Name[j] <= 'Z' {
					replace = append(replace, '_')
					replace = append(replace, c+32)
				} else {
					replace = append(replace, c)
				}
			}
		}
		nameMap[string(replace)] = fi.Name
	}

	slice.Set(reflect.MakeSlice(slice.Type(), len(querys), len(querys)))

	for i, mp := range querys {
		x := reflect.New(elemType)
		for fn, v := range mp {
			keyFn := nameMap[fn]
			tp := fieldMap[keyFn]
			switch tp.Kind() {
			case reflect.Int:
				intVal, _ := strconv.Atoi(string(v))
				x.Elem().FieldByName(keyFn).SetInt(int64(intVal))
			case reflect.String:
				x.Elem().FieldByName(keyFn).SetString(string(v))
			default:
				fmt.Println("convert err!")
			}
		}

		slice.Index(i).Set(x.Elem())
	}

}
