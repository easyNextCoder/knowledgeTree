package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func emptySliceWork() {
	var x []int
	x = append(x, 999)
	fmt.Println("var 申请", x[0]) //ok

	y := []int{}
	y = append(y, 888)
	fmt.Println(":= 申请", y[0]) //ok

	//解析slice的结构
	fmt.Printf("x直接地址(%p) x直接地址取地址(%p) y直接地址取地址(%p) y地址(%p)", x, &x, y, &y)
	xsh := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	ysh := (*reflect.SliceHeader)(unsafe.Pointer(&y))
	fmt.Printf("xslice结构: %x, %d, %d yslice结构: %x, %d, %d", xsh.Data, xsh.Len, xsh.Cap, ysh.Data, ysh.Len, ysh.Cap)
}
