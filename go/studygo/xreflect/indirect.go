package xreflect

import (
	"fmt"
	"reflect"
)

type India struct {
	a int
	b []int
}

func indirectWork() {

	//type是struct的时候返回struct
	var va India
	t := reflect.ValueOf(va)
	rt := reflect.Indirect(t)
	fmt.Println(rt)

	//type是struct ptr的时候返回struct
	var vb India
	tb := reflect.ValueOf(&vb)
	rtb := reflect.Indirect(tb)
	fmt.Println(rtb)

}
