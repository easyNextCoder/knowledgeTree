package xreflect

import (
	"fmt"
	"reflect"
)

type (
	s struct {
		int
		string
		intVal int
	}
)

func buildin_func() {
	//v := reflect.Value{}
	var int_i = 0
	ty := reflect.TypeOf(int_i)
	fmt.Println(ty)

	var arr_int []int
	v := reflect.TypeOf(arr_int)
	fmt.Printf("arr %v  item of arr %v\n", v, v.Elem())

	var arr_struct []s
	v = reflect.TypeOf(arr_struct)
	fmt.Printf("arr_struct %v item of arr_struct %v\n", v, v.Elem())
	var b s
	b.int = 10
	fmt.Printf("b.p(%p) b.int.p(%p) b.string.p(%p) b.intVal.p(%p)\n", &b, &b.int, &b.string, &b.intVal)

	r := reflect.ValueOf(b)
	rr := r

	fmt.Printf("reflect.ValueOf %p %p %p\n", &b, &r, &rr)
	//vl := reflect.ValueOf(i)
	//if vl.Kind() == reflect.Array {
	//
	//}
}
