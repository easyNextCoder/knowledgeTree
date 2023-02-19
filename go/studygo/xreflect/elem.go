package xreflect

import (
	"fmt"
	"reflect"
)

type XElem struct {
	a int
	b string
	c []int
	d [10]string
}

func typeElem() {
	var va XElem
	fmt.Println(reflect.ValueOf(va).Elem())
}

func valueElem() {
	var va XElem
	fmt.Println(reflect.TypeOf(&va).Elem())
}
