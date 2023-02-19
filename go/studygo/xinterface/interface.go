package xinterface

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

//新
type xCommandRsp interface {
	Work()
}

type xCommandRspChild struct {
	data [10]int
}

func (xCommandRspChild) Work() {

}

type xFinalRsp struct {
	rsp chan xCommandRsp
}

func testRsp() {
	v := xFinalRsp{
		rsp: make(chan xCommandRsp),
	}
	go func() {
		val := <-v.rsp
		fmt.Println(val)
	}()

	v.rsp <- &xCommandRspChild{data: [10]int{1, 2, 3}}

	var cm xCommandRspChild
	cm.Work()

}
