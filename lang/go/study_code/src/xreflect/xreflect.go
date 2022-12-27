package xreflect

import (
	"fmt"
	"reflect"
)

func xreflect() {
	x := struct {
		v  [10]int
		vp []int
	}{}

	fmt.Println(reflect.ValueOf(x))
	fmt.Println(reflect.ValueOf(x.v))
	fmt.Println(reflect.ValueOf(x.vp))
	//fmt.Println(reflect.ValueOf(x.vp).Elem()) //value的elem相当于解引用，而如果value不是引用则会panic
	//fmt.Println(reflect.ValueOf(x.vp).Elem()) //value的elem相当于解引用，而如果value不是引用则会panic
	fmt.Println(reflect.New(reflect.TypeOf(x)).Elem())
	fmt.Println("willl")

	y := reflect.ValueOf(x)
	fmt.Println(y)
	//ytype := y.Type().Elem()//type的elem() 如果type不是几种arr， slice等类型panic
	//fmt.Println(ytype)

	//总结
	//Value.Elem() 解指针的引用
	//Type.Elem() 返回内部元素的类型
}
