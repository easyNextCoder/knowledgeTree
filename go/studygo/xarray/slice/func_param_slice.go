package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func appendOri(ori []int) {
	ori = append(ori, 1)

	fmt.Println("sub function")

	fmt.Printf("%d %d %p %p %v\n", len(ori), cap(ori), ori, &ori, ori)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&ori))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	fmt.Println("append", ori)
}

func sliceCopyWork() {
	fmt.Println("main function")
	ori := make([]int, 0, 100)
	//100拷贝时候底层data指针相同，0的时候append之后底层data指针不同
	appendOri(ori)

	fmt.Printf("%d %d %p %p %v\n", len(ori), cap(ori), ori, &ori, ori)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&ori))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	fmt.Println("after append", ori)
}

func passNilSlice(x []*int) {
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("passNilSlice done")
}

func workPassNilSlice() {
	passNilSlice(nil)
}
