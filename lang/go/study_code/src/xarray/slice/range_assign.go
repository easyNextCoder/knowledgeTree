package xarray

import (
	"fmt"
	"reflect"
	"unsafe"
)

func MakeCopyAppendWork() {
	a := make([]int, 0, 100)
	b := a[:]
	b = append(b, []int{1, 2, 3}...)
	fmt.Printf("%d %d %p %p %v\n", len(a), cap(a), a, &a, a)
	fmt.Printf("%d %d %p %p %v\n", len(b), cap(b), b, &b, b)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("%x, %d, %d\n", cb.Data, cb.Cap, cb.Len)
}

func LiteralCopyAppendWork() {
	a := []int{0, 0}
	b := a[:]
	b = append(b, []int{1, 2, 3}...)
	fmt.Printf("%d %d %p %p %v\n", len(a), cap(a), a, &a, a)
	fmt.Printf("%d %d %p %p %v\n", len(b), cap(b), b, &b, b)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("%x, %d, %d\n", cb.Data, cb.Cap, cb.Len)
}

func copyAssignment() {
	a := make([]int, 2, 10)
	fmt.Printf("p(%p) by make\n", a)
	//a = []int{1, 2, 3}
	//fmt.Printf("p(%p) by literal\n", a)

	b := a[:]
	b[0] = 100

	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("%x, %d, %d\n", cb.Data, cb.Cap, cb.Len)
}

func makeAndLiteralAlloc() {
	a := make([]int, 0, 100)
	b := []int{1, 2, 3}
	c := []int{1, 2, 3}

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Println(ca.Data, ca.Cap, ca.Len)

	cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Println(cb.Data, cb.Cap, cb.Len)

	cc := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	fmt.Println(cc.Data, cc.Cap, cc.Len)

	for i := range []int{1, 2, 3} {
		fmt.Println("index:", i)
		a = append(a, []int{0, 0, 0, 0}...)
		ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
		fmt.Println(ca.Data, ca.Cap, ca.Len)

		b = append(b, []int{0, 0, 0, 0}...)
		cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		fmt.Println(cb.Data, cb.Cap, cb.Len)

	}

}
