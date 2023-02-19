package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

//Notice:
//append可能会生成新数组

func sliceCut() {
	a := make([]int, 0, 100)

	a = append(a, []int{1, 2, 3}...)
	b := a[2:] //b的内部指针不再与a相同，而是向后移动对应的字节数
	fmt.Printf("%d %d %p %p %v\n", len(a), cap(a), a, &a, a)
	fmt.Printf("%d %d %p %p %v\n", len(b), cap(b), b, &b, b)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("%v %x, %d, %d\n", ca, ca.Data, ca.Cap, ca.Len)

	cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("%v %x, %d, %d\n", cb, cb.Data, cb.Cap, cb.Len)
}

func MakeCopyAppendWork() {
	a := make([]int, 0, 100)
	b := a[:]
	b = append(b, []int{1, 2, 3}...) //a b 内部的数据的指针仍然相同

	fmt.Printf("%d %d %p %p %v\n", len(a), cap(a), a, &a, a)
	fmt.Printf("%d %d %p %p %v\n", len(b), cap(b), b, &b, b)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("%x, %d, %d\n", cb.Data, cb.Cap, cb.Len)
}

func LiteralCopyAppendWork() {
	a := []int{0, 0} //len == caps == 2
	b := a[:]
	b = append(b, []int{1, 2, 3}...) //append改变了a的长度，因此生成新的数组
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
		a = append(a, []int{0, 0, 0, 0}...) //append并未改变caps所以，指针一只不变
		ca := (*reflect.SliceHeader)(unsafe.Pointer(&a))
		fmt.Println(ca.Data, ca.Cap, ca.Len)

		b = append(b, []int{0, 0, 0, 0}...) //append改变了数组容量因此，内部重新申请内存，产生新的地址
		cb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		fmt.Println(cb.Data, cb.Cap, cb.Len)

	}

}
