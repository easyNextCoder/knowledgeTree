package xvar

import (
	"fmt"
	"reflect"
	"unsafe"
)

func xvar() {
	var x []int
	var y [10]int
	fmt.Println(reflect.TypeOf(&x), reflect.TypeOf(y))
	cc := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	fmt.Println(cc.Data, cc.Cap, cc.Len)

	x = append(x, 10)
	cca := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	fmt.Println(cca.Data, cca.Cap, cca.Len)

	x = make([]int, 0)
	ccn := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	fmt.Println(ccn.Data, ccn.Cap, ccn.Len)

	z := y[1:5]
	z[1] = 5
	fmt.Println(y, z)
	a := y
	a[0] = 10
	fmt.Println(y, a)

	x = append(x, 11)
	x = append(x, 12)
	b := x
	fmt.Println(x, b)
	b[0] = 99
	fmt.Println(x, b)
	c := b[1:2]
	c[0] = 999
	fmt.Println(x, b, c)

	var d = struct {
		arr  [10]int
		parr []int
	}{}
	fmt.Println(d.parr, &d.parr)
	d.parr = append(d.parr, 12)
	fmt.Println(d)

}
