package xinterface

import (
	"fmt"
	"reflect"
	"time"
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

func changeInterface(rsp interface{}) {
	oy := new(OtherYork)
	oy.Data = []int{1, 2, 3, 4, 5, 6, 7}
	*(rsp.(*interface{})) = oy
}

func fillValue(from *Command) {
	rspType := reflect.TypeOf(from.rsp).Elem()
	rsp := reflect.New(rspType).Interface()
	changeInterface(&rsp) //interface{}的指针仍然是interface{}类型可以通过.(*interface{})来恢复
	from.rsp = rsp

}

func workOnChangeInterface() {
	fillValue(UserCommand)
	fmt.Println(UserCommand.rsp)
}

//测试cmd{rsp:chan<-struct}结构

type xCommand struct {
	data interface{}
	rsp  chan xCommandRsp
	prsp chan *xCommandRsp
}

type xInputData struct {
	CommandId   int
	CommandName string
}

type xCommandRsp interface {
	GetData()
}

type xCommandRspChild struct {
	data      [10]int
	childName string
}

func (self xCommandRspChild) GetData() {
	fmt.Println("i am xCommandRspChild, i am GetData", self.data, self.childName)
}

func (self *xCommandRspChild) Hello() {
	fmt.Println("i am xCommandRspChild, i am Hello", self.childName)
}

func executeCommand(cmd *xCommand) {
	fmt.Println("executing command...")
	val := cmd.data.(xInputData)
	fmt.Printf("input data: commandId(%d) commandName(%s)\n", val.CommandId, val.CommandName)
	res := &xCommandRspChild{data: [10]int{1, 2, 3}, childName: "execute done!"}
	fmt.Printf("res data(%p) childName(%p)\n", &res.data, &res.childName)
	cmd.rsp <- res

}
func executeCommandRspStruct(cmd *xCommand) {
	fmt.Println("executing command...")
	val := cmd.data.(xInputData)
	fmt.Printf("input data: commandId(%d) commandName(%s)\n", val.CommandId, val.CommandName)
	res := xCommandRspChild{data: [10]int{1, 2, 3}, childName: "execute done!"}
	fmt.Printf("res data(%p) childName(%p)\n", &res.data, &res.childName)
	cmd.rsp <- res

}

func readCommandRsp(cmd *xCommand) interface{} {
	val := <-cmd.rsp
	return val
}

func testRsp() { //rsp<-struct

	cmd := &xCommand{
		data: xInputData{
			CommandId:   1001,
			CommandName: "登陆",
		},
		rsp:  make(chan xCommandRsp),
		prsp: nil,
	}

	go executeCommandRspStruct(cmd)

	i := readCommandRsp(cmd)

	obj := i.(xCommandRspChild)

	fmt.Printf("readCommandRsp: %+v\ngot data(%p) childName(%p)\n", i, &obj.data, &obj.childName)

	time.Sleep(time.Second * 2)
}

func testRsp2() {
	cmd := &xCommand{
		data: xInputData{
			CommandId:   1002,
			CommandName: "签到",
		},
		rsp:  make(chan xCommandRsp),
		prsp: nil,
	}

	go executeCommand(cmd)

	i := readCommandRsp(cmd)

	obj := i.(*xCommandRspChild)

	obj.GetData()

	fmt.Printf("readCommandRsp: %+v\ngot data(%p) childName(%p)\n", i, &obj.data, &obj.childName)

	time.Sleep(time.Second * 2)
}
