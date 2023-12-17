package xreflect

import (
	"fmt"
	"reflect"
)

type (
	ReflectCar struct {
		int
		string
		intVal int
	}
)

// TypeOfFunc TypeOf 返回interface{}
func TypeOfFunc() {
	//v := reflect.Value{}
	var intI = 0
	ty := reflect.TypeOf(intI)
	fmt.Println("type:", ty)

	var arrInt []int
	v := reflect.TypeOf(arrInt)
	fmt.Println("type:", v, "elem:", v.Elem())

	var arrStruct []ReflectCar
	v = reflect.TypeOf(arrStruct)
	fmt.Println("type:", v, "elem:", v.Elem())

	var b ReflectCar
	b.int = 10
	fmt.Printf("ptr: b(%p) b.int(%p) b.string(%p) b.intVal(%p)\n", &b, &b.int, &b.string, &b.intVal)

}

// ValueOfFunc ValueOf 返回struct
func ValueOfFunc() {

	var b ReflectCar
	r := reflect.ValueOf(b)
	rr := r

	tp := reflect.TypeOf(r)

	fmt.Printf("reflect.ValueOf %+v %p %p %p %s\n", r, &b, &r, &rr, tp)

}
