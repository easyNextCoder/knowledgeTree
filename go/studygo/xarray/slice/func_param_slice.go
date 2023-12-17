package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func appendOri(ori []int) {

	fmt.Printf("main sub start:%d %d %p %p %v\n", len(ori), cap(ori), ori, &ori, ori)

	ori = append(ori, 1)

	fmt.Printf("main sub:%d %d %p %p %v\n", len(ori), cap(ori), ori, &ori, ori)

	ca := (*reflect.SliceHeader)(unsafe.Pointer(&ori))
	fmt.Printf("%x, %d, %d\n", ca.Data, ca.Cap, ca.Len)

	fmt.Println("append", ori)
}

func sliceCopyWork() {

	ori := make([]int, 0, 0) //100拷贝时候底层data指针相同，0的时候append之后底层data指针不同

	fmt.Printf("main origin: %d %d %p %p %v\n", len(ori), cap(ori), ori, &ori, ori)
	appendOri(ori)

	fmt.Printf("main end:%d %d %p %p %v\n", len(ori), cap(ori), ori, &ori, ori)

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
