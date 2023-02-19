package xreflect

import (
	"fmt"
	"reflect"
)

type Indir struct {
	a int
	b []int
}

func indirectWork() {
	//type是struct的时候返回struct
	var va Indir
	t := reflect.ValueOf(va)
	rt := reflect.Indirect(t)
	fmt.Println(rt)

	//type是struct ptr的时候返回struct
	var vb Indir
	tb := reflect.ValueOf(&vb)
	rtb := reflect.Indirect(tb)
	fmt.Println(rtb)

}
