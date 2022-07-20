package main

import (
	"fmt"
	"reflect"
)

type CommandWrapper struct {
	Data    interface{}
	Result  int
	Message string
}

type Command struct {
	rsp interface{}
}

var UserCommand = &Command{
	rsp: &struct {
		userInfo string
		id       int
	}{
		userInfo: "xuyongkang",
		id:       1,
	},
}

func fillValue(form *Command) {
	rspType := reflect.TypeOf(form.rsp).Elem()
	rsp := reflect.New(rspType).Interface()
	changeInterface(rsp)

}

func changeInterface(rsp interface{}) {
	oy := new(CommandWrapper)
	oy.Data = rsp
	//oy.Data = []int{1, 2, 3, 4, 5, 6, 7}
	//oy.Data = json.Unmarshal()
}

func workOnChangeInterface() {
	fillValue(UserCommand)
	fmt.Println(UserCommand.rsp)
}
