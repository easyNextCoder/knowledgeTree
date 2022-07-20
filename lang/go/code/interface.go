package main

import (
	"fmt"
	"reflect"
)

type OtherYork struct {
	Data interface{}
}

type Command struct {
	rsp interface{}
}

var UserCommand = &Command{
	rsp: &struct {
		userInfo string
		id       int
	}{},
}

func fillValue(form *Command) {
	rspType := reflect.TypeOf(form.rsp).Elem()
	rsp := reflect.New(rspType).Interface()
	changeInterface(rsp)

}

func changeInterface(rsp interface{}) {
	oy := new(OtherYork)
	oy.Data = rsp
	oy.Data = []int{1, 2, 3, 4, 5, 6, 7}
}

func workOnChangeInterface() {
	fillValue(UserCommand)
	fmt.Println(UserCommand.rsp)
}
